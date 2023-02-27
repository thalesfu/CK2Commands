package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RomsdalCounty interface {
    feud.County
    BHelland赫兰() 	feud.Barony
    BHustad许斯塔() 	feud.Barony
    BMolde莫尔德() 	feud.Barony
    BRodven勒德文() 	feud.Barony
    BSund_romsdal松() 	feud.Barony
    BVeoy维伊() 	feud.Barony
    BVoll沃尔() 	feud.Barony
}

type 鲁姆斯达尔RomsdalCounty struct {
	feud.BaseCounty
	赫兰Helland 	feud.Barony
	许斯塔Hustad 	feud.Barony
	莫尔德Molde 	feud.Barony
	勒德文Rodven 	feud.Barony
	松Sund_romsdal 	feud.Barony
	维伊Veoy 	feud.Barony
	沃尔Voll 	feud.Barony
}

func (c *鲁姆斯达尔RomsdalCounty) BHelland赫兰() feud.Barony {
	return c.赫兰Helland
}
    
func (c *鲁姆斯达尔RomsdalCounty) BHustad许斯塔() feud.Barony {
	return c.许斯塔Hustad
}
    
func (c *鲁姆斯达尔RomsdalCounty) BMolde莫尔德() feud.Barony {
	return c.莫尔德Molde
}
    
func (c *鲁姆斯达尔RomsdalCounty) BRodven勒德文() feud.Barony {
	return c.勒德文Rodven
}
    
func (c *鲁姆斯达尔RomsdalCounty) BSund_romsdal松() feud.Barony {
	return c.松Sund_romsdal
}
    
func (c *鲁姆斯达尔RomsdalCounty) BVeoy维伊() feud.Barony {
	return c.维伊Veoy
}
    
func (c *鲁姆斯达尔RomsdalCounty) BVoll沃尔() feud.Barony {
	return c.沃尔Voll
}
    
var CRomsdal鲁姆斯达尔 RomsdalCounty = &鲁姆斯达尔RomsdalCounty{}

func init() {
	f := CRomsdal鲁姆斯达尔.(*鲁姆斯达尔RomsdalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1622",
		Title:     "romsdal",
		TitleName: "鲁姆斯达尔",
		TitleCode: "c_romsdal",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫兰Helland = BHelland赫兰
	f.赫兰Helland.SetParent(f)
	
	f.许斯塔Hustad = BHustad许斯塔
	f.许斯塔Hustad.SetParent(f)
	
	f.莫尔德Molde = BMolde莫尔德
	f.莫尔德Molde.SetParent(f)
	
	f.勒德文Rodven = BRodven勒德文
	f.勒德文Rodven.SetParent(f)
	
	f.松Sund_romsdal = BSund_romsdal松
	f.松Sund_romsdal.SetParent(f)
	
	f.维伊Veoy = BVeoy维伊
	f.维伊Veoy.SetParent(f)
	
	f.沃尔Voll = BVoll沃尔
	f.沃尔Voll.SetParent(f)
	
}
