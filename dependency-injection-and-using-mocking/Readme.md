# Go Interfaces

Go has a static type system, so we have to explicitly state types when creating functions, eg - that the function argument is a string

```Go
func WriteGreeting(name string, f *os.File) error {
  greeting := Greeting(name)
  _, err := f.Write([]byte(greeting))
  if err != nil {
    return err
  }
  return nil
}
```

Above function writes a greeting to a file. Now if we wish to write a greeting to a *strings.Builder we'll need to have another function. Both os.File and strings.Builder have the Write method, and we only care that we are given something that has a Write method and our code will continue to work. Any type that has Write method can be passed to WriteGreeting!

```Go
Write(p []byte) (int, error)
```

```Go
type Writer interface {
  Write([]byte) (int, error)
}

func WriteGreeting(name string, w Writer) error {
    ....
}
```

Go uses duck typing - "If it looks like a duck, and it quacks like a duck, then it is a duck”.
If it looks like a Writer interface, has the methods defined by the Writer interface, then it is a Writer interface implementation. This isn't true for all statisticaly tyoed languages - for instance, in Java you need to explicitly state that a class implements an interface, otherwise the compiler won’t let you use it as an interface implementation.

Getting into more concrete details, any type that has every method defined by an interface will implement that interface. If a type has more methods than the interface defines that is okay too, but those additional methods will NOT be accessible while inside a function that accepts the interface.

Every method defined by the interface must be implemented for a type to implement an interface. The following example will result in a compiler error because our Demo type doesn’t define all of the methods defined by BigInterface:

```Go
type BigInterface interface {
  A()
  B()
  C()
}

type Demo struct {}

func (Demo) A() {}
func (Demo) B() {}

func Test(bi BigInterface) {}

func main() {
  var demo Demo
  Test(demo) // errors because Demo doesn't implement the C() method.
}
```
