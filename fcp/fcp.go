package fcp

import (
	"os"
	"fmt"
	"bytes"
	"github.com/virtao/GoTypeBytes"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)


type FcpCanvas struct {
	Scene 		*widgets.QGraphicsScene
	View  		*widgets.QGraphicsView	
	}

type FullColorPallete struct {
	*widgets.QGraphicsWidget
	filename 	string
	pal			*pallete
	over		bool
	}

type FcpCell struct {
	*widgets.QGraphicsWidget
	index		int
	over		bool
	r, g, b		int
	}

type pallete struct {
	r, g, b		[]int
	}

func NewFcpCanvas() *FcpCanvas {
	
	fcpcan := &FcpCanvas{
		Scene: 		widgets.NewQGraphicsScene(nil),
		View: 		widgets.NewQGraphicsView(nil),
		}
	
	fcpcan.Scene.ConnectKeyPressEvent(fcpcan.keyPressEvent)
	fcpcan.Scene.ConnectWheelEvent(fcpcan.wheelEvent)
	
	fcpcan.View.ConnectResizeEvent(fcpcan.resizeEvent)
	
	fcpcan.Scene.AddItem(NewFullColorPallete())
	fcpcan.View.SetScene(fcpcan.Scene)

	fcpcan.View.Show()
	
	return fcpcan
	
	}

func NewFullColorPallete() *widgets.QGraphicsWidget {

	cp := &FullColorPallete{
		QGraphicsWidget: widgets.NewQGraphicsWidget(nil, 0), // se, 0) type widget
		filename: "C:/800.pal",
		}

	cp.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, 1) // Fixed Expanding

	cp.SetAcceptHoverEvents(true)
	cp.ConnectPaint(cp.Paint)
	cp.ConnectHoverEnterEvent(cp.HoverEnter)
	cp.ConnectHoverLeaveEvent(cp.HoverLeave)

	layout := widgets.NewQGraphicsGridLayout(nil) // cp)
	layout.SetSpacing(0)
	
	cp.pal = LoadPalette(cp.filename)
	
	x, y := 8, 16

	base := 10.0

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			layout.SetColumnFixedWidth(j, 10*base)
		}
		layout.SetRowFixedHeight(i, 5*base)
}
	z := 0
	
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
				c1 := NewFcpCell(cp, z)
				layout.AddItem2(c1, i, j, 1) // align left
				z = z + 1
		}
	}

	cp.SetLayout(layout)
	return cp.QGraphicsWidget
}

func NewFcpCell(cp *FullColorPallete, idx int) *widgets.QGraphicsWidget {

	fc := &FcpCell {
		QGraphicsWidget: widgets.NewQGraphicsWidget(nil, 0), //cp, 0) // type widget
		index: idx,
		r: cp.pal.r[idx], 
		g: cp.pal.g[idx], 
		b: cp.pal.b[idx],
		}

	fc.SetAcceptHoverEvents(true)
	fc.ConnectPaint(fc.Paint)
	fc.ConnectHoverEnterEvent(fc.HoverEnter)
	fc.ConnectHoverLeaveEvent(fc.HoverLeave)

	return fc.QGraphicsWidget
}

// fcp cell
func (fc *FcpCell) Paint(p *gui.QPainter, o *widgets.QStyleOptionGraphicsItem, w *widgets.QWidget) {

	var font = gui.NewQFont2("verdana", 7, 1, false)
	p.SetFont(font)

	var qpf = core.NewQPointF3(1.0, 1.0)

	color := gui.NewQColor3(fc.r, fc.g, fc.b, 255) // r, g, b, a
	var brush = gui.NewQBrush3(color, 1)
	
	var path = gui.NewQPainterPath()
	
	rf := core.NewQRectF5(o.Rect())
	path.AddRect(rf)
	
	p.DrawPath(path)
	p.FillPath(path, brush)

	if fc.over {

		p.DrawText(qpf, "r"+strconv.Itoa(fc.r)+"g"+strconv.Itoa(fc.g)+"b"+strconv.Itoa(fc.b))

	} else {

		p.DrawText(qpf, strconv.Itoa(fc.index))

	}
}

func (fc *FcpCell) HoverEnter(e *widgets.QGraphicsSceneHoverEvent) {
	fc.over = true
	e.Widget().Update3(e.Widget().Rect())
}

func (fc *FcpCell) HoverLeave(e *widgets.QGraphicsSceneHoverEvent) {
	fc.over = false
	e.Widget().Update3(e.Widget().Rect())
}

// color pallete
func (cp *FullColorPallete) Paint(p *gui.QPainter, o *widgets.QStyleOptionGraphicsItem, w *widgets.QWidget) {

	var font = gui.NewQFont2("verdana", 20, 1, false)
	p.SetFont(font)

	var qpf = core.NewQPointF3(1.0, 1.0)

	//p.DrawRect2(0,0,160,160) 
	p.DrawRect3(o.Rect())

	if cp.over {

		p.DrawText(qpf, "3")

	} else {

		p.DrawText(qpf, "4")

	}
}

func (cp *FullColorPallete) HoverEnter(e *widgets.QGraphicsSceneHoverEvent) {
	cp.over = true
	e.Widget().Update3(e.Widget().Rect())
}

func (cp *FullColorPallete) HoverLeave(e *widgets.QGraphicsSceneHoverEvent) {
	cp.over = false
	e.Widget().Update3(e.Widget().Rect())
}


// fcp canvas
func (fcpcan *FcpCanvas) keyPressEvent(e *gui.QKeyEvent) {

	if e.Modifiers() == core.Qt__ControlModifier {
		switch int32(e.Key()) {
		case int32(core.Qt__Key_Equal):
			fcpcan.View.Scale(1.25, 1.25)

		case int32(core.Qt__Key_Minus):
			fcpcan.View.Scale(0.8, 0.8)
		}
	}

}

func (fcpcan *FcpCanvas) wheelEvent(e *widgets.QGraphicsSceneWheelEvent) {
	if e.Modifiers() == core.Qt__ControlModifier {
		if e.Delta() > 0 {
			fcpcan.View.Scale(1.25, 1.25)
		} else {
			fcpcan.View.Scale(0.8, 0.8)
		}
	}
}

func (fcpcan *FcpCanvas) resizeEvent(e *gui.QResizeEvent) {
	
	fcpcan.View.FitInView(fcpcan.Scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio)
	
	}
	
func LoadPalette(filename string) *pallete {
	
	e := "0"
	
	file, err := os.Open(filename) // For read access.
	if err != nil {
		e = "1"
	}
	
	//statusbar.ShowMessage("palette found!", 0)
	
	fi, err := os.Stat(filename)
	if err != nil {
		e = "2"
	}

	data := make([]byte, fi.Size())
	fmt.Println("fi.Size:", fi.Size())

	count, err := file.Read(data)
	file.Close()
	if err != nil {
		e = "3"
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	var bt bytes.Buffer

	bt.Write(data[:count])
	
	var r, g, b []int
	
	
	for ; count > 0; count-- {

		r = append(r, typeBytes.BytesToInt(bt.Next(1)))

		g = append(g, typeBytes.BytesToInt(bt.Next(1)))

		b = append(b, typeBytes.BytesToInt(bt.Next(1)))
		
		bt.Next(3)

	}
	
	p := &pallete{
		r: r,
		g: g,
		b: b,
		}
	
	
	
	fmt.Println(e)
	//statusbar.ShowMessage(e, 0)
	
	return p
}
