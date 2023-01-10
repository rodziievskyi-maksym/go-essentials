package main

import (
	"fmt"
	"math"
)

// Interface it is a contract that obligate his follower to implement it behavior

type netWorkProblems struct {
	massage string
	code    int
}

func (np netWorkProblems) Error() string {
	return fmt.Sprintf("network error! message: %s, code: %d", np.massage, np.code)
}

func handleErr(err error) {
	fmt.Println(err.Error())
}

func defineError() {
	np := netWorkProblems{
		massage: "we received a problem",
		code:    404,
	}

	handleErr(np)
}

//////

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func stringCaller() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}

	b := Book{
		Title:  "All about Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}

	Print(a)
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

////// Multiple behaviors in an Interface

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func AreaCaller() {
	circle := Circle{Radius: 10}
	PrintArea(circle)

	square := Square{
		Width:  10,
		Height: 5,
	}
	PrintArea(square)

	l := Less(circle, square)
	fmt.Printf("%+v is the smallest\n", l)

}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("Area of %s is %.2f\n", s.String(), s.Area())
}
