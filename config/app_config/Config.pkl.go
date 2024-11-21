// Code generated from Pkl module `Config`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/yavurb/mobility-payments/config/app_config/environment"
	"github.com/yavurb/mobility-payments/config/app_config/loglevel"
)

// Configuration for the application.
type Config struct {
	Host string `pkl:"host"`

	Port uint16 `pkl:"port"`

	Cors *Cors `pkl:"cors"`

	HttpAuth *HttpAuth `pkl:"httpAuth"`

	Database *DatabaseConfig `pkl:"database"`

	LogLevel loglevel.LogLevel `pkl:"logLevel"`

	Environment environment.Environment `pkl:"environment"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Config
func LoadFromPath(ctx context.Context, path string) (ret *Config, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Config
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Config, error) {
	var ret Config
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
