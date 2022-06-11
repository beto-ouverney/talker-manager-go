package talkerrepository

import (
	"os"
)

func readJSON() ([]byte, error) {
	jsonFile, err := os.ReadFile("./talkers.json")
	return jsonFile, err
}
