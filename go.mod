module learning

go 1.18

//github.com/armstrongli/go-bmi v0.0.1
require github.com/spf13/cobra v1.5.0

//replace go.tools => ../go.tools
require github.com/armstrongli/go-bmi v0.0.1

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
