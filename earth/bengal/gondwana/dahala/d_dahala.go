package dahala

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/dahala/chauragarh"
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/dahala/tripuri"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DahalaDuke interface {
	feud.Duke
	CChauragarh招罗迦罗() chauragarh.ChauragarhCounty
	CTripuri帝利城() tripuri.TripuriCounty
}

type 陀诃罗DahalaDuke struct {
	feud.BaseDuke
	招罗迦罗Chauragarh chauragarh.ChauragarhCounty
	帝利城Tripuri     tripuri.TripuriCounty
}

func (d *陀诃罗DahalaDuke) CChauragarh招罗迦罗() chauragarh.ChauragarhCounty {
	return d.招罗迦罗Chauragarh
}

func (d *陀诃罗DahalaDuke) CTripuri帝利城() tripuri.TripuriCounty {
	return d.帝利城Tripuri
}

var DDahala陀诃罗 DahalaDuke = &陀诃罗DahalaDuke{}

func init() {
	f := DDahala陀诃罗.(*陀诃罗DahalaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dahala",
		TitleName: "陀诃罗",
		TitleCode: "d_dahala",
		Counties:  map[string]feud.County{},
	}

	f.招罗迦罗Chauragarh = chauragarh.CChauragarh招罗迦罗
	f.招罗迦罗Chauragarh.SetParent(f)

	f.帝利城Tripuri = tripuri.CTripuri帝利城
	f.帝利城Tripuri.SetParent(f)

}
