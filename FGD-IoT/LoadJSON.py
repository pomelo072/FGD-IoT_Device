import ujson


def LoadJSON():
    try:
        config = open("config.json", "r", encoding="utf-8")
        loaded_config = ujson.load(config)
        WIFISSID = loaded_config['WIFISSID']
        WIFIPassWord = loaded_config['WIFIPassWord']
        ReloadTime = loaded_config['ReloadTime']
    finally:
        config.close()
    print("WIFI Config: {0} @ {1}    ReloadTime: {2}".format(WIFISSID, WIFIPassWord, ReloadTime))
    return WIFISSID, WIFIPassWord, ReloadTime
