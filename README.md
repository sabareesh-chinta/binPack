bp3d
====

3D Bin Packing implementation based on [this paper](https://github.com/bom-d-van/binpacking/blob/master/erick_dube_507-034.pdf). The code is based on [binpacking by gedex](https://github.com/gedex/bp3d) but
modified to identify a single bin that fits all the items when one exists.

## Install

```
go get github.com/schinta-heidi/bp3d
```

## Usage

```
p := bp3d.NewPacker()

// Add bins.
p.AddBin(bp3d.NewBin("Small Bin", 10, 15, 20, 100))
p.AddBin(bp3d.NewBin("Medium Bin", 100, 150, 200, 1000))

// Add items.
p.AddItem(bp3d.NewItem("Item 1", 2, 2, 1, 2))
p.AddItem(bp3d.NewItem("Item 2", 3, 3, 2, 3))

// Pack items to bins.
out := p.Pack()
if out == nil {
	log.Fatal(out)
}

out is the smallest bin that can accommodate all the input items

```

See [`example/example.go`](./example/example.go)

## Credit

* hhttps://github.com/bom-d-van/binpacking/blob/master/erick_dube_507-034.pdf
* https://github.com/gedex/bp3d
* https://github.com/bom-d-van/binpacking

## License

[MIT](./LICENSE)
