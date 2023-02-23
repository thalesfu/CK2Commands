package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedakCounty interface {
	feud.County
	BAnanthagiri阿难多耆厘() feud.Barony
	BBaramgala婆兰迦罗() feud.Barony
	BChikalthana支迦罗傥那() feud.Barony
	BDomakonda度摩军荼() feud.Barony
	BKohir科赫尔() feud.Barony
	BMardi玛地() feud.Barony
	BMedak弥陀迦() feud.Barony
}

type 弥陀迦MedakCounty struct {
	feud.BaseCounty
	阿难多耆厘Ananthagiri feud.Barony
	婆兰迦罗Baramgala    feud.Barony
	支迦罗傥那Chikalthana feud.Barony
	度摩军荼Domakonda    feud.Barony
	科赫尔Kohir         feud.Barony
	玛地Mardi          feud.Barony
	弥陀迦Medak         feud.Barony
}

func (c *弥陀迦MedakCounty) BAnanthagiri阿难多耆厘() feud.Barony {
	return c.阿难多耆厘Ananthagiri
}

func (c *弥陀迦MedakCounty) BBaramgala婆兰迦罗() feud.Barony {
	return c.婆兰迦罗Baramgala
}

func (c *弥陀迦MedakCounty) BChikalthana支迦罗傥那() feud.Barony {
	return c.支迦罗傥那Chikalthana
}

func (c *弥陀迦MedakCounty) BDomakonda度摩军荼() feud.Barony {
	return c.度摩军荼Domakonda
}

func (c *弥陀迦MedakCounty) BKohir科赫尔() feud.Barony {
	return c.科赫尔Kohir
}

func (c *弥陀迦MedakCounty) BMardi玛地() feud.Barony {
	return c.玛地Mardi
}

func (c *弥陀迦MedakCounty) BMedak弥陀迦() feud.Barony {
	return c.弥陀迦Medak
}

var CMedak弥陀迦 MedakCounty = &弥陀迦MedakCounty{}

func init() {
	f := CMedak弥陀迦.(*弥陀迦MedakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1257",
		Title:     "medak",
		TitleName: "弥陀迦",
		TitleCode: "c_medak",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿难多耆厘Ananthagiri = BAnanthagiri阿难多耆厘
	f.阿难多耆厘Ananthagiri.SetParent(f)

	f.婆兰迦罗Baramgala = BBaramgala婆兰迦罗
	f.婆兰迦罗Baramgala.SetParent(f)

	f.支迦罗傥那Chikalthana = BChikalthana支迦罗傥那
	f.支迦罗傥那Chikalthana.SetParent(f)

	f.度摩军荼Domakonda = BDomakonda度摩军荼
	f.度摩军荼Domakonda.SetParent(f)

	f.科赫尔Kohir = BKohir科赫尔
	f.科赫尔Kohir.SetParent(f)

	f.玛地Mardi = BMardi玛地
	f.玛地Mardi.SetParent(f)

	f.弥陀迦Medak = BMedak弥陀迦
	f.弥陀迦Medak.SetParent(f)

}
