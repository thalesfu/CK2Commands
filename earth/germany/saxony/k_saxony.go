package saxony

import (
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/angria"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/bremen"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/brunswick"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaxonyKingdom interface {
	feud.Kingdom
	DAngria盎格利亚() angria.AngriaDuke
	DBremen不来梅() bremen.BremenDuke
	DBrunswick布伦瑞克() brunswick.BrunswickDuke
	DSaxony萨克森() saxony.SaxonyDuke
}

type 萨克森SaxonyKingdom struct {
	feud.BaseKingdom
	盎格利亚Angria    angria.AngriaDuke
	不来梅Bremen     bremen.BremenDuke
	布伦瑞克Brunswick brunswick.BrunswickDuke
	萨克森Saxony     saxony.SaxonyDuke
}

func (k *萨克森SaxonyKingdom) DAngria盎格利亚() angria.AngriaDuke {
	return k.盎格利亚Angria
}

func (k *萨克森SaxonyKingdom) DBremen不来梅() bremen.BremenDuke {
	return k.不来梅Bremen
}

func (k *萨克森SaxonyKingdom) DBrunswick布伦瑞克() brunswick.BrunswickDuke {
	return k.布伦瑞克Brunswick
}

func (k *萨克森SaxonyKingdom) DSaxony萨克森() saxony.SaxonyDuke {
	return k.萨克森Saxony
}

var KSaxony萨克森 SaxonyKingdom = &萨克森SaxonyKingdom{}

func init() {
	f := KSaxony萨克森.(*萨克森SaxonyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "saxony",
		TitleName: "萨克森",
		TitleCode: "k_saxony",
		Dukes:     map[string]feud.Duke{},
	}

	f.盎格利亚Angria = angria.DAngria盎格利亚
	f.盎格利亚Angria.SetParent(f)

	f.不来梅Bremen = bremen.DBremen不来梅
	f.不来梅Bremen.SetParent(f)

	f.布伦瑞克Brunswick = brunswick.DBrunswick布伦瑞克
	f.布伦瑞克Brunswick.SetParent(f)

	f.萨克森Saxony = saxony.DSaxony萨克森
	f.萨克森Saxony.SetParent(f)

}
