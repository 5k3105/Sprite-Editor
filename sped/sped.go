package sped

import(
	//"strconv"
	"local/spc/spritelib"
	
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	)

type SpriteCanvas struct {
	Scene 	*widgets.QGraphicsScene
	View  	*widgets.QGraphicsView
	Editor	*SpriteEditor
	}

type SpriteEditor struct {
	*widgets.QGraphicsWidget
	over		bool
	Cells		*treemap.Map
	}

type SpriteCell struct {
	*widgets.QGraphicsWidget
	Index		int
	over		bool
	R, G, B		int
	
	font  *gui.QFont
	qpf   *core.QPointF
	Brush *gui.QBrush
	path  *gui.QPainterPath	
	
	
	}


func NewSpriteCanvas() *SpriteCanvas {
	
	spcan := &SpriteCanvas{
		Scene: 	widgets.NewQGraphicsScene(nil),
		View: 	widgets.NewQGraphicsView(nil),
		Editor: NewSpriteEditor(),
		}
	
	spcan.Scene.ConnectKeyPressEvent(spcan.keyPressEvent)
	spcan.Scene.ConnectWheelEvent(spcan.wheelEvent)
	
	spcan.View.ConnectResizeEvent(spcan.resizeEvent)
	
	spcan.Scene.AddItem(spcan.Editor.QGraphicsWidget)

	spcan.View.SetScene(spcan.Scene)
	
	spcan.View.Show()
	
	return spcan
	
	}


func NewSpriteEditor() *SpriteEditor  {
	
	se := &SpriteEditor{
		QGraphicsWidget: 	widgets.NewQGraphicsWidget(nil, 0), // type widget
		Cells:				treemap.NewWithIntComparator(),
	}
	
	se.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, 1) // Fixed Expanding

	se.SetAcceptHoverEvents(true)
	se.ConnectPaint(se.Paint)
	se.ConnectHoverEnterEvent(se.HoverEnter)
	se.ConnectHoverLeaveEvent(se.HoverLeave)

	layout := widgets.NewQGraphicsGridLayout(nil) // se)
	layout.SetSpacing(0)
	
	sl := spritelib.NewSpriteLib()
	sl.NewAtariSprite(16, 32)
	sl.NewAtariSprite(16, 32)

	_, x, y := sl.GetSprite(1)

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
				scell := NewSpriteCell(se, z)
				se.Cells.Put(z, scell)
				layout.AddItem2(scell, i, j, 1)
				z = z + 1
		}
	}

	se.SetLayout(layout)
	return se

	}
	

func NewSpriteCell(parent *SpriteEditor, idx int) *SpriteCell { //*widgets.QGraphicsWidget {

	sc := &SpriteCell{
	QGraphicsWidget: widgets.NewQGraphicsWidget(nil, 0), // parent, 0) // type widget
	Index: idx,
	
		font: gui.NewQFont2("verdana", 7, 1, false),
		qpf:  core.NewQPointF3(1.0, 1.0),
		path: gui.NewQPainterPath(),	
	}

	var color = gui.NewQColor3(sc.R, sc.G, sc.B, 255) // r, g, b, a
	sc.Brush = gui.NewQBrush3(color, 1)

	sc.SetAcceptHoverEvents(true)
	sc.ConnectPaint(sc.Paint)
	sc.ConnectHoverEnterEvent(sc.HoverEnter)
	sc.ConnectHoverLeaveEvent(sc.HoverLeave)

	return sc //.QGraphicsWidget
}

// sprite cell
func (sc *SpriteCell) Paint(p *gui.QPainter, o *widgets.QStyleOptionGraphicsItem, w *widgets.QWidget) {


	if sc.path.IsEmpty() {
		sc.path.AddRect(core.NewQRectF5(o.Rect())) //will only be called once, with the first paint event (as the painter is always the same for each cell)
	}

	p.SetFont(sc.font)
	p.DrawPath(sc.path)
	p.FillPath(sc.path, sc.Brush)

	if sc.over {

		p.DrawText(sc.qpf, "")
		//p.DrawText(qpf, "r"+strconv.Itoa(sc.R)+"g"+strconv.Itoa(sc.G)+"b"+strconv.Itoa(sc.B))

	} else {
		p.DrawText(sc.qpf, "")
		//p.DrawText(qpf, strconv.Itoa(sc.Index))

	}



	//var font = gui.NewQFont2("verdana", 7, 1, false)
	//p.SetFont(font)

	//var qpf = core.NewQPointF3(1.0, 1.0)

	//color := gui.NewQColor3(sc.R, sc.G, sc.B, 255) // 255) // r, g, b, a
	//var brush = gui.NewQBrush3(color, 1)
	
	//var path = gui.NewQPainterPath()
	
	//rf := core.NewQRectF5(o.Rect())
	//path.AddRect(rf)
	
	//p.DrawPath(path)
	//p.FillPath(path, brush)

	//if sc.over {
		
		//p.DrawText(qpf,"")
		////p.DrawText(qpf, "r"+strconv.Itoa(sc.R)+"g"+strconv.Itoa(sc.G)+"b"+strconv.Itoa(sc.B))

	//} else {
		//p.DrawText(qpf,"")
		////p.DrawText(qpf, strconv.Itoa(sc.Index))

	//}
	
}

func (sc *SpriteCell) HoverEnter(e *widgets.QGraphicsSceneHoverEvent) {
	sc.over = true
	e.Widget().Update3(e.Widget().Rect())
}

func (sc *SpriteCell) HoverLeave(e *widgets.QGraphicsSceneHoverEvent) {
	sc.over = false
	e.Widget().Update3(e.Widget().Rect())
}

// sprite editor
func (se *SpriteEditor) Paint(p *gui.QPainter, o *widgets.QStyleOptionGraphicsItem, w *widgets.QWidget) {

	var font = gui.NewQFont2("verdana", 20, 1, false)
	p.SetFont(font)

	var qpf = core.NewQPointF3(1.0, 1.0)

	p.DrawRect3(o.Rect())

	if se.over {

		p.DrawText(qpf, "3")

	} else {

		p.DrawText(qpf, "4")

	}
}

func (se *SpriteEditor) HoverEnter(e *widgets.QGraphicsSceneHoverEvent) {
	se.over = true
	e.Widget().Update3(e.Widget().Rect())
}

func (se *SpriteEditor) HoverLeave(e *widgets.QGraphicsSceneHoverEvent) {
	se.over = false
	e.Widget().Update3(e.Widget().Rect())
}


// sprite canvas
func (spcan *SpriteCanvas) keyPressEvent(e *gui.QKeyEvent) {

	if e.Modifiers() == core.Qt__ControlModifier {
		switch int32(e.Key()) {
		case int32(core.Qt__Key_Equal):
			spcan.View.Scale(1.25, 1.25)

		case int32(core.Qt__Key_Minus):
			spcan.View.Scale(0.8, 0.8)
		}
	}

	//if e.Key() == int(core.Qt__Key_Escape) {
	//	c.ClearScene()
	//	c.AddCells()
	//}
}

func (spcan *SpriteCanvas) wheelEvent(e *widgets.QGraphicsSceneWheelEvent) {
	if e.Modifiers() == core.Qt__ControlModifier {
		if e.Delta() > 0 {
			spcan.View.Scale(1.25, 1.25)
		} else {
			spcan.View.Scale(0.8, 0.8)
		}
	}
}

func (spcan *SpriteCanvas) resizeEvent(e *gui.QResizeEvent) {
	
	spcan.View.FitInView(spcan.Scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio)
	//se.BoundingRect()
	//spcan.View.Scale(1, -1)
	
	}


