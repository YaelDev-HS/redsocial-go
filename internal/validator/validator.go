package validator

import "regexp"

var (
	EmailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

type Validator struct {
	errors map[string]string
}

func New() *Validator {
	return &Validator{
		errors: map[string]string{},
	}
}

func (v *Validator) AddError(key, value string) {
	_, ok := v.errors[key]

	if !ok {
		v.errors[key] = value
	}
}

func (v *Validator) Check(ok bool, key, value string) {
	if !ok {
		v.AddError(key, value)
	}
}

func (v *Validator) Match(value string, match string) bool {
	ok, err := regexp.MatchString(value, match)

	if err != nil {
		return false
	}

	return ok
}

func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

func (v *Validator) Errors() map[string]string {
	return v.errors
}
