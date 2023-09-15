import cv2
import numpy as np

def detect_color_proportions(image_path, colors):
    image = cv2.imread(image_path)
    hsv_image = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

    color_proportions = {}
    total_pixels = image.shape[0] * image.shape[1]

    for color_name, color_range in colors.items():
        lower_color = np.array(color_range[0], dtype=np.uint8)
        upper_color = np.array(color_range[1], dtype=np.uint8)

        # 使用 cv2.inRange 函式創建了一個遮罩 (mask)
        # 以過濾出在顏色範圍內的像素。這可以將圖像中的目標顏色部分標記為白色，其他部分標記為黑色。
        mask = cv2.inRange(hsv_image, lower_color, upper_color)
        # 使用 cv2.countNonZero 函式計算出遮罩中白色像素的數量，也就是符合顏色範圍的像素數量。
        color_pixels = cv2.countNonZero(mask)
        # 將該顏色的比例 (符合顏色範圍的像素數量除以總像素數量) 存儲到 color_proportions 字典中。
        color_proportions[color_name] = color_pixels / total_pixels * 100

        # 將該顏色的像素標記為白色，其他像素標記為黑色
        result_image = np.zeros_like(image)
        result_image[mask > 0] = (255, 255, 255)  # 白色
        result_image[mask == 0] = (0, 0, 0)  # 黑色

        # 儲存結果圖片
        cv2.imwrite(f"{color_name}_result.jpg", result_image)

    return color_proportions

if __name__ == "__main__":
    image_path = "test2.jpg"
    colors = {
        # "red": [(0, 50, 20), (8, 255, 255)],
        "orange": [(0, 110, 100), (100, 255, 255)],
        "blue":  [(100, 100, 0), (255, 255, 255)],
        # "yellow": [(20, 100, 100), (30, 255, 255)]
    }

    result = detect_color_proportions(image_path, colors)
    print(result)