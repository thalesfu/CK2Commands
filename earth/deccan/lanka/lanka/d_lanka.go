package lanka

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/lanka/rohana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LankaDuke interface {
	feud.Duke
	CRohana卢诃拿() rohana.RohanaCounty
}

type 楞迦LankaDuke struct {
	feud.BaseDuke
	卢诃拿Rohana rohana.RohanaCounty
}

func (d *楞迦LankaDuke) CRohana卢诃拿() rohana.RohanaCounty {
	return d.卢诃拿Rohana
}

var DLanka楞迦 LankaDuke = &楞迦LankaDuke{}

func init() {
	f := DLanka楞迦.(*楞迦LankaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lanka",
		TitleName: "楞迦",
		TitleCode: "d_lanka",
		Counties:  map[string]feud.County{},
	}

	f.卢诃拿Rohana = rohana.CRohana卢诃拿
	f.卢诃拿Rohana.SetParent(f)

}
