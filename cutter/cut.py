import argparse
import ffmpeg

# python3 cut.py -f cars.mp4 -t "[[0, 10 ], [30, 40]]"  
parser = argparse.ArgumentParser(description='Резак видосов по timestamp')
parser.add_argument('-f', '--file', type=str, help='Путь до исходного видео')
parser.add_argument('-t', '--timestamps', type=str, help='Массив временных отрезков в формате "[[начало_1, конец_1], [начало_2, конец_2] ...]"')
args = parser.parse_args()

if args.file is None or args.timestamps is None:
    print("Необходимо указать путь до файла и временные отрезки.")
    exit()

video_path = args.file
timestamps = eval(args.timestamps)

for idx, timestamp in enumerate(timestamps):
    start_time = timestamp[0]
    end_time = timestamp[1]
    output_file = f'video_clip_{idx+1}.mp4'

    stream = ffmpeg.input(video_path, ss=start_time, to=end_time)
    stream = ffmpeg.output(stream, output_file)
    ffmpeg.run(stream, overwrite_output=True, quiet=True)


