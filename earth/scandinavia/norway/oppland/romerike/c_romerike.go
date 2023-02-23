package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RomerikeCounty interface {
	feud.County
	BEidsvoll埃兹伏尔() feud.Barony
	BFetsund费聪() feud.Barony
	BHurdal许达尔() feud.Barony
	BJessheim杰斯海姆() feud.Barony
	BNannestad南内斯塔() feud.Barony
	BSolheim索尔海姆() feud.Barony
	BSorumsand瑟吕姆桑() feud.Barony
}

type 鲁默里克RomerikeCounty struct {
	feud.BaseCounty
	埃兹伏尔Eidsvoll  feud.Barony
	费聪Fetsund     feud.Barony
	许达尔Hurdal     feud.Barony
	杰斯海姆Jessheim  feud.Barony
	南内斯塔Nannestad feud.Barony
	索尔海姆Solheim   feud.Barony
	瑟吕姆桑Sorumsand feud.Barony
}

func (c *鲁默里克RomerikeCounty) BEidsvoll埃兹伏尔() feud.Barony {
	return c.埃兹伏尔Eidsvoll
}

func (c *鲁默里克RomerikeCounty) BFetsund费聪() feud.Barony {
	return c.费聪Fetsund
}

func (c *鲁默里克RomerikeCounty) BHurdal许达尔() feud.Barony {
	return c.许达尔Hurdal
}

func (c *鲁默里克RomerikeCounty) BJessheim杰斯海姆() feud.Barony {
	return c.杰斯海姆Jessheim
}

func (c *鲁默里克RomerikeCounty) BNannestad南内斯塔() feud.Barony {
	return c.南内斯塔Nannestad
}

func (c *鲁默里克RomerikeCounty) BSolheim索尔海姆() feud.Barony {
	return c.索尔海姆Solheim
}

func (c *鲁默里克RomerikeCounty) BSorumsand瑟吕姆桑() feud.Barony {
	return c.瑟吕姆桑Sorumsand
}

var CRomerike鲁默里克 RomerikeCounty = &鲁默里克RomerikeCounty{}

func init() {
	f := CRomerike鲁默里克.(*鲁默里克RomerikeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1620",
		Title:     "romerike",
		TitleName: "鲁默里克",
		TitleCode: "c_romerike",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃兹伏尔Eidsvoll = BEidsvoll埃兹伏尔
	f.埃兹伏尔Eidsvoll.SetParent(f)

	f.费聪Fetsund = BFetsund费聪
	f.费聪Fetsund.SetParent(f)

	f.许达尔Hurdal = BHurdal许达尔
	f.许达尔Hurdal.SetParent(f)

	f.杰斯海姆Jessheim = BJessheim杰斯海姆
	f.杰斯海姆Jessheim.SetParent(f)

	f.南内斯塔Nannestad = BNannestad南内斯塔
	f.南内斯塔Nannestad.SetParent(f)

	f.索尔海姆Solheim = BSolheim索尔海姆
	f.索尔海姆Solheim.SetParent(f)

	f.瑟吕姆桑Sorumsand = BSorumsand瑟吕姆桑
	f.瑟吕姆桑Sorumsand.SetParent(f)

}
