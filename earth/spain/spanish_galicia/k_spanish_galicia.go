package spanish_galicia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/spanish_galicia/galicia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Spanish_galiciaKingdom interface {
    feud.Kingdom
    DGalicia加利西亚() 	galicia.GaliciaDuke
}

type 加利西亚Spanish_galiciaKingdom struct {
	feud.BaseKingdom
	加利西亚Galicia 	galicia.GaliciaDuke
}

func (k *加利西亚Spanish_galiciaKingdom) DGalicia加利西亚() galicia.GaliciaDuke {
	return k.加利西亚Galicia
}
    
var KSpanish_galicia加利西亚 Spanish_galiciaKingdom = &加利西亚Spanish_galiciaKingdom{}

func init() {
	f := KSpanish_galicia加利西亚.(*加利西亚Spanish_galiciaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "spanish_galicia",
		TitleName: "加利西亚",
		TitleCode: "k_spanish_galicia",
		Dukes:  map[string]feud.Duke{},
	}

	f.加利西亚Galicia = galicia.DGalicia加利西亚
	f.加利西亚Galicia.SetParent(f)
	
}
