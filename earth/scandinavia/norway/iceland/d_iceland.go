package iceland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/iceland/austisland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/iceland/nordlendingafjordungur"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/iceland/sunnlendingafjordungur"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/iceland/vestisland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IcelandDuke interface {
	feud.Duke
	CAustisland东冰岛() austisland.AustislandCounty
	CNordlendingafjordungur北冰岛() nordlendingafjordungur.NordlendingafjordungurCounty
	CSunnlendingafjordungur南冰岛() sunnlendingafjordungur.SunnlendingafjordungurCounty
	CVestisland西冰岛() vestisland.VestislandCounty
}

type 冰岛IcelandDuke struct {
	feud.BaseDuke
	东冰岛Austisland             austisland.AustislandCounty
	北冰岛Nordlendingafjordungur nordlendingafjordungur.NordlendingafjordungurCounty
	南冰岛Sunnlendingafjordungur sunnlendingafjordungur.SunnlendingafjordungurCounty
	西冰岛Vestisland             vestisland.VestislandCounty
}

func (d *冰岛IcelandDuke) CAustisland东冰岛() austisland.AustislandCounty {
	return d.东冰岛Austisland
}

func (d *冰岛IcelandDuke) CNordlendingafjordungur北冰岛() nordlendingafjordungur.NordlendingafjordungurCounty {
	return d.北冰岛Nordlendingafjordungur
}

func (d *冰岛IcelandDuke) CSunnlendingafjordungur南冰岛() sunnlendingafjordungur.SunnlendingafjordungurCounty {
	return d.南冰岛Sunnlendingafjordungur
}

func (d *冰岛IcelandDuke) CVestisland西冰岛() vestisland.VestislandCounty {
	return d.西冰岛Vestisland
}

var DIceland冰岛 IcelandDuke = &冰岛IcelandDuke{}

func init() {
	f := DIceland冰岛.(*冰岛IcelandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "iceland",
		TitleName: "冰岛",
		TitleCode: "d_iceland",
		Counties:  map[string]feud.County{},
	}

	f.东冰岛Austisland = austisland.CAustisland东冰岛
	f.东冰岛Austisland.SetParent(f)

	f.北冰岛Nordlendingafjordungur = nordlendingafjordungur.CNordlendingafjordungur北冰岛
	f.北冰岛Nordlendingafjordungur.SetParent(f)

	f.南冰岛Sunnlendingafjordungur = sunnlendingafjordungur.CSunnlendingafjordungur南冰岛
	f.南冰岛Sunnlendingafjordungur.SetParent(f)

	f.西冰岛Vestisland = vestisland.CVestisland西冰岛
	f.西冰岛Vestisland.SetParent(f)

}
