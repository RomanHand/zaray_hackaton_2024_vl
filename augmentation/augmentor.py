import os
from PIL import Image, ImageFilter
import random
from tqdm import tqdm

import Augmentor

OUTPUT_DIR = "augmentation/images/output"
SOURCE_DIR = "augmentation/images"

p = Augmentor.Pipeline(SOURCE_DIR)
p.random_distortion(probability=0.4, grid_width=100, grid_height=100, magnitude=1)
p.flip_left_right(probability=0.4)
p.sample(30)

image_files = [
    f for f in os.listdir(OUTPUT_DIR) if f.endswith((".jpg", ".png", ".jpeg"))
]
images_to_blur = random.sample(image_files, len(image_files) // 3)

with tqdm(total=len(images_to_blur), desc="Bluring Images", unit=" Samples") as progress_bar:
    for image_file in images_to_blur:
        image_path = os.path.join(OUTPUT_DIR, image_file)
        image = Image.open(image_path)
        blurred_image = image.filter(ImageFilter.GaussianBlur(radius=round(random.uniform(0.5,5), 2)))
        blurred_image.save(image_path)
        progress_bar.update(1)
