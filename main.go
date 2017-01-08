package main

import (
	"os"
	"strconv"
	
	"local/spc/fcp"
	"local/spc/sped"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
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
	
	//ConnectSpriteEvents(spcan)

	// full color pallete panel
	var dfcp = widgets.NewQDockWidget("Full Color Pallete", window, 0)
	window.AddDockWidget(core.Qt__RightDockWidgetArea, dfcp)
	var palette = fcp.NewFcpCanvas() // window
	dfcp.SetWidget(palette.View)
	
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
		
		sc.ConnectMousePressEvent(SpriteCellClick(sc)) // QGraphicsWidget
		
		}

	}


func SpriteCellClick(sc *sped.SpriteCell) func (event *widgets.QGraphicsSceneMouseEvent){ return func (event *widgets.QGraphicsSceneMouseEvent) {
	
	statusbar.ShowMessage(strconv.Itoa(sc.Index), 0)
	sc.MousePressEventDefault(event)
	
	}
}

	
