package baluchistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/sistan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaluchistanKingdom interface {
	feud.Kingdom
	DBaluchistan弥兰() baluchistan.BaluchistanDuke
	DSistan昔思丹() sistan.SistanDuke
}

type 昔思丹BaluchistanKingdom struct {
	feud.BaseKingdom
	弥兰Baluchistan baluchistan.BaluchistanDuke
	昔思丹Sistan     sistan.SistanDuke
}

func (k *昔思丹BaluchistanKingdom) DBaluchistan弥兰() baluchistan.BaluchistanDuke {
	return k.弥兰Baluchistan
}

func (k *昔思丹BaluchistanKingdom) DSistan昔思丹() sistan.SistanDuke {
	return k.昔思丹Sistan
}

var KBaluchistan昔思丹 BaluchistanKingdom = &昔思丹BaluchistanKingdom{}

func init() {
	f := KBaluchistan昔思丹.(*昔思丹BaluchistanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "baluchistan",
		TitleName: "昔思丹",
		TitleCode: "k_baluchistan",
		Dukes:     map[string]feud.Duke{},
	}

	f.弥兰Baluchistan = baluchistan.DBaluchistan弥兰
	f.弥兰Baluchistan.SetParent(f)

	f.昔思丹Sistan = sistan.DSistan昔思丹
	f.昔思丹Sistan.SetParent(f)

}
