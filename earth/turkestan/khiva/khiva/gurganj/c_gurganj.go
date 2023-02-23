package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurganjCounty interface {
	feud.County
	BKaradepe卡拉捷佩() feud.Barony
	BKaraoy卡拉奥伊() feud.Barony
	BKhodjali霍卓里() feud.Barony
	BKyrkmolla基尔克莫拉() feud.Barony
	BNukus努库斯() feud.Barony
	BUrgench玉龙杰赤() feud.Barony
	BUzynsuv乌津苏夫() feud.Barony
}

type 玉龙杰赤GurganjCounty struct {
	feud.BaseCounty
	卡拉捷佩Karadepe   feud.Barony
	卡拉奥伊Karaoy     feud.Barony
	霍卓里Khodjali    feud.Barony
	基尔克莫拉Kyrkmolla feud.Barony
	努库斯Nukus       feud.Barony
	玉龙杰赤Urgench    feud.Barony
	乌津苏夫Uzynsuv    feud.Barony
}

func (c *玉龙杰赤GurganjCounty) BKaradepe卡拉捷佩() feud.Barony {
	return c.卡拉捷佩Karadepe
}

func (c *玉龙杰赤GurganjCounty) BKaraoy卡拉奥伊() feud.Barony {
	return c.卡拉奥伊Karaoy
}

func (c *玉龙杰赤GurganjCounty) BKhodjali霍卓里() feud.Barony {
	return c.霍卓里Khodjali
}

func (c *玉龙杰赤GurganjCounty) BKyrkmolla基尔克莫拉() feud.Barony {
	return c.基尔克莫拉Kyrkmolla
}

func (c *玉龙杰赤GurganjCounty) BNukus努库斯() feud.Barony {
	return c.努库斯Nukus
}

func (c *玉龙杰赤GurganjCounty) BUrgench玉龙杰赤() feud.Barony {
	return c.玉龙杰赤Urgench
}

func (c *玉龙杰赤GurganjCounty) BUzynsuv乌津苏夫() feud.Barony {
	return c.乌津苏夫Uzynsuv
}

var CGurganj玉龙杰赤 GurganjCounty = &玉龙杰赤GurganjCounty{}

func init() {
	f := CGurganj玉龙杰赤.(*玉龙杰赤GurganjCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1791",
		Title:     "gurganj",
		TitleName: "玉龙杰赤",
		TitleCode: "c_gurganj",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉捷佩Karadepe = BKaradepe卡拉捷佩
	f.卡拉捷佩Karadepe.SetParent(f)

	f.卡拉奥伊Karaoy = BKaraoy卡拉奥伊
	f.卡拉奥伊Karaoy.SetParent(f)

	f.霍卓里Khodjali = BKhodjali霍卓里
	f.霍卓里Khodjali.SetParent(f)

	f.基尔克莫拉Kyrkmolla = BKyrkmolla基尔克莫拉
	f.基尔克莫拉Kyrkmolla.SetParent(f)

	f.努库斯Nukus = BNukus努库斯
	f.努库斯Nukus.SetParent(f)

	f.玉龙杰赤Urgench = BUrgench玉龙杰赤
	f.玉龙杰赤Urgench.SetParent(f)

	f.乌津苏夫Uzynsuv = BUzynsuv乌津苏夫
	f.乌津苏夫Uzynsuv.SetParent(f)

}
