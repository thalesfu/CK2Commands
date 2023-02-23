package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlcalaCounty interface {
	feud.County
	BAlcala阿尔卡拉() feud.Barony
	BComplutum孔普卢图姆() feud.Barony
	BGuadalajara瓜达拉哈拉() feud.Barony
	BSegontia锡古恩萨() feud.Barony
	BSiguenza锡古恩萨() feud.Barony
	BVillavieja维拉韦利亚() feud.Barony
}

type 阿尔卡拉AlcalaCounty struct {
	feud.BaseCounty
	阿尔卡拉Alcala       feud.Barony
	孔普卢图姆Complutum   feud.Barony
	瓜达拉哈拉Guadalajara feud.Barony
	锡古恩萨Segontia     feud.Barony
	锡古恩萨Siguenza     feud.Barony
	维拉韦利亚Villavieja  feud.Barony
}

func (c *阿尔卡拉AlcalaCounty) BAlcala阿尔卡拉() feud.Barony {
	return c.阿尔卡拉Alcala
}

func (c *阿尔卡拉AlcalaCounty) BComplutum孔普卢图姆() feud.Barony {
	return c.孔普卢图姆Complutum
}

func (c *阿尔卡拉AlcalaCounty) BGuadalajara瓜达拉哈拉() feud.Barony {
	return c.瓜达拉哈拉Guadalajara
}

func (c *阿尔卡拉AlcalaCounty) BSegontia锡古恩萨() feud.Barony {
	return c.锡古恩萨Segontia
}

func (c *阿尔卡拉AlcalaCounty) BSiguenza锡古恩萨() feud.Barony {
	return c.锡古恩萨Siguenza
}

func (c *阿尔卡拉AlcalaCounty) BVillavieja维拉韦利亚() feud.Barony {
	return c.维拉韦利亚Villavieja
}

var CAlcala阿尔卡拉 AlcalaCounty = &阿尔卡拉AlcalaCounty{}

func init() {
	f := CAlcala阿尔卡拉.(*阿尔卡拉AlcalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1978",
		Title:     "alcala",
		TitleName: "阿尔卡拉",
		TitleCode: "c_alcala",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔卡拉Alcala = BAlcala阿尔卡拉
	f.阿尔卡拉Alcala.SetParent(f)

	f.孔普卢图姆Complutum = BComplutum孔普卢图姆
	f.孔普卢图姆Complutum.SetParent(f)

	f.瓜达拉哈拉Guadalajara = BGuadalajara瓜达拉哈拉
	f.瓜达拉哈拉Guadalajara.SetParent(f)

	f.锡古恩萨Segontia = BSegontia锡古恩萨
	f.锡古恩萨Segontia.SetParent(f)

	f.锡古恩萨Siguenza = BSiguenza锡古恩萨
	f.锡古恩萨Siguenza.SetParent(f)

	f.维拉韦利亚Villavieja = BVillavieja维拉韦利亚
	f.维拉韦利亚Villavieja.SetParent(f)

}
