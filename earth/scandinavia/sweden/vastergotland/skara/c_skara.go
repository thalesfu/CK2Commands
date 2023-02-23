package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SkaraCounty interface {
	feud.County
	BAxevalla阿克塞瓦拉() feud.Barony
	BForsholm福什霍尔姆() feud.Barony
	BGalakvist耶拉克维斯特() feud.Barony
	BHusaby许萨比() feud.Barony
	BSkara斯卡拉() feud.Barony
	BVarnhem瓦恩赫姆() feud.Barony
	BYmseborg于姆瑟堡() feud.Barony
}

type 斯卡拉SkaraCounty struct {
	feud.BaseCounty
	阿克塞瓦拉Axevalla   feud.Barony
	福什霍尔姆Forsholm   feud.Barony
	耶拉克维斯特Galakvist feud.Barony
	许萨比Husaby       feud.Barony
	斯卡拉Skara        feud.Barony
	瓦恩赫姆Varnhem     feud.Barony
	于姆瑟堡Ymseborg    feud.Barony
}

func (c *斯卡拉SkaraCounty) BAxevalla阿克塞瓦拉() feud.Barony {
	return c.阿克塞瓦拉Axevalla
}

func (c *斯卡拉SkaraCounty) BForsholm福什霍尔姆() feud.Barony {
	return c.福什霍尔姆Forsholm
}

func (c *斯卡拉SkaraCounty) BGalakvist耶拉克维斯特() feud.Barony {
	return c.耶拉克维斯特Galakvist
}

func (c *斯卡拉SkaraCounty) BHusaby许萨比() feud.Barony {
	return c.许萨比Husaby
}

func (c *斯卡拉SkaraCounty) BSkara斯卡拉() feud.Barony {
	return c.斯卡拉Skara
}

func (c *斯卡拉SkaraCounty) BVarnhem瓦恩赫姆() feud.Barony {
	return c.瓦恩赫姆Varnhem
}

func (c *斯卡拉SkaraCounty) BYmseborg于姆瑟堡() feud.Barony {
	return c.于姆瑟堡Ymseborg
}

var CSkara斯卡拉 SkaraCounty = &斯卡拉SkaraCounty{}

func init() {
	f := CSkara斯卡拉.(*斯卡拉SkaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1706",
		Title:     "skara",
		TitleName: "斯卡拉",
		TitleCode: "c_skara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克塞瓦拉Axevalla = BAxevalla阿克塞瓦拉
	f.阿克塞瓦拉Axevalla.SetParent(f)

	f.福什霍尔姆Forsholm = BForsholm福什霍尔姆
	f.福什霍尔姆Forsholm.SetParent(f)

	f.耶拉克维斯特Galakvist = BGalakvist耶拉克维斯特
	f.耶拉克维斯特Galakvist.SetParent(f)

	f.许萨比Husaby = BHusaby许萨比
	f.许萨比Husaby.SetParent(f)

	f.斯卡拉Skara = BSkara斯卡拉
	f.斯卡拉Skara.SetParent(f)

	f.瓦恩赫姆Varnhem = BVarnhem瓦恩赫姆
	f.瓦恩赫姆Varnhem.SetParent(f)

	f.于姆瑟堡Ymseborg = BYmseborg于姆瑟堡
	f.于姆瑟堡Ymseborg.SetParent(f)

}
