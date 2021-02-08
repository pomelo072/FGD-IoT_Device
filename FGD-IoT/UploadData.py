import ujson
import urequests
import SensorController


def UploadData(Data):
    SensorController.Display_LCD1602("Data Uploading..", 3)
    url = "https://iot.pomelo072.top/FGD/upload?device_id=" + str(Data[0]) + "&temperature=" + str(Data[1]) + "&humidity=" + str(Data[2]) + "&light=" + str(Data[3]) + "&solid=" + str(Data[4])
    response = urequests.get(url)
    print(response.json())
    SensorController.Display_LCD1602("Data Uploading..\nSuccess.", 1)
