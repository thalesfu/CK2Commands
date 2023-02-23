package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WurzburgCounty interface {
	feud.County
	BHammelburg哈默尔堡() feud.Barony
	BMainderheim明德尔海姆() feud.Barony
	BMarienberg马林贝格() feud.Barony
	BSchwarzenberg施瓦岑贝格() feud.Barony
	BSchweinfurt施韦因富特() feud.Barony
	BTheiheim泰尔海姆() feud.Barony
	BWurzburg维尔茨堡() feud.Barony
}

type 维尔茨堡WurzburgCounty struct {
	feud.BaseCounty
	哈默尔堡Hammelburg     feud.Barony
	明德尔海姆Mainderheim   feud.Barony
	马林贝格Marienberg     feud.Barony
	施瓦岑贝格Schwarzenberg feud.Barony
	施韦因富特Schweinfurt   feud.Barony
	泰尔海姆Theiheim       feud.Barony
	维尔茨堡Wurzburg       feud.Barony
}

func (c *维尔茨堡WurzburgCounty) BHammelburg哈默尔堡() feud.Barony {
	return c.哈默尔堡Hammelburg
}

func (c *维尔茨堡WurzburgCounty) BMainderheim明德尔海姆() feud.Barony {
	return c.明德尔海姆Mainderheim
}

func (c *维尔茨堡WurzburgCounty) BMarienberg马林贝格() feud.Barony {
	return c.马林贝格Marienberg
}

func (c *维尔茨堡WurzburgCounty) BSchwarzenberg施瓦岑贝格() feud.Barony {
	return c.施瓦岑贝格Schwarzenberg
}

func (c *维尔茨堡WurzburgCounty) BSchweinfurt施韦因富特() feud.Barony {
	return c.施韦因富特Schweinfurt
}

func (c *维尔茨堡WurzburgCounty) BTheiheim泰尔海姆() feud.Barony {
	return c.泰尔海姆Theiheim
}

func (c *维尔茨堡WurzburgCounty) BWurzburg维尔茨堡() feud.Barony {
	return c.维尔茨堡Wurzburg
}

var CWurzburg维尔茨堡 WurzburgCounty = &维尔茨堡WurzburgCounty{}

func init() {
	f := CWurzburg维尔茨堡.(*维尔茨堡WurzburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "254",
		Title:     "wurzburg",
		TitleName: "维尔茨堡",
		TitleCode: "c_wurzburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈默尔堡Hammelburg = BHammelburg哈默尔堡
	f.哈默尔堡Hammelburg.SetParent(f)

	f.明德尔海姆Mainderheim = BMainderheim明德尔海姆
	f.明德尔海姆Mainderheim.SetParent(f)

	f.马林贝格Marienberg = BMarienberg马林贝格
	f.马林贝格Marienberg.SetParent(f)

	f.施瓦岑贝格Schwarzenberg = BSchwarzenberg施瓦岑贝格
	f.施瓦岑贝格Schwarzenberg.SetParent(f)

	f.施韦因富特Schweinfurt = BSchweinfurt施韦因富特
	f.施韦因富特Schweinfurt.SetParent(f)

	f.泰尔海姆Theiheim = BTheiheim泰尔海姆
	f.泰尔海姆Theiheim.SetParent(f)

	f.维尔茨堡Wurzburg = BWurzburg维尔茨堡
	f.维尔茨堡Wurzburg.SetParent(f)

}
