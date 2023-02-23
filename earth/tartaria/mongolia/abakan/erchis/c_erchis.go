package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ErchisCounty interface {
	feud.County
	BAqsw阿克苏() feud.Barony
	BEkibastuz埃基巴斯图兹() feud.Barony
	BErchis曳咥() feud.Barony
	BErtis也儿的石() feud.Barony
	BKoktobe科克托别() feud.Barony
	BKorya科尔雅() feud.Barony
	BSemey塞米伊() feud.Barony
}

type 克季ErchisCounty struct {
	feud.BaseCounty
	阿克苏Aqsw         feud.Barony
	埃基巴斯图兹Ekibastuz feud.Barony
	曳咥Erchis        feud.Barony
	也儿的石Ertis       feud.Barony
	科克托别Koktobe     feud.Barony
	科尔雅Korya        feud.Barony
	塞米伊Semey        feud.Barony
}

func (c *克季ErchisCounty) BAqsw阿克苏() feud.Barony {
	return c.阿克苏Aqsw
}

func (c *克季ErchisCounty) BEkibastuz埃基巴斯图兹() feud.Barony {
	return c.埃基巴斯图兹Ekibastuz
}

func (c *克季ErchisCounty) BErchis曳咥() feud.Barony {
	return c.曳咥Erchis
}

func (c *克季ErchisCounty) BErtis也儿的石() feud.Barony {
	return c.也儿的石Ertis
}

func (c *克季ErchisCounty) BKoktobe科克托别() feud.Barony {
	return c.科克托别Koktobe
}

func (c *克季ErchisCounty) BKorya科尔雅() feud.Barony {
	return c.科尔雅Korya
}

func (c *克季ErchisCounty) BSemey塞米伊() feud.Barony {
	return c.塞米伊Semey
}

var CErchis克季 ErchisCounty = &克季ErchisCounty{}

func init() {
	f := CErchis克季.(*克季ErchisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1463",
		Title:     "erchis",
		TitleName: "克季",
		TitleCode: "c_erchis",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克苏Aqsw = BAqsw阿克苏
	f.阿克苏Aqsw.SetParent(f)

	f.埃基巴斯图兹Ekibastuz = BEkibastuz埃基巴斯图兹
	f.埃基巴斯图兹Ekibastuz.SetParent(f)

	f.曳咥Erchis = BErchis曳咥
	f.曳咥Erchis.SetParent(f)

	f.也儿的石Ertis = BErtis也儿的石
	f.也儿的石Ertis.SetParent(f)

	f.科克托别Koktobe = BKoktobe科克托别
	f.科克托别Koktobe.SetParent(f)

	f.科尔雅Korya = BKorya科尔雅
	f.科尔雅Korya.SetParent(f)

	f.塞米伊Semey = BSemey塞米伊
	f.塞米伊Semey.SetParent(f)

}
