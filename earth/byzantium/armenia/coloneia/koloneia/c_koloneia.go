package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KoloneiaCounty interface {
	feud.County
	BAkn埃金() feud.Barony
	BAnatolnikopolis尼科波利斯() feud.Barony
	BCeltzene凯尔齐尼() feud.Barony
	BGamakh凯马赫() feud.Barony
	BGerjanis盖尔贾尼斯() feud.Barony
	BKoloneia科洛尼亚() feud.Barony
	BMazaka马扎卡() feud.Barony
	BTephrice特拉赫里斯() feud.Barony
}

type 科洛尼亚KoloneiaCounty struct {
	feud.BaseCounty
	埃金Akn                feud.Barony
	尼科波利斯Anatolnikopolis feud.Barony
	凯尔齐尼Celtzene         feud.Barony
	凯马赫Gamakh            feud.Barony
	盖尔贾尼斯Gerjanis        feud.Barony
	科洛尼亚Koloneia         feud.Barony
	马扎卡Mazaka            feud.Barony
	特拉赫里斯Tephrice        feud.Barony
}

func (c *科洛尼亚KoloneiaCounty) BAkn埃金() feud.Barony {
	return c.埃金Akn
}

func (c *科洛尼亚KoloneiaCounty) BAnatolnikopolis尼科波利斯() feud.Barony {
	return c.尼科波利斯Anatolnikopolis
}

func (c *科洛尼亚KoloneiaCounty) BCeltzene凯尔齐尼() feud.Barony {
	return c.凯尔齐尼Celtzene
}

func (c *科洛尼亚KoloneiaCounty) BGamakh凯马赫() feud.Barony {
	return c.凯马赫Gamakh
}

func (c *科洛尼亚KoloneiaCounty) BGerjanis盖尔贾尼斯() feud.Barony {
	return c.盖尔贾尼斯Gerjanis
}

func (c *科洛尼亚KoloneiaCounty) BKoloneia科洛尼亚() feud.Barony {
	return c.科洛尼亚Koloneia
}

func (c *科洛尼亚KoloneiaCounty) BMazaka马扎卡() feud.Barony {
	return c.马扎卡Mazaka
}

func (c *科洛尼亚KoloneiaCounty) BTephrice特拉赫里斯() feud.Barony {
	return c.特拉赫里斯Tephrice
}

var CKoloneia科洛尼亚 KoloneiaCounty = &科洛尼亚KoloneiaCounty{}

func init() {
	f := CKoloneia科洛尼亚.(*科洛尼亚KoloneiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "706",
		Title:     "koloneia",
		TitleName: "科洛尼亚",
		TitleCode: "c_koloneia",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃金Akn = BAkn埃金
	f.埃金Akn.SetParent(f)

	f.尼科波利斯Anatolnikopolis = BAnatolnikopolis尼科波利斯
	f.尼科波利斯Anatolnikopolis.SetParent(f)

	f.凯尔齐尼Celtzene = BCeltzene凯尔齐尼
	f.凯尔齐尼Celtzene.SetParent(f)

	f.凯马赫Gamakh = BGamakh凯马赫
	f.凯马赫Gamakh.SetParent(f)

	f.盖尔贾尼斯Gerjanis = BGerjanis盖尔贾尼斯
	f.盖尔贾尼斯Gerjanis.SetParent(f)

	f.科洛尼亚Koloneia = BKoloneia科洛尼亚
	f.科洛尼亚Koloneia.SetParent(f)

	f.马扎卡Mazaka = BMazaka马扎卡
	f.马扎卡Mazaka.SetParent(f)

	f.特拉赫里斯Tephrice = BTephrice特拉赫里斯
	f.特拉赫里斯Tephrice.SetParent(f)

}
