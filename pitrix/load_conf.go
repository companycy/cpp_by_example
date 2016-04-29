package main

import (
	"fmt"
	ini "github.com/vaughan0/go-ini"
	"io/ioutil"
	"strings"
)

const (
	HYPER             = "hyper"
	SNAPSHOT          = "snapshot"
	SEED              = "seed"
	FEATURE_HYPERNODE = "feature_hypernode"
	ROLE              = "role"
)

var (
	global_conf = map[string]map[string]string{}
)

func CollectConf(f string, section map[string]string) {
	global_conf[f] = section
}

func GetNodesByType(nodeType string) []string {
	var nodes []string
	for f, settings := range global_conf {
		// fmt.Printf("f: %s, %s\n", f, settings[FEATURE_HYPERNODE])
		switch nodeType {
		case HYPER:
			if strings.Contains(settings[FEATURE_HYPERNODE], "on") {
				nodes = append(nodes, f)
			}
		default:
			if strings.Contains(settings[ROLE], nodeType) {
				nodes = append(nodes, f)
			}
		}
	}
	return nodes
}

func main() {
	absPath := "/home/bjcheny/Documents/pitrix/pitrix-tools/first_box/pitrix/conf/settings_pek3/conf/settings/"
	files, _ := ioutil.ReadDir(absPath)
	for _, f := range files {
		if !f.Mode().IsRegular() { // backup dir
			continue
		}

		// type File map[string]Section
		file, err := ini.LoadFile(absPath + f.Name())
		if err != nil {
			fmt.Errorf("error: %v", err)
			continue
		}

		// fmt.Printf("setting: %s\n", f.Name())

		for _, section := range file {
			// fmt.Printf("section name: %s\n", name)
			// for key, value := range section { // section content
			// 	fmt.Printf("\t%s => %s\n", key, value)
			// }
			// type Section map[string]string
			CollectConf(f.Name(), section)
		}

	}

	// for k, v := range global_conf {
	// 	fmt.Printf("%s => %s\n", k, v)
	// }

	types := []string{
		HYPER, SNAPSHOT, SEED,
	}
	for i := 0; i < len(types); i++ {
		nodes := GetNodesByType(types[i])
		fmt.Println(nodes)
	}

}
