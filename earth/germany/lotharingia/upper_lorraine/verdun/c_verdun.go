package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VerdunCounty interface {
    feud.County
    BBarlecomte巴勒孔特() 	feud.Barony
    BDoaumont杜奥蒙() 	feud.Barony
    BEtain埃坦() 	feud.Barony
    BGrandpre格朗普雷() 	feud.Barony
    BLonguyon隆吉永() 	feud.Barony
    BStenay斯特奈() 	feud.Barony
    BVerdun凡尔登() 	feud.Barony
}

type 凡尔登VerdunCounty struct {
	feud.BaseCounty
	巴勒孔特Barlecomte 	feud.Barony
	杜奥蒙Doaumont 	feud.Barony
	埃坦Etain 	feud.Barony
	格朗普雷Grandpre 	feud.Barony
	隆吉永Longuyon 	feud.Barony
	斯特奈Stenay 	feud.Barony
	凡尔登Verdun 	feud.Barony
}

func (c *凡尔登VerdunCounty) BBarlecomte巴勒孔特() feud.Barony {
	return c.巴勒孔特Barlecomte
}
    
func (c *凡尔登VerdunCounty) BDoaumont杜奥蒙() feud.Barony {
	return c.杜奥蒙Doaumont
}
    
func (c *凡尔登VerdunCounty) BEtain埃坦() feud.Barony {
	return c.埃坦Etain
}
    
func (c *凡尔登VerdunCounty) BGrandpre格朗普雷() feud.Barony {
	return c.格朗普雷Grandpre
}
    
func (c *凡尔登VerdunCounty) BLonguyon隆吉永() feud.Barony {
	return c.隆吉永Longuyon
}
    
func (c *凡尔登VerdunCounty) BStenay斯特奈() feud.Barony {
	return c.斯特奈Stenay
}
    
func (c *凡尔登VerdunCounty) BVerdun凡尔登() feud.Barony {
	return c.凡尔登Verdun
}
    
var CVerdun凡尔登 VerdunCounty = &凡尔登VerdunCounty{}

func init() {
	f := CVerdun凡尔登.(*凡尔登VerdunCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "129",
		Title:     "verdun",
		TitleName: "凡尔登",
		TitleCode: "c_verdun",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴勒孔特Barlecomte = BBarlecomte巴勒孔特
	f.巴勒孔特Barlecomte.SetParent(f)
	
	f.杜奥蒙Doaumont = BDoaumont杜奥蒙
	f.杜奥蒙Doaumont.SetParent(f)
	
	f.埃坦Etain = BEtain埃坦
	f.埃坦Etain.SetParent(f)
	
	f.格朗普雷Grandpre = BGrandpre格朗普雷
	f.格朗普雷Grandpre.SetParent(f)
	
	f.隆吉永Longuyon = BLonguyon隆吉永
	f.隆吉永Longuyon.SetParent(f)
	
	f.斯特奈Stenay = BStenay斯特奈
	f.斯特奈Stenay.SetParent(f)
	
	f.凡尔登Verdun = BVerdun凡尔登
	f.凡尔登Verdun.SetParent(f)
	
}
