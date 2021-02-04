import network
import machine
import ntptime


def do_connect(essid, password):
    sta_if = network.WLAN(network.STA_IF)
    if not sta_if.isconnected():
        print('connecting to network...')
        sta_if.active(True)
        sta_if.connect(essid, password)
        while not sta_if.isconnected():
            pass
    print("WIFI connected.")
    print('network config:', sta_if.ifconfig())

    ntptime.settime()
    rtc = machine.RTC()
    print(rtc.datetime())
