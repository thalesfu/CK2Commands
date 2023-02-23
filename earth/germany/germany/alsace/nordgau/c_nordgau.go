package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NordgauCounty interface {
	feud.County
	BBrumath布吕马特() feud.Barony
	BEgisheim埃吉斯海姆() feud.Barony
	BErstein埃尔施泰因() feud.Barony
	BLauterburg劳特堡() feud.Barony
	BMolsheim莫尔塞姆() feud.Barony
	BSchlettstadt施莱特施塔特() feud.Barony
	BSelz塞尔茨() feud.Barony
	BStrassburg斯特拉斯堡() feud.Barony
}

type 诺德高NordgauCounty struct {
	feud.BaseCounty
	布吕马特Brumath        feud.Barony
	埃吉斯海姆Egisheim      feud.Barony
	埃尔施泰因Erstein       feud.Barony
	劳特堡Lauterburg      feud.Barony
	莫尔塞姆Molsheim       feud.Barony
	施莱特施塔特Schlettstadt feud.Barony
	塞尔茨Selz            feud.Barony
	斯特拉斯堡Strassburg    feud.Barony
}

func (c *诺德高NordgauCounty) BBrumath布吕马特() feud.Barony {
	return c.布吕马特Brumath
}

func (c *诺德高NordgauCounty) BEgisheim埃吉斯海姆() feud.Barony {
	return c.埃吉斯海姆Egisheim
}

func (c *诺德高NordgauCounty) BErstein埃尔施泰因() feud.Barony {
	return c.埃尔施泰因Erstein
}

func (c *诺德高NordgauCounty) BLauterburg劳特堡() feud.Barony {
	return c.劳特堡Lauterburg
}

func (c *诺德高NordgauCounty) BMolsheim莫尔塞姆() feud.Barony {
	return c.莫尔塞姆Molsheim
}

func (c *诺德高NordgauCounty) BSchlettstadt施莱特施塔特() feud.Barony {
	return c.施莱特施塔特Schlettstadt
}

func (c *诺德高NordgauCounty) BSelz塞尔茨() feud.Barony {
	return c.塞尔茨Selz
}

func (c *诺德高NordgauCounty) BStrassburg斯特拉斯堡() feud.Barony {
	return c.斯特拉斯堡Strassburg
}

var CNordgau诺德高 NordgauCounty = &诺德高NordgauCounty{}

func init() {
	f := CNordgau诺德高.(*诺德高NordgauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "126",
		Title:     "nordgau",
		TitleName: "诺德高",
		TitleCode: "c_nordgau",
		Baronies:  map[string]feud.Barony{},
	}

	f.布吕马特Brumath = BBrumath布吕马特
	f.布吕马特Brumath.SetParent(f)

	f.埃吉斯海姆Egisheim = BEgisheim埃吉斯海姆
	f.埃吉斯海姆Egisheim.SetParent(f)

	f.埃尔施泰因Erstein = BErstein埃尔施泰因
	f.埃尔施泰因Erstein.SetParent(f)

	f.劳特堡Lauterburg = BLauterburg劳特堡
	f.劳特堡Lauterburg.SetParent(f)

	f.莫尔塞姆Molsheim = BMolsheim莫尔塞姆
	f.莫尔塞姆Molsheim.SetParent(f)

	f.施莱特施塔特Schlettstadt = BSchlettstadt施莱特施塔特
	f.施莱特施塔特Schlettstadt.SetParent(f)

	f.塞尔茨Selz = BSelz塞尔茨
	f.塞尔茨Selz.SetParent(f)

	f.斯特拉斯堡Strassburg = BStrassburg斯特拉斯堡
	f.斯特拉斯堡Strassburg.SetParent(f)

}
