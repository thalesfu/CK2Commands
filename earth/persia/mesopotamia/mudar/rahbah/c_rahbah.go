package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RahbahCounty interface {
	feud.County
	BAnah阿纳海() feud.Barony
	BBusan布桑() feud.Barony
	BNimreh尼梅赫() feud.Barony
	BQummoualad喀穆喀拉德() feud.Barony
	BRahbah拉赫巴() feud.Barony
	BSalah萨拉赫() feud.Barony
	BTaraba塔拉巴() feud.Barony
	BTlilin特利林() feud.Barony
}

type 拉赫巴RahbahCounty struct {
	feud.BaseCounty
	阿纳海Anah         feud.Barony
	布桑Busan         feud.Barony
	尼梅赫Nimreh       feud.Barony
	喀穆喀拉德Qummoualad feud.Barony
	拉赫巴Rahbah       feud.Barony
	萨拉赫Salah        feud.Barony
	塔拉巴Taraba       feud.Barony
	特利林Tlilin       feud.Barony
}

func (c *拉赫巴RahbahCounty) BAnah阿纳海() feud.Barony {
	return c.阿纳海Anah
}

func (c *拉赫巴RahbahCounty) BBusan布桑() feud.Barony {
	return c.布桑Busan
}

func (c *拉赫巴RahbahCounty) BNimreh尼梅赫() feud.Barony {
	return c.尼梅赫Nimreh
}

func (c *拉赫巴RahbahCounty) BQummoualad喀穆喀拉德() feud.Barony {
	return c.喀穆喀拉德Qummoualad
}

func (c *拉赫巴RahbahCounty) BRahbah拉赫巴() feud.Barony {
	return c.拉赫巴Rahbah
}

func (c *拉赫巴RahbahCounty) BSalah萨拉赫() feud.Barony {
	return c.萨拉赫Salah
}

func (c *拉赫巴RahbahCounty) BTaraba塔拉巴() feud.Barony {
	return c.塔拉巴Taraba
}

func (c *拉赫巴RahbahCounty) BTlilin特利林() feud.Barony {
	return c.特利林Tlilin
}

var CRahbah拉赫巴 RahbahCounty = &拉赫巴RahbahCounty{}

func init() {
	f := CRahbah拉赫巴.(*拉赫巴RahbahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "712",
		Title:     "rahbah",
		TitleName: "拉赫巴",
		TitleCode: "c_rahbah",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿纳海Anah = BAnah阿纳海
	f.阿纳海Anah.SetParent(f)

	f.布桑Busan = BBusan布桑
	f.布桑Busan.SetParent(f)

	f.尼梅赫Nimreh = BNimreh尼梅赫
	f.尼梅赫Nimreh.SetParent(f)

	f.喀穆喀拉德Qummoualad = BQummoualad喀穆喀拉德
	f.喀穆喀拉德Qummoualad.SetParent(f)

	f.拉赫巴Rahbah = BRahbah拉赫巴
	f.拉赫巴Rahbah.SetParent(f)

	f.萨拉赫Salah = BSalah萨拉赫
	f.萨拉赫Salah.SetParent(f)

	f.塔拉巴Taraba = BTaraba塔拉巴
	f.塔拉巴Taraba.SetParent(f)

	f.特利林Tlilin = BTlilin特利林
	f.特利林Tlilin.SetParent(f)

}
