package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UsturtCounty interface {
	feud.County
	BAkkuduk阿克_库都克() feud.Barony
	BAksu阿克苏() feud.Barony
	BBailjar拜伊加尔() feud.Barony
	BBarsakelmos巴尔萨_克尔梅斯() feud.Barony
	BBussaga布萨加() feud.Barony
	BKaramola卡拉莫拉() feud.Barony
	BSengirkum先吉尔_库姆() feud.Barony
	BSumbe孙别() feud.Barony
}

type 乌斯秋尔特UsturtCounty struct {
	feud.BaseCounty
	阿克_库都克Akkuduk       feud.Barony
	阿克苏Aksu             feud.Barony
	拜伊加尔Bailjar         feud.Barony
	巴尔萨_克尔梅斯Barsakelmos feud.Barony
	布萨加Bussaga          feud.Barony
	卡拉莫拉Karamola        feud.Barony
	先吉尔_库姆Sengirkum     feud.Barony
	孙别Sumbe             feud.Barony
}

func (c *乌斯秋尔特UsturtCounty) BAkkuduk阿克_库都克() feud.Barony {
	return c.阿克_库都克Akkuduk
}

func (c *乌斯秋尔特UsturtCounty) BAksu阿克苏() feud.Barony {
	return c.阿克苏Aksu
}

func (c *乌斯秋尔特UsturtCounty) BBailjar拜伊加尔() feud.Barony {
	return c.拜伊加尔Bailjar
}

func (c *乌斯秋尔特UsturtCounty) BBarsakelmos巴尔萨_克尔梅斯() feud.Barony {
	return c.巴尔萨_克尔梅斯Barsakelmos
}

func (c *乌斯秋尔特UsturtCounty) BBussaga布萨加() feud.Barony {
	return c.布萨加Bussaga
}

func (c *乌斯秋尔特UsturtCounty) BKaramola卡拉莫拉() feud.Barony {
	return c.卡拉莫拉Karamola
}

func (c *乌斯秋尔特UsturtCounty) BSengirkum先吉尔_库姆() feud.Barony {
	return c.先吉尔_库姆Sengirkum
}

func (c *乌斯秋尔特UsturtCounty) BSumbe孙别() feud.Barony {
	return c.孙别Sumbe
}

var CUsturt乌斯秋尔特 UsturtCounty = &乌斯秋尔特UsturtCounty{}

func init() {
	f := CUsturt乌斯秋尔特.(*乌斯秋尔特UsturtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "625",
		Title:     "usturt",
		TitleName: "乌斯秋尔特",
		TitleCode: "c_usturt",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克_库都克Akkuduk = BAkkuduk阿克_库都克
	f.阿克_库都克Akkuduk.SetParent(f)

	f.阿克苏Aksu = BAksu阿克苏
	f.阿克苏Aksu.SetParent(f)

	f.拜伊加尔Bailjar = BBailjar拜伊加尔
	f.拜伊加尔Bailjar.SetParent(f)

	f.巴尔萨_克尔梅斯Barsakelmos = BBarsakelmos巴尔萨_克尔梅斯
	f.巴尔萨_克尔梅斯Barsakelmos.SetParent(f)

	f.布萨加Bussaga = BBussaga布萨加
	f.布萨加Bussaga.SetParent(f)

	f.卡拉莫拉Karamola = BKaramola卡拉莫拉
	f.卡拉莫拉Karamola.SetParent(f)

	f.先吉尔_库姆Sengirkum = BSengirkum先吉尔_库姆
	f.先吉尔_库姆Sengirkum.SetParent(f)

	f.孙别Sumbe = BSumbe孙别
	f.孙别Sumbe.SetParent(f)

}
