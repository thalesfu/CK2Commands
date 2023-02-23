package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VengipuraCounty interface {
	feud.County
	BBandarulanka班达鲁兰卡() feud.Barony
	BDwarakatirumala达瓦拉卡蒂鲁马拉() feud.Barony
	BEluru艾鲁卢() feud.Barony
	BKondapalli贡德伯莱() feud.Barony
	BNarasapuram纳勒瑟布尔姆() feud.Barony
	BPalakollu波罗科鲁() feud.Barony
	BVengipura瓶耆罗() feud.Barony
}

type 瓶耆罗VengipuraCounty struct {
	feud.BaseCounty
	班达鲁兰卡Bandarulanka       feud.Barony
	达瓦拉卡蒂鲁马拉Dwarakatirumala feud.Barony
	艾鲁卢Eluru                feud.Barony
	贡德伯莱Kondapalli          feud.Barony
	纳勒瑟布尔姆Narasapuram       feud.Barony
	波罗科鲁Palakollu           feud.Barony
	瓶耆罗Vengipura            feud.Barony
}

func (c *瓶耆罗VengipuraCounty) BBandarulanka班达鲁兰卡() feud.Barony {
	return c.班达鲁兰卡Bandarulanka
}

func (c *瓶耆罗VengipuraCounty) BDwarakatirumala达瓦拉卡蒂鲁马拉() feud.Barony {
	return c.达瓦拉卡蒂鲁马拉Dwarakatirumala
}

func (c *瓶耆罗VengipuraCounty) BEluru艾鲁卢() feud.Barony {
	return c.艾鲁卢Eluru
}

func (c *瓶耆罗VengipuraCounty) BKondapalli贡德伯莱() feud.Barony {
	return c.贡德伯莱Kondapalli
}

func (c *瓶耆罗VengipuraCounty) BNarasapuram纳勒瑟布尔姆() feud.Barony {
	return c.纳勒瑟布尔姆Narasapuram
}

func (c *瓶耆罗VengipuraCounty) BPalakollu波罗科鲁() feud.Barony {
	return c.波罗科鲁Palakollu
}

func (c *瓶耆罗VengipuraCounty) BVengipura瓶耆罗() feud.Barony {
	return c.瓶耆罗Vengipura
}

var CVengipura瓶耆罗 VengipuraCounty = &瓶耆罗VengipuraCounty{}

func init() {
	f := CVengipura瓶耆罗.(*瓶耆罗VengipuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1123",
		Title:     "vengipura",
		TitleName: "瓶耆罗",
		TitleCode: "c_vengipura",
		Baronies:  map[string]feud.Barony{},
	}

	f.班达鲁兰卡Bandarulanka = BBandarulanka班达鲁兰卡
	f.班达鲁兰卡Bandarulanka.SetParent(f)

	f.达瓦拉卡蒂鲁马拉Dwarakatirumala = BDwarakatirumala达瓦拉卡蒂鲁马拉
	f.达瓦拉卡蒂鲁马拉Dwarakatirumala.SetParent(f)

	f.艾鲁卢Eluru = BEluru艾鲁卢
	f.艾鲁卢Eluru.SetParent(f)

	f.贡德伯莱Kondapalli = BKondapalli贡德伯莱
	f.贡德伯莱Kondapalli.SetParent(f)

	f.纳勒瑟布尔姆Narasapuram = BNarasapuram纳勒瑟布尔姆
	f.纳勒瑟布尔姆Narasapuram.SetParent(f)

	f.波罗科鲁Palakollu = BPalakollu波罗科鲁
	f.波罗科鲁Palakollu.SetParent(f)

	f.瓶耆罗Vengipura = BVengipura瓶耆罗
	f.瓶耆罗Vengipura.SetParent(f)

}
