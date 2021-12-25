package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mohammadiahmad/shop/internal/app"
	"github.com/mohammadiahmad/shop/internal/cart_storage"
	"github.com/mohammadiahmad/shop/internal/readisearch"
	"io/ioutil"
)

type Config struct {
	Storage readisearch.Config
	Server app.Config
	Redis  cart_storage.Config
}

const (
	Path = "./configs/"
)

func Load() *Config {
	config := Config{}
	k := koanf.New(".")
	parser := yaml.Parser()
	files, err := ioutil.ReadDir(Path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		filePath := Path + f.Name()

		if err := k.Load(file.Provider(filePath), parser); err != nil {
			fmt.Printf("Error in load config file: %s\n", err)
		}

	}

	err = k.Unmarshal("", &config)
	if err != nil {
		fmt.Printf("Error in unmarshal yml files: %s\n", err)
		return nil
	}

	fmt.Printf("configuration: %s\n\n", spew.Sdump(config))

	return &config

}
