package portugal

import (
	"github.com/thalesfu/CK2Commands/earth/spain/portugal/porto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PortugalKingdom interface {
	feud.Kingdom
	DPorto波图斯卡莱() porto.PortoDuke
}

type 葡萄牙PortugalKingdom struct {
	feud.BaseKingdom
	波图斯卡莱Porto porto.PortoDuke
}

func (k *葡萄牙PortugalKingdom) DPorto波图斯卡莱() porto.PortoDuke {
	return k.波图斯卡莱Porto
}

var KPortugal葡萄牙 PortugalKingdom = &葡萄牙PortugalKingdom{}

func init() {
	f := KPortugal葡萄牙.(*葡萄牙PortugalKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "portugal",
		TitleName: "葡萄牙",
		TitleCode: "k_portugal",
		Dukes:     map[string]feud.Duke{},
	}

	f.波图斯卡莱Porto = porto.DPorto波图斯卡莱
	f.波图斯卡莱Porto.SetParent(f)

}
