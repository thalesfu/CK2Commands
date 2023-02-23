package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VladimirCounty interface {
	feud.County
	BBogolyubovo博戈柳博沃() feud.Barony
	BGorokhovets戈罗霍韦茨() feud.Barony
	BKovrov科夫罗夫() feud.Barony
	BNikologory尼科洛戈雷() feud.Barony
	BRochdestvenskoye罗日杰斯特文斯科耶() feud.Barony
	BViazniki维亚兹尼基() feud.Barony
	BVladimir弗拉基米尔() feud.Barony
}

type 弗拉基米尔VladimirCounty struct {
	feud.BaseCounty
	博戈柳博沃Bogolyubovo          feud.Barony
	戈罗霍韦茨Gorokhovets          feud.Barony
	科夫罗夫Kovrov                feud.Barony
	尼科洛戈雷Nikologory           feud.Barony
	罗日杰斯特文斯科耶Rochdestvenskoye feud.Barony
	维亚兹尼基Viazniki             feud.Barony
	弗拉基米尔Vladimir             feud.Barony
}

func (c *弗拉基米尔VladimirCounty) BBogolyubovo博戈柳博沃() feud.Barony {
	return c.博戈柳博沃Bogolyubovo
}

func (c *弗拉基米尔VladimirCounty) BGorokhovets戈罗霍韦茨() feud.Barony {
	return c.戈罗霍韦茨Gorokhovets
}

func (c *弗拉基米尔VladimirCounty) BKovrov科夫罗夫() feud.Barony {
	return c.科夫罗夫Kovrov
}

func (c *弗拉基米尔VladimirCounty) BNikologory尼科洛戈雷() feud.Barony {
	return c.尼科洛戈雷Nikologory
}

func (c *弗拉基米尔VladimirCounty) BRochdestvenskoye罗日杰斯特文斯科耶() feud.Barony {
	return c.罗日杰斯特文斯科耶Rochdestvenskoye
}

func (c *弗拉基米尔VladimirCounty) BViazniki维亚兹尼基() feud.Barony {
	return c.维亚兹尼基Viazniki
}

func (c *弗拉基米尔VladimirCounty) BVladimir弗拉基米尔() feud.Barony {
	return c.弗拉基米尔Vladimir
}

var CVladimir弗拉基米尔 VladimirCounty = &弗拉基米尔VladimirCounty{}

func init() {
	f := CVladimir弗拉基米尔.(*弗拉基米尔VladimirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "582",
		Title:     "vladimir",
		TitleName: "弗拉基米尔",
		TitleCode: "c_vladimir",
		Baronies:  map[string]feud.Barony{},
	}

	f.博戈柳博沃Bogolyubovo = BBogolyubovo博戈柳博沃
	f.博戈柳博沃Bogolyubovo.SetParent(f)

	f.戈罗霍韦茨Gorokhovets = BGorokhovets戈罗霍韦茨
	f.戈罗霍韦茨Gorokhovets.SetParent(f)

	f.科夫罗夫Kovrov = BKovrov科夫罗夫
	f.科夫罗夫Kovrov.SetParent(f)

	f.尼科洛戈雷Nikologory = BNikologory尼科洛戈雷
	f.尼科洛戈雷Nikologory.SetParent(f)

	f.罗日杰斯特文斯科耶Rochdestvenskoye = BRochdestvenskoye罗日杰斯特文斯科耶
	f.罗日杰斯特文斯科耶Rochdestvenskoye.SetParent(f)

	f.维亚兹尼基Viazniki = BViazniki维亚兹尼基
	f.维亚兹尼基Viazniki.SetParent(f)

	f.弗拉基米尔Vladimir = BVladimir弗拉基米尔
	f.弗拉基米尔Vladimir.SetParent(f)

}
