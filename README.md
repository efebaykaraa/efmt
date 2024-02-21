# Basicly an extended fmt! #

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
![Output](https://github.com/efexplose/efmt/.assets/prefix.png)
Error handeling:
```go
message := "unexpected"
if message != "expected" {
  p.Errorf("Unexpected message: %s", message)
}
```
![Output](https://github.com/efexplose/efmt/.assets/error.png)

## Colors ##
You can deal with colors easyly
```go
p.Println(efmt.Yellow, "Listening on: ", "127.0.0.1:1111")
```
![Output](https://github.com/efexplose/efmt/.assets/color.png)
The code above automatically detects the header by spliting the string from both '\n' and ':'.

