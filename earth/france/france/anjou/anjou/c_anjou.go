package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnjouCounty interface {
    feud.County
    BAngers昂热() 	feud.Barony
    BChateaugontier贡捷堡() 	feud.Barony
    BGraon格朗() 	feud.Barony
    BLude吕德() 	feud.Barony
    BMontsoreau蒙索罗() 	feud.Barony
    BTreves特雷沃() 	feud.Barony
    BVihiers维耶() 	feud.Barony
}

type 安茹AnjouCounty struct {
	feud.BaseCounty
	昂热Angers 	feud.Barony
	贡捷堡Chateaugontier 	feud.Barony
	格朗Graon 	feud.Barony
	吕德Lude 	feud.Barony
	蒙索罗Montsoreau 	feud.Barony
	特雷沃Treves 	feud.Barony
	维耶Vihiers 	feud.Barony
}

func (c *安茹AnjouCounty) BAngers昂热() feud.Barony {
	return c.昂热Angers
}
    
func (c *安茹AnjouCounty) BChateaugontier贡捷堡() feud.Barony {
	return c.贡捷堡Chateaugontier
}
    
func (c *安茹AnjouCounty) BGraon格朗() feud.Barony {
	return c.格朗Graon
}
    
func (c *安茹AnjouCounty) BLude吕德() feud.Barony {
	return c.吕德Lude
}
    
func (c *安茹AnjouCounty) BMontsoreau蒙索罗() feud.Barony {
	return c.蒙索罗Montsoreau
}
    
func (c *安茹AnjouCounty) BTreves特雷沃() feud.Barony {
	return c.特雷沃Treves
}
    
func (c *安茹AnjouCounty) BVihiers维耶() feud.Barony {
	return c.维耶Vihiers
}
    
var CAnjou安茹 AnjouCounty = &安茹AnjouCounty{}

func init() {
	f := CAnjou安茹.(*安茹AnjouCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "107",
		Title:     "anjou",
		TitleName: "安茹",
		TitleCode: "c_anjou",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂热Angers = BAngers昂热
	f.昂热Angers.SetParent(f)
	
	f.贡捷堡Chateaugontier = BChateaugontier贡捷堡
	f.贡捷堡Chateaugontier.SetParent(f)
	
	f.格朗Graon = BGraon格朗
	f.格朗Graon.SetParent(f)
	
	f.吕德Lude = BLude吕德
	f.吕德Lude.SetParent(f)
	
	f.蒙索罗Montsoreau = BMontsoreau蒙索罗
	f.蒙索罗Montsoreau.SetParent(f)
	
	f.特雷沃Treves = BTreves特雷沃
	f.特雷沃Treves.SetParent(f)
	
	f.维耶Vihiers = BVihiers维耶
	f.维耶Vihiers.SetParent(f)
	
}
