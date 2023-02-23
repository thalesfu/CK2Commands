package cordoba

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/cordoba/calatrava"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/cordoba/cordoba"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/cordoba/jaen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CordobaDuke interface {
	feud.Duke
	CCalatrava卡拉特拉瓦() calatrava.CalatravaCounty
	CCordoba科尔多瓦() cordoba.CordobaCounty
	CJaen哈恩() jaen.JaenCounty
}

type 科尔多瓦CordobaDuke struct {
	feud.BaseDuke
	卡拉特拉瓦Calatrava calatrava.CalatravaCounty
	科尔多瓦Cordoba    cordoba.CordobaCounty
	哈恩Jaen         jaen.JaenCounty
}

func (d *科尔多瓦CordobaDuke) CCalatrava卡拉特拉瓦() calatrava.CalatravaCounty {
	return d.卡拉特拉瓦Calatrava
}

func (d *科尔多瓦CordobaDuke) CCordoba科尔多瓦() cordoba.CordobaCounty {
	return d.科尔多瓦Cordoba
}

func (d *科尔多瓦CordobaDuke) CJaen哈恩() jaen.JaenCounty {
	return d.哈恩Jaen
}

var DCordoba科尔多瓦 CordobaDuke = &科尔多瓦CordobaDuke{}

func init() {
	f := DCordoba科尔多瓦.(*科尔多瓦CordobaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cordoba",
		TitleName: "科尔多瓦",
		TitleCode: "d_cordoba",
		Counties:  map[string]feud.County{},
	}

	f.卡拉特拉瓦Calatrava = calatrava.CCalatrava卡拉特拉瓦
	f.卡拉特拉瓦Calatrava.SetParent(f)

	f.科尔多瓦Cordoba = cordoba.CCordoba科尔多瓦
	f.科尔多瓦Cordoba.SetParent(f)

	f.哈恩Jaen = jaen.CJaen哈恩
	f.哈恩Jaen.SetParent(f)

}
