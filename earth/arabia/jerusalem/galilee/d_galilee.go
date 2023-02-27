package galilee

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/galilee/beirut"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/galilee/safed"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/galilee/tiberias"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/galilee/tyrus"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalileeDuke interface {
    feud.Duke
    CBeirut贝鲁特() 	beirut.BeirutCounty
    CSafed采法特() 	safed.SafedCounty
    CTiberias太巴列() 	tiberias.TiberiasCounty
    CTyrus推罗() 	tyrus.TyrusCounty
}

type 加利利GalileeDuke struct {
	feud.BaseDuke
	贝鲁特Beirut 	beirut.BeirutCounty
	采法特Safed 	safed.SafedCounty
	太巴列Tiberias 	tiberias.TiberiasCounty
	推罗Tyrus 	tyrus.TyrusCounty
}

func (d *加利利GalileeDuke) CBeirut贝鲁特() beirut.BeirutCounty {
	return d.贝鲁特Beirut
}
    
func (d *加利利GalileeDuke) CSafed采法特() safed.SafedCounty {
	return d.采法特Safed
}
    
func (d *加利利GalileeDuke) CTiberias太巴列() tiberias.TiberiasCounty {
	return d.太巴列Tiberias
}
    
func (d *加利利GalileeDuke) CTyrus推罗() tyrus.TyrusCounty {
	return d.推罗Tyrus
}
    
var DGalilee加利利 GalileeDuke = &加利利GalileeDuke{}

func init() {
	f := DGalilee加利利.(*加利利GalileeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "galilee",
		TitleName: "加利利",
		TitleCode: "d_galilee",
		Counties:  map[string]feud.County{},
	}

	f.贝鲁特Beirut = beirut.CBeirut贝鲁特
	f.贝鲁特Beirut.SetParent(f)
	
	f.采法特Safed = safed.CSafed采法特
	f.采法特Safed.SetParent(f)
	
	f.太巴列Tiberias = tiberias.CTiberias太巴列
	f.太巴列Tiberias.SetParent(f)
	
	f.推罗Tyrus = tyrus.CTyrus推罗
	f.推罗Tyrus.SetParent(f)
	
}
