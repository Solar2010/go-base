package main 

import "fmt"

func main(){
  var countryCaptialMap map[string]string
  countryCaptialMap = make(map[string]string)

  countryCaptialMap ["France"] = "巴黎"
  countryCaptialMap ["Italy"] = "罗马"
  countryCaptialMap ["Japan"] = "东京"
  countryCaptialMap ["India"] = "新德里"

  for country := range countryCaptialMap {
    fmt.Println(country, "首都是", countryCaptialMap [country])
  }
  

  captial, ok := countryCaptialMap ["美国"]
  
  if (ok) {
    fmt.Println("美国的首都是",captial)
  } else {
    fmt.Println("美国的首都不存在")
  }
}
