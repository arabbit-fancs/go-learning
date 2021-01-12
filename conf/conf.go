package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var y = getYaml()
var Flavor = "production"

func Env(key string) string {
	if _, ok := y["production"].(map[string]interface{})[key]; ok {
		return y[Flavor].(map[string]interface{})[key].(string)
	}
	panic(fmt.Sprintf("%s: [%s] is notfund", Flavor, key))
	return ""
}

func getYaml() map[interface{}]interface{} {
	var buf, err = ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	return m
}
