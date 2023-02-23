package punjab

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/gandhara"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/lahore"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/trigarta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PunjabKingdom interface {
	feud.Kingdom
	DGandhara犍陀罗() gandhara.GandharaDuke
	DLahore拉合尔() lahore.LahoreDuke
	DMultan木尔坦() multan.MultanDuke
	DTrigarta三穴() trigarta.TrigartaDuke
}

type 五河PunjabKingdom struct {
	feud.BaseKingdom
	犍陀罗Gandhara gandhara.GandharaDuke
	拉合尔Lahore   lahore.LahoreDuke
	木尔坦Multan   multan.MultanDuke
	三穴Trigarta  trigarta.TrigartaDuke
}

func (k *五河PunjabKingdom) DGandhara犍陀罗() gandhara.GandharaDuke {
	return k.犍陀罗Gandhara
}

func (k *五河PunjabKingdom) DLahore拉合尔() lahore.LahoreDuke {
	return k.拉合尔Lahore
}

func (k *五河PunjabKingdom) DMultan木尔坦() multan.MultanDuke {
	return k.木尔坦Multan
}

func (k *五河PunjabKingdom) DTrigarta三穴() trigarta.TrigartaDuke {
	return k.三穴Trigarta
}

var KPunjab五河 PunjabKingdom = &五河PunjabKingdom{}

func init() {
	f := KPunjab五河.(*五河PunjabKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "punjab",
		TitleName: "五河",
		TitleCode: "k_punjab",
		Dukes:     map[string]feud.Duke{},
	}

	f.犍陀罗Gandhara = gandhara.DGandhara犍陀罗
	f.犍陀罗Gandhara.SetParent(f)

	f.拉合尔Lahore = lahore.DLahore拉合尔
	f.拉合尔Lahore.SetParent(f)

	f.木尔坦Multan = multan.DMultan木尔坦
	f.木尔坦Multan.SetParent(f)

	f.三穴Trigarta = trigarta.DTrigarta三穴
	f.三穴Trigarta.SetParent(f)

}
