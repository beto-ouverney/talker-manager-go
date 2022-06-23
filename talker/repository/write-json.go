package talkerrepository

import (
	"os"
)

func writeJSON(jsonFile []byte) (err error) {
	f, err := os.Create("./talkers.json")
	if err == nil {
		_, err = f.Write(jsonFile)
	}
	defer f.Close()
	return
}
