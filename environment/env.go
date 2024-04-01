package environment

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	AppName    string
	AppVersion string

	LogLevel string

	RestListenAddr string
	Domain         string
}

func Get() (Env, error) {
	var err error

	if err = load(); err != nil {
		return Env{}, err
	}

	var appName string
	if os.Getenv("APP_NAME") == "" {
		return Env{}, fmt.Errorf("app name is required")
	} else {
		appName = os.Getenv("APP_NAME")
	}

	var appVersion string
	if os.Getenv("APP_VERSION") == "" {
		return Env{}, fmt.Errorf("app version is required")
	} else {
		appVersion = os.Getenv("APP_VERSION")
	}

	var logLevel string
	if os.Getenv("LOG_LEVEL") == "" {
		logLevel = "INFO"
	} else {
		logLevel = os.Getenv("LOG_LEVEL")
	}

	var restListenAddr string
	if os.Getenv("REST_LISTEN_ADDR") == "" {
		restListenAddr = "localhost:8090"
	} else {
		restListenAddr = os.Getenv("REST_LISTEN_ADDR")
	}

	var domain string
	if os.Getenv("DOMAIN") == "" {
		return Env{}, fmt.Errorf("domain is required")
	} else {
		domain = os.Getenv("DOMAIN")
	}

	return Env{
		AppName:        appName,
		AppVersion:     appVersion,
		LogLevel:       logLevel,
		RestListenAddr: restListenAddr,
		Domain:         domain,
	}, nil
}

var (
	envLoaded = false
	mu        sync.Mutex
)

func load() error {
	mu.Lock()
	defer mu.Unlock()
	if envLoaded {
		return nil
	}

	//定位到根目录的.env文件
	_, f, _, _ := runtime.Caller(0) // 当前执行的文件，即此文件environment/env.go
	basepath := filepath.Dir(f)
	envFile := path.Join(basepath, "../.env")
	_ = godotenv.Load(envFile)

	//从os读取环境变量，忽略error
	_ = godotenv.Load()
	envLoaded = true
	return nil
}
