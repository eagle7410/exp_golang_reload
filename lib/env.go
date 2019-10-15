package lib

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"reflect"
	"strings"
)

type env struct {
	WorkDir,
	TimeZone,
	Place string
	IsDev, IsCorsAllowCredentials bool
	AllowedMethods                []string
}

func (i *env) GetLinkIsDev() *bool {
	return &i.IsDev
}
func (i *env) GetLinkAllowedMethods() *[]string {
	return &i.AllowedMethods
}
func (i *env) GetLinkIsCorsAllowCredentials() *bool {
	return &i.IsCorsAllowCredentials
}

func (i *env) Init() error {
	i.AllowedMethods = []string{"GET", "POST", "OPTIONS"}
	pwd, err := os.Getwd()

	fmt.Println(pwd)

	if err != nil {
		return err
	}

	i.WorkDir = pwd

	envPath := path.Join(pwd, ".env")

	if _, err := os.Stat(envPath); err == nil {
		fmt.Println("Env load from file")
		err := godotenv.Load(envPath)

		if err != nil {
			return err
		}
	}

	props := map[string]bool{
		"TimeZone": false,
		"Place":    true,
	}

	for prop, isRequired := range props {

		v := os.Getenv(prop)

		if isRequired == true && v == "" {
			return errors.New("Bad " + prop)
		}

		reflect.ValueOf(i).Elem().FieldByName(prop).SetString(v)
	}

	if strings.ToLower(os.Getenv("isDev")) == "true" {
		i.IsDev = true
	}

	if strings.ToLower(os.Getenv("IsCorsAllowCredentials")) == "true" {
		i.IsCorsAllowCredentials = true
	}

	if i.TimeZone == "" {
		i.TimeZone = "Europe/London"
	}

	os.Setenv("TZ", i.TimeZone)

	return nil
}

var ENV env
