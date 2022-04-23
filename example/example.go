package main

import (
	"fmt"
	"log"

	"github.com/schinta-heidi/bp3d"
)

func main() {
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Bin 1", 28, 16, 12, 100))
	p.AddBin(bp3d.NewBin("Bin 2", 14, 12, 10, 100))
	p.AddBin(bp3d.NewBin("Bin 3", 10, 10, 10, 100))
	p.AddBin(bp3d.NewBin("Bin 4", 10, 6, 4, 100))
	p.AddBin(bp3d.NewBin("Bin 5", 18, 12, 12, 100))
	p.AddBin(bp3d.NewBin("Bin 6", 20, 16, 16, 100))
	p.AddBin(bp3d.NewBin("Bin 7", 2, 1, 2, 20))
	p.AddBin(bp3d.NewBin("Bin 8", 4, 3, 5, 25))
	p.AddBin(bp3d.NewBin("Bin 9", 7, 7, 7, 50))

	// p.AddBin(bp3d.NewBin("Medium Bin", 100, 150, 200, 1000))

	// Add items.
	p.AddItem(bp3d.NewItem("Item 1", 5, 4, 2, 2))
	p.AddItem(bp3d.NewItem("Item 2", 3, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Item 3", 1, 8, 2, 4))
	p.AddItem(bp3d.NewItem("Item 4", 6, 3, 5, 5))
	p.AddItem(bp3d.NewItem("Item 5", 5, 4, 2, 2))
	p.AddItem(bp3d.NewItem("Item 6", 3, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Item 7", 1, 8, 2, 4))
	p.AddItem(bp3d.NewItem("Item 8", 6, 3, 5, 5))
	p.AddItem(bp3d.NewItem("Item 9", 5, 4, 2, 2))
	p.AddItem(bp3d.NewItem("Item 10", 3, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Item 11", 1, 8, 2, 4))
	p.AddItem(bp3d.NewItem("Item 12", 6, 3, 5, 5))
	p.AddItem(bp3d.NewItem("Item 13", 5, 4, 2, 2))
	p.AddItem(bp3d.NewItem("Item 14", 3, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Item 15", 1, 8, 2, 4))
	p.AddItem(bp3d.NewItem("Item 16", 6, 3, 5, 5))
	p.AddItem(bp3d.NewItem("Item 17", 10, 8, 5, 5))

	// Pack items to bins.
	out := p.Pack()
	if out == nil {
		log.Fatal(out)
	}

	// Will output:
	//
	// Bin 2(14x12x10, max_weight:100)
	// packed items:
	//   Item 17(10x8x5, weight: 5) pos(0,0,0) rt(RotationType_WHD (w,h,d))
	//   Item 16(6x3x5, weight: 5) pos(10,0,0) rt(RotationType_HWD (h,w,d))
	//   Item 4(6x3x5, weight: 5) pos(0,8,0) rt(RotationType_WHD (w,h,d))
	//   Item 12(6x3x5, weight: 5) pos(6,8,0) rt(RotationType_WHD (w,h,d))
	//   Item 8(6x3x5, weight: 5) pos(0,0,5) rt(RotationType_WHD (w,h,d))
	//   Item 13(5x4x2, weight: 2) pos(12,8,0) rt(RotationType_DHW (d,h,w))
	//   Item 5(5x4x2, weight: 2) pos(6,0,5) rt(RotationType_WHD (w,h,d))
	//   Item 9(5x4x2, weight: 2) pos(11,0,5) rt(RotationType_DHW (d,h,w))
	//   Item 1(5x4x2, weight: 2) pos(0,3,5) rt(RotationType_WHD (w,h,d))
	//   Item 10(3x3x2, weight: 3) pos(6,4,5) rt(RotationType_WHD (w,h,d))
	//   Item 6(3x3x2, weight: 3) pos(9,4,5) rt(RotationType_WHD (w,h,d))
	//   Item 14(3x3x2, weight: 3) pos(12,4,5) rt(RotationType_DHW (d,h,w))
	//   Item 2(3x3x2, weight: 3) pos(0,7,5) rt(RotationType_WHD (w,h,d))
	//   Item 7(1x8x2, weight: 4) pos(5,3,5) rt(RotationType_WHD (w,h,d))
	//   Item 3(1x8x2, weight: 4) pos(0,11,0) rt(RotationType_HWD (h,w,d))
	//   Item 15(1x8x2, weight: 4) pos(6,7,5) rt(RotationType_HWD (h,w,d))
	//   Item 11(1x8x2, weight: 4) pos(5,11,5) rt(RotationType_HWD (h,w,d))
	displayPacked(out)
}

func displayPacked(bin *bp3d.Bin) {
	fmt.Println(bin)
	fmt.Println(" packed items:")
	for _, i := range bin.Items {
		fmt.Println("  ", i)
	}
}
