# Basicly a better version of fmt #

> [!WARNING]
> This package works great but isn't maintained a lot. So, if you'll use it. Remember that you will fix the bugs you'll face with.

## PREFIX ##
Imlementations for prefix handling.
Define the printer:
```go
p := efmt.NewPrinter("[Server] ")
```
Printing:
```go
p.Println("Listening on: ", "127.0.0.1:1111")
```
![Output](https://github.com/efexplose/efmt/assets/52001980/977b60aa-bff8-4d58-ba73-b99b68e47cb1)

Error handeling:
```go
var err error
message := "unexpected"
if message != "expected" {
  err := p.Errorf("Unexpected message: %s", message)
  p.Println(err)
}
```
![Output](https://github.com/efexplose/efmt/assets/52001980/2e526a47-0a6f-4381-bfe9-e1351abe9bfd)


## Colors ##
Add 'c' at the end of the original name of the function
```go
p.Printlnc(efmt.Yellow, "Listening on: ", "127.0.0.1:1111")
```
![Output](https://github.com/efexplose/efmt/assets/52001980/5834b852-eaa2-40e5-987a-63c6ea3e1226)

The code above automatically detects the header by spliting the string from both '\n' and ':'.
