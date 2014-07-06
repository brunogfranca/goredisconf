package goredisconf

import (
	"fmt"
	"strings"
	"errors"
	"github.com/garyburd/redigo/redis"
)

type Question struct {
	Key string
	Question string
	Default_value string
	Is_password bool
}


func ReadConfig(namespace string) (map[string]string, error){
	c, err := redis.Dial("tcp", ":6379")
	config := make(map[string]string)
	if err != nil {
		return config, errors.New("redis connection error")
	}
	defer c.Close()
	keys, err := redis.Strings(c.Do("KEYS", fmt.Sprintf("%s.*", namespace)))
	
	if err != nil {
		return config, errors.New("no keys found")
	}
	for _, key := range keys {
		value, err := redis.String(c.Do("GET", key))
		if err != nil {
			return config, errors.New(fmt.Sprintf("key error %s", key))
		}
		config[strings.Replace(key, fmt.Sprintf("%s.", namespace), "", -1)] = value
	}
	return config, nil
}


func SetConfig(namespace string, questions []Question) {
	// TODO transfer redis connection to another method
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	for _,question := range questions {
		var i string
	    fmt.Println(question.Question)
	    fmt.Scan(&i)
	    c.Do("SET", fmt.Sprintf("%s.%s", namespace, question.Key), i)
	}
}


func main() {
	// var questions []Question
	// q := Question{key:"mongodb_host", question:"Enter the value for mongodb_host", default_value:"127.0.0.1"}
	// questions = append(questions, q)
	// SetConfig("teste", questions)
}