package main

import (
	"fmt"
	"strings"
	"github.com/garyburd/redigo/redis"
)

type Question struct {
	key string
	question string
	default_value string
	is_password bool
}


func ReadConfig(namespace string) (map[string]string){
	c, err := redis.Dial("tcp", ":6379")
	config := make(map[string]string)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	keys, err := redis.Strings(c.Do("KEYS", fmt.Sprintf("%s.*", namespace)))
	
	if err != nil {
		fmt.Println("no keys found")
	}
	for _, key := range keys {
		value, err := redis.String(c.Do("GET", key))
		if err != nil {
			fmt.Println("key error", key)
		}
		config[strings.Replace(key, fmt.Sprintf("%s.", namespace), "", -1)] = value
	}
	return config
}


func SetConfig(namespace string, questions []Question) {
	// TODO transfer redis connection to another method
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	for _,question := range questions {
		var i string
	    fmt.Println(question.question)
	    fmt.Scan(&i)
	    c.Do("SET", fmt.Sprintf("%s.%s", namespace, question.key), i)
	}
}


func main() {
	// var questions []Question
	// q := Question{key:"mongodb_host", question:"Enter the value for mongodb_host", default_value:"127.0.0.1"}
	// questions = append(questions, q)
	// SetConfig("teste", questions)
}