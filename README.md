# go-lazy
A basic struct generator for lazy people....

# The why
One day at work I found myself having to implement a bunch of structs related to some crazy big layouts. Some were not that big but some had 30+ attributes. 

On that day I came back home and searched for a tool that could generate GO structs like rails does for ruby, but all I could find were a few solutions that converted code from JSON to GO structs and that wasn't going to help my case.

Then I decided to implement my own struct generator for Go so that I could make things flow faster at work.

This is what I came up with.

# How to use
First you have to build the tool with 

```shell
$ go build go-lazy
```

Then you have to run the tool and pass a few things.

```shell
$ ./go-lazy somePackageName someStructName argument0:int argument1:somecrazyargumenttype
```

and it will produce the following code:

```go
package somePackageName

type someStructName struct {
	argument0 int                  	`json:"argument0"`
	argument1 somecrazyargumenttype	`json:"argument1"`
}
```

Yeeah, already indented.

# What still needs to be done...

Improve error handling, as now it has none.

Improve user feedback, as right now it just breaks and doesn't explain why.

Let the user pick where they want to generate the file
...
