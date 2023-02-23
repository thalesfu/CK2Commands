package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EsztergomCounty interface {
	feud.County
	BEsztergom埃斯泰尔戈姆() feud.Barony
	BKakath考考特() feud.Barony
	BKomarom科马罗姆() feud.Barony
	BNagyigmand瑙吉格曼德() feud.Barony
	BNemesocsa奈迈绍乔() feud.Barony
	BOgylla欧焦洛() feud.Barony
	BTatabanya陶陶巴尼奥() feud.Barony
}

type 埃斯泰尔戈姆EsztergomCounty struct {
	feud.BaseCounty
	埃斯泰尔戈姆Esztergom feud.Barony
	考考特Kakath       feud.Barony
	科马罗姆Komarom     feud.Barony
	瑙吉格曼德Nagyigmand feud.Barony
	奈迈绍乔Nemesocsa   feud.Barony
	欧焦洛Ogylla       feud.Barony
	陶陶巴尼奥Tatabanya  feud.Barony
}

func (c *埃斯泰尔戈姆EsztergomCounty) BEsztergom埃斯泰尔戈姆() feud.Barony {
	return c.埃斯泰尔戈姆Esztergom
}

func (c *埃斯泰尔戈姆EsztergomCounty) BKakath考考特() feud.Barony {
	return c.考考特Kakath
}

func (c *埃斯泰尔戈姆EsztergomCounty) BKomarom科马罗姆() feud.Barony {
	return c.科马罗姆Komarom
}

func (c *埃斯泰尔戈姆EsztergomCounty) BNagyigmand瑙吉格曼德() feud.Barony {
	return c.瑙吉格曼德Nagyigmand
}

func (c *埃斯泰尔戈姆EsztergomCounty) BNemesocsa奈迈绍乔() feud.Barony {
	return c.奈迈绍乔Nemesocsa
}

func (c *埃斯泰尔戈姆EsztergomCounty) BOgylla欧焦洛() feud.Barony {
	return c.欧焦洛Ogylla
}

func (c *埃斯泰尔戈姆EsztergomCounty) BTatabanya陶陶巴尼奥() feud.Barony {
	return c.陶陶巴尼奥Tatabanya
}

var CEsztergom埃斯泰尔戈姆 EsztergomCounty = &埃斯泰尔戈姆EsztergomCounty{}

func init() {
	f := CEsztergom埃斯泰尔戈姆.(*埃斯泰尔戈姆EsztergomCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "444",
		Title:     "esztergom",
		TitleName: "埃斯泰尔戈姆",
		TitleCode: "c_esztergom",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃斯泰尔戈姆Esztergom = BEsztergom埃斯泰尔戈姆
	f.埃斯泰尔戈姆Esztergom.SetParent(f)

	f.考考特Kakath = BKakath考考特
	f.考考特Kakath.SetParent(f)

	f.科马罗姆Komarom = BKomarom科马罗姆
	f.科马罗姆Komarom.SetParent(f)

	f.瑙吉格曼德Nagyigmand = BNagyigmand瑙吉格曼德
	f.瑙吉格曼德Nagyigmand.SetParent(f)

	f.奈迈绍乔Nemesocsa = BNemesocsa奈迈绍乔
	f.奈迈绍乔Nemesocsa.SetParent(f)

	f.欧焦洛Ogylla = BOgylla欧焦洛
	f.欧焦洛Ogylla.SetParent(f)

	f.陶陶巴尼奥Tatabanya = BTatabanya陶陶巴尼奥
	f.陶陶巴尼奥Tatabanya.SetParent(f)

}
