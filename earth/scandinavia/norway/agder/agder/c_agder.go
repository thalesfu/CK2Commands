package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AgderCounty interface {
	feud.County
	BFlekkefjord弗莱克菲尤尔() feud.Barony
	BGrimstad格里姆斯塔() feud.Barony
	BHolt霍尔特() feud.Barony
	BHorga霍尔加() feud.Barony
	BHylestad许勒斯塔() feud.Barony
	BIveland伊夫兰() feud.Barony
	BSirdal西尔达尔() feud.Barony
	BVisedal维瑟达尔() feud.Barony
}

type 阿格德尔AgderCounty struct {
	feud.BaseCounty
	弗莱克菲尤尔Flekkefjord feud.Barony
	格里姆斯塔Grimstad     feud.Barony
	霍尔特Holt           feud.Barony
	霍尔加Horga          feud.Barony
	许勒斯塔Hylestad      feud.Barony
	伊夫兰Iveland        feud.Barony
	西尔达尔Sirdal        feud.Barony
	维瑟达尔Visedal       feud.Barony
}

func (c *阿格德尔AgderCounty) BFlekkefjord弗莱克菲尤尔() feud.Barony {
	return c.弗莱克菲尤尔Flekkefjord
}

func (c *阿格德尔AgderCounty) BGrimstad格里姆斯塔() feud.Barony {
	return c.格里姆斯塔Grimstad
}

func (c *阿格德尔AgderCounty) BHolt霍尔特() feud.Barony {
	return c.霍尔特Holt
}

func (c *阿格德尔AgderCounty) BHorga霍尔加() feud.Barony {
	return c.霍尔加Horga
}

func (c *阿格德尔AgderCounty) BHylestad许勒斯塔() feud.Barony {
	return c.许勒斯塔Hylestad
}

func (c *阿格德尔AgderCounty) BIveland伊夫兰() feud.Barony {
	return c.伊夫兰Iveland
}

func (c *阿格德尔AgderCounty) BSirdal西尔达尔() feud.Barony {
	return c.西尔达尔Sirdal
}

func (c *阿格德尔AgderCounty) BVisedal维瑟达尔() feud.Barony {
	return c.维瑟达尔Visedal
}

var CAgder阿格德尔 AgderCounty = &阿格德尔AgderCounty{}

func init() {
	f := CAgder阿格德尔.(*阿格德尔AgderCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "268",
		Title:     "agder",
		TitleName: "阿格德尔",
		TitleCode: "c_agder",
		Baronies:  map[string]feud.Barony{},
	}

	f.弗莱克菲尤尔Flekkefjord = BFlekkefjord弗莱克菲尤尔
	f.弗莱克菲尤尔Flekkefjord.SetParent(f)

	f.格里姆斯塔Grimstad = BGrimstad格里姆斯塔
	f.格里姆斯塔Grimstad.SetParent(f)

	f.霍尔特Holt = BHolt霍尔特
	f.霍尔特Holt.SetParent(f)

	f.霍尔加Horga = BHorga霍尔加
	f.霍尔加Horga.SetParent(f)

	f.许勒斯塔Hylestad = BHylestad许勒斯塔
	f.许勒斯塔Hylestad.SetParent(f)

	f.伊夫兰Iveland = BIveland伊夫兰
	f.伊夫兰Iveland.SetParent(f)

	f.西尔达尔Sirdal = BSirdal西尔达尔
	f.西尔达尔Sirdal.SetParent(f)

	f.维瑟达尔Visedal = BVisedal维瑟达尔
	f.维瑟达尔Visedal.SetParent(f)

}
