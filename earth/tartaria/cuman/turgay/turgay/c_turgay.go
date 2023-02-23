package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurgayCounty interface {
	feud.County
	BDralyk德拉雷克() feud.Barony
	BKarakal卡拉卡尔() feud.Barony
	BMankal曼卡尔() feud.Barony
	BTurgay图尔盖() feud.Barony
	BUrlyki乌尔雷基() feud.Barony
	BZhailyk扎伊雷克() feud.Barony
}

type 图尔盖TurgayCounty struct {
	feud.BaseCounty
	德拉雷克Dralyk  feud.Barony
	卡拉卡尔Karakal feud.Barony
	曼卡尔Mankal   feud.Barony
	图尔盖Turgay   feud.Barony
	乌尔雷基Urlyki  feud.Barony
	扎伊雷克Zhailyk feud.Barony
}

func (c *图尔盖TurgayCounty) BDralyk德拉雷克() feud.Barony {
	return c.德拉雷克Dralyk
}

func (c *图尔盖TurgayCounty) BKarakal卡拉卡尔() feud.Barony {
	return c.卡拉卡尔Karakal
}

func (c *图尔盖TurgayCounty) BMankal曼卡尔() feud.Barony {
	return c.曼卡尔Mankal
}

func (c *图尔盖TurgayCounty) BTurgay图尔盖() feud.Barony {
	return c.图尔盖Turgay
}

func (c *图尔盖TurgayCounty) BUrlyki乌尔雷基() feud.Barony {
	return c.乌尔雷基Urlyki
}

func (c *图尔盖TurgayCounty) BZhailyk扎伊雷克() feud.Barony {
	return c.扎伊雷克Zhailyk
}

var CTurgay图尔盖 TurgayCounty = &图尔盖TurgayCounty{}

func init() {
	f := CTurgay图尔盖.(*图尔盖TurgayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1273",
		Title:     "turgay",
		TitleName: "图尔盖",
		TitleCode: "c_turgay",
		Baronies:  map[string]feud.Barony{},
	}

	f.德拉雷克Dralyk = BDralyk德拉雷克
	f.德拉雷克Dralyk.SetParent(f)

	f.卡拉卡尔Karakal = BKarakal卡拉卡尔
	f.卡拉卡尔Karakal.SetParent(f)

	f.曼卡尔Mankal = BMankal曼卡尔
	f.曼卡尔Mankal.SetParent(f)

	f.图尔盖Turgay = BTurgay图尔盖
	f.图尔盖Turgay.SetParent(f)

	f.乌尔雷基Urlyki = BUrlyki乌尔雷基
	f.乌尔雷基Urlyki.SetParent(f)

	f.扎伊雷克Zhailyk = BZhailyk扎伊雷克
	f.扎伊雷克Zhailyk.SetParent(f)

}
