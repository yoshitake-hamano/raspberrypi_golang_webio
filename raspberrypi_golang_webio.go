// File: raspberrypi_ruby_webio.go - last edit
// yoshitake 15-Jul-2015

package main

import (
  "time"
  "io/ioutil"
  "net/http"
)

const GPIO_INDEX            = "4"
const GPIO_ROOT             = "/sys/class/gpio"
const GPIO_EXPORT           = GPIO_ROOT   + "/export"
const GPIO_TARGET           = GPIO_ROOT   + "/gpio" + GPIO_INDEX
const GPIO_TARGET_DIRECTION = GPIO_TARGET + "/direction"
const GPIO_TARGET_VALUE     = GPIO_TARGET + "/value"

func onHandler(w http.ResponseWriter, r *http.Request) {
  ioutil.WriteFile(GPIO_TARGET_VALUE, []byte("1"), 0644)
}

func offHandler(w http.ResponseWriter, r *http.Request) {
  ioutil.WriteFile(GPIO_TARGET_VALUE, []byte("0"), 0644)
}

func main() {
  ioutil.WriteFile(GPIO_EXPORT, []byte(GPIO_INDEX), 0644)
  time.Sleep(time.Second)
  ioutil.WriteFile(GPIO_TARGET_DIRECTION, []byte("out"), 0644)

  http.HandleFunc("/on",  onHandler)
  http.HandleFunc("/off", offHandler)
  http.ListenAndServe(":8000", nil)
}

// Log
// 15-Jul-2015 yoshitake Created.
