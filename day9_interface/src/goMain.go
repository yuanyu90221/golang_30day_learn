package main

import (
	"fmt"
	"io"
)

type Engine interface {
	Start()
	Stop()
}

type CarEngine struct {
}

func (c CarEngine) Start() {
	fmt.Println("Car engine is started")
}

func (c CarEngine) Stop() {
	fmt.Println("Car engine is stoped")
}

type TrainEngine struct {
}

func (t TrainEngine) Start() {
	fmt.Println("Train engine is started")
}

func (t TrainEngine) Stop() {
	fmt.Println("Train engine is stoped")
}
func Stopping(e Engine) {
	e.Stop()
}
func Starting(e Engine) {
	e.Start()
}
func main() {
	carEngine := CarEngine{}
	trainEngine := TrainEngine{}
	engines := []Engine{
		carEngine, trainEngine,
	}

	for _, engine := range engines {
		Starting(engine)
		Stopping(engine)
	}
	duckTypeDemo()
	duckTypeDemo1()
	mainDemo()
	serviceDone()
	circleDemo()
}

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("Write has be invoked")
	return 0, nil
}

func (s *Socket) Close() error {
	// Socket close method
	fmt.Println("Close has be invoked")
	return nil
}

func usingWriter(writer io.Writer) {
	writer.Write(nil)
}
func usingCloser(closer io.Closer) {
	closer.Close()
}
func duckTypeDemo() {
	s := new(Socket)
	usingWriter(s)
	usingCloser(s)
}

type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func duckTypeDemo1() {
	var wc io.WriteCloser = new(device)

	wc.Write(nil)

	wc.Close()

	var writeOnly io.Writer = new(device)

	writeOnly.Write(nil)

}

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird can fly")
}

func (b *bird) Walk() {
	fmt.Println("bird can walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig can walk")
}

func mainDemo() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {
		switch obj := obj.(type) {
		case Flyer:
			fmt.Printf("name: %s\n", name)
			obj.Fly()
		case Walker:
			fmt.Printf("name: %s\n", name)
			obj.Walk()
		}
	}
}

type Service interface {
	Start()
	Log(string)
}

type Logger struct{}

func (g *Logger) Log(l string) {
	fmt.Println(l)
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {
	fmt.Println("game service start")
}

func serviceDone() {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}

type Shape interface {
	Area() int64
}

type Rectangle struct {
	width, height int64
}

func NewRectangle(width, height int64) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) Area() int64 {
	return r.width * r.height
}

type Circle struct {
	radius int64
}

func NewCircle(radius int64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Area() int64 {
	return c.radius * c.radius
}

func circleDemo() {
	r := NewRectangle(10, 5)
	c := NewCircle(5)
	s := []Shape{r, c}
	for _, shape := range s {
		fmt.Println(shape.Area())
	}
}
