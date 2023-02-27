package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CinarcaCounty interface {
    feud.County
    BAjaccio阿雅克肖() 	feud.Barony
    BBastelicaccia巴斯特利卡西亚() 	feud.Barony
    BBonifacio博尼法乔() 	feud.Barony
    BCinarca奇纳尔卡() 	feud.Barony
    BFilitosa菲利托萨() 	feud.Barony
    BPortevecchio韦基奥港() 	feud.Barony
    BPropriano普罗普里亚诺() 	feud.Barony
    BSartene萨尔泰讷() 	feud.Barony
}

type 奇纳尔卡CinarcaCounty struct {
	feud.BaseCounty
	阿雅克肖Ajaccio 	feud.Barony
	巴斯特利卡西亚Bastelicaccia 	feud.Barony
	博尼法乔Bonifacio 	feud.Barony
	奇纳尔卡Cinarca 	feud.Barony
	菲利托萨Filitosa 	feud.Barony
	韦基奥港Portevecchio 	feud.Barony
	普罗普里亚诺Propriano 	feud.Barony
	萨尔泰讷Sartene 	feud.Barony
}

func (c *奇纳尔卡CinarcaCounty) BAjaccio阿雅克肖() feud.Barony {
	return c.阿雅克肖Ajaccio
}
    
func (c *奇纳尔卡CinarcaCounty) BBastelicaccia巴斯特利卡西亚() feud.Barony {
	return c.巴斯特利卡西亚Bastelicaccia
}
    
func (c *奇纳尔卡CinarcaCounty) BBonifacio博尼法乔() feud.Barony {
	return c.博尼法乔Bonifacio
}
    
func (c *奇纳尔卡CinarcaCounty) BCinarca奇纳尔卡() feud.Barony {
	return c.奇纳尔卡Cinarca
}
    
func (c *奇纳尔卡CinarcaCounty) BFilitosa菲利托萨() feud.Barony {
	return c.菲利托萨Filitosa
}
    
func (c *奇纳尔卡CinarcaCounty) BPortevecchio韦基奥港() feud.Barony {
	return c.韦基奥港Portevecchio
}
    
func (c *奇纳尔卡CinarcaCounty) BPropriano普罗普里亚诺() feud.Barony {
	return c.普罗普里亚诺Propriano
}
    
func (c *奇纳尔卡CinarcaCounty) BSartene萨尔泰讷() feud.Barony {
	return c.萨尔泰讷Sartene
}
    
var CCinarca奇纳尔卡 CinarcaCounty = &奇纳尔卡CinarcaCounty{}

func init() {
	f := CCinarca奇纳尔卡.(*奇纳尔卡CinarcaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1577",
		Title:     "cinarca",
		TitleName: "奇纳尔卡",
		TitleCode: "c_cinarca",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿雅克肖Ajaccio = BAjaccio阿雅克肖
	f.阿雅克肖Ajaccio.SetParent(f)
	
	f.巴斯特利卡西亚Bastelicaccia = BBastelicaccia巴斯特利卡西亚
	f.巴斯特利卡西亚Bastelicaccia.SetParent(f)
	
	f.博尼法乔Bonifacio = BBonifacio博尼法乔
	f.博尼法乔Bonifacio.SetParent(f)
	
	f.奇纳尔卡Cinarca = BCinarca奇纳尔卡
	f.奇纳尔卡Cinarca.SetParent(f)
	
	f.菲利托萨Filitosa = BFilitosa菲利托萨
	f.菲利托萨Filitosa.SetParent(f)
	
	f.韦基奥港Portevecchio = BPortevecchio韦基奥港
	f.韦基奥港Portevecchio.SetParent(f)
	
	f.普罗普里亚诺Propriano = BPropriano普罗普里亚诺
	f.普罗普里亚诺Propriano.SetParent(f)
	
	f.萨尔泰讷Sartene = BSartene萨尔泰讷
	f.萨尔泰讷Sartene.SetParent(f)
	
}
