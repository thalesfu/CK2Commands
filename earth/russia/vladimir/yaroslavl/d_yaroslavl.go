package yaroslavl

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/yaroslavl/chud"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/yaroslavl/kostroma"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/yaroslavl/vologda"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YaroslavlDuke interface {
	feud.Duke
	CChud楚德() chud.ChudCounty
	CKostroma科斯特罗马() kostroma.KostromaCounty
	CVologda沃洛格达() vologda.VologdaCounty
}

type 沃洛格达YaroslavlDuke struct {
	feud.BaseDuke
	楚德Chud        chud.ChudCounty
	科斯特罗马Kostroma kostroma.KostromaCounty
	沃洛格达Vologda   vologda.VologdaCounty
}

func (d *沃洛格达YaroslavlDuke) CChud楚德() chud.ChudCounty {
	return d.楚德Chud
}

func (d *沃洛格达YaroslavlDuke) CKostroma科斯特罗马() kostroma.KostromaCounty {
	return d.科斯特罗马Kostroma
}

func (d *沃洛格达YaroslavlDuke) CVologda沃洛格达() vologda.VologdaCounty {
	return d.沃洛格达Vologda
}

var DYaroslavl沃洛格达 YaroslavlDuke = &沃洛格达YaroslavlDuke{}

func init() {
	f := DYaroslavl沃洛格达.(*沃洛格达YaroslavlDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yaroslavl",
		TitleName: "沃洛格达",
		TitleCode: "d_yaroslavl",
		Counties:  map[string]feud.County{},
	}

	f.楚德Chud = chud.CChud楚德
	f.楚德Chud.SetParent(f)

	f.科斯特罗马Kostroma = kostroma.CKostroma科斯特罗马
	f.科斯特罗马Kostroma.SetParent(f)

	f.沃洛格达Vologda = vologda.CVologda沃洛格达
	f.沃洛格达Vologda.SetParent(f)

}
