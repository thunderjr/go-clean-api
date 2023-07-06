package helpers

import "reflect"

func invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func InvokeWithError(any interface{}, name string, args ...interface{}) error {
	result := invoke(any, name, args...)

	if len(result) > 0 {
		if err, ok := result[0].Interface().(error); ok {
			return err
		}
	}

	return nil
}
