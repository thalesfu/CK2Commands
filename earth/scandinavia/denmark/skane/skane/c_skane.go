package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SkaneCounty interface {
	feud.County
	BDalby达尔比() feud.Barony
	BHelsingborg赫尔辛堡() feud.Barony
	BHerrevad赫雷瓦德() feud.Barony
	BLillohus利勒胡斯() feud.Barony
	BLund隆德() feud.Barony
	BMalmo马尔默() feud.Barony
	BTrelleborg特雷勒堡() feud.Barony
	BUppakra乌波克拉() feud.Barony
	BYstad于斯塔德() feud.Barony
}

type 斯堪尼亚SkaneCounty struct {
	feud.BaseCounty
	达尔比Dalby        feud.Barony
	赫尔辛堡Helsingborg feud.Barony
	赫雷瓦德Herrevad    feud.Barony
	利勒胡斯Lillohus    feud.Barony
	隆德Lund          feud.Barony
	马尔默Malmo        feud.Barony
	特雷勒堡Trelleborg  feud.Barony
	乌波克拉Uppakra     feud.Barony
	于斯塔德Ystad       feud.Barony
}

func (c *斯堪尼亚SkaneCounty) BDalby达尔比() feud.Barony {
	return c.达尔比Dalby
}

func (c *斯堪尼亚SkaneCounty) BHelsingborg赫尔辛堡() feud.Barony {
	return c.赫尔辛堡Helsingborg
}

func (c *斯堪尼亚SkaneCounty) BHerrevad赫雷瓦德() feud.Barony {
	return c.赫雷瓦德Herrevad
}

func (c *斯堪尼亚SkaneCounty) BLillohus利勒胡斯() feud.Barony {
	return c.利勒胡斯Lillohus
}

func (c *斯堪尼亚SkaneCounty) BLund隆德() feud.Barony {
	return c.隆德Lund
}

func (c *斯堪尼亚SkaneCounty) BMalmo马尔默() feud.Barony {
	return c.马尔默Malmo
}

func (c *斯堪尼亚SkaneCounty) BTrelleborg特雷勒堡() feud.Barony {
	return c.特雷勒堡Trelleborg
}

func (c *斯堪尼亚SkaneCounty) BUppakra乌波克拉() feud.Barony {
	return c.乌波克拉Uppakra
}

func (c *斯堪尼亚SkaneCounty) BYstad于斯塔德() feud.Barony {
	return c.于斯塔德Ystad
}

var CSkane斯堪尼亚 SkaneCounty = &斯堪尼亚SkaneCounty{}

func init() {
	f := CSkane斯堪尼亚.(*斯堪尼亚SkaneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "303",
		Title:     "skane",
		TitleName: "斯堪尼亚",
		TitleCode: "c_skane",
		Baronies:  map[string]feud.Barony{},
	}

	f.达尔比Dalby = BDalby达尔比
	f.达尔比Dalby.SetParent(f)

	f.赫尔辛堡Helsingborg = BHelsingborg赫尔辛堡
	f.赫尔辛堡Helsingborg.SetParent(f)

	f.赫雷瓦德Herrevad = BHerrevad赫雷瓦德
	f.赫雷瓦德Herrevad.SetParent(f)

	f.利勒胡斯Lillohus = BLillohus利勒胡斯
	f.利勒胡斯Lillohus.SetParent(f)

	f.隆德Lund = BLund隆德
	f.隆德Lund.SetParent(f)

	f.马尔默Malmo = BMalmo马尔默
	f.马尔默Malmo.SetParent(f)

	f.特雷勒堡Trelleborg = BTrelleborg特雷勒堡
	f.特雷勒堡Trelleborg.SetParent(f)

	f.乌波克拉Uppakra = BUppakra乌波克拉
	f.乌波克拉Uppakra.SetParent(f)

	f.于斯塔德Ystad = BYstad于斯塔德
	f.于斯塔德Ystad.SetParent(f)

}
