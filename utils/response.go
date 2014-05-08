package utils

type Response struct {
	Ok   bool
	Key  string
	Data interface{}
}

//返回JSON格式消息
func JsonMessage(ok bool, k, d string) *Response {
	var r = new(Response)
	r.Ok = ok
	r.Key = k
	r.Data = d
	return r
}

//返回JSON格式对象
func JsonResult(ok bool, k string, d interface{}) *Response {
	var r = new(Response)
	r.Ok = ok
	r.Key = k
	r.Data = d
	return r
}
