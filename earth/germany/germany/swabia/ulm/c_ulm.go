package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UlmCounty interface {
	feud.County
	BBiberach比伯拉赫() feud.Barony
	BErbach埃尔巴赫() feud.Barony
	BGoppingen格平根() feud.Barony
	BIsny伊斯尼() feud.Barony
	BMemmingen梅明根() feud.Barony
	BTeck泰克() feud.Barony
	BUlm乌尔姆() feud.Barony
	BZwiefalten茨维法尔滕() feud.Barony
}

type 乌尔姆UlmCounty struct {
	feud.BaseCounty
	比伯拉赫Biberach    feud.Barony
	埃尔巴赫Erbach      feud.Barony
	格平根Goppingen    feud.Barony
	伊斯尼Isny         feud.Barony
	梅明根Memmingen    feud.Barony
	泰克Teck          feud.Barony
	乌尔姆Ulm          feud.Barony
	茨维法尔滕Zwiefalten feud.Barony
}

func (c *乌尔姆UlmCounty) BBiberach比伯拉赫() feud.Barony {
	return c.比伯拉赫Biberach
}

func (c *乌尔姆UlmCounty) BErbach埃尔巴赫() feud.Barony {
	return c.埃尔巴赫Erbach
}

func (c *乌尔姆UlmCounty) BGoppingen格平根() feud.Barony {
	return c.格平根Goppingen
}

func (c *乌尔姆UlmCounty) BIsny伊斯尼() feud.Barony {
	return c.伊斯尼Isny
}

func (c *乌尔姆UlmCounty) BMemmingen梅明根() feud.Barony {
	return c.梅明根Memmingen
}

func (c *乌尔姆UlmCounty) BTeck泰克() feud.Barony {
	return c.泰克Teck
}

func (c *乌尔姆UlmCounty) BUlm乌尔姆() feud.Barony {
	return c.乌尔姆Ulm
}

func (c *乌尔姆UlmCounty) BZwiefalten茨维法尔滕() feud.Barony {
	return c.茨维法尔滕Zwiefalten
}

var CUlm乌尔姆 UlmCounty = &乌尔姆UlmCounty{}

func init() {
	f := CUlm乌尔姆.(*乌尔姆UlmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "252",
		Title:     "ulm",
		TitleName: "乌尔姆",
		TitleCode: "c_ulm",
		Baronies:  map[string]feud.Barony{},
	}

	f.比伯拉赫Biberach = BBiberach比伯拉赫
	f.比伯拉赫Biberach.SetParent(f)

	f.埃尔巴赫Erbach = BErbach埃尔巴赫
	f.埃尔巴赫Erbach.SetParent(f)

	f.格平根Goppingen = BGoppingen格平根
	f.格平根Goppingen.SetParent(f)

	f.伊斯尼Isny = BIsny伊斯尼
	f.伊斯尼Isny.SetParent(f)

	f.梅明根Memmingen = BMemmingen梅明根
	f.梅明根Memmingen.SetParent(f)

	f.泰克Teck = BTeck泰克
	f.泰克Teck.SetParent(f)

	f.乌尔姆Ulm = BUlm乌尔姆
	f.乌尔姆Ulm.SetParent(f)

	f.茨维法尔滕Zwiefalten = BZwiefalten茨维法尔滕
	f.茨维法尔滕Zwiefalten.SetParent(f)

}
