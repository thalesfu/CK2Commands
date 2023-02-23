package burgundy

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurgundyKingdom interface {
	feud.Kingdom
	DDauphine多菲内() dauphine.DauphineDuke
	DProvence普罗旺斯() provence.ProvenceDuke
	DSavoie萨伏依() savoie.SavoieDuke
}

type 勃艮第BurgundyKingdom struct {
	feud.BaseKingdom
	多菲内Dauphine  dauphine.DauphineDuke
	普罗旺斯Provence provence.ProvenceDuke
	萨伏依Savoie    savoie.SavoieDuke
}

func (k *勃艮第BurgundyKingdom) DDauphine多菲内() dauphine.DauphineDuke {
	return k.多菲内Dauphine
}

func (k *勃艮第BurgundyKingdom) DProvence普罗旺斯() provence.ProvenceDuke {
	return k.普罗旺斯Provence
}

func (k *勃艮第BurgundyKingdom) DSavoie萨伏依() savoie.SavoieDuke {
	return k.萨伏依Savoie
}

var KBurgundy勃艮第 BurgundyKingdom = &勃艮第BurgundyKingdom{}

func init() {
	f := KBurgundy勃艮第.(*勃艮第BurgundyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "burgundy",
		TitleName: "勃艮第",
		TitleCode: "k_burgundy",
		Dukes:     map[string]feud.Duke{},
	}

	f.多菲内Dauphine = dauphine.DDauphine多菲内
	f.多菲内Dauphine.SetParent(f)

	f.普罗旺斯Provence = provence.DProvence普罗旺斯
	f.普罗旺斯Provence.SetParent(f)

	f.萨伏依Savoie = savoie.DSavoie萨伏依
	f.萨伏依Savoie.SetParent(f)

}
