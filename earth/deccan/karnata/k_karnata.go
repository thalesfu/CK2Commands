package karnata

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/gangavadi"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/kalyani"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/nulambavadi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarnataKingdom interface {
	feud.Kingdom
	DGangavadi恒伽婆稚() gangavadi.GangavadiDuke
	DKalyani迦梨耶尼() kalyani.KalyaniDuke
	DNulambavadi奴蓝婆婆稚() nulambavadi.NulambavadiDuke
}

type 羯罗拏吒KarnataKingdom struct {
	feud.BaseKingdom
	恒伽婆稚Gangavadi    gangavadi.GangavadiDuke
	迦梨耶尼Kalyani      kalyani.KalyaniDuke
	奴蓝婆婆稚Nulambavadi nulambavadi.NulambavadiDuke
}

func (k *羯罗拏吒KarnataKingdom) DGangavadi恒伽婆稚() gangavadi.GangavadiDuke {
	return k.恒伽婆稚Gangavadi
}

func (k *羯罗拏吒KarnataKingdom) DKalyani迦梨耶尼() kalyani.KalyaniDuke {
	return k.迦梨耶尼Kalyani
}

func (k *羯罗拏吒KarnataKingdom) DNulambavadi奴蓝婆婆稚() nulambavadi.NulambavadiDuke {
	return k.奴蓝婆婆稚Nulambavadi
}

var KKarnata羯罗拏吒 KarnataKingdom = &羯罗拏吒KarnataKingdom{}

func init() {
	f := KKarnata羯罗拏吒.(*羯罗拏吒KarnataKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "karnata",
		TitleName: "羯罗拏吒",
		TitleCode: "k_karnata",
		Dukes:     map[string]feud.Duke{},
	}

	f.恒伽婆稚Gangavadi = gangavadi.DGangavadi恒伽婆稚
	f.恒伽婆稚Gangavadi.SetParent(f)

	f.迦梨耶尼Kalyani = kalyani.DKalyani迦梨耶尼
	f.迦梨耶尼Kalyani.SetParent(f)

	f.奴蓝婆婆稚Nulambavadi = nulambavadi.DNulambavadi奴蓝婆婆稚
	f.奴蓝婆婆稚Nulambavadi.SetParent(f)

}
