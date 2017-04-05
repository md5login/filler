# filler [![Go Report Card](https://goreportcard.com/badge/github.com/md5login/filler)](https://goreportcard.com/report/github.com/yaronsumel/filler) [![Build Status](https://travis-ci.org/md5login/filler.svg?branch=master)](https://travis-ci.org/md5login/filler) [![GoDoc](https://godoc.org/github.com/md5login/filler?status.svg)](https://godoc.org/github.com/md5login/filler)
###### small util to fill gaps in your structs 

Installation
------
```bash
$ go get github.com/yaronsumel/filler
```

[Working Example](https://github.com/md5login/filler/blob/master/example/example.go)

Usage
------

##### fill tag

###### `fill:"[FillerName:OptionalValue]"`
###### `fill:"[User:UserId]"` - Fill current filed with the "User" Filler and UserId value
###### `fill:"[SayHello]"` = Fill current with "SayHello" Filler Without any value 


###### Add the `fill` tag in your model
```go
type Model struct {
	UserId   bson.ObjectId 
	FieldA   string        `fill:"SayHelloFiller"`
	UserName string        `fill:"UserNameFiller:UserId"`
}
```
###### Register the fillers
```go
	filler.RegFiller(filler.Filler{
		Tag: "UserNameFiller",
		Fn: func(value interface{}) (interface{}, error) {
			return "this is the user name", nil
		},
	})

	filler.RegFiller(filler.Filler{
		Tag: "SayHelloFiller",
		Fn: func(value interface{}) (interface{}, error) {
			return "Hello", nil
		},
	})
```

###### and Fill
```go
	filler.Fill(&m)
```

###### Add the 'defaults' tag in your model
```go
type Model struct {
	UserId   bson.ObjectId 
	FieldA   string        `defaults:"This is a field"`
	Age	 int64         `defaults:"18"`
}
```
###### Fill the defaul values
```go
	filler.Defaults(&m)
```

> ##### Forked from [Filler](https://github.com/yaronsumel/filler/) by [YaronSumel](https://twitter.com/yaronsumel) #####
