import machine


def deepsleep(time):
    rtc = machine.RTC()
    rtc.irq(trigger=rtc.ALARM0, wake=machine.DEEPSLEEP)

    if machine.reset_cause() == machine.DEEPSLEEP_RESET:
        print('woke from a deep sleep')

    rtc.alarm(rtc.ALARM0, time * 1000)

    machine.deepsleep()
