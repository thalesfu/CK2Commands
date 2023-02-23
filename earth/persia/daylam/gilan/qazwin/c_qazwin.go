package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QazwinCounty interface {
	feud.County
	BAbeyek阿卜耶克() feud.Barony
	BAhmadabad艾哈迈达巴德() feud.Barony
	BAlvand阿尔万德() feud.Barony
	BAvaj阿瓦吉() feud.Barony
	BBuinzahra布津扎拉() feud.Barony
	BLambsar拉莫萨尔() feud.Barony
	BQazwin加兹温() feud.Barony
	BTakestan塔凯斯坦() feud.Barony
}

type 加兹温QazwinCounty struct {
	feud.BaseCounty
	阿卜耶克Abeyek      feud.Barony
	艾哈迈达巴德Ahmadabad feud.Barony
	阿尔万德Alvand      feud.Barony
	阿瓦吉Avaj         feud.Barony
	布津扎拉Buinzahra   feud.Barony
	拉莫萨尔Lambsar     feud.Barony
	加兹温Qazwin       feud.Barony
	塔凯斯坦Takestan    feud.Barony
}

func (c *加兹温QazwinCounty) BAbeyek阿卜耶克() feud.Barony {
	return c.阿卜耶克Abeyek
}

func (c *加兹温QazwinCounty) BAhmadabad艾哈迈达巴德() feud.Barony {
	return c.艾哈迈达巴德Ahmadabad
}

func (c *加兹温QazwinCounty) BAlvand阿尔万德() feud.Barony {
	return c.阿尔万德Alvand
}

func (c *加兹温QazwinCounty) BAvaj阿瓦吉() feud.Barony {
	return c.阿瓦吉Avaj
}

func (c *加兹温QazwinCounty) BBuinzahra布津扎拉() feud.Barony {
	return c.布津扎拉Buinzahra
}

func (c *加兹温QazwinCounty) BLambsar拉莫萨尔() feud.Barony {
	return c.拉莫萨尔Lambsar
}

func (c *加兹温QazwinCounty) BQazwin加兹温() feud.Barony {
	return c.加兹温Qazwin
}

func (c *加兹温QazwinCounty) BTakestan塔凯斯坦() feud.Barony {
	return c.塔凯斯坦Takestan
}

var CQazwin加兹温 QazwinCounty = &加兹温QazwinCounty{}

func init() {
	f := CQazwin加兹温.(*加兹温QazwinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "664",
		Title:     "qazwin",
		TitleName: "加兹温",
		TitleCode: "c_qazwin",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜耶克Abeyek = BAbeyek阿卜耶克
	f.阿卜耶克Abeyek.SetParent(f)

	f.艾哈迈达巴德Ahmadabad = BAhmadabad艾哈迈达巴德
	f.艾哈迈达巴德Ahmadabad.SetParent(f)

	f.阿尔万德Alvand = BAlvand阿尔万德
	f.阿尔万德Alvand.SetParent(f)

	f.阿瓦吉Avaj = BAvaj阿瓦吉
	f.阿瓦吉Avaj.SetParent(f)

	f.布津扎拉Buinzahra = BBuinzahra布津扎拉
	f.布津扎拉Buinzahra.SetParent(f)

	f.拉莫萨尔Lambsar = BLambsar拉莫萨尔
	f.拉莫萨尔Lambsar.SetParent(f)

	f.加兹温Qazwin = BQazwin加兹温
	f.加兹温Qazwin.SetParent(f)

	f.塔凯斯坦Takestan = BTakestan塔凯斯坦
	f.塔凯斯坦Takestan.SetParent(f)

}
