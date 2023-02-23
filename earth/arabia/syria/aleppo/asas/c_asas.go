package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsasCounty interface {
	feud.County
	BAsas阿萨斯() feud.Barony
	BKallinikos卡利尼克尤姆() feud.Barony
	BRakka腊卡() feud.Barony
	BResafa鲁萨费() feud.Barony
	BSouriya叙利亚() feud.Barony
	BTabqa塔布卡() feud.Barony
	BTalabyad泰勒艾卜耶德() feud.Barony
	BZaazouaa扎佐乌() feud.Barony
}

type 阿萨斯AsasCounty struct {
	feud.BaseCounty
	阿萨斯Asas          feud.Barony
	卡利尼克尤姆Kallinikos feud.Barony
	腊卡Rakka          feud.Barony
	鲁萨费Resafa        feud.Barony
	叙利亚Souriya       feud.Barony
	塔布卡Tabqa         feud.Barony
	泰勒艾卜耶德Talabyad   feud.Barony
	扎佐乌Zaazouaa      feud.Barony
}

func (c *阿萨斯AsasCounty) BAsas阿萨斯() feud.Barony {
	return c.阿萨斯Asas
}

func (c *阿萨斯AsasCounty) BKallinikos卡利尼克尤姆() feud.Barony {
	return c.卡利尼克尤姆Kallinikos
}

func (c *阿萨斯AsasCounty) BRakka腊卡() feud.Barony {
	return c.腊卡Rakka
}

func (c *阿萨斯AsasCounty) BResafa鲁萨费() feud.Barony {
	return c.鲁萨费Resafa
}

func (c *阿萨斯AsasCounty) BSouriya叙利亚() feud.Barony {
	return c.叙利亚Souriya
}

func (c *阿萨斯AsasCounty) BTabqa塔布卡() feud.Barony {
	return c.塔布卡Tabqa
}

func (c *阿萨斯AsasCounty) BTalabyad泰勒艾卜耶德() feud.Barony {
	return c.泰勒艾卜耶德Talabyad
}

func (c *阿萨斯AsasCounty) BZaazouaa扎佐乌() feud.Barony {
	return c.扎佐乌Zaazouaa
}

var CAsas阿萨斯 AsasCounty = &阿萨斯AsasCounty{}

func init() {
	f := CAsas阿萨斯.(*阿萨斯AsasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "709",
		Title:     "asas",
		TitleName: "阿萨斯",
		TitleCode: "c_asas",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿萨斯Asas = BAsas阿萨斯
	f.阿萨斯Asas.SetParent(f)

	f.卡利尼克尤姆Kallinikos = BKallinikos卡利尼克尤姆
	f.卡利尼克尤姆Kallinikos.SetParent(f)

	f.腊卡Rakka = BRakka腊卡
	f.腊卡Rakka.SetParent(f)

	f.鲁萨费Resafa = BResafa鲁萨费
	f.鲁萨费Resafa.SetParent(f)

	f.叙利亚Souriya = BSouriya叙利亚
	f.叙利亚Souriya.SetParent(f)

	f.塔布卡Tabqa = BTabqa塔布卡
	f.塔布卡Tabqa.SetParent(f)

	f.泰勒艾卜耶德Talabyad = BTalabyad泰勒艾卜耶德
	f.泰勒艾卜耶德Talabyad.SetParent(f)

	f.扎佐乌Zaazouaa = BZaazouaa扎佐乌
	f.扎佐乌Zaazouaa.SetParent(f)

}
