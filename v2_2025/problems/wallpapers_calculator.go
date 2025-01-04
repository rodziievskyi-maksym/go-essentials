package main

import "fmt"

type RoomSquareMeters struct {
	//float32 can store a value with a precision of 7 digits
	//float64 can store a value with a precision of 15 digits
	length float32
	width  float32
}

func (r RoomSquareMeters) Area() float32 {
	return r.length * r.width
}

func (r RoomSquareMeters) Perimeter() float32 {
	return 2*r.length + 2*r.width
}

type RectangularRoomSquareMeters struct {
	*RoomSquareMeters
	wallHeight float32
}

type WallPaperRolls struct {
	*RectangularRoomSquareMeters
	rollWidth  float32
	rollLength float32
	price      float32
}

func (w WallPaperRolls) RollsNeeded() float32 {
	return w.SurfaceArea() / (w.rollWidth * w.rollLength)
}

func (w WallPaperRolls) TotalCost() float32 {
	return w.RollsNeeded() * w.price
}

func (r RectangularRoomSquareMeters) SurfaceArea() float32 {
	return r.RoomSquareMeters.Perimeter() * r.wallHeight
}

func main() {
	var rollWidth, rollLength, price float32
	fmt.Println("Enter the roll width and length and the price:")
	if _, err := fmt.Scan(&rollWidth, &rollLength, &price); err != nil {
		fmt.Println(err)
		return
	}

	roomSquareMeters := RoomSquareMeters{3.45, 3.12}

	room := RectangularRoomSquareMeters{
		RoomSquareMeters: &roomSquareMeters,
		wallHeight:       2.65,
	}

	wallPaperRolls := WallPaperRolls{
		&room,
		rollWidth,
		rollLength,
		price,
	}

	fmt.Printf("Room's Square Meters: %.2fm2\n", room.RoomSquareMeters.Area())
	fmt.Printf("Room's Surface Area: %.2fm2\n", room.SurfaceArea())
	fmt.Printf("Wallpaper Rolls Needed: %.2f\n", wallPaperRolls.RollsNeeded())
	fmt.Printf("Total Cost: %.2f\n", wallPaperRolls.TotalCost())
}
