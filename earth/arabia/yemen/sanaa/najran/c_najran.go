package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NajranCounty interface {
	feud.County
	BHasakat哈萨卡特() feud.Barony
	BHuth侯斯() feud.Barony
	BNajran纳季兰() feud.Barony
	BRibakhah里巴哈() feud.Barony
	BSada萨达() feud.Barony
}

type 纳季兰NajranCounty struct {
	feud.BaseCounty
	哈萨卡特Hasakat feud.Barony
	侯斯Huth      feud.Barony
	纳季兰Najran   feud.Barony
	里巴哈Ribakhah feud.Barony
	萨达Sada      feud.Barony
}

func (c *纳季兰NajranCounty) BHasakat哈萨卡特() feud.Barony {
	return c.哈萨卡特Hasakat
}

func (c *纳季兰NajranCounty) BHuth侯斯() feud.Barony {
	return c.侯斯Huth
}

func (c *纳季兰NajranCounty) BNajran纳季兰() feud.Barony {
	return c.纳季兰Najran
}

func (c *纳季兰NajranCounty) BRibakhah里巴哈() feud.Barony {
	return c.里巴哈Ribakhah
}

func (c *纳季兰NajranCounty) BSada萨达() feud.Barony {
	return c.萨达Sada
}

var CNajran纳季兰 NajranCounty = &纳季兰NajranCounty{}

func init() {
	f := CNajran纳季兰.(*纳季兰NajranCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1533",
		Title:     "najran",
		TitleName: "纳季兰",
		TitleCode: "c_najran",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈萨卡特Hasakat = BHasakat哈萨卡特
	f.哈萨卡特Hasakat.SetParent(f)

	f.侯斯Huth = BHuth侯斯
	f.侯斯Huth.SetParent(f)

	f.纳季兰Najran = BNajran纳季兰
	f.纳季兰Najran.SetParent(f)

	f.里巴哈Ribakhah = BRibakhah里巴哈
	f.里巴哈Ribakhah.SetParent(f)

	f.萨达Sada = BSada萨达
	f.萨达Sada.SetParent(f)

}
