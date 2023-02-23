package iraq

import (
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/baghdad"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/basra"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/samarra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IraqKingdom interface {
	feud.Kingdom
	DBaghdad巴格达() baghdad.BaghdadDuke
	DBasra巴士拉() basra.BasraDuke
	DSamarra萨迈拉() samarra.SamarraDuke
}

type 伊拉克IraqKingdom struct {
	feud.BaseKingdom
	巴格达Baghdad baghdad.BaghdadDuke
	巴士拉Basra   basra.BasraDuke
	萨迈拉Samarra samarra.SamarraDuke
}

func (k *伊拉克IraqKingdom) DBaghdad巴格达() baghdad.BaghdadDuke {
	return k.巴格达Baghdad
}

func (k *伊拉克IraqKingdom) DBasra巴士拉() basra.BasraDuke {
	return k.巴士拉Basra
}

func (k *伊拉克IraqKingdom) DSamarra萨迈拉() samarra.SamarraDuke {
	return k.萨迈拉Samarra
}

var KIraq伊拉克 IraqKingdom = &伊拉克IraqKingdom{}

func init() {
	f := KIraq伊拉克.(*伊拉克IraqKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "iraq",
		TitleName: "伊拉克",
		TitleCode: "k_iraq",
		Dukes:     map[string]feud.Duke{},
	}

	f.巴格达Baghdad = baghdad.DBaghdad巴格达
	f.巴格达Baghdad.SetParent(f)

	f.巴士拉Basra = basra.DBasra巴士拉
	f.巴士拉Basra.SetParent(f)

	f.萨迈拉Samarra = samarra.DSamarra萨迈拉
	f.萨迈拉Samarra.SetParent(f)

}
