package corsica

import (
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/corsica/cinarca"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/corsica/corsica"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CorsicaDuke interface {
	feud.Duke
	CCinarca奇纳尔卡() cinarca.CinarcaCounty
	CCorsica科西嘉() corsica.CorsicaCounty
}

type 科西嘉CorsicaDuke struct {
	feud.BaseDuke
	奇纳尔卡Cinarca cinarca.CinarcaCounty
	科西嘉Corsica  corsica.CorsicaCounty
}

func (d *科西嘉CorsicaDuke) CCinarca奇纳尔卡() cinarca.CinarcaCounty {
	return d.奇纳尔卡Cinarca
}

func (d *科西嘉CorsicaDuke) CCorsica科西嘉() corsica.CorsicaCounty {
	return d.科西嘉Corsica
}

var DCorsica科西嘉 CorsicaDuke = &科西嘉CorsicaDuke{}

func init() {
	f := DCorsica科西嘉.(*科西嘉CorsicaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "corsica",
		TitleName: "科西嘉",
		TitleCode: "d_corsica",
		Counties:  map[string]feud.County{},
	}

	f.奇纳尔卡Cinarca = cinarca.CCinarca奇纳尔卡
	f.奇纳尔卡Cinarca.SetParent(f)

	f.科西嘉Corsica = corsica.CCorsica科西嘉
	f.科西嘉Corsica.SetParent(f)

}
