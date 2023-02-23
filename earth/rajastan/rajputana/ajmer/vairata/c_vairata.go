package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VairataCounty interface {
	feud.County
	BAmer阿墨尔() feud.Barony
	BDhosi都斯() feud.Barony
	BHarshnath曷利沙那他() feud.Barony
	BJeenmata耆那摩多() feud.Barony
	BJhunjhunu均均乌恩() feud.Barony
	BKhatushyam佉住商摩() feud.Barony
	BSanganer桑格内尔() feud.Barony
	BVairata吠罗多() feud.Barony
}

type 吠罗多VairataCounty struct {
	feud.BaseCounty
	阿墨尔Amer        feud.Barony
	都斯Dhosi        feud.Barony
	曷利沙那他Harshnath feud.Barony
	耆那摩多Jeenmata   feud.Barony
	均均乌恩Jhunjhunu  feud.Barony
	佉住商摩Khatushyam feud.Barony
	桑格内尔Sanganer   feud.Barony
	吠罗多Vairata     feud.Barony
}

func (c *吠罗多VairataCounty) BAmer阿墨尔() feud.Barony {
	return c.阿墨尔Amer
}

func (c *吠罗多VairataCounty) BDhosi都斯() feud.Barony {
	return c.都斯Dhosi
}

func (c *吠罗多VairataCounty) BHarshnath曷利沙那他() feud.Barony {
	return c.曷利沙那他Harshnath
}

func (c *吠罗多VairataCounty) BJeenmata耆那摩多() feud.Barony {
	return c.耆那摩多Jeenmata
}

func (c *吠罗多VairataCounty) BJhunjhunu均均乌恩() feud.Barony {
	return c.均均乌恩Jhunjhunu
}

func (c *吠罗多VairataCounty) BKhatushyam佉住商摩() feud.Barony {
	return c.佉住商摩Khatushyam
}

func (c *吠罗多VairataCounty) BSanganer桑格内尔() feud.Barony {
	return c.桑格内尔Sanganer
}

func (c *吠罗多VairataCounty) BVairata吠罗多() feud.Barony {
	return c.吠罗多Vairata
}

var CVairata吠罗多 VairataCounty = &吠罗多VairataCounty{}

func init() {
	f := CVairata吠罗多.(*吠罗多VairataCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1353",
		Title:     "vairata",
		TitleName: "吠罗多",
		TitleCode: "c_vairata",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿墨尔Amer = BAmer阿墨尔
	f.阿墨尔Amer.SetParent(f)

	f.都斯Dhosi = BDhosi都斯
	f.都斯Dhosi.SetParent(f)

	f.曷利沙那他Harshnath = BHarshnath曷利沙那他
	f.曷利沙那他Harshnath.SetParent(f)

	f.耆那摩多Jeenmata = BJeenmata耆那摩多
	f.耆那摩多Jeenmata.SetParent(f)

	f.均均乌恩Jhunjhunu = BJhunjhunu均均乌恩
	f.均均乌恩Jhunjhunu.SetParent(f)

	f.佉住商摩Khatushyam = BKhatushyam佉住商摩
	f.佉住商摩Khatushyam.SetParent(f)

	f.桑格内尔Sanganer = BSanganer桑格内尔
	f.桑格内尔Sanganer.SetParent(f)

	f.吠罗多Vairata = BVairata吠罗多
	f.吠罗多Vairata.SetParent(f)

}
