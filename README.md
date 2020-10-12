# go-node

Run Javascript in Go using Node.js.

## Installing

```
go get -u github.com/tidwall/go-node
```

## Examples

Create a Node.js VM and run some Javascript.

```go
vm := node.New(nil)
vm.Run("iLoveTheJS = true");
youLoveIt := vm.Run("iLoveTheJS")
println(youLoveIt.String()) 

// output: 
// true
```

The `Run` function runs any Javascript code and returns a `Value` type that has
two methods, `String` and `Error`. If the `Run` call caused a Javascript
error then the `Error` method will contain that error. Otherwise, the `String`
method will contain the Javascript return value.

```go
vm := node.New(nil)
vm.Run(`
  function greatScott(){ 
    console.log('1.21 gigawatts?');
    return 'Delorean go vroom';
  }
`);
v := vm.Run("greatScott()");
if err := v.Error(); err != nil {
    log.Fatal(err)
}
println(v.String())

// output:
// 1.21 gigawatts?
// Deloran go vroom
```

You can "emit" messages from the Javascript to the Go. It's easy. Just use
the Options.OnEmit method in Go, and call the `emit()` function from JS.

```go
opts := node.Options {
    OnEmit: func(thing string) {
        println(thing)
    },
}
vm := node.New(&opts)
v := vm.Run("emit('a thing')")
if err := v.Error(); err != nil{
    log.Fatal(err)
}

// output: 
// a thing
```

Yay 🎉. Have fun now.
