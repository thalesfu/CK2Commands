package socotra

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/socotra/socotra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SocotraDuke interface {
    feud.Duke
    CSocotra索科特拉() 	socotra.SocotraCounty
}

type 索科特拉SocotraDuke struct {
	feud.BaseDuke
	索科特拉Socotra 	socotra.SocotraCounty
}

func (d *索科特拉SocotraDuke) CSocotra索科特拉() socotra.SocotraCounty {
	return d.索科特拉Socotra
}
    
var DSocotra索科特拉 SocotraDuke = &索科特拉SocotraDuke{}

func init() {
	f := DSocotra索科特拉.(*索科特拉SocotraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "socotra",
		TitleName: "索科特拉",
		TitleCode: "d_socotra",
		Counties:  map[string]feud.County{},
	}

	f.索科特拉Socotra = socotra.CSocotra索科特拉
	f.索科特拉Socotra.SetParent(f)
	
}
