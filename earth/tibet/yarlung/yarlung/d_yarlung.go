package yarlung

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/yarlung/lhunze"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/yarlung/nedong"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/yarlung/taktse"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YarlungDuke interface {
	feud.Duke
	CLhunze隆子() lhunze.LhunzeCounty
	CNedong乃东() nedong.NedongCounty
	CTaktse达孜() taktse.TaktseCounty
}

type 雅砻YarlungDuke struct {
	feud.BaseDuke
	隆子Lhunze lhunze.LhunzeCounty
	乃东Nedong nedong.NedongCounty
	达孜Taktse taktse.TaktseCounty
}

func (d *雅砻YarlungDuke) CLhunze隆子() lhunze.LhunzeCounty {
	return d.隆子Lhunze
}

func (d *雅砻YarlungDuke) CNedong乃东() nedong.NedongCounty {
	return d.乃东Nedong
}

func (d *雅砻YarlungDuke) CTaktse达孜() taktse.TaktseCounty {
	return d.达孜Taktse
}

var DYarlung雅砻 YarlungDuke = &雅砻YarlungDuke{}

func init() {
	f := DYarlung雅砻.(*雅砻YarlungDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yarlung",
		TitleName: "雅砻",
		TitleCode: "d_yarlung",
		Counties:  map[string]feud.County{},
	}

	f.隆子Lhunze = lhunze.CLhunze隆子
	f.隆子Lhunze.SetParent(f)

	f.乃东Nedong = nedong.CNedong乃东
	f.乃东Nedong.SetParent(f)

	f.达孜Taktse = taktse.CTaktse达孜
	f.达孜Taktse.SetParent(f)

}
