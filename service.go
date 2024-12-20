package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type Config struct {
	REST     *RESTConfig         `yaml:"rest"`
	Services []*EndpointCategory `yaml:"services"`
}

type RESTConfig struct {
	Hostname string `yaml:"hostname"`
	Port     uint16 `yaml:"port"`
}

type EndpointCategory struct {
	Name      string      `yaml:"name"`
	Secret    string      `yaml:"secret"`
	Endpoints []*Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Key      string `yaml:"key"`
	Endpoint string `yaml:"endpoint"`
	Disabled bool   `yaml:"disabled"`
}

var cache map[string][]byte

func (c *Config) ReadConfig(ConfigFilepath string) error {
	log.Infof("Reading configuration from %s", ConfigFilepath)
	buffer, err := os.ReadFile(ConfigFilepath)
	if err != nil {
		return fmt.Errorf("failed to read config: %s", err.Error())
	}

	err = yaml.Unmarshal(buffer, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}

	return nil
}

func BuildCache(services []*EndpointCategory) error {
	cache = make(map[string][]byte)

	for _, service := range services {

		var pairs string = ""

		for _, endpoint := range service.Endpoints {
			if endpoint.Disabled {
				continue
			}
			pairs += fmt.Sprintf("{'%s': '%s'},", endpoint.Key, endpoint.Endpoint)
		}

		if len(pairs) > 0 {
			cache[service.Secret] = []byte(fmt.Sprintf("[%s]", pairs[:len(pairs)-1]))
		}
	}

	if len(cache) == 0 {
		return fmt.Errorf("empty cache")
	}

	return nil
}

func RunService(c *cli.Context) error {
	log.SetLevel(log.TraceLevel)

	config := new(Config)
	if err := config.ReadConfig(ConfigFilePath); err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}
	if config.REST == nil {
		log.Fatalf("Missing REST config")
	}

	BuildCache(config.Services)

	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", Handle)

	log.Infof("Starting webserver on %s:%d", config.REST.Hostname, config.REST.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", config.REST.Hostname, config.REST.Port), r)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Traceln("Handle()")
	req, ok := cache[r.Header.Get("X-Webgate-Request")]
	if !ok {
		log.Debugf("Request missing header")
		return
	}
	w.Write(req)
}
