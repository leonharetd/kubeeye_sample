package main

import (
	// "os"
	// "fmt"
	// "strings"
	// "path/filepath"
	// "io/ioutil"
	"github.com/kubesphere/kubeeye/cmd"
	kc "github.com/leonharetd/kubeeye_sample/cmd"
	krr "github.com/leonharetd/kubeeye_sample/regorules"
	kfr "github.com/leonharetd/kubeeye_sample/funcrules"
)

func main() {
	cmd := cmd.NewKubeEyeCommand().WithFuncRule(kfr.FuncTestRule{}).WithRegoRule(krr.CustomEmbRegoRules).WithCommand(kc.TestCmd).DO()
	cmd.Execute()
	// pathabs, err := filepath.Abs("./kubeeye_sample/regorules")
	// if err != nil {
	// 	panic(err)
	// }
	// if strings.HasSuffix(pathabs, "/") == false {
	// 	pathabs += "/"
	// }
	// files, err := ioutil.ReadDir(pathabs)
	// if err != nil {
	// 	fmt.Println("Failed to read the dir of rego rule files.")
	// 	os.Exit(1)
	// }
	// for _, f := range files {
	// 	fmt.Println(filepath.Abs(f.Name()))
	// }
}
