package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VaranasiCounty interface {
	feud.County
	BAleppi亚力皮() feud.Barony
	BBallia婆罗利阿() feud.Barony
	BBhadohi巴多希() feud.Barony
	BSarnath鹿野苑() feud.Barony
	BVaranasi波罗奈() feud.Barony
}

type 波罗奈VaranasiCounty struct {
	feud.BaseCounty
	亚力皮Aleppi   feud.Barony
	婆罗利阿Ballia  feud.Barony
	巴多希Bhadohi  feud.Barony
	鹿野苑Sarnath  feud.Barony
	波罗奈Varanasi feud.Barony
}

func (c *波罗奈VaranasiCounty) BAleppi亚力皮() feud.Barony {
	return c.亚力皮Aleppi
}

func (c *波罗奈VaranasiCounty) BBallia婆罗利阿() feud.Barony {
	return c.婆罗利阿Ballia
}

func (c *波罗奈VaranasiCounty) BBhadohi巴多希() feud.Barony {
	return c.巴多希Bhadohi
}

func (c *波罗奈VaranasiCounty) BSarnath鹿野苑() feud.Barony {
	return c.鹿野苑Sarnath
}

func (c *波罗奈VaranasiCounty) BVaranasi波罗奈() feud.Barony {
	return c.波罗奈Varanasi
}

var CVaranasi波罗奈 VaranasiCounty = &波罗奈VaranasiCounty{}

func init() {
	f := CVaranasi波罗奈.(*波罗奈VaranasiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1163",
		Title:     "varanasi",
		TitleName: "波罗奈",
		TitleCode: "c_varanasi",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚力皮Aleppi = BAleppi亚力皮
	f.亚力皮Aleppi.SetParent(f)

	f.婆罗利阿Ballia = BBallia婆罗利阿
	f.婆罗利阿Ballia.SetParent(f)

	f.巴多希Bhadohi = BBhadohi巴多希
	f.巴多希Bhadohi.SetParent(f)

	f.鹿野苑Sarnath = BSarnath鹿野苑
	f.鹿野苑Sarnath.SetParent(f)

	f.波罗奈Varanasi = BVaranasi波罗奈
	f.波罗奈Varanasi.SetParent(f)

}
