// # yay: http://squarism.com/2014/10/13/yaml-go/
// # yay2: https://stackoverflow.com/questions/28682439/go-parse-yaml-file
package main

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
  // "reflect"
)

type conf struct {
  NAME string `yaml:"name"`
  LOCATION string `yaml:"location"`
  TIMEZONE string `yaml:"timezone"`
  BAM struct{
    SKILLS struct {
      SCIENCE []map[string][]string `yaml:"Science"`
      AWS []map[string][]string `yaml:"AWS"`
      TEACHING []map[string]string `yaml:"Teaching"`
    } //`yaml:"Science"`
    LANGUAGES []string `yaml:"languages"`
    FRAMEWORKS []string `yaml:"frameworks/libraries/concepts/api"`
  }
  SAM map[string][]string `yaml:"sam"`
  // BAM map[string][]string
  // SAM string `yaml:"sam"`

}

func (c *conf) getConf() *conf {
  yamlFile, err := ioutil.ReadFile("dblake.yml")
  // yamlFile, err := ioutil.ReadFile("tmp.yml")
  if err != nil {
      log.Printf("yamlFile.Get err   #%v ", err)
  }
  err = yaml.UnmarshalStrict(yamlFile, c)
  if err != nil {
      log.Fatalf("Unmarshal: %v", err)
  }

  return c
}

func main() {
  var c conf
  c.getConf()
  // fmt.Println(reflect.TypeOf(c))
  fmt.Println(c.NAME)
  fmt.Println(c.LOCATION)
  fmt.Println(c.TIMEZONE)
  fmt.Println(c.SAM)
  fmt.Println(c.BAM)
  fmt.Println(c)
  fmt.Println()
  fmt.Println(c.BAM.SKILLS.AWS)
}
