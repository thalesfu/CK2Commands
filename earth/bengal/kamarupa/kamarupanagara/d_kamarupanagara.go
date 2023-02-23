package kamarupanagara

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/kamarupanagara/kamarupanagara"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/kamarupanagara/kamatapur"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/kamarupanagara/sikkim"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KamarupanagaraDuke interface {
	feud.Duke
	CKamarupanagara迦摩缕波城() kamarupanagara.KamarupanagaraCounty
	CKamatapur迦摩多城() kamatapur.KamatapurCounty
	CSikkim锡金() sikkim.SikkimCounty
}

type 迦摩缕波城KamarupanagaraDuke struct {
	feud.BaseDuke
	迦摩缕波城Kamarupanagara kamarupanagara.KamarupanagaraCounty
	迦摩多城Kamatapur       kamatapur.KamatapurCounty
	锡金Sikkim            sikkim.SikkimCounty
}

func (d *迦摩缕波城KamarupanagaraDuke) CKamarupanagara迦摩缕波城() kamarupanagara.KamarupanagaraCounty {
	return d.迦摩缕波城Kamarupanagara
}

func (d *迦摩缕波城KamarupanagaraDuke) CKamatapur迦摩多城() kamatapur.KamatapurCounty {
	return d.迦摩多城Kamatapur
}

func (d *迦摩缕波城KamarupanagaraDuke) CSikkim锡金() sikkim.SikkimCounty {
	return d.锡金Sikkim
}

var DKamarupanagara迦摩缕波城 KamarupanagaraDuke = &迦摩缕波城KamarupanagaraDuke{}

func init() {
	f := DKamarupanagara迦摩缕波城.(*迦摩缕波城KamarupanagaraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kamarupanagara",
		TitleName: "迦摩缕波城",
		TitleCode: "d_kamarupanagara",
		Counties:  map[string]feud.County{},
	}

	f.迦摩缕波城Kamarupanagara = kamarupanagara.CKamarupanagara迦摩缕波城
	f.迦摩缕波城Kamarupanagara.SetParent(f)

	f.迦摩多城Kamatapur = kamatapur.CKamatapur迦摩多城
	f.迦摩多城Kamatapur.SetParent(f)

	f.锡金Sikkim = sikkim.CSikkim锡金
	f.锡金Sikkim.SetParent(f)

}
