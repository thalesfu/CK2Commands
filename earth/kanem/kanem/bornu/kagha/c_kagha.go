package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaghaCounty interface {
	feud.County
	BAmcaka阿姆察卡() feud.Barony
	BAtemaga阿特马加() feud.Barony
	BBauni包尼() feud.Barony
	BKotoko科托科() feud.Barony
	BNgala恩加拉() feud.Barony
}

type 卡加KaghaCounty struct {
	feud.BaseCounty
	阿姆察卡Amcaka  feud.Barony
	阿特马加Atemaga feud.Barony
	包尼Bauni     feud.Barony
	科托科Kotoko   feud.Barony
	恩加拉Ngala    feud.Barony
}

func (c *卡加KaghaCounty) BAmcaka阿姆察卡() feud.Barony {
	return c.阿姆察卡Amcaka
}

func (c *卡加KaghaCounty) BAtemaga阿特马加() feud.Barony {
	return c.阿特马加Atemaga
}

func (c *卡加KaghaCounty) BBauni包尼() feud.Barony {
	return c.包尼Bauni
}

func (c *卡加KaghaCounty) BKotoko科托科() feud.Barony {
	return c.科托科Kotoko
}

func (c *卡加KaghaCounty) BNgala恩加拉() feud.Barony {
	return c.恩加拉Ngala
}

var CKagha卡加 KaghaCounty = &卡加KaghaCounty{}

func init() {
	f := CKagha卡加.(*卡加KaghaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1742",
		Title:     "kagha",
		TitleName: "卡加",
		TitleCode: "c_kagha",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆察卡Amcaka = BAmcaka阿姆察卡
	f.阿姆察卡Amcaka.SetParent(f)

	f.阿特马加Atemaga = BAtemaga阿特马加
	f.阿特马加Atemaga.SetParent(f)

	f.包尼Bauni = BBauni包尼
	f.包尼Bauni.SetParent(f)

	f.科托科Kotoko = BKotoko科托科
	f.科托科Kotoko.SetParent(f)

	f.恩加拉Ngala = BNgala恩加拉
	f.恩加拉Ngala.SetParent(f)

}
