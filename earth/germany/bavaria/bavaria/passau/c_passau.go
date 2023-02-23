package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PassauCounty interface {
	feud.County
	BFormbach福尔姆巴赫() feud.Barony
	BFreyung弗赖翁() feud.Barony
	BOrtenburg奥尔滕堡() feud.Barony
	BPassau帕绍() feud.Barony
	BRaab拉布() feud.Barony
	BSchaumberg绍姆堡() feud.Barony
	BUlrichsberg乌尔里希斯贝格() feud.Barony
}

type 帕绍PassauCounty struct {
	feud.BaseCounty
	福尔姆巴赫Formbach      feud.Barony
	弗赖翁Freyung         feud.Barony
	奥尔滕堡Ortenburg      feud.Barony
	帕绍Passau           feud.Barony
	拉布Raab             feud.Barony
	绍姆堡Schaumberg      feud.Barony
	乌尔里希斯贝格Ulrichsberg feud.Barony
}

func (c *帕绍PassauCounty) BFormbach福尔姆巴赫() feud.Barony {
	return c.福尔姆巴赫Formbach
}

func (c *帕绍PassauCounty) BFreyung弗赖翁() feud.Barony {
	return c.弗赖翁Freyung
}

func (c *帕绍PassauCounty) BOrtenburg奥尔滕堡() feud.Barony {
	return c.奥尔滕堡Ortenburg
}

func (c *帕绍PassauCounty) BPassau帕绍() feud.Barony {
	return c.帕绍Passau
}

func (c *帕绍PassauCounty) BRaab拉布() feud.Barony {
	return c.拉布Raab
}

func (c *帕绍PassauCounty) BSchaumberg绍姆堡() feud.Barony {
	return c.绍姆堡Schaumberg
}

func (c *帕绍PassauCounty) BUlrichsberg乌尔里希斯贝格() feud.Barony {
	return c.乌尔里希斯贝格Ulrichsberg
}

var CPassau帕绍 PassauCounty = &帕绍PassauCounty{}

func init() {
	f := CPassau帕绍.(*帕绍PassauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "447",
		Title:     "passau",
		TitleName: "帕绍",
		TitleCode: "c_passau",
		Baronies:  map[string]feud.Barony{},
	}

	f.福尔姆巴赫Formbach = BFormbach福尔姆巴赫
	f.福尔姆巴赫Formbach.SetParent(f)

	f.弗赖翁Freyung = BFreyung弗赖翁
	f.弗赖翁Freyung.SetParent(f)

	f.奥尔滕堡Ortenburg = BOrtenburg奥尔滕堡
	f.奥尔滕堡Ortenburg.SetParent(f)

	f.帕绍Passau = BPassau帕绍
	f.帕绍Passau.SetParent(f)

	f.拉布Raab = BRaab拉布
	f.拉布Raab.SetParent(f)

	f.绍姆堡Schaumberg = BSchaumberg绍姆堡
	f.绍姆堡Schaumberg.SetParent(f)

	f.乌尔里希斯贝格Ulrichsberg = BUlrichsberg乌尔里希斯贝格
	f.乌尔里希斯贝格Ulrichsberg.SetParent(f)

}
