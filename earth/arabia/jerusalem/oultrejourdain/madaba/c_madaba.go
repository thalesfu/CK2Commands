package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MadabaCounty interface {
	feud.County
	BBilal比拉勒() feud.Barony
	BMadaba马德巴() feud.Barony
	BMuwaqqar穆瓦盖尔() feud.Barony
	BQastal盖斯泰勒() feud.Barony
	BSahab萨哈布() feud.Barony
	BSamhal萨姆哈拉() feud.Barony
}

type 马德巴MadabaCounty struct {
	feud.BaseCounty
	比拉勒Bilal     feud.Barony
	马德巴Madaba    feud.Barony
	穆瓦盖尔Muwaqqar feud.Barony
	盖斯泰勒Qastal   feud.Barony
	萨哈布Sahab     feud.Barony
	萨姆哈拉Samhal   feud.Barony
}

func (c *马德巴MadabaCounty) BBilal比拉勒() feud.Barony {
	return c.比拉勒Bilal
}

func (c *马德巴MadabaCounty) BMadaba马德巴() feud.Barony {
	return c.马德巴Madaba
}

func (c *马德巴MadabaCounty) BMuwaqqar穆瓦盖尔() feud.Barony {
	return c.穆瓦盖尔Muwaqqar
}

func (c *马德巴MadabaCounty) BQastal盖斯泰勒() feud.Barony {
	return c.盖斯泰勒Qastal
}

func (c *马德巴MadabaCounty) BSahab萨哈布() feud.Barony {
	return c.萨哈布Sahab
}

func (c *马德巴MadabaCounty) BSamhal萨姆哈拉() feud.Barony {
	return c.萨姆哈拉Samhal
}

var CMadaba马德巴 MadabaCounty = &马德巴MadabaCounty{}

func init() {
	f := CMadaba马德巴.(*马德巴MadabaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "723",
		Title:     "madaba",
		TitleName: "马德巴",
		TitleCode: "c_madaba",
		Baronies:  map[string]feud.Barony{},
	}

	f.比拉勒Bilal = BBilal比拉勒
	f.比拉勒Bilal.SetParent(f)

	f.马德巴Madaba = BMadaba马德巴
	f.马德巴Madaba.SetParent(f)

	f.穆瓦盖尔Muwaqqar = BMuwaqqar穆瓦盖尔
	f.穆瓦盖尔Muwaqqar.SetParent(f)

	f.盖斯泰勒Qastal = BQastal盖斯泰勒
	f.盖斯泰勒Qastal.SetParent(f)

	f.萨哈布Sahab = BSahab萨哈布
	f.萨哈布Sahab.SetParent(f)

	f.萨姆哈拉Samhal = BSamhal萨姆哈拉
	f.萨姆哈拉Samhal.SetParent(f)

}
