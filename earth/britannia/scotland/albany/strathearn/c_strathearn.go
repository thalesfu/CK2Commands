package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StrathearnCounty interface {
	feud.County
	BAuchterarder奥赫特拉德() feud.Barony
	BCrieff克里夫() feud.Barony
	BDoune杜恩() feud.Barony
	BDunblane邓布兰() feud.Barony
	BInchaffray因彻法莱() feud.Barony
	BKenmore肯莫尔() feud.Barony
	BMadderty马德第() feud.Barony
	BTullibardine图里巴丁() feud.Barony
}

type 斯特拉森StrathearnCounty struct {
	feud.BaseCounty
	奥赫特拉德Auchterarder feud.Barony
	克里夫Crieff         feud.Barony
	杜恩Doune           feud.Barony
	邓布兰Dunblane       feud.Barony
	因彻法莱Inchaffray    feud.Barony
	肯莫尔Kenmore        feud.Barony
	马德第Madderty       feud.Barony
	图里巴丁Tullibardine  feud.Barony
}

func (c *斯特拉森StrathearnCounty) BAuchterarder奥赫特拉德() feud.Barony {
	return c.奥赫特拉德Auchterarder
}

func (c *斯特拉森StrathearnCounty) BCrieff克里夫() feud.Barony {
	return c.克里夫Crieff
}

func (c *斯特拉森StrathearnCounty) BDoune杜恩() feud.Barony {
	return c.杜恩Doune
}

func (c *斯特拉森StrathearnCounty) BDunblane邓布兰() feud.Barony {
	return c.邓布兰Dunblane
}

func (c *斯特拉森StrathearnCounty) BInchaffray因彻法莱() feud.Barony {
	return c.因彻法莱Inchaffray
}

func (c *斯特拉森StrathearnCounty) BKenmore肯莫尔() feud.Barony {
	return c.肯莫尔Kenmore
}

func (c *斯特拉森StrathearnCounty) BMadderty马德第() feud.Barony {
	return c.马德第Madderty
}

func (c *斯特拉森StrathearnCounty) BTullibardine图里巴丁() feud.Barony {
	return c.图里巴丁Tullibardine
}

var CStrathearn斯特拉森 StrathearnCounty = &斯特拉森StrathearnCounty{}

func init() {
	f := CStrathearn斯特拉森.(*斯特拉森StrathearnCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "42",
		Title:     "strathearn",
		TitleName: "斯特拉森",
		TitleCode: "c_strathearn",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥赫特拉德Auchterarder = BAuchterarder奥赫特拉德
	f.奥赫特拉德Auchterarder.SetParent(f)

	f.克里夫Crieff = BCrieff克里夫
	f.克里夫Crieff.SetParent(f)

	f.杜恩Doune = BDoune杜恩
	f.杜恩Doune.SetParent(f)

	f.邓布兰Dunblane = BDunblane邓布兰
	f.邓布兰Dunblane.SetParent(f)

	f.因彻法莱Inchaffray = BInchaffray因彻法莱
	f.因彻法莱Inchaffray.SetParent(f)

	f.肯莫尔Kenmore = BKenmore肯莫尔
	f.肯莫尔Kenmore.SetParent(f)

	f.马德第Madderty = BMadderty马德第
	f.马德第Madderty.SetParent(f)

	f.图里巴丁Tullibardine = BTullibardine图里巴丁
	f.图里巴丁Tullibardine.SetParent(f)

}
