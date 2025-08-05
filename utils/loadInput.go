package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadInput() (Config, error) {
	var config Config

	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("erro ao ler entrada do GoIt: %w", err)
	}

	return config, nil
}
