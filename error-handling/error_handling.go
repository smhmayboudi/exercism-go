package erratum

import (
	"errors"
)

// Use opens a resource opener
func Use(opener ResourceOpener, input string) (err error) {
	counter := 0
	var resource Resource
	for {
		resource, err = opener()
		if err != nil {
			if transErr, ok := err.(TransientError); ok {
				counter++
				if counter > 3 {
					return transErr
				}
				continue
			} else {
				return err
			}
		}
		break
	}

	defer resource.Close()
	defer func(res Resource) {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe.inner
			} else {
				err = errors.New("meh")
			}
		}
	}(resource)

	resource.Frob(input)
	return err
}
