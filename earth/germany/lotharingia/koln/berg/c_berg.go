package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BergCounty interface {
	feud.County
	BBarmen巴门() feud.Barony
	BBerg贝格() feud.Barony
	BDusseldorf杜塞尔多夫() feud.Barony
	BElberfeld埃尔伯费尔德() feud.Barony
	BThrotmanni特罗特曼尼() feud.Barony
	BWerth韦特() feud.Barony
}

type 贝格BergCounty struct {
	feud.BaseCounty
	巴门Barmen        feud.Barony
	贝格Berg          feud.Barony
	杜塞尔多夫Dusseldorf feud.Barony
	埃尔伯费尔德Elberfeld feud.Barony
	特罗特曼尼Throtmanni feud.Barony
	韦特Werth         feud.Barony
}

func (c *贝格BergCounty) BBarmen巴门() feud.Barony {
	return c.巴门Barmen
}

func (c *贝格BergCounty) BBerg贝格() feud.Barony {
	return c.贝格Berg
}

func (c *贝格BergCounty) BDusseldorf杜塞尔多夫() feud.Barony {
	return c.杜塞尔多夫Dusseldorf
}

func (c *贝格BergCounty) BElberfeld埃尔伯费尔德() feud.Barony {
	return c.埃尔伯费尔德Elberfeld
}

func (c *贝格BergCounty) BThrotmanni特罗特曼尼() feud.Barony {
	return c.特罗特曼尼Throtmanni
}

func (c *贝格BergCounty) BWerth韦特() feud.Barony {
	return c.韦特Werth
}

var CBerg贝格 BergCounty = &贝格BergCounty{}

func init() {
	f := CBerg贝格.(*贝格BergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1983",
		Title:     "berg",
		TitleName: "贝格",
		TitleCode: "c_berg",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴门Barmen = BBarmen巴门
	f.巴门Barmen.SetParent(f)

	f.贝格Berg = BBerg贝格
	f.贝格Berg.SetParent(f)

	f.杜塞尔多夫Dusseldorf = BDusseldorf杜塞尔多夫
	f.杜塞尔多夫Dusseldorf.SetParent(f)

	f.埃尔伯费尔德Elberfeld = BElberfeld埃尔伯费尔德
	f.埃尔伯费尔德Elberfeld.SetParent(f)

	f.特罗特曼尼Throtmanni = BThrotmanni特罗特曼尼
	f.特罗特曼尼Throtmanni.SetParent(f)

	f.韦特Werth = BWerth韦特
	f.韦特Werth.SetParent(f)

}
