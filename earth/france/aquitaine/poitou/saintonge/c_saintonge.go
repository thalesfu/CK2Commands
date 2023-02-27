package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaintongeCounty interface {
    feud.County
    BAulnay欧奈() 	feud.Barony
    BMontguyon蒙吉永() 	feud.Barony
    BRoyan鲁瓦扬() 	feud.Barony
    BSaintes桑特() 	feud.Barony
    BStjeandangely圣让当热利() 	feud.Barony
    BTaillebourg塔耶堡() 	feud.Barony
    BTonnay托奈() 	feud.Barony
    BVilleneuve维尔讷夫() 	feud.Barony
}

type 圣通日SaintongeCounty struct {
	feud.BaseCounty
	欧奈Aulnay 	feud.Barony
	蒙吉永Montguyon 	feud.Barony
	鲁瓦扬Royan 	feud.Barony
	桑特Saintes 	feud.Barony
	圣让当热利Stjeandangely 	feud.Barony
	塔耶堡Taillebourg 	feud.Barony
	托奈Tonnay 	feud.Barony
	维尔讷夫Villeneuve 	feud.Barony
}

func (c *圣通日SaintongeCounty) BAulnay欧奈() feud.Barony {
	return c.欧奈Aulnay
}
    
func (c *圣通日SaintongeCounty) BMontguyon蒙吉永() feud.Barony {
	return c.蒙吉永Montguyon
}
    
func (c *圣通日SaintongeCounty) BRoyan鲁瓦扬() feud.Barony {
	return c.鲁瓦扬Royan
}
    
func (c *圣通日SaintongeCounty) BSaintes桑特() feud.Barony {
	return c.桑特Saintes
}
    
func (c *圣通日SaintongeCounty) BStjeandangely圣让当热利() feud.Barony {
	return c.圣让当热利Stjeandangely
}
    
func (c *圣通日SaintongeCounty) BTaillebourg塔耶堡() feud.Barony {
	return c.塔耶堡Taillebourg
}
    
func (c *圣通日SaintongeCounty) BTonnay托奈() feud.Barony {
	return c.托奈Tonnay
}
    
func (c *圣通日SaintongeCounty) BVilleneuve维尔讷夫() feud.Barony {
	return c.维尔讷夫Villeneuve
}
    
var CSaintonge圣通日 SaintongeCounty = &圣通日SaintongeCounty{}

func init() {
	f := CSaintonge圣通日.(*圣通日SaintongeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "143",
		Title:     "saintonge",
		TitleName: "圣通日",
		TitleCode: "c_saintonge",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧奈Aulnay = BAulnay欧奈
	f.欧奈Aulnay.SetParent(f)
	
	f.蒙吉永Montguyon = BMontguyon蒙吉永
	f.蒙吉永Montguyon.SetParent(f)
	
	f.鲁瓦扬Royan = BRoyan鲁瓦扬
	f.鲁瓦扬Royan.SetParent(f)
	
	f.桑特Saintes = BSaintes桑特
	f.桑特Saintes.SetParent(f)
	
	f.圣让当热利Stjeandangely = BStjeandangely圣让当热利
	f.圣让当热利Stjeandangely.SetParent(f)
	
	f.塔耶堡Taillebourg = BTaillebourg塔耶堡
	f.塔耶堡Taillebourg.SetParent(f)
	
	f.托奈Tonnay = BTonnay托奈
	f.托奈Tonnay.SetParent(f)
	
	f.维尔讷夫Villeneuve = BVilleneuve维尔讷夫
	f.维尔讷夫Villeneuve.SetParent(f)
	
}
