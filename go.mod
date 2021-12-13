module github.com/leonharetd/kubeeye_sample

go 1.16

require (
	github.com/leonharetd/kubeeye v0.4.0
	github.com/spf13/cobra v1.2.1
)

replace k8s.io/klog/v2 v2.30.0 => k8s.io/klog/v2 v2.0.0

replace github.com/googleapis/gnostic v0.5.7 => github.com/googleapis/gnostic v0.3.1

replace sigs.k8s.io/controller-runtime v0.10.3 => sigs.k8s.io/controller-runtime v0.6.3

replace github.com/leonharetd/kubeeye v0.4.0 => /Users/leonharetd/WorkSpace/github/newkubeeye/kubeeye

replace github.com/kubesphere/kubeeye v0.4.0 => /Users/leonharetd/WorkSpace/github/newkubeeye/kubeeye
