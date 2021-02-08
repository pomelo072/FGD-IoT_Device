# FGD-IoT_Device

新手的尝试, 有新想法或者更好的改进请提交Issue或PR.

### 文档位置

[设备端](https://github.com/pomelo072/FGD-IoT_Device/tree/main/FGD-IoT)

[服务端](https://github.com/pomelo072/FGD-IoT_Device/tree/main/IoT-server)

[初期设计想法](https://github.com/pomelo072/FGD-IoT_Device/blob/main/FGD-IoT%20Design.md)

[PCB](https://github.com/pomelo072/FGD-IoT_Device/blob/main/Schematic_FGD-IoT.pdf)

## FGD-IoT

需要在项目文件中加上一个配置文件`config.json`:

```json
{
  "WIFISSID": "ABC",
  "WIFIPassWord": "12345678",
  "ReloadTime": 300
}
```

分别对应:

- WIFI的SSID
- WIFI的密码
- 深度睡眠唤醒时间

## IoT-server

菜鸡简单的写了个``HTTP API`, Go写的, Web框架`iris`

### 配置文件

在项目根目录也需要一个配置文件`config.json`

```json
{
  "Port": ":8086",
  "DBUserName": "ABC",
  "DBPassword": "12345678",
  "DBIp": "127.0.0.1",
  "DBPort": "3306",
  "DBName": "abc"
}
```

分别对应配置:

- 程序通信端口
- 数据库用户名
- 数据库用户密码
- 数据库IP
- 数据库端口
- 数据库名

### API接口

接口不多就不单独写文档了

#### 数据上传

`/FGD/upload`

- Method: `GET`

- `param`:

  ```
  ?device_id=1  // 设备ID
  &temperature=20  // 温度值
  &humidity=50  // 湿度值
  &light=500  // 光强值
  &solid=200  // 土壤湿度值
  ```

- 返回值

  ```json
  {
  	"Code": 200,
      "Msg": "SUCCESS",
      "Data": "Success."
  }
  ```

#### 读取最后一次测量值

`/FGD/details`

- Method: `GET`

- `param`:

  ```
  ?DeviceID=1 // 设备ID
  ```

- 返回值

  ```json
  {
      "Code":200,
      "Msg":"SUCCESS",
      "Data":{
          "ID":19,
          "DEVICEID":"1",
          "RECORDTIME":"2021-02-08 21:36:09",
          "TEMPERATURE":"25",
          "HUMIDITY":"61",
          "LIGHTINTENSITY":"142",
          "SOLIDHUMIDITY":"132"
      }
  }
  ```

#### 读取最近一天的测量值

`/FGD/list`

- Method: `GET`

- `param`:

  ```
  ?DeviceID=1 // 设备ID
  ```

- 返回值

  ```json
  {
      "Code":200,
      "Msg":"SUCCESS",
      "Data":[
          {
              "device_id":"1",
          	"hum_id_ity":"55",
              "id":1,
              "lightintensity":"19",
              "recordtime":"2021-02-08 17:26:38",
              "sol_id_hum_id_ity":"275",
              "temperature":"25"
          },
          {
              "device_id":"1",
              "hum_id_ity":"61",
              "id":2,
              "lightintensity":"140",
              "recordtime":"2021-02-08 21:09:01",
              "sol_id_hum_id_ity":"138",
              "temperature":"25"
          }
      ]
  }
  ```

  

## End

时间匆忙写得拉跨, 新手操作也实属是菜.

乐意听取各位指教