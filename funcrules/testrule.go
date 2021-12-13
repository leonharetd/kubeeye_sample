package funcrules

import (
	"fmt"
	"strconv"

	kube "github.com/leonharetd/kubeeye/pkg/funcrules"
)

var _ kube.FuncRule = FuncTestRule{}

type FuncTestRule struct{}

func (cer FuncTestRule) Exec() kube.ValidateResults {
	output := kube.ValidateResults{ValidateResults: make([]kube.ResultReceiver, 0)}
	var certExpiresOutput kube.ResultReceiver
	for i := range []int{1, 2, 3, 4} {
		certExpiresOutput.Name = fmt.Sprint("test", strconv.Itoa(i))
		certExpiresOutput.Type = "testExpire"
		certExpiresOutput.Message = []string{strconv.Itoa(i), "expire"}
		output.ValidateResults = append(output.ValidateResults, certExpiresOutput)
	}
	return output
}
 