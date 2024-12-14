# JSON 2 Go

This is a command line utility to convert JSON into a Go struct similar to the json2pojo utility for converting JSON to Java classes.

## Why does this exist?

I am a self-taught developer who would like to be able to show that I do know how to write code even if I do not have a degree in Computer Science.  This project is meant to do just that.  It is also a way for me to learn how to write things in Go.

I decided I should write this utility because I was not happy with the way online utilities were structuring the outputs.  I wanted something closer to what json2pojo produces where any inner structs are defined as separate types.  This makes them easier to read and understand as well as modify if needed.

## This ain't pretty...

Since this the first thing I've written in Go, I doubt the code will be considered pretty by any stretch of the imagination.  I am only hoping that it will function the way I intend it to.  The real challenge will be to not let the ADHD kick in before I've finished.  With any luck, the Autism hyper focus will help me stay on task... even if that means eating doesn't always happen--the things I do for the sake of a random idea had while sitting in my pajamas drinking coffee.

## Testing Standards

In order to ensure that testing actually happens, following a test first design is preferred.  This means that the first thing written for any function or feature is the unit test to prove that it works.  This should be a simple use case for this since the bulk of what is needed is text formatting into a very specific syntax.
