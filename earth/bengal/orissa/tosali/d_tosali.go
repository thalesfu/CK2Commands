package tosali

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/tosali/kataka"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/tosali/khijjingakota"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/tosali/viraja"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TosaliDuke interface {
	feud.Duke
	CKataka羯磔迦() kataka.KatakaCounty
	CKhijjingakota契承伽拘吒() khijjingakota.KhijjingakotaCounty
	CViraja毗罗阇() viraja.VirajaCounty
}

type 睹舍离TosaliDuke struct {
	feud.BaseDuke
	羯磔迦Kataka          kataka.KatakaCounty
	契承伽拘吒Khijjingakota khijjingakota.KhijjingakotaCounty
	毗罗阇Viraja          viraja.VirajaCounty
}

func (d *睹舍离TosaliDuke) CKataka羯磔迦() kataka.KatakaCounty {
	return d.羯磔迦Kataka
}

func (d *睹舍离TosaliDuke) CKhijjingakota契承伽拘吒() khijjingakota.KhijjingakotaCounty {
	return d.契承伽拘吒Khijjingakota
}

func (d *睹舍离TosaliDuke) CViraja毗罗阇() viraja.VirajaCounty {
	return d.毗罗阇Viraja
}

var DTosali睹舍离 TosaliDuke = &睹舍离TosaliDuke{}

func init() {
	f := DTosali睹舍离.(*睹舍离TosaliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tosali",
		TitleName: "睹舍离",
		TitleCode: "d_tosali",
		Counties:  map[string]feud.County{},
	}

	f.羯磔迦Kataka = kataka.CKataka羯磔迦
	f.羯磔迦Kataka.SetParent(f)

	f.契承伽拘吒Khijjingakota = khijjingakota.CKhijjingakota契承伽拘吒
	f.契承伽拘吒Khijjingakota.SetParent(f)

	f.毗罗阇Viraja = viraja.CViraja毗罗阇
	f.毗罗阇Viraja.SetParent(f)

}
