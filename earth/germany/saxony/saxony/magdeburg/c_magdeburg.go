package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MagdeburgCounty interface {
	feud.County
	BArnstein阿恩施泰因() feud.Barony
	BBernburg贝恩堡() feud.Barony
	BHaldensleben哈尔登斯莱本() feud.Barony
	BMagdeburg马格德堡() feud.Barony
	BMansfeld曼斯费尔德() feud.Barony
	BQuerfurt奎尔富特() feud.Barony
}

type 马格德堡MagdeburgCounty struct {
	feud.BaseCounty
	阿恩施泰因Arnstein      feud.Barony
	贝恩堡Bernburg        feud.Barony
	哈尔登斯莱本Haldensleben feud.Barony
	马格德堡Magdeburg      feud.Barony
	曼斯费尔德Mansfeld      feud.Barony
	奎尔富特Querfurt       feud.Barony
}

func (c *马格德堡MagdeburgCounty) BArnstein阿恩施泰因() feud.Barony {
	return c.阿恩施泰因Arnstein
}

func (c *马格德堡MagdeburgCounty) BBernburg贝恩堡() feud.Barony {
	return c.贝恩堡Bernburg
}

func (c *马格德堡MagdeburgCounty) BHaldensleben哈尔登斯莱本() feud.Barony {
	return c.哈尔登斯莱本Haldensleben
}

func (c *马格德堡MagdeburgCounty) BMagdeburg马格德堡() feud.Barony {
	return c.马格德堡Magdeburg
}

func (c *马格德堡MagdeburgCounty) BMansfeld曼斯费尔德() feud.Barony {
	return c.曼斯费尔德Mansfeld
}

func (c *马格德堡MagdeburgCounty) BQuerfurt奎尔富特() feud.Barony {
	return c.奎尔富特Querfurt
}

var CMagdeburg马格德堡 MagdeburgCounty = &马格德堡MagdeburgCounty{}

func init() {
	f := CMagdeburg马格德堡.(*马格德堡MagdeburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1982",
		Title:     "magdeburg",
		TitleName: "马格德堡",
		TitleCode: "c_magdeburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿恩施泰因Arnstein = BArnstein阿恩施泰因
	f.阿恩施泰因Arnstein.SetParent(f)

	f.贝恩堡Bernburg = BBernburg贝恩堡
	f.贝恩堡Bernburg.SetParent(f)

	f.哈尔登斯莱本Haldensleben = BHaldensleben哈尔登斯莱本
	f.哈尔登斯莱本Haldensleben.SetParent(f)

	f.马格德堡Magdeburg = BMagdeburg马格德堡
	f.马格德堡Magdeburg.SetParent(f)

	f.曼斯费尔德Mansfeld = BMansfeld曼斯费尔德
	f.曼斯费尔德Mansfeld.SetParent(f)

	f.奎尔富特Querfurt = BQuerfurt奎尔富特
	f.奎尔富特Querfurt.SetParent(f)

}
