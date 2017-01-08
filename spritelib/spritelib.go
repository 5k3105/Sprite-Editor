package spritelib

import (
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/virtao/GoTypeBytes"
	"io/ioutil"
	"os"
)

const (
	commodoreSprite = 1
	atariSprite     = 2
	astrocadeSprite = 3
)

type SpriteLib struct {
	Sprites    *treemap.Map // []AtariSprite
	SpriteDimX *treemap.Map // []int
	SpriteDimY *treemap.Map // []int
	lastid     int
}

type Sprite struct {
	spritetype int
	color      []byte
	pixel      []byte // 16 x ...
}

var autonum int

//func main() {

//sl := NewSpriteLib()
//sl.NewAtariSprite(16,32)
//sl.NewAtariSprite(16,32)

//DisplayAtariSprite(sl,1)

//sl.SaveSpriteLib("dat1")

//sl2 := NewSpriteLib()

//sl2.LoadSpriteLib("dat1")

//}

func DisplayAtariSprite(sl *SpriteLib, spid int) {

	sprite, x, y := sl.GetSprite(spid)
	sp := 0
	cp := 0

	sprite.pixel[1] = 1
	sprite.color[20] = 5

	nc := GetSpriteType(sprite.spritetype)

	for i := 0; i < y; i++ {
		fmt.Println(i, ": ", sprite.pixel[sp:sp+x], sprite.color[cp:cp+nc])
		sp = sp + x
		cp = cp + nc

	}

}

func (sl *SpriteLib) GetSprite(spid int) (*Sprite, int, int) {

	spriteif, _ := sl.Sprites.Get(spid)
	sprite := spriteif.(*Sprite)
	xif, _ := sl.SpriteDimX.Get(spid)
	dx := xif.(int)
	yif, _ := sl.SpriteDimY.Get(spid)
	dy := yif.(int)

	return sprite, dx, dy

}

func NewSpriteLib() *SpriteLib {

	sl := &SpriteLib{
		Sprites:    treemap.NewWithIntComparator(),
		SpriteDimX: treemap.NewWithIntComparator(),
		SpriteDimY: treemap.NewWithIntComparator(),
	}

	return sl

}

func (sl *SpriteLib) NewAtariSprite(xdim int, ydim int) {

	sprite := NewAtariSprite(xdim, ydim)
	ni := newint()
	sl.Sprites.Put(ni, sprite)
	sl.SpriteDimX.Put(ni, xdim)
	sl.SpriteDimY.Put(ni, ydim)
}

func NewAtariSprite(xdim int, ydim int) *Sprite {

	sprite := &Sprite{
		spritetype: atariSprite,
		color:      make([]byte, 3*ydim), // 3 colors
		pixel:      make([]byte, xdim*ydim),
	}

	return sprite

}

func (sl *SpriteLib) LoadSpriteLib(filename string) {

	file, err := os.Open(filename) // For read access.
	if err != nil {
		fmt.Println("fuck")
	}

	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Println("goto hell")
	}

	data := make([]byte, fi.Size())
	fmt.Println("fi.Size:", fi.Size())

	count, err := file.Read(data)
	file.Close()
	if err != nil {
		fmt.Println("your file fukd up")
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	var b bytes.Buffer

	b.Write(data[:count])

	size := typeBytes.BytesToInt(b.Next(8))
	fmt.Println("size:", size)

	for ; size > 0; size-- {

		dimx := typeBytes.BytesToInt(b.Next(8))
		fmt.Println("dimx:", dimx)

		dimy := typeBytes.BytesToInt(b.Next(8))
		fmt.Println("dimy:", dimy)

		spritetype := typeBytes.BytesToInt(b.Next(8))
		nc := GetSpriteType(spritetype)

		color := b.Next(nc * dimy)
		fmt.Println(color)

		pixel := b.Next(dimx * dimy)
		fmt.Println(pixel)

		sprite := &Sprite{
			color: color,
			pixel: pixel,
		}

		ni := newint()
		sl.Sprites.Put(ni, sprite)

	}
}

func (sl *SpriteLib) SaveSpriteLib(filename string) {

	var b bytes.Buffer

	size := sl.Sprites.Size()

	b.Write(typeBytes.IntToBytes(size)) // size 1

	itx := sl.SpriteDimX.Iterator() // dimx 16
	ity := sl.SpriteDimY.Iterator() // dimy 32
	its := sl.Sprites.Iterator()    // pixel and color

	itx.Next()
	ity.Next()
	its.Next()

	for ; size > 0; size-- {

		b.Write(typeBytes.IntToBytes(itx.Value().(int)))
		b.Write(typeBytes.IntToBytes(ity.Value().(int)))

		b.Write(typeBytes.IntToBytes(its.Value().(*Sprite).spritetype))
		b.Write(its.Value().(*Sprite).color)
		b.Write(its.Value().(*Sprite).pixel)

		itx.Next()
		ity.Next()
		its.Next()

	}

	err := ioutil.WriteFile(filename, b.Bytes(), 0644)
	check(err)

}

func GetSpriteType(spritetype int) int { // returns number of colors

	var ncolors int
	switch spritetype {
	case atariSprite:
		ncolors = 3
	}

	return ncolors
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newint() int {

	autonum = autonum + 1
	return autonum

}

//it := sl.SpriteDimX.Iterator()	// dimx 16
//for it.Next() {
//b.Write([]byte(typeBytes.IntToBytes(it.Value().(int))))
//}

//it = sl.SpriteDimY.Iterator()	// dimy 32
//for it.Next() {
//b.Write([]byte(typeBytes.IntToBytes(it.Value().(int))))
//}

//it = sl.Sprites.Iterator()	// color
//for it.Next() {
//b.Write(it.Value().(*AtariSprite).color)
//}

//it = sl.Sprites.Iterator()	// pixel
//for it.Next() {
//b.Write(it.Value().(*AtariSprite).pixel)
//}
