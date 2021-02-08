import network
import machine
import SensorController


def do_connect(essid, password):
    sta_if = network.WLAN(network.STA_IF)
    if not sta_if.isconnected():
        SensorController.Display_LCD1602('connecting to\nnetwork...', 2)
        sta_if.active(True)
        sta_if.connect(essid, password)
        while not sta_if.isconnected():
            pass
    SensorController.Display_LCD1602("WIFI connected.", 1)
    SensorController.Display_LCD1602('network config\n' + sta_if.ifconfig()[0], 2)



