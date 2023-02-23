package qamdo

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/qamdo/banbar"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/qamdo/bome"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/qamdo/qamdo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QamdoDuke interface {
	feud.Duke
	CBanbar边坝() banbar.BanbarCounty
	CBome波密() bome.BomeCounty
	CQamdo察木多() qamdo.QamdoCounty
}

type 察木多QamdoDuke struct {
	feud.BaseDuke
	边坝Banbar banbar.BanbarCounty
	波密Bome   bome.BomeCounty
	察木多Qamdo qamdo.QamdoCounty
}

func (d *察木多QamdoDuke) CBanbar边坝() banbar.BanbarCounty {
	return d.边坝Banbar
}

func (d *察木多QamdoDuke) CBome波密() bome.BomeCounty {
	return d.波密Bome
}

func (d *察木多QamdoDuke) CQamdo察木多() qamdo.QamdoCounty {
	return d.察木多Qamdo
}

var DQamdo察木多 QamdoDuke = &察木多QamdoDuke{}

func init() {
	f := DQamdo察木多.(*察木多QamdoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "qamdo",
		TitleName: "察木多",
		TitleCode: "d_qamdo",
		Counties:  map[string]feud.County{},
	}

	f.边坝Banbar = banbar.CBanbar边坝
	f.边坝Banbar.SetParent(f)

	f.波密Bome = bome.CBome波密
	f.波密Bome.SetParent(f)

	f.察木多Qamdo = qamdo.CQamdo察木多
	f.察木多Qamdo.SetParent(f)

}
