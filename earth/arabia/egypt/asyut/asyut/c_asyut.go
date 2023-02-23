package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsyutCounty interface {
	feud.County
	BAsyut艾斯尤特() feud.Barony
	BBeitkhallaf拜特哈利拉夫() feud.Barony
	BDurunka杜鲁卡() feud.Barony
	BKusai古西耶() feud.Barony
	BMeir米尔() feud.Barony
	BSuhaj苏哈吉() feud.Barony
	BWannina万尼亚() feud.Barony
}

type 艾斯尤特AsyutCounty struct {
	feud.BaseCounty
	艾斯尤特Asyut         feud.Barony
	拜特哈利拉夫Beitkhallaf feud.Barony
	杜鲁卡Durunka        feud.Barony
	古西耶Kusai          feud.Barony
	米尔Meir            feud.Barony
	苏哈吉Suhaj          feud.Barony
	万尼亚Wannina        feud.Barony
}

func (c *艾斯尤特AsyutCounty) BAsyut艾斯尤特() feud.Barony {
	return c.艾斯尤特Asyut
}

func (c *艾斯尤特AsyutCounty) BBeitkhallaf拜特哈利拉夫() feud.Barony {
	return c.拜特哈利拉夫Beitkhallaf
}

func (c *艾斯尤特AsyutCounty) BDurunka杜鲁卡() feud.Barony {
	return c.杜鲁卡Durunka
}

func (c *艾斯尤特AsyutCounty) BKusai古西耶() feud.Barony {
	return c.古西耶Kusai
}

func (c *艾斯尤特AsyutCounty) BMeir米尔() feud.Barony {
	return c.米尔Meir
}

func (c *艾斯尤特AsyutCounty) BSuhaj苏哈吉() feud.Barony {
	return c.苏哈吉Suhaj
}

func (c *艾斯尤特AsyutCounty) BWannina万尼亚() feud.Barony {
	return c.万尼亚Wannina
}

var CAsyut艾斯尤特 AsyutCounty = &艾斯尤特AsyutCounty{}

func init() {
	f := CAsyut艾斯尤特.(*艾斯尤特AsyutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "795",
		Title:     "asyut",
		TitleName: "艾斯尤特",
		TitleCode: "c_asyut",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾斯尤特Asyut = BAsyut艾斯尤特
	f.艾斯尤特Asyut.SetParent(f)

	f.拜特哈利拉夫Beitkhallaf = BBeitkhallaf拜特哈利拉夫
	f.拜特哈利拉夫Beitkhallaf.SetParent(f)

	f.杜鲁卡Durunka = BDurunka杜鲁卡
	f.杜鲁卡Durunka.SetParent(f)

	f.古西耶Kusai = BKusai古西耶
	f.古西耶Kusai.SetParent(f)

	f.米尔Meir = BMeir米尔
	f.米尔Meir.SetParent(f)

	f.苏哈吉Suhaj = BSuhaj苏哈吉
	f.苏哈吉Suhaj.SetParent(f)

	f.万尼亚Wannina = BWannina万尼亚
	f.万尼亚Wannina.SetParent(f)

}
