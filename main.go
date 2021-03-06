package main

import (
	"os"
	"strconv"
	
	"local/spc/fcp"
	"local/spc/sped"
	
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	r,g,b		int
	statusbar 	*widgets.QStatusBar
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Main Window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Sprite Editor")

	// Main Widget
	//mw := widgets.NewQWidget(nil, 0)

	// Statusbar
	statusbar = widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)

	// sprite editor panel
	var dsped = widgets.NewQDockWidget("Sprite Editor", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dsped)
	var spcan = sped.NewSpriteCanvas() // window
	dsped.SetWidget(spcan.View)
	
	ConnectSpriteEvents(spcan)

	// full color pallete panel
	var dfcp = widgets.NewQDockWidget("Full Color Pallete", window, 0)
	window.AddDockWidget(core.Qt__RightDockWidgetArea, dfcp)
	var fcpcan = fcp.NewFcpCanvas() // window
	dfcp.SetWidget(fcpcan.View)
	
	ConnectFcpEvents(fcpcan)
	
	statusbar.ShowMessage(core.QCoreApplication_ApplicationDirPath(), 0)

	// Set Central Widget
	//window.SetCentralWidget(mw)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.ShowMaximized()
	widgets.QApplication_Exec()
}

func ConnectSpriteEvents(spcan *sped.SpriteCanvas) {

	it := spcan.Editor.Cells.Iterator()
	for it.Next() {
		sc := it.Value().(*sped.SpriteCell)
		
		sc.ConnectMousePressEvent(SpriteCellClick(sc, spcan)) // QGraphicsWidget
		
		}
	}


func SpriteCellClick(sc *sped.SpriteCell, spcan *sped.SpriteCanvas) func (event *widgets.QGraphicsSceneMouseEvent){ return func (event *widgets.QGraphicsSceneMouseEvent) {
	
	sc.R, sc.G, sc.B = r,g,b
	var color = gui.NewQColor3(sc.R, sc.G, sc.B, 255) // r, g, b, a
	sc.Brush.SetColor(color)	
	
	spcan.Scene.Update(spcan.Scene.ItemsBoundingRect())
	
	statusbar.ShowMessage(strconv.Itoa(sc.Index), 0)
	sc.MousePressEventDefault(event)
	
	}
}

func ConnectFcpEvents(fcpcan *fcp.FcpCanvas) {
	
	it := fcpcan.Pallete.Cells.Iterator()
	for it.Next() {
		fc := it.Value().(*fcp.FcpCell)
		
		fc.ConnectMousePressEvent(ColorCellClick(fc)) // QGraphicsWidget
		
		}

	}

func ColorCellClick(fc *fcp.FcpCell) func (event *widgets.QGraphicsSceneMouseEvent){ return func (event *widgets.QGraphicsSceneMouseEvent) {
	
	r,g,b = fc.R, fc.G, fc.B
	statusbar.ShowMessage(strconv.Itoa(fc.Index), 0)
	fc.MousePressEventDefault(event)
	
	}
}
