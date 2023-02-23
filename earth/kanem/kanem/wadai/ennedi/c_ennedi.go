package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EnnediCounty interface {
	feud.County
	BBandarbi班达尔比() feud.Barony
	BKebkabiya凯布卡比亚() feud.Barony
	BKobbai科拜() feud.Barony
	BSabu萨布() feud.Barony
	BShaiqab谢加卜() feud.Barony
}

type 恩内迪EnnediCounty struct {
	feud.BaseCounty
	班达尔比Bandarbi   feud.Barony
	凯布卡比亚Kebkabiya feud.Barony
	科拜Kobbai       feud.Barony
	萨布Sabu         feud.Barony
	谢加卜Shaiqab     feud.Barony
}

func (c *恩内迪EnnediCounty) BBandarbi班达尔比() feud.Barony {
	return c.班达尔比Bandarbi
}

func (c *恩内迪EnnediCounty) BKebkabiya凯布卡比亚() feud.Barony {
	return c.凯布卡比亚Kebkabiya
}

func (c *恩内迪EnnediCounty) BKobbai科拜() feud.Barony {
	return c.科拜Kobbai
}

func (c *恩内迪EnnediCounty) BSabu萨布() feud.Barony {
	return c.萨布Sabu
}

func (c *恩内迪EnnediCounty) BShaiqab谢加卜() feud.Barony {
	return c.谢加卜Shaiqab
}

var CEnnedi恩内迪 EnnediCounty = &恩内迪EnnediCounty{}

func init() {
	f := CEnnedi恩内迪.(*恩内迪EnnediCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1771",
		Title:     "ennedi",
		TitleName: "恩内迪",
		TitleCode: "c_ennedi",
		Baronies:  map[string]feud.Barony{},
	}

	f.班达尔比Bandarbi = BBandarbi班达尔比
	f.班达尔比Bandarbi.SetParent(f)

	f.凯布卡比亚Kebkabiya = BKebkabiya凯布卡比亚
	f.凯布卡比亚Kebkabiya.SetParent(f)

	f.科拜Kobbai = BKobbai科拜
	f.科拜Kobbai.SetParent(f)

	f.萨布Sabu = BSabu萨布
	f.萨布Sabu.SetParent(f)

	f.谢加卜Shaiqab = BShaiqab谢加卜
	f.谢加卜Shaiqab.SetParent(f)

}
