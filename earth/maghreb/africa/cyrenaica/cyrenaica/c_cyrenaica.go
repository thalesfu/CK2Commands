package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CyrenaicaCounty interface {
	feud.County
	BAlqubbah拱北() feud.Barony
	BAltamimi泰米米() feud.Barony
	BAthrun阿特伦() feud.Barony
	BCyrene昔兰尼() feud.Barony
	BDerna德尔纳() feud.Barony
	BKhadra卡哈达拉() feud.Barony
	BMarsasusah苏塞港() feud.Barony
	BMukhayta穆哈亚特() feud.Barony
}

type 昔兰尼加CyrenaicaCounty struct {
	feud.BaseCounty
	拱北Alqubbah    feud.Barony
	泰米米Altamimi   feud.Barony
	阿特伦Athrun     feud.Barony
	昔兰尼Cyrene     feud.Barony
	德尔纳Derna      feud.Barony
	卡哈达拉Khadra    feud.Barony
	苏塞港Marsasusah feud.Barony
	穆哈亚特Mukhayta  feud.Barony
}

func (c *昔兰尼加CyrenaicaCounty) BAlqubbah拱北() feud.Barony {
	return c.拱北Alqubbah
}

func (c *昔兰尼加CyrenaicaCounty) BAltamimi泰米米() feud.Barony {
	return c.泰米米Altamimi
}

func (c *昔兰尼加CyrenaicaCounty) BAthrun阿特伦() feud.Barony {
	return c.阿特伦Athrun
}

func (c *昔兰尼加CyrenaicaCounty) BCyrene昔兰尼() feud.Barony {
	return c.昔兰尼Cyrene
}

func (c *昔兰尼加CyrenaicaCounty) BDerna德尔纳() feud.Barony {
	return c.德尔纳Derna
}

func (c *昔兰尼加CyrenaicaCounty) BKhadra卡哈达拉() feud.Barony {
	return c.卡哈达拉Khadra
}

func (c *昔兰尼加CyrenaicaCounty) BMarsasusah苏塞港() feud.Barony {
	return c.苏塞港Marsasusah
}

func (c *昔兰尼加CyrenaicaCounty) BMukhayta穆哈亚特() feud.Barony {
	return c.穆哈亚特Mukhayta
}

var CCyrenaica昔兰尼加 CyrenaicaCounty = &昔兰尼加CyrenaicaCounty{}

func init() {
	f := CCyrenaica昔兰尼加.(*昔兰尼加CyrenaicaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "806",
		Title:     "cyrenaica",
		TitleName: "昔兰尼加",
		TitleCode: "c_cyrenaica",
		Baronies:  map[string]feud.Barony{},
	}

	f.拱北Alqubbah = BAlqubbah拱北
	f.拱北Alqubbah.SetParent(f)

	f.泰米米Altamimi = BAltamimi泰米米
	f.泰米米Altamimi.SetParent(f)

	f.阿特伦Athrun = BAthrun阿特伦
	f.阿特伦Athrun.SetParent(f)

	f.昔兰尼Cyrene = BCyrene昔兰尼
	f.昔兰尼Cyrene.SetParent(f)

	f.德尔纳Derna = BDerna德尔纳
	f.德尔纳Derna.SetParent(f)

	f.卡哈达拉Khadra = BKhadra卡哈达拉
	f.卡哈达拉Khadra.SetParent(f)

	f.苏塞港Marsasusah = BMarsasusah苏塞港
	f.苏塞港Marsasusah.SetParent(f)

	f.穆哈亚特Mukhayta = BMukhayta穆哈亚特
	f.穆哈亚特Mukhayta.SetParent(f)

}
