package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SripathaCounty interface {
	feud.County
	BAbanhagari阿班和伽尼() feud.Barony
	BBayana婆耶那() feud.Barony
	BDhavalapuri陀伐罗补梨() feud.Barony
	BDurpa补波() feud.Barony
	BFirozpur费罗兹补() feud.Barony
	BSripatha室利波他() feud.Barony
}

type 室利波他SripathaCounty struct {
	feud.BaseCounty
	阿班和伽尼Abanhagari  feud.Barony
	婆耶那Bayana        feud.Barony
	陀伐罗补梨Dhavalapuri feud.Barony
	补波Durpa          feud.Barony
	费罗兹补Firozpur     feud.Barony
	室利波他Sripatha     feud.Barony
}

func (c *室利波他SripathaCounty) BAbanhagari阿班和伽尼() feud.Barony {
	return c.阿班和伽尼Abanhagari
}

func (c *室利波他SripathaCounty) BBayana婆耶那() feud.Barony {
	return c.婆耶那Bayana
}

func (c *室利波他SripathaCounty) BDhavalapuri陀伐罗补梨() feud.Barony {
	return c.陀伐罗补梨Dhavalapuri
}

func (c *室利波他SripathaCounty) BDurpa补波() feud.Barony {
	return c.补波Durpa
}

func (c *室利波他SripathaCounty) BFirozpur费罗兹补() feud.Barony {
	return c.费罗兹补Firozpur
}

func (c *室利波他SripathaCounty) BSripatha室利波他() feud.Barony {
	return c.室利波他Sripatha
}

var CSripatha室利波他 SripathaCounty = &室利波他SripathaCounty{}

func init() {
	f := CSripatha室利波他.(*室利波他SripathaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1357",
		Title:     "sripatha",
		TitleName: "室利波他",
		TitleCode: "c_sripatha",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿班和伽尼Abanhagari = BAbanhagari阿班和伽尼
	f.阿班和伽尼Abanhagari.SetParent(f)

	f.婆耶那Bayana = BBayana婆耶那
	f.婆耶那Bayana.SetParent(f)

	f.陀伐罗补梨Dhavalapuri = BDhavalapuri陀伐罗补梨
	f.陀伐罗补梨Dhavalapuri.SetParent(f)

	f.补波Durpa = BDurpa补波
	f.补波Durpa.SetParent(f)

	f.费罗兹补Firozpur = BFirozpur费罗兹补
	f.费罗兹补Firozpur.SetParent(f)

	f.室利波他Sripatha = BSripatha室利波他
	f.室利波他Sripatha.SetParent(f)

}
