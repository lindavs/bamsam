package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type conf struct {
    BAM string `yaml:"bam"`
    SAM string `yaml:"sam"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("dblake.yml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf
    c.getConf()

    fmt.Println(c)
}
