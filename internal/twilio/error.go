package twilio

import "fmt"

type InvalidParamError struct {
	Parameter string
}

func (i InvalidParamError) Error() string {
	return fmt.Sprintf("invalid parameter: %s", i.Parameter)
}
