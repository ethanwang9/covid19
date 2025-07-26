package global

// name: 请求返回数据结构
// author: Ethan.Wang
// desc:

// MsgBack 消息返回结构体
type MsgBack struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RedisLogin Redis Token存储设计
type RedisLogin struct {
	UUID     string `json:"uuid"`
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Level    int    `json:"level"`
	Location string `json:"location"`
	IsLogin  bool   `json:"is_login"`
	IsToken  bool   `json:"is_token"`
	IsStop   bool   `json:"is_stop"`
}
