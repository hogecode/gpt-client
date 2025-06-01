package application

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

const configFile = "config.json"

type Configuration struct {
    ApiKey    string `json:"api_key"`
    OutputDir string `json:"output_dir"`
}

func LoadConfig() (*Configuration, error) {
    file, err := ioutil.ReadFile(configFile)
    if err != nil {
        return nil, fmt.Errorf("unable to read config file: %v", err)
    }

    var config Configuration
    err = json.Unmarshal(file, &config)
    if err != nil {
        return nil, fmt.Errorf("unable to unmarshal config: %v", err)
    }
    return &config, nil
}

func SaveApiKey(apiKey string) error {
    config, err := LoadConfig()
    if err != nil {
        config = &Configuration{}
    }

    config.ApiKey = apiKey
    return saveConfig(config)
}

func SaveOutputDir(outputDir string) error {
    config, err := LoadConfig()
    if err != nil {
        config = &Configuration{}
    }

    config.OutputDir = outputDir
    return saveConfig(config)
}

func saveConfig(config *Configuration) error {
    file, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return fmt.Errorf("unable to marshal config: %v", err)
    }

    err = ioutil.WriteFile(configFile, file, 0644)
    if err != nil {
        return fmt.Errorf("unable to write config file: %v", err)
    }

    return nil
}
