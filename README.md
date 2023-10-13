# raylib

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mewspring/raylib)

The raylib project implements window creation, event handling and image drawing using [raylib](https://www.raylib.com/) version 4.5.

## Dependencies

```bash
pacman -Sy raylib
```

## Examples

### input_events

The [input_events](https://github.com/mewspring/raylib/blob/master/examples/input_events/main.go#L33) command demonstrates how to create a window and handle input events using [window.Open](http://godoc.org/github.com/mewspring/raylib/window#Open) and [Window.PollEvent](http://godoc.org/github.com/mewspring/raylib/window#Window.PollEvent).

```bash
go install -v github.com/mewspring/raylib/examples/input_events@master
```

![Screenshot - input_events](https://raw.githubusercontent.com/mewspring/raylib/master/examples/input_events/input_events.png)

### tiny

The [tiny](https://github.com/mewspring/raylib/blob/master/examples/tiny/main.go#L60) command demonstrates how to render images onto the window using [Window.Draw](http://godoc.org/github.com/mewspring/raylib/window#Window.Draw) and [Window.DrawRect](http://godoc.org/github.com/mewspring/raylib/window#Window.DrawRect). It also gives an example of a basic event loop.

```bash
go install -v github.com/mewspring/raylib/examples/tiny@master
```

![Screenshot - tiny](https://raw.githubusercontent.com/mewspring/raylib/master/examples/tiny/tiny.png)
