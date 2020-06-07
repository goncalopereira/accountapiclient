package test

import (
	"io/ioutil"
	"os"
)

const integrationDataPath = "../../test/integration/data/"

//ReadJSON returns byte array based on relative file path.
func ReadJSON(path string) []byte {
	jsonFile, err := os.Open(integrationDataPath + path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("could not read test data")
	}

	return bytes
}
