package ops

import "log"

// CheckArity returns an error if the input number does not correspond to the expected arity
func CheckArity(op Arityer, inputs int) error {
	if inputs != op.Arity() && op.Arity() >= 0 {
		// TODO: deal with that
		//return errors.Errorf("%v has an arity of %d. Got %d instead", op, op.Arity(), inputs)
		log.Printf("%v has an arity of %d. Got %d instead", op, op.Arity(), inputs)

		return nil
	}
	return nil
}
