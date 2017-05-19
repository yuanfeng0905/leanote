package lea

import (
    "encoding/json"
    "fmt"
    "github.com/revel/revel"
)

func Log(i ...interface{}) {
    revel.INFO.Println(i...)
}

func LogW(i ...interface{}) {
    revel.WARN.Println(i...)
}

func Log1(key, i interface{}) {
    revel.INFO.Println(key, i)
}

func LogJ(i interface{}) {
    b, _ := json.MarshalIndent(i, "", " ")
    revel.INFO.Println(string(b))
}

// 为test用
func L(i interface{}) {
    fmt.Println(i)
}

func LJ(i interface{}) {
    b, _ := json.MarshalIndent(i, "", " ")
    fmt.Println(string(b))
}
