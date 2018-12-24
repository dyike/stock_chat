package config

import (
    "os"
    "encoding/json"
)

type Config struct {
    Wechat struct {
        WorkId string  `json:"workid"`
        Secret string `json:"secret"`
        AgentId int64 `json:"agentid"`
    } `json:"wechat"`
    Stock struct {
        Time struct {
            US struct {
                Start string `json:"start"`
                End   string `json:"end"`
            } `json:"us"`
            HK struct {
                Start string `json:"start"`
                End   string `json:"end"`
            }  `json:"hk"`
        } `json:"time"`
        List []struct {
            Name string `json:"name"`
            ExpectIn float64 `json:"expect_in"`
            ExpectOut float64 `json:"expect_out"`
        } `json:"list"`
    } `json:"stock"`
}


func LoadConfig(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        return config, err
    }
    jsonParse := json.NewDecoder(configFile)
    err = jsonParse.Decode(&config)
    return config, err
}