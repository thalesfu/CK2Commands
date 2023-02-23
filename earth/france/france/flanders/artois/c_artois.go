package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArtoisCounty interface {
	feud.County
	BArras阿拉斯() feud.Barony
	BBapaume巴波姆() feud.Barony
	BBeaumetz博梅斯() feud.Barony
	BBethune贝蒂讷() feud.Barony
	BBouvines布汶() feud.Barony
	BLens朗斯() feud.Barony
	BLille里尔() feud.Barony
	BTerwaan泰鲁阿讷() feud.Barony
}

type 阿图瓦ArtoisCounty struct {
	feud.BaseCounty
	阿拉斯Arras    feud.Barony
	巴波姆Bapaume  feud.Barony
	博梅斯Beaumetz feud.Barony
	贝蒂讷Bethune  feud.Barony
	布汶Bouvines  feud.Barony
	朗斯Lens      feud.Barony
	里尔Lille     feud.Barony
	泰鲁阿讷Terwaan feud.Barony
}

func (c *阿图瓦ArtoisCounty) BArras阿拉斯() feud.Barony {
	return c.阿拉斯Arras
}

func (c *阿图瓦ArtoisCounty) BBapaume巴波姆() feud.Barony {
	return c.巴波姆Bapaume
}

func (c *阿图瓦ArtoisCounty) BBeaumetz博梅斯() feud.Barony {
	return c.博梅斯Beaumetz
}

func (c *阿图瓦ArtoisCounty) BBethune贝蒂讷() feud.Barony {
	return c.贝蒂讷Bethune
}

func (c *阿图瓦ArtoisCounty) BBouvines布汶() feud.Barony {
	return c.布汶Bouvines
}

func (c *阿图瓦ArtoisCounty) BLens朗斯() feud.Barony {
	return c.朗斯Lens
}

func (c *阿图瓦ArtoisCounty) BLille里尔() feud.Barony {
	return c.里尔Lille
}

func (c *阿图瓦ArtoisCounty) BTerwaan泰鲁阿讷() feud.Barony {
	return c.泰鲁阿讷Terwaan
}

var CArtois阿图瓦 ArtoisCounty = &阿图瓦ArtoisCounty{}

func init() {
	f := CArtois阿图瓦.(*阿图瓦ArtoisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "77",
		Title:     "artois",
		TitleName: "阿图瓦",
		TitleCode: "c_artois",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉斯Arras = BArras阿拉斯
	f.阿拉斯Arras.SetParent(f)

	f.巴波姆Bapaume = BBapaume巴波姆
	f.巴波姆Bapaume.SetParent(f)

	f.博梅斯Beaumetz = BBeaumetz博梅斯
	f.博梅斯Beaumetz.SetParent(f)

	f.贝蒂讷Bethune = BBethune贝蒂讷
	f.贝蒂讷Bethune.SetParent(f)

	f.布汶Bouvines = BBouvines布汶
	f.布汶Bouvines.SetParent(f)

	f.朗斯Lens = BLens朗斯
	f.朗斯Lens.SetParent(f)

	f.里尔Lille = BLille里尔
	f.里尔Lille.SetParent(f)

	f.泰鲁阿讷Terwaan = BTerwaan泰鲁阿讷
	f.泰鲁阿讷Terwaan.SetParent(f)

}
