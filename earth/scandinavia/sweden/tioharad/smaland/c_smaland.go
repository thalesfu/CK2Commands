package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SmalandCounty interface {
	feud.County
	BAlvesta阿尔维斯塔() feud.Barony
	BAringsas阿灵索斯() feud.Barony
	BBergkvara贝里克瓦拉() feud.Barony
	BFuruby富鲁比() feud.Barony
	BKronoberg克鲁努贝里() feud.Barony
	BSkir希尔() feud.Barony
	BVaxjo韦克赫() feud.Barony
}

type 韦伦德SmalandCounty struct {
	feud.BaseCounty
	阿尔维斯塔Alvesta   feud.Barony
	阿灵索斯Aringsas   feud.Barony
	贝里克瓦拉Bergkvara feud.Barony
	富鲁比Furuby      feud.Barony
	克鲁努贝里Kronoberg feud.Barony
	希尔Skir         feud.Barony
	韦克赫Vaxjo       feud.Barony
}

func (c *韦伦德SmalandCounty) BAlvesta阿尔维斯塔() feud.Barony {
	return c.阿尔维斯塔Alvesta
}

func (c *韦伦德SmalandCounty) BAringsas阿灵索斯() feud.Barony {
	return c.阿灵索斯Aringsas
}

func (c *韦伦德SmalandCounty) BBergkvara贝里克瓦拉() feud.Barony {
	return c.贝里克瓦拉Bergkvara
}

func (c *韦伦德SmalandCounty) BFuruby富鲁比() feud.Barony {
	return c.富鲁比Furuby
}

func (c *韦伦德SmalandCounty) BKronoberg克鲁努贝里() feud.Barony {
	return c.克鲁努贝里Kronoberg
}

func (c *韦伦德SmalandCounty) BSkir希尔() feud.Barony {
	return c.希尔Skir
}

func (c *韦伦德SmalandCounty) BVaxjo韦克赫() feud.Barony {
	return c.韦克赫Vaxjo
}

var CSmaland韦伦德 SmalandCounty = &韦伦德SmalandCounty{}

func init() {
	f := CSmaland韦伦德.(*韦伦德SmalandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "298",
		Title:     "smaland",
		TitleName: "韦伦德",
		TitleCode: "c_smaland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔维斯塔Alvesta = BAlvesta阿尔维斯塔
	f.阿尔维斯塔Alvesta.SetParent(f)

	f.阿灵索斯Aringsas = BAringsas阿灵索斯
	f.阿灵索斯Aringsas.SetParent(f)

	f.贝里克瓦拉Bergkvara = BBergkvara贝里克瓦拉
	f.贝里克瓦拉Bergkvara.SetParent(f)

	f.富鲁比Furuby = BFuruby富鲁比
	f.富鲁比Furuby.SetParent(f)

	f.克鲁努贝里Kronoberg = BKronoberg克鲁努贝里
	f.克鲁努贝里Kronoberg.SetParent(f)

	f.希尔Skir = BSkir希尔
	f.希尔Skir.SetParent(f)

	f.韦克赫Vaxjo = BVaxjo韦克赫
	f.韦克赫Vaxjo.SetParent(f)

}
