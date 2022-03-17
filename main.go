package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/cmxjs/fund/src/bark"
	"github.com/cmxjs/fund/src/fund"
)

type Config struct {
	Host     string   `json:"host"`
	Key      string   `json:"key"`
	Fundcode []string `json:"fundcode"`
}

var (
	Config_file string = "./config.json"
	Wg                 = sync.WaitGroup{}
	Jz_info            = sync.Map{}
)

var init_config = []byte(`{
  "key":"",
  "host":"api.day.app",
  "fundcode":[
    "000001"
  ]
}`)

func create_config() (err error) {
	f, err := os.Create(Config_file)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(init_config); err != nil {
		return err
	}
	return nil
}

func init() {
	if _, err := os.Stat(Config_file); err != nil {
		if err := create_config(); err != nil {
			log.Fatal(err)
		}
		log.Printf("create default config file success, path: '%s'\n", Config_file)
		os.Exit(0)
	}
}

func parse_json(file string, c *Config) (err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return json.Unmarshal(data, c)
}

func fetchFundInfo() {
	c := Config{}
	if err := parse_json(Config_file, &c); err != nil {
		log.Fatal(err)
	}
	for _, v := range c.Fundcode {
		Wg.Add(1)
		go func(_v string) {
			defer Wg.Done()

			jz_data, err := fund.GetFundJz(_v)
			if err != nil {
				log.Println("fund:", _v, "err:", err)
			} else {
				Jz_info.Store(_v, (*jz_data))
			}
		}(v)
	}
	Wg.Wait()

	//conversion to type fund.SortBygszzl
	var data fund.SortBygszzl
	Jz_info.Range(func(key, value interface{}) bool {
		data = append(data, value.(map[string]interface{}))
		return true
	})
	sort.Sort(data)

	body := ""
	for _, v := range data {
		body += fmt.Sprintln(v["name"], v["gszzl"])
	}
	bark.Send(c.Host, c.Key, "Fund", body, "fund")
}

func main() {
	fetchFundInfo()
}
