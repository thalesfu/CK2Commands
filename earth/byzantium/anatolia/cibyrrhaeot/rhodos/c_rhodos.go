package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RhodosCounty interface {
	feud.County
	BHaraki哈拉基() feud.Barony
	BIalysos伊阿吕索斯() feud.Barony
	BKarpathos卡尔帕索斯() feud.Barony
	BKos科斯() feud.Barony
	BKoskinou科斯基努() feud.Barony
	BLindos林佐斯() feud.Barony
	BPefkos佩夫科斯() feud.Barony
	BRhodos罗德岛() feud.Barony
}

type 罗德岛RhodosCounty struct {
	feud.BaseCounty
	哈拉基Haraki      feud.Barony
	伊阿吕索斯Ialysos   feud.Barony
	卡尔帕索斯Karpathos feud.Barony
	科斯Kos          feud.Barony
	科斯基努Koskinou   feud.Barony
	林佐斯Lindos      feud.Barony
	佩夫科斯Pefkos     feud.Barony
	罗德岛Rhodos      feud.Barony
}

func (c *罗德岛RhodosCounty) BHaraki哈拉基() feud.Barony {
	return c.哈拉基Haraki
}

func (c *罗德岛RhodosCounty) BIalysos伊阿吕索斯() feud.Barony {
	return c.伊阿吕索斯Ialysos
}

func (c *罗德岛RhodosCounty) BKarpathos卡尔帕索斯() feud.Barony {
	return c.卡尔帕索斯Karpathos
}

func (c *罗德岛RhodosCounty) BKos科斯() feud.Barony {
	return c.科斯Kos
}

func (c *罗德岛RhodosCounty) BKoskinou科斯基努() feud.Barony {
	return c.科斯基努Koskinou
}

func (c *罗德岛RhodosCounty) BLindos林佐斯() feud.Barony {
	return c.林佐斯Lindos
}

func (c *罗德岛RhodosCounty) BPefkos佩夫科斯() feud.Barony {
	return c.佩夫科斯Pefkos
}

func (c *罗德岛RhodosCounty) BRhodos罗德岛() feud.Barony {
	return c.罗德岛Rhodos
}

var CRhodos罗德岛 RhodosCounty = &罗德岛RhodosCounty{}

func init() {
	f := CRhodos罗德岛.(*罗德岛RhodosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "483",
		Title:     "rhodos",
		TitleName: "罗德岛",
		TitleCode: "c_rhodos",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈拉基Haraki = BHaraki哈拉基
	f.哈拉基Haraki.SetParent(f)

	f.伊阿吕索斯Ialysos = BIalysos伊阿吕索斯
	f.伊阿吕索斯Ialysos.SetParent(f)

	f.卡尔帕索斯Karpathos = BKarpathos卡尔帕索斯
	f.卡尔帕索斯Karpathos.SetParent(f)

	f.科斯Kos = BKos科斯
	f.科斯Kos.SetParent(f)

	f.科斯基努Koskinou = BKoskinou科斯基努
	f.科斯基努Koskinou.SetParent(f)

	f.林佐斯Lindos = BLindos林佐斯
	f.林佐斯Lindos.SetParent(f)

	f.佩夫科斯Pefkos = BPefkos佩夫科斯
	f.佩夫科斯Pefkos.SetParent(f)

	f.罗德岛Rhodos = BRhodos罗德岛
	f.罗德岛Rhodos.SetParent(f)

}
