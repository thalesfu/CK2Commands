package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SoliCounty interface {
	feud.County
	BBrod布罗德() feud.Barony
	BDoboj多博伊() feud.Barony
	BMaglaj马格拉伊() feud.Barony
	BModrica莫德里查() feud.Barony
	BSrebrenik斯雷布雷尼克() feud.Barony
	BUsice乌日采() feud.Barony
}

type 乌索拉SoliCounty struct {
	feud.BaseCounty
	布罗德Brod         feud.Barony
	多博伊Doboj        feud.Barony
	马格拉伊Maglaj      feud.Barony
	莫德里查Modrica     feud.Barony
	斯雷布雷尼克Srebrenik feud.Barony
	乌日采Usice        feud.Barony
}

func (c *乌索拉SoliCounty) BBrod布罗德() feud.Barony {
	return c.布罗德Brod
}

func (c *乌索拉SoliCounty) BDoboj多博伊() feud.Barony {
	return c.多博伊Doboj
}

func (c *乌索拉SoliCounty) BMaglaj马格拉伊() feud.Barony {
	return c.马格拉伊Maglaj
}

func (c *乌索拉SoliCounty) BModrica莫德里查() feud.Barony {
	return c.莫德里查Modrica
}

func (c *乌索拉SoliCounty) BSrebrenik斯雷布雷尼克() feud.Barony {
	return c.斯雷布雷尼克Srebrenik
}

func (c *乌索拉SoliCounty) BUsice乌日采() feud.Barony {
	return c.乌日采Usice
}

var CSoli乌索拉 SoliCounty = &乌索拉SoliCounty{}

func init() {
	f := CSoli乌索拉.(*乌索拉SoliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1970",
		Title:     "soli",
		TitleName: "乌索拉",
		TitleCode: "c_soli",
		Baronies:  map[string]feud.Barony{},
	}

	f.布罗德Brod = BBrod布罗德
	f.布罗德Brod.SetParent(f)

	f.多博伊Doboj = BDoboj多博伊
	f.多博伊Doboj.SetParent(f)

	f.马格拉伊Maglaj = BMaglaj马格拉伊
	f.马格拉伊Maglaj.SetParent(f)

	f.莫德里查Modrica = BModrica莫德里查
	f.莫德里查Modrica.SetParent(f)

	f.斯雷布雷尼克Srebrenik = BSrebrenik斯雷布雷尼克
	f.斯雷布雷尼克Srebrenik.SetParent(f)

	f.乌日采Usice = BUsice乌日采
	f.乌日采Usice.SetParent(f)

}
