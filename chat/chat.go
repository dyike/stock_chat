package chat

import (
    "fmt"
    "time"
    "github.com/dyike/httpClient"
    "encoding/json"
)

type ChatClient struct {
    workId   string
    secret   string
    agentid  int64
}

type ChatToken struct {
    token   string
    expire  int64
}

var Token = &ChatToken{
    token: "",
    expire: 0,
}

func New(workId string, secret string, agentid int64) *ChatClient {
    return &ChatClient{
        workId: workId,
        secret: secret,
        agentid: agentid,
    }
}

func store(token ChatToken) {
    Token = &token
}

func (cc *ChatClient) getToken() (token ChatToken, err error) {
    tokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", cc.workId, cc.secret)
    timeStamp := time.Now().Unix()
    result, err := httpClient.DoRequest(httpClient.Request{
        Method: "GET",
        URL:    tokenUrl,
    })
    var res = struct {
        ErrCode     int    `json:"errcode"`
        ErrMsg      string `json:"errmsg"`
        AccessToken string `json:"access_token"`
        ExpiresIn   int64  `json:"expires_in"`
    }{}
    if err == nil {
        err = json.Unmarshal(result, &res)
        if err == nil {
            token.token = res.AccessToken
            token.expire = timeStamp + res.ExpiresIn - 120
            store(token)
            return token, nil
        }
    }
    return token, err
}



func (cc *ChatClient) SendMessage(content map[string]interface{}) {
    var accessToken string
    if Token.expire < time.Now().Unix() {
        token, err := cc.getToken()
        if err != nil {
            return
        }
        accessToken = token.token
    } else {
        accessToken = Token.token
    }
    content["agentid"] = cc.agentid
    sendMsgUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken)
    body, _ := json.Marshal(content)
    result, err := httpClient.DoRequest(httpClient.Request{
        Method: "POST",
        URL:    sendMsgUrl,
        Body:   body,
    })
    var res = map[string]interface{}{}
    err = json.Unmarshal(result, &res)
    if err == nil {
        return
    }
}
