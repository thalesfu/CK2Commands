package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SmolenskCounty interface {
    feud.County
    BGnyozdovo格涅兹多沃() 	feud.Barony
    BKrasny克拉斯内() 	feud.Barony
    BNetrizovo涅特里佐沃() 	feud.Barony
    BRay赖() 	feud.Barony
    BSmolensk斯摩棱斯克() 	feud.Barony
    BVelizh韦利日() 	feud.Barony
    BYartsevo亚尔采沃() 	feud.Barony
}

type 斯摩棱斯克SmolenskCounty struct {
	feud.BaseCounty
	格涅兹多沃Gnyozdovo 	feud.Barony
	克拉斯内Krasny 	feud.Barony
	涅特里佐沃Netrizovo 	feud.Barony
	赖Ray 	feud.Barony
	斯摩棱斯克Smolensk 	feud.Barony
	韦利日Velizh 	feud.Barony
	亚尔采沃Yartsevo 	feud.Barony
}

func (c *斯摩棱斯克SmolenskCounty) BGnyozdovo格涅兹多沃() feud.Barony {
	return c.格涅兹多沃Gnyozdovo
}
    
func (c *斯摩棱斯克SmolenskCounty) BKrasny克拉斯内() feud.Barony {
	return c.克拉斯内Krasny
}
    
func (c *斯摩棱斯克SmolenskCounty) BNetrizovo涅特里佐沃() feud.Barony {
	return c.涅特里佐沃Netrizovo
}
    
func (c *斯摩棱斯克SmolenskCounty) BRay赖() feud.Barony {
	return c.赖Ray
}
    
func (c *斯摩棱斯克SmolenskCounty) BSmolensk斯摩棱斯克() feud.Barony {
	return c.斯摩棱斯克Smolensk
}
    
func (c *斯摩棱斯克SmolenskCounty) BVelizh韦利日() feud.Barony {
	return c.韦利日Velizh
}
    
func (c *斯摩棱斯克SmolenskCounty) BYartsevo亚尔采沃() feud.Barony {
	return c.亚尔采沃Yartsevo
}
    
var CSmolensk斯摩棱斯克 SmolenskCounty = &斯摩棱斯克SmolenskCounty{}

func init() {
	f := CSmolensk斯摩棱斯克.(*斯摩棱斯克SmolenskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "568",
		Title:     "smolensk",
		TitleName: "斯摩棱斯克",
		TitleCode: "c_smolensk",
		Baronies:  map[string]feud.Barony{},
	}

	f.格涅兹多沃Gnyozdovo = BGnyozdovo格涅兹多沃
	f.格涅兹多沃Gnyozdovo.SetParent(f)
	
	f.克拉斯内Krasny = BKrasny克拉斯内
	f.克拉斯内Krasny.SetParent(f)
	
	f.涅特里佐沃Netrizovo = BNetrizovo涅特里佐沃
	f.涅特里佐沃Netrizovo.SetParent(f)
	
	f.赖Ray = BRay赖
	f.赖Ray.SetParent(f)
	
	f.斯摩棱斯克Smolensk = BSmolensk斯摩棱斯克
	f.斯摩棱斯克Smolensk.SetParent(f)
	
	f.韦利日Velizh = BVelizh韦利日
	f.韦利日Velizh.SetParent(f)
	
	f.亚尔采沃Yartsevo = BYartsevo亚尔采沃
	f.亚尔采沃Yartsevo.SetParent(f)
	
}
