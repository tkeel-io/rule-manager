package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	mathRand "math/rand"
	"net"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type prefix string

func DeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

const (
	RuleID   prefix = "rule"
	ActionID prefix = "ac"
)

const _IDTemplate = "iot-%s-%s"

// MyStringList 将[]string定义为MyStringList类型
type MyStringList []string

// Len 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// Swap 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// SortStringSlice sort slice.
func SortStringSlice(s []string) []string {
	ss := MyStringList(s)
	sort.Sort(ss)
	return []string(ss)
}

func GenerateUrlChronusDB(hosts []string, database string) []string {

	endpoints := make([]string, 0)
	for _, host := range hosts {

		// u := url.URL{
		// 	Scheme: "http",
		// 	//Opaque      string    // encoded opaque data
		// 	//User: url.UserPassword(user, passwd),
		// 	Host: host,
		// 	//Path        string    // path (relative paths may omit leading slash)
		// 	//RawPath     string    // encoded path hint (see EscapedPath method); added in Go 1.5
		// 	//ForceQuery  bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
		// 	//RawQuery    string    // encoded query values, without '?'
		// 	//Fragment    string    // fragment for references, without '#'
		// 	//RawFragment string    // encoded fragment hint (see EscapedFragment method); added in Go 1.15
		// }
		// q := u.Query()
		// q.Set("database", database)
		// u.RawQuery = q.Encode()
		// endpoints = append(endpoints, u.String())

		endpoints = append(endpoints, fmt.Sprintf("http://%s?database=%s", host, database))
	}
	return endpoints
}

func GenerateUrlsChronusDB(hosts []string, user, password, database string) []string {

	endpoints := make([]string, 0)
	for _, host := range hosts {

		u := url.URL{
			Scheme: "http",
			//Opaque      string    // encoded opaque data
			//User: url.UserPassword(user, passwd),
			Host: host,
			User: url.UserPassword(user, password),
			Path: url.PathEscape(database), //        string    // path (relative paths may omit leading slash)
			//RawPath     string    // encoded path hint (see EscapedPath method); added in Go 1.5
			//ForceQuery  bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
			//RawQuery    string    // encoded query values, without '?'
			//Fragment    string    // fragment for references, without '#'
			//RawFragment string    // encoded fragment hint (see EscapedFragment method); added in Go 1.15
		}
		//q := u.Query()
		//q.Set("database", database)
		//u.RawQuery = q.Encode()
		endpoints = append(endpoints, u.String())
	}
	return endpoints
}

func GenerateUrlKafka(host, user, passwd, topic string) string {

	return fmt.Sprintf("kafka://%s/%s/qingcloud", host, url.PathEscape(topic))
}

func GenerateUrlMysql(endpoints []string, user, passwd, db string) []string {
	urls := []string{}
	for _, endpoint := range endpoints {
		urls = append(urls, fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, passwd, endpoint, db, "charset=utf8&parseTime=True&loc=Asia%2FShanghai"))
	}
	return urls
}

func GenerateUrlPostgresql(endpoints []string, user, passwd, db string) []string {
	urls := []string{}
	for _, endpoint := range endpoints {
		str := strings.Split(endpoint, ":")
		if len(str) != 2 {
			continue
		}
		urls = append(urls, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", str[0], str[1], user, passwd, db))
	}
	return urls
}

func GenerateUrlRedis(endpoints []string, user, passwd, db string) []string {
	urls := []string{}
	for _, endpoint := range endpoints {
		url := url.URL{
			Scheme: "redis",
			Host:   endpoint,
			User:   url.UserPassword(user, passwd),
			Path:   url.PathEscape(db),
		}
		urls = append(urls, url.String())
	}
	return urls
}

func CheckHost(hosts []string) bool {
	for _, host := range hosts {
		p := strings.Split(host, ":")
		if len(p) != 2 {
			return false
		}
		//check ip
		if nil == net.ParseIP(p[0]) {
			return false
		}
		//check port
		if port, err := strconv.ParseInt(p[1], 10, 63); nil != err {
			return false
		} else if port >= 65535 {
			return false
		}
	}
	return true
}

func MapCat(m1, m2 map[string]interface{}) map[string]interface{} {

	if nil == m1 {
		m1 = make(map[string]interface{})
	}
	for key, value := range m2 {
		m1[key] = value
	}
	return m1
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandString(len int) string {
	b := make([]byte, len, len)
	for i := 0; i < len; i++ {
		source := mathRand.NewSource(time.Now().UnixNano())
		index := mathRand.New(source).Intn(len)
		b[i] = chars[index]
	}
	b[0] = RandomWord()
	return string(b)
}

func RandomWord() byte {
	words := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	source := mathRand.NewSource(time.Now().UnixNano())
	index := mathRand.New(source).Intn(len(words))
	return words[index]
}
