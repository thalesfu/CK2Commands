package loire

import (
	"github.com/thalesfu/CK2Commands/earth/france/brittany/loire/nantes"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/loire/rennes"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/loire/retz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LoireDuke interface {
	feud.Duke
	CNantes南特() nantes.NantesCounty
	CRennes雷恩() rennes.RennesCounty
	CRetz雷茨() retz.RetzCounty
}

type 上布列塔尼LoireDuke struct {
	feud.BaseDuke
	南特Nantes nantes.NantesCounty
	雷恩Rennes rennes.RennesCounty
	雷茨Retz   retz.RetzCounty
}

func (d *上布列塔尼LoireDuke) CNantes南特() nantes.NantesCounty {
	return d.南特Nantes
}

func (d *上布列塔尼LoireDuke) CRennes雷恩() rennes.RennesCounty {
	return d.雷恩Rennes
}

func (d *上布列塔尼LoireDuke) CRetz雷茨() retz.RetzCounty {
	return d.雷茨Retz
}

var DLoire上布列塔尼 LoireDuke = &上布列塔尼LoireDuke{}

func init() {
	f := DLoire上布列塔尼.(*上布列塔尼LoireDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "loire",
		TitleName: "上布列塔尼",
		TitleCode: "d_loire",
		Counties:  map[string]feud.County{},
	}

	f.南特Nantes = nantes.CNantes南特
	f.南特Nantes.SetParent(f)

	f.雷恩Rennes = rennes.CRennes雷恩
	f.雷恩Rennes.SetParent(f)

	f.雷茨Retz = retz.CRetz雷茨
	f.雷茨Retz.SetParent(f)

}
