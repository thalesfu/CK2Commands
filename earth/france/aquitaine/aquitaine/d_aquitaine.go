package aquitaine

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/aquitaine/agen"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/aquitaine/angouleme"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/aquitaine/bordeaux"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/aquitaine/perigord"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AquitaineDuke interface {
	feud.Duke
	CAgen阿让() agen.AgenCounty
	CAngouleme昂古莱姆() angouleme.AngoulemeCounty
	CBordeaux波尔多() bordeaux.BordeauxCounty
	CPerigord佩里戈尔() perigord.PerigordCounty
}

type 阿基坦AquitaineDuke struct {
	feud.BaseDuke
	阿让Agen        agen.AgenCounty
	昂古莱姆Angouleme angouleme.AngoulemeCounty
	波尔多Bordeaux   bordeaux.BordeauxCounty
	佩里戈尔Perigord  perigord.PerigordCounty
}

func (d *阿基坦AquitaineDuke) CAgen阿让() agen.AgenCounty {
	return d.阿让Agen
}

func (d *阿基坦AquitaineDuke) CAngouleme昂古莱姆() angouleme.AngoulemeCounty {
	return d.昂古莱姆Angouleme
}

func (d *阿基坦AquitaineDuke) CBordeaux波尔多() bordeaux.BordeauxCounty {
	return d.波尔多Bordeaux
}

func (d *阿基坦AquitaineDuke) CPerigord佩里戈尔() perigord.PerigordCounty {
	return d.佩里戈尔Perigord
}

var DAquitaine阿基坦 AquitaineDuke = &阿基坦AquitaineDuke{}

func init() {
	f := DAquitaine阿基坦.(*阿基坦AquitaineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aquitaine",
		TitleName: "阿基坦",
		TitleCode: "d_aquitaine",
		Counties:  map[string]feud.County{},
	}

	f.阿让Agen = agen.CAgen阿让
	f.阿让Agen.SetParent(f)

	f.昂古莱姆Angouleme = angouleme.CAngouleme昂古莱姆
	f.昂古莱姆Angouleme.SetParent(f)

	f.波尔多Bordeaux = bordeaux.CBordeaux波尔多
	f.波尔多Bordeaux.SetParent(f)

	f.佩里戈尔Perigord = perigord.CPerigord佩里戈尔
	f.佩里戈尔Perigord.SetParent(f)

}
