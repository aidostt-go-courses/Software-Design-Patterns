package main

import (
	_ "errors"
	"fmt"
)

type Device interface {
	getModel() string
	getPrice() int
}

type Computer struct {
}

func (a *Computer) getModel() string {
	return "Computer"
}

func (a *Computer) getPrice() int {
	return 100
}

type Laptop struct {
}

func (l *Laptop) getModel() string {
	return "Computer"
}

func (l *Laptop) getPrice() int {
	return 100
}

type GraphicalCard struct {
	device Device
}

func (g *GraphicalCard) getModel() string {
	deviceModel := g.device.getModel()
	return deviceModel + ", GraphicalCard - RTX 4090 8GB"
}
func (g *GraphicalCard) getPrice() int {
	devicePrice := g.device.getPrice()
	return devicePrice + 1200
}

type Processor struct {
	device Device
}

func (p *Processor) getModel() string {
	deviceModel := p.device.getModel()
	return deviceModel + ", CPU - intel core i7 12700i"
}

func (p *Processor) getPrice() int {
	devicePrice := p.device.getPrice()
	return devicePrice + 800
}

type RAM struct {
	device Device
}

func (r *RAM) GetModel() string {
	deviceModel := r.device.getModel()
	return deviceModel + ", Ram - HyperX 16GB"
}

func (r *RAM) GetPrice() int {
	devicePrice := r.device.getPrice()
	return devicePrice + 400
}

type Camera struct {
	device Device
}

func (c *Camera) GetModel() string {
	deviceModel := c.device.getModel()
	return deviceModel + ", Camera - Logitech UltraHD"
}

func (c *Camera) GetPrice() int {
	devicePrice := c.device.getPrice()
	return devicePrice + 300
}

func main() {
	laptop := &Laptop{}

	laptopGraphicalCard := &GraphicalCard{
		device: laptop,
	}

	laptopGraphicalCardProcessor := &Processor{
		device: laptopGraphicalCard,
	}

	computer := &Computer{}

	computerGraphicalCard := &GraphicalCard{
		device: computer,
	}

	computerGraphicalCardProcessor := &Processor{
		device: computerGraphicalCard,
	}

	computerGraphicalCardProcessorCamera := &Camera{
		device: computerGraphicalCardProcessor,
	}

	fmt.Printf("%v \n price - %v ", computerGraphicalCardProcessorCamera.GetModel(), computerGraphicalCardProcessorCamera.GetPrice())
	fmt.Println()
	fmt.Printf("%v \n price - %v ", laptopGraphicalCardProcessor.getModel(), laptopGraphicalCardProcessor.getPrice())

}
