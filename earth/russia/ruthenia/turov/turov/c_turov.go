package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurovCounty interface {
	feud.County
	BAlsany奥利沙内() feud.Barony
	BFiadory费多雷() feud.Barony
	BSarny萨尔内() feud.Barony
	BStachava斯塔霍沃() feud.Barony
	BStolin斯托林() feud.Barony
	BTurov图罗夫() feud.Barony
}

type 图罗夫TurovCounty struct {
	feud.BaseCounty
	奥利沙内Alsany   feud.Barony
	费多雷Fiadory   feud.Barony
	萨尔内Sarny     feud.Barony
	斯塔霍沃Stachava feud.Barony
	斯托林Stolin    feud.Barony
	图罗夫Turov     feud.Barony
}

func (c *图罗夫TurovCounty) BAlsany奥利沙内() feud.Barony {
	return c.奥利沙内Alsany
}

func (c *图罗夫TurovCounty) BFiadory费多雷() feud.Barony {
	return c.费多雷Fiadory
}

func (c *图罗夫TurovCounty) BSarny萨尔内() feud.Barony {
	return c.萨尔内Sarny
}

func (c *图罗夫TurovCounty) BStachava斯塔霍沃() feud.Barony {
	return c.斯塔霍沃Stachava
}

func (c *图罗夫TurovCounty) BStolin斯托林() feud.Barony {
	return c.斯托林Stolin
}

func (c *图罗夫TurovCounty) BTurov图罗夫() feud.Barony {
	return c.图罗夫Turov
}

var CTurov图罗夫 TurovCounty = &图罗夫TurovCounty{}

func init() {
	f := CTurov图罗夫.(*图罗夫TurovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "552",
		Title:     "turov",
		TitleName: "图罗夫",
		TitleCode: "c_turov",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥利沙内Alsany = BAlsany奥利沙内
	f.奥利沙内Alsany.SetParent(f)

	f.费多雷Fiadory = BFiadory费多雷
	f.费多雷Fiadory.SetParent(f)

	f.萨尔内Sarny = BSarny萨尔内
	f.萨尔内Sarny.SetParent(f)

	f.斯塔霍沃Stachava = BStachava斯塔霍沃
	f.斯塔霍沃Stachava.SetParent(f)

	f.斯托林Stolin = BStolin斯托林
	f.斯托林Stolin.SetParent(f)

	f.图罗夫Turov = BTurov图罗夫
	f.图罗夫Turov.SetParent(f)

}
