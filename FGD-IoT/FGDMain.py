import ConnectWIFI
import LoadJSON
import SensorController
import UploadData
import DeepSleep
import machine


def main():
    try:
        WIFISSID, WIFIPassWord, ReloadTime = LoadJSON.LoadJSON()
        LCD = machine.Pin(13, machine.Pin.OUT)
        LCD.on()
        ConnectWIFI.do_connect(WIFISSID, WIFIPassWord)
        SensorData = SensorController.SensorController()
        UploadData.UploadData(SensorData)
        LCD.off()
    except:
        machine.reset()
    else:
        DeepSleep.deepsleep(300)


