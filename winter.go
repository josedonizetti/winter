package winter

import "reflect"

type Container struct {
  objects []interface{}
}

func NewContainer() *Container {
  objects := []interface{}{}
  return &Container{objects}
}

func (c *Container) add(object interface{}) {
  c.objects = append(c.objects, object)
}

func (c *Container) execute(function interface{}) {
    functionType := reflect.TypeOf(function)
    functionValue := reflect.ValueOf(function)

    if (functionType.Kind() != reflect.Func) {
      panic("Must be a function")
    }

    args := functionType.NumIn()

    if args == 0 {
      functionValue.Call([]reflect.Value{})
    } else {
      values := []reflect.Value{}

      for i := 0; i < args; i++ {
        value := c.findValue(functionType.In(i))
        values = append(values, value)
      }

      functionValue.Call(values)
    }
}

func (c *Container) findValue(param reflect.Type) reflect.Value {
  for _, object := range c.objects {
    if param.Kind() == reflect.TypeOf(object).Kind() {
      return reflect.ValueOf(object)
    }
  }

  panic("Container must manage dependency ???")
}
