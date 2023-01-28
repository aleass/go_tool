package main

func main() {
	res := Redis.Keys("user:token:*")
	data, err := res.Result()
	if err != nil {
		return
	}
	for _, v := range data {
		_res := Redis.Get(v)
		_res2 := Redis.TTL(v).Val()
		var ttl float64
		var unit = "s"
		if ttl = _res2.Seconds(); ttl > 3600 {
			ttl = _res2.Hours()
			unit = "h"
		}
		println(v, int64(ttl), unit)
		println(_res.Val(), "\r\n\r\n")
	}
}
