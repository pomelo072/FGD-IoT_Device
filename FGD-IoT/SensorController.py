import machine
import bh1750fvi
from time import sleep_ms, ticks_ms
import dht
from i2c_lcd1602 import I2cLcd


def Sensor_DHT11():
    sensor = dht.DHT11(machine.Pin(2))
    sensor.measure()
    t = sensor.temperature()
    h = sensor.humidity()
    return t, h


def Sensor_GY3o2():
    i2csensor = machine.I2C(scl=machine.Pin(5), sda=machine.Pin(4))
    result = bh1750fvi.sample(i2csensor)
    return result


def Sensor_YL69():
    YL69 = machine.ADC(0)
    solid_humidity = YL69.read()
    return solid_humidity


def Display_LCD1602(word, t):
    i2c = machine.I2C(scl=machine.Pin(12), sda=machine.Pin(14), freq=400000)
    lcd = I2cLcd(i2c, 0x27, 2, 16)
    lcd.putstr(word)
    sleep_ms(t * 1000)
    lcd.clear()


def SensorController():
    t, h = Sensor_DHT11()
    lx = Sensor_GY3o2()
    solid_humidity = Sensor_YL69()
    SensorData = [1, t, h,  lx, solid_humidity]
    first_msg = "temperature:" + str(t) + "\nhumidity:" + str(h) + "%"
    second_msg = "Light intensity\n" + str(lx) + " lx"
    third_msg = "solid_humidity\n" + str(solid_humidity)
    Display_LCD1602(first_msg, 2)
    Display_LCD1602(second_msg, 2)
    Display_LCD1602(third_msg, 2)
    return SensorData
