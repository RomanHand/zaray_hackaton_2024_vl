import cv2
import sys
import os
import json
from ultralytics import YOLO
from math import floor
from functools import reduce
import torch


class Config:
    """
    Конфигурация для работы с фреймами
        - CONF_THRESHOLD: Порог вероятности фиксации нарушения
        - MIN_FRAME_AGE: Количество прошедших кадров для обнаружения нарушения
        - FPS_SPLIT: Количество кадров на 1 секунду видео
        - MIN_INTERVAL_DISTANCE - Минимальное расстояние в секундах между 
        интервалами, чтобы считать их разными
        - LIFE_TIME: Допустимый интервал утечки ошибки из кадра
    """

    def __init__(self, conf_threshold=0.5, min_frame_age=3, fps_split=3, life_time=2):
        self.CONF_THRESHOLD = conf_threshold
        self.MIN_FRAME_AGE = min_frame_age
        self.FPS_SPLIT = fps_split
        self.MIN_INTERVAL_DISTANCE = 5
        self.LIFE_TIME = life_time


def glue_last_intervals(intervals: list):
    intervals[-2][1] = intervals[-1][0]
    intervals.remove(intervals[-1])


def write_previews(video_name, cap, fps, intervals):
    os.makedirs('previews', exist_ok=True)
    for i, interval in enumerate(intervals):
        cap.set(cv2.CAP_PROP_POS_FRAMES, interval[0] * fps)
        _, frame = cap.read()
        out_path = f"previews/{video_name}_{i}.png"
        cv2.imwrite(out_path, frame)


def format_seconds(seconds):
    return f"{floor(seconds) // 60:02}:{floor(seconds) % 60:02}"


def process_file(video_name, path_source, model, config: Config):
    cap = cv2.VideoCapture(path_source)

    frame_age = 0
    frame_pos = 0
    prev_detected = 0

    frame_count = cap.get(cv2.CAP_PROP_FRAME_COUNT)
    fps = cap.get(cv2.CAP_PROP_FPS)
    frames_increment = fps / config.FPS_SPLIT
    print(f'TOTAL: {frame_count}')

    intervals = []

    while frame_pos < floor(frame_count):
        cap.set(cv2.CAP_PROP_POS_FRAMES, frame_pos)
        _, frame = cap.read()
        results = model.predict(frame, verbose=False)

        confs = reduce(lambda prev, cur: prev + cur.boxes.conf.detach().cpu().tolist(), results, [])
        if len(list(filter(lambda conf: conf > config.CONF_THRESHOLD, confs))) >= 1:
            frame_age += 1
            prev_detected = frame_pos
        elif frame_age >= config.MIN_FRAME_AGE and (frame_pos - prev_detected) / fps > config.LIFE_TIME:
            interval_end = (frame_pos - frames_increment) / fps
            intervals[-1].append(interval_end)
            prev_detected = False
            frame_age = 0

            if len(intervals) > 1 and intervals[-1][0] - intervals[-2][1] <= config.MIN_INTERVAL_DISTANCE:
                glue_last_intervals(intervals)

        if frame_age >= config.MIN_FRAME_AGE and (len(intervals) == 0 or len(intervals[-1]) != 1):
            interval_start = (frame_pos - (frame_age - 1) * frames_increment) / fps
            intervals.append([interval_start])

        frame_pos += frames_increment

    if len(intervals) > 1 and len(intervals[-1]) == 1:
        intervals[-1].append(frame_count / fps)

    for i in range(len(intervals)):
        intervals[i] = list(map(format_seconds, intervals[i]))

    write_previews(video_name, cap, fps, intervals)
    cap.release()

    return intervals


def write_json(intervals: list, video_name, path_source):
    violations = list()

    for i in range(len(intervals)):
        violation = dict()
        violation["preview"] = f"previews/{video_name}_{i}.png"
        violation["start"] = round(intervals[i][0], 2)
        violation["end"] = round(intervals[i][1], 2)
        violations.append(violation)

    json_data = {"name": video_name, "path_source": path_source, "violations": violations}

    with open('data.json', 'w') as f:
        json.dump(json_data, f)


def main(video_name, path_source, model_path, config_args=None):
    if config_args is None:
        config = Config()
    else:
        for key in config_args.keys():
            config_args[key] = float(config_args[key])
        config = Config(*config)

    device = 'cuda' if torch.cuda.is_available() else 'cpu'
    model = YOLO(model_path).to(device)

    intervals = process_file(video_name, path_source, model, config)
    write_json(intervals, video_name, path_source)


if __name__ == '__main__':
    main(*map(str, sys.argv[1:4]), 
        **dict(arg.split('=') for arg in sys.argv[5:])
    )
