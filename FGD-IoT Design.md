# 花草检测系统

新人练手, 主要在硬件方面和网络方面.

## 设计要素

`初次设计` 主要完成硬件方面的设计, 能够对花草进行探测并显示结果. 网络方面在考虑

## 硬件

1. MCU: `NodeMCU` 基于ESP8266 `Wifi`模组
   - 串口芯片: `CH340` 
   - 固件刷`microPython` : [microPython库](http://docs.micropython.org/en/latest/library/index.html)
2. 传感器
   - 光照强度传感器: `BH1750`
   - 温湿度传感器: `DHT11`
   - 土壤湿度传感器: 电阻式
   - 显示模块: `LCD1602_I2C`
3. 若干LED,220k电阻 \* 1, 110k电阻 \* 1, 跳线

## 网络

1. 物理层: `ESP8266` `WIFI`模组, 支持`802.11 b/g/n` , 支持`STA/AP`模式, 内置`TCP/IP`协议栈
2. 网络层: 正常走`IPv4`吧, 有心思再去捣鼓`IPv6`
3. 应用层: 考虑`MQTT`, 有现成服务可搭建. 得捣鼓一下协议内容.

## 电源需求

暂时不考虑

后期可能会上1000mAh左右的锂电池, 春节后再考虑

## 安全性

待定

## 应用程序

待定

数据上云后需要有应用的支持, 得找前端的同学帮帮忙.