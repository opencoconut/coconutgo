package coconut

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Config struct {
	ApiVersion string
	Source     string
	Webhook    string
	Conf       string
	Outputs    Outputs
	Vars       Vars
}

type Outputs map[string]string
type Vars map[string]string

func (c Config) String() (string, error) {
	confFile := c.Conf
	conf := []string{}

	if confFile != "" {
		confContent, err := ioutil.ReadFile(confFile)
		if err != nil {
			return "", err
		}

		conf = strings.Split(strings.TrimSpace(string(confContent)), "\n")
	}

	if vars := c.Vars; vars != nil {
		for v, val := range vars {
			conf = append(conf, fmt.Sprintf("var %s = %s", v, val))
		}
	}

	if apiVersion := c.ApiVersion; apiVersion != "" {
		conf = append(conf, fmt.Sprintf("set api_version = %s", apiVersion))
	}

	if source := c.Source; source != "" {
		conf = append(conf, fmt.Sprintf("set source = %s", source))
	}

	if webhook := c.Webhook; webhook != "" {
		conf = append(conf, fmt.Sprintf("set webhook = %s", webhook))
	}

	if outputs := c.Outputs; outputs != nil {
		for o, url := range outputs {
			conf = append(conf, fmt.Sprintf("-> %s = %s", o, url))
		}
	}

	newConf := []string{}

	confVars := filter(conf, func(v string) bool {
		return strings.HasPrefix(v, "var")
	})
	sort.Strings(confVars)
	newConf = append(newConf, strings.Join(confVars, "\n"), "")

	confSettings := filter(conf, func(v string) bool {
		return strings.HasPrefix(v, "set")
	})
	sort.Strings(confSettings)
	newConf = append(newConf, strings.Join(confSettings, "\n"), "")

	confOutputs := filter(conf, func(v string) bool {
		return strings.HasPrefix(v, "->")
	})
	sort.Strings(confOutputs)
	newConf = append(newConf, strings.Join(confOutputs, "\n"))

	return strings.TrimSpace(strings.Join(newConf, "\n")), nil
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
