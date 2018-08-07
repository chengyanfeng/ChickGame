package util

import (
	"github.com/astaxie/beego/cache"
	"fmt"
	"crypto/md5"
	"time"
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"strconv"
	"github.com/astaxie/beego"
	"encoding/json"
)
var localCache cache.Cache
type P map[string]interface{}
func InitCache() {
	c, err := cache.NewCache("memory", `{"interval":60}`)
	//c, err := cache.NewCache("file", `{"CachePath":"./dhcache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err != nil {
		Error(err)
	} else {
		localCache = c
	}
}

// 缓存接口，存 S("key", value)，取 S("key")
func S(key string, p ...interface{}) (v interface{}) {
	md5 := Md5(key)
	if len(p) == 0 {
		return localCache.Get(md5)
	} else {
		if len(p) == 2 {
			var ttl int64
			switch p[1].(type) {
			case int:
				ttl = int64(p[1].(int))
			case int64:
				ttl = p[1].(int64)
			}
			localCache.Put(md5, p[0], time.Duration(ttl)*time.Second)
		} else if len(p) == 1 {
			localCache.Put(md5, p[0], 24*time.Hour)
		}
		return p[0]
	}
}
//删除缓存中的缓存grade
func Del(key string)(log string) {

	md5 := Md5(key)
	err:=localCache.Delete(md5)
	fmt.Println("err:")
	fmt.Println(err)
	if err!=nil{
		return "nil"
	}
	return "ok"
}
func Md5(s ...interface{}) (r string) {
	return Hash("md5", s...)
}
func Hash(algorithm string, s ...interface{}) (r string) {
	var h hash.Hash
	switch algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha2", "sha256":
		h = sha256.New()
	}
	for _, value := range s {
		switch value.(type) {
		case []byte:
			h.Write(value.([]byte))
		default:
			h.Write([]byte(ToString(value)))
		}
	}
	r = hex.EncodeToString(h.Sum(nil))
	return
}
func ToString(v interface{}) string {
	if v != nil {
		switch v.(type) {
		case bson.ObjectId:
			return v.(bson.ObjectId).Hex()
		case []byte:
			return string(v.([]byte))
		case *P, P:
			var p P
			switch v.(type) {
			case *P:
				if v.(*P) != nil {
					p = *v.(*P)
				}
			case P:
				p = v.(P)
			}
			var keys []string
			for k := range p {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			r := "P{"
			for _, k := range keys {
				r = JoinStr(r, k, ":", p[k], " ")
			}
			r = JoinStr(r, "}")
			return r
		case int64:
			return strconv.FormatInt(v.(int64), 10)
		default:
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}
func JoinStr(val ...interface{}) (r string) {
	for _, v := range val {
		r += ToString(v)
	}
	return
}
// 记录err信息
func Error(v ...interface{}) {
	beego.Error(v)
}
//string 转P
func JsonDecode(b []byte) (p *map[string]interface{}) {
	p = &map[string]interface{}{}
	err := json.Unmarshal(b, p)
	if err != nil {
		Error("JsonDecode", string(b), err)
	}
	return
}