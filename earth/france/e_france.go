package france

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine"
	"github.com/thalesfu/CK2Commands/earth/france/brittany"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy"
	"github.com/thalesfu/CK2Commands/earth/france/france"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FranceEmpire interface {
	feud.Empire
	KAquitaine阿基坦() aquitaine.AquitaineKingdom
	KBrittany布列塔尼() brittany.BrittanyKingdom
	KBurgundy勃艮第() burgundy.BurgundyKingdom
	KFrance法兰西() france.FranceKingdom
}

type 法兰克FranceEmpire struct {
	feud.BaseEmpire
	阿基坦Aquitaine aquitaine.AquitaineKingdom
	布列塔尼Brittany brittany.BrittanyKingdom
	勃艮第Burgundy  burgundy.BurgundyKingdom
	法兰西France    france.FranceKingdom
}

func (e *法兰克FranceEmpire) KAquitaine阿基坦() aquitaine.AquitaineKingdom {
	return e.阿基坦Aquitaine
}

func (e *法兰克FranceEmpire) KBrittany布列塔尼() brittany.BrittanyKingdom {
	return e.布列塔尼Brittany
}

func (e *法兰克FranceEmpire) KBurgundy勃艮第() burgundy.BurgundyKingdom {
	return e.勃艮第Burgundy
}

func (e *法兰克FranceEmpire) KFrance法兰西() france.FranceKingdom {
	return e.法兰西France
}

var EFrance法兰克 FranceEmpire = &法兰克FranceEmpire{}

func init() {
	f := EFrance法兰克.(*法兰克FranceEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "france",
		TitleName: "法兰克",
		TitleCode: "e_france",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿基坦Aquitaine = aquitaine.KAquitaine阿基坦
	f.阿基坦Aquitaine.SetParent(f)
	f.布列塔尼Brittany = brittany.KBrittany布列塔尼
	f.布列塔尼Brittany.SetParent(f)
	f.勃艮第Burgundy = burgundy.KBurgundy勃艮第
	f.勃艮第Burgundy.SetParent(f)
	f.法兰西France = france.KFrance法兰西
	f.法兰西France.SetParent(f)
}
