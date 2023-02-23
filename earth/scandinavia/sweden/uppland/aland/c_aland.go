package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlandCounty interface {
	feud.County
	BEckero埃克勒() feud.Barony
	BFinstrom芬斯特伦() feud.Barony
	BGeta耶塔() feud.Barony
	BGodby哥德比() feud.Barony
	BKastelholm卡斯特霍姆() feud.Barony
	BKumlinge库姆灵厄() feud.Barony
	BSaltvik萨尔特维克() feud.Barony
	BSund松德() feud.Barony
}

type 奥兰AlandCounty struct {
	feud.BaseCounty
	埃克勒Eckero       feud.Barony
	芬斯特伦Finstrom    feud.Barony
	耶塔Geta          feud.Barony
	哥德比Godby        feud.Barony
	卡斯特霍姆Kastelholm feud.Barony
	库姆灵厄Kumlinge    feud.Barony
	萨尔特维克Saltvik    feud.Barony
	松德Sund          feud.Barony
}

func (c *奥兰AlandCounty) BEckero埃克勒() feud.Barony {
	return c.埃克勒Eckero
}

func (c *奥兰AlandCounty) BFinstrom芬斯特伦() feud.Barony {
	return c.芬斯特伦Finstrom
}

func (c *奥兰AlandCounty) BGeta耶塔() feud.Barony {
	return c.耶塔Geta
}

func (c *奥兰AlandCounty) BGodby哥德比() feud.Barony {
	return c.哥德比Godby
}

func (c *奥兰AlandCounty) BKastelholm卡斯特霍姆() feud.Barony {
	return c.卡斯特霍姆Kastelholm
}

func (c *奥兰AlandCounty) BKumlinge库姆灵厄() feud.Barony {
	return c.库姆灵厄Kumlinge
}

func (c *奥兰AlandCounty) BSaltvik萨尔特维克() feud.Barony {
	return c.萨尔特维克Saltvik
}

func (c *奥兰AlandCounty) BSund松德() feud.Barony {
	return c.松德Sund
}

var CAland奥兰 AlandCounty = &奥兰AlandCounty{}

func init() {
	f := CAland奥兰.(*奥兰AlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "291",
		Title:     "aland",
		TitleName: "奥兰",
		TitleCode: "c_aland",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃克勒Eckero = BEckero埃克勒
	f.埃克勒Eckero.SetParent(f)

	f.芬斯特伦Finstrom = BFinstrom芬斯特伦
	f.芬斯特伦Finstrom.SetParent(f)

	f.耶塔Geta = BGeta耶塔
	f.耶塔Geta.SetParent(f)

	f.哥德比Godby = BGodby哥德比
	f.哥德比Godby.SetParent(f)

	f.卡斯特霍姆Kastelholm = BKastelholm卡斯特霍姆
	f.卡斯特霍姆Kastelholm.SetParent(f)

	f.库姆灵厄Kumlinge = BKumlinge库姆灵厄
	f.库姆灵厄Kumlinge.SetParent(f)

	f.萨尔特维克Saltvik = BSaltvik萨尔特维克
	f.萨尔特维克Saltvik.SetParent(f)

	f.松德Sund = BSund松德
	f.松德Sund.SetParent(f)

}
