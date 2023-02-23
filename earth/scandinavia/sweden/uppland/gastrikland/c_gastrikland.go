package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GastriklandCounty interface {
	feud.County
	BArsunda奥松达() feud.Barony
	BFarnebo费尔内布() feud.Barony
	BHamrange哈姆罗耶() feud.Barony
	BHedesunda海德松达() feud.Barony
	BHille希勒() feud.Barony
	BOckelbo奥克尔布() feud.Barony
	BTorsaker托绍克尔() feud.Barony
	BValbo瓦尔博() feud.Barony
}

type 耶斯特里克兰GastriklandCounty struct {
	feud.BaseCounty
	奥松达Arsunda    feud.Barony
	费尔内布Farnebo   feud.Barony
	哈姆罗耶Hamrange  feud.Barony
	海德松达Hedesunda feud.Barony
	希勒Hille       feud.Barony
	奥克尔布Ockelbo   feud.Barony
	托绍克尔Torsaker  feud.Barony
	瓦尔博Valbo      feud.Barony
}

func (c *耶斯特里克兰GastriklandCounty) BArsunda奥松达() feud.Barony {
	return c.奥松达Arsunda
}

func (c *耶斯特里克兰GastriklandCounty) BFarnebo费尔内布() feud.Barony {
	return c.费尔内布Farnebo
}

func (c *耶斯特里克兰GastriklandCounty) BHamrange哈姆罗耶() feud.Barony {
	return c.哈姆罗耶Hamrange
}

func (c *耶斯特里克兰GastriklandCounty) BHedesunda海德松达() feud.Barony {
	return c.海德松达Hedesunda
}

func (c *耶斯特里克兰GastriklandCounty) BHille希勒() feud.Barony {
	return c.希勒Hille
}

func (c *耶斯特里克兰GastriklandCounty) BOckelbo奥克尔布() feud.Barony {
	return c.奥克尔布Ockelbo
}

func (c *耶斯特里克兰GastriklandCounty) BTorsaker托绍克尔() feud.Barony {
	return c.托绍克尔Torsaker
}

func (c *耶斯特里克兰GastriklandCounty) BValbo瓦尔博() feud.Barony {
	return c.瓦尔博Valbo
}

var CGastrikland耶斯特里克兰 GastriklandCounty = &耶斯特里克兰GastriklandCounty{}

func init() {
	f := CGastrikland耶斯特里克兰.(*耶斯特里克兰GastriklandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "286",
		Title:     "gastrikland",
		TitleName: "耶斯特里克兰",
		TitleCode: "c_gastrikland",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥松达Arsunda = BArsunda奥松达
	f.奥松达Arsunda.SetParent(f)

	f.费尔内布Farnebo = BFarnebo费尔内布
	f.费尔内布Farnebo.SetParent(f)

	f.哈姆罗耶Hamrange = BHamrange哈姆罗耶
	f.哈姆罗耶Hamrange.SetParent(f)

	f.海德松达Hedesunda = BHedesunda海德松达
	f.海德松达Hedesunda.SetParent(f)

	f.希勒Hille = BHille希勒
	f.希勒Hille.SetParent(f)

	f.奥克尔布Ockelbo = BOckelbo奥克尔布
	f.奥克尔布Ockelbo.SetParent(f)

	f.托绍克尔Torsaker = BTorsaker托绍克尔
	f.托绍克尔Torsaker.SetParent(f)

	f.瓦尔博Valbo = BValbo瓦尔博
	f.瓦尔博Valbo.SetParent(f)

}
