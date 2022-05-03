package main

import (
	"fmt"
	"github/headfirstgo/gadget"
	// "log"
	// "github/headfirstgo/mypkg"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

type ComedyError string
func (c ComedyError) Error() string {
	return string(c)
}

type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

type CoffeePot string
func (c CoffeePot) String() string {
	return string(c) + " coffe pot"
}

type Gallons float64
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(Whistle("Toyco Canary"))
	// fmt.Println(Gallons(12.09248342))
	// fmt.Println(Liters(12.09248342))
	// fmt.Println(Milliliters(12.09248342))

	// coffeePot := CoffeePot("LuxBrew")
	// fmt.Print(coffeePot, "\n")
	// fmt.Println(coffeePot)
	// fmt.Printf("%s", coffeePot)

	// var err error = checkTemperature(121.379, 100.0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var err error
	// err = ComedyError("What's a programmer's favorite beer? Logger!")
	// fmt.Println(err)

	// TryOut(gadget.TapeRecorder{})
	// TryOut(gadget.TapePlayer{})

	// var noiseMaker NoiseMaker = Robot("Botco Ambler")
	// noiseMaker.MakeSound()
	// var robot Robot = noiseMaker.(Robot)
	// robot.Walk()
	// play(Whistle("Toyco Canary"))
	// play(Horn("Toyco Blaster"))
	// play(Robot("Botco Ambler"))

	// var toy NoiseMaker
	// toy = Whistle("Toyco Canary")
	// toy.MakeSound()
	// toy = Horn("Toyco Blaster")
	// toy.MakeSound()

	// var value mypkg.MyInterface
	// value = mypkg.MyType(5)
	// value.MethodWithoutParameters()
	// value.MethodWithParameters(127.3)
	// fmt.Println(value.MethodWithReturnValue())

	// mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	// var player Player = gadget.TapePlayer{}
	// playList(player, mixtape)
	// player = gadget.TapeRecorder{}
	// playList(player, mixtape)
}