package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CharolaisCounty interface {
	feud.County
	BCharolles沙罗勒() feud.Barony
	BDigoine迪关() feud.Barony
	BJoncy容西() feud.Barony
	BMontstvincent蒙圣万桑() feud.Barony
	BParay帕雷() feud.Barony
	BPerrecy佩雷西() feud.Barony
	BSemurenbrionnais瑟米尔昂布里奥奈() feud.Barony
	BToulonsurarroux阿鲁河畔土伦() feud.Barony
}

type 沙罗莱CharolaisCounty struct {
	feud.BaseCounty
	沙罗勒Charolles             feud.Barony
	迪关Digoine                feud.Barony
	容西Joncy                  feud.Barony
	蒙圣万桑Montstvincent        feud.Barony
	帕雷Paray                  feud.Barony
	佩雷西Perrecy               feud.Barony
	瑟米尔昂布里奥奈Semurenbrionnais feud.Barony
	阿鲁河畔土伦Toulonsurarroux    feud.Barony
}

func (c *沙罗莱CharolaisCounty) BCharolles沙罗勒() feud.Barony {
	return c.沙罗勒Charolles
}

func (c *沙罗莱CharolaisCounty) BDigoine迪关() feud.Barony {
	return c.迪关Digoine
}

func (c *沙罗莱CharolaisCounty) BJoncy容西() feud.Barony {
	return c.容西Joncy
}

func (c *沙罗莱CharolaisCounty) BMontstvincent蒙圣万桑() feud.Barony {
	return c.蒙圣万桑Montstvincent
}

func (c *沙罗莱CharolaisCounty) BParay帕雷() feud.Barony {
	return c.帕雷Paray
}

func (c *沙罗莱CharolaisCounty) BPerrecy佩雷西() feud.Barony {
	return c.佩雷西Perrecy
}

func (c *沙罗莱CharolaisCounty) BSemurenbrionnais瑟米尔昂布里奥奈() feud.Barony {
	return c.瑟米尔昂布里奥奈Semurenbrionnais
}

func (c *沙罗莱CharolaisCounty) BToulonsurarroux阿鲁河畔土伦() feud.Barony {
	return c.阿鲁河畔土伦Toulonsurarroux
}

var CCharolais沙罗莱 CharolaisCounty = &沙罗莱CharolaisCounty{}

func init() {
	f := CCharolais沙罗莱.(*沙罗莱CharolaisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "226",
		Title:     "charolais",
		TitleName: "沙罗莱",
		TitleCode: "c_charolais",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙罗勒Charolles = BCharolles沙罗勒
	f.沙罗勒Charolles.SetParent(f)

	f.迪关Digoine = BDigoine迪关
	f.迪关Digoine.SetParent(f)

	f.容西Joncy = BJoncy容西
	f.容西Joncy.SetParent(f)

	f.蒙圣万桑Montstvincent = BMontstvincent蒙圣万桑
	f.蒙圣万桑Montstvincent.SetParent(f)

	f.帕雷Paray = BParay帕雷
	f.帕雷Paray.SetParent(f)

	f.佩雷西Perrecy = BPerrecy佩雷西
	f.佩雷西Perrecy.SetParent(f)

	f.瑟米尔昂布里奥奈Semurenbrionnais = BSemurenbrionnais瑟米尔昂布里奥奈
	f.瑟米尔昂布里奥奈Semurenbrionnais.SetParent(f)

	f.阿鲁河畔土伦Toulonsurarroux = BToulonsurarroux阿鲁河畔土伦
	f.阿鲁河畔土伦Toulonsurarroux.SetParent(f)

}
