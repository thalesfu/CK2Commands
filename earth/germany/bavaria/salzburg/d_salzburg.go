package salzburg

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/salzburg/pongau"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/salzburg/salzburg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SalzburgDuke interface {
	feud.Duke
	CPongau蓬高() pongau.PongauCounty
	CSalzburg萨尔茨堡() salzburg.SalzburgCounty
}

type 萨尔茨堡SalzburgDuke struct {
	feud.BaseDuke
	蓬高Pongau     pongau.PongauCounty
	萨尔茨堡Salzburg salzburg.SalzburgCounty
}

func (d *萨尔茨堡SalzburgDuke) CPongau蓬高() pongau.PongauCounty {
	return d.蓬高Pongau
}

func (d *萨尔茨堡SalzburgDuke) CSalzburg萨尔茨堡() salzburg.SalzburgCounty {
	return d.萨尔茨堡Salzburg
}

var DSalzburg萨尔茨堡 SalzburgDuke = &萨尔茨堡SalzburgDuke{}

func init() {
	f := DSalzburg萨尔茨堡.(*萨尔茨堡SalzburgDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "salzburg",
		TitleName: "萨尔茨堡",
		TitleCode: "d_salzburg",
		Counties:  map[string]feud.County{},
	}

	f.蓬高Pongau = pongau.CPongau蓬高
	f.蓬高Pongau.SetParent(f)

	f.萨尔茨堡Salzburg = salzburg.CSalzburg萨尔茨堡
	f.萨尔茨堡Salzburg.SetParent(f)

}
