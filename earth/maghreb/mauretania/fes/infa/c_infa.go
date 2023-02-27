package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type InfaCounty interface {
    feud.County
    BAzemmour艾宰穆尔() 	feud.Barony
    BBerrechid拜赖希德() 	feud.Barony
    BInfa安法() 	feud.Barony
    BKhouribga胡里卜盖() 	feud.Barony
    BOulmes乌勒马斯() 	feud.Barony
    BRabat拉巴特() 	feud.Barony
    BSale塞拉() 	feud.Barony
    BTiflet提夫莱特() 	feud.Barony
}

type 安法InfaCounty struct {
	feud.BaseCounty
	艾宰穆尔Azemmour 	feud.Barony
	拜赖希德Berrechid 	feud.Barony
	安法Infa 	feud.Barony
	胡里卜盖Khouribga 	feud.Barony
	乌勒马斯Oulmes 	feud.Barony
	拉巴特Rabat 	feud.Barony
	塞拉Sale 	feud.Barony
	提夫莱特Tiflet 	feud.Barony
}

func (c *安法InfaCounty) BAzemmour艾宰穆尔() feud.Barony {
	return c.艾宰穆尔Azemmour
}
    
func (c *安法InfaCounty) BBerrechid拜赖希德() feud.Barony {
	return c.拜赖希德Berrechid
}
    
func (c *安法InfaCounty) BInfa安法() feud.Barony {
	return c.安法Infa
}
    
func (c *安法InfaCounty) BKhouribga胡里卜盖() feud.Barony {
	return c.胡里卜盖Khouribga
}
    
func (c *安法InfaCounty) BOulmes乌勒马斯() feud.Barony {
	return c.乌勒马斯Oulmes
}
    
func (c *安法InfaCounty) BRabat拉巴特() feud.Barony {
	return c.拉巴特Rabat
}
    
func (c *安法InfaCounty) BSale塞拉() feud.Barony {
	return c.塞拉Sale
}
    
func (c *安法InfaCounty) BTiflet提夫莱特() feud.Barony {
	return c.提夫莱特Tiflet
}
    
var CInfa安法 InfaCounty = &安法InfaCounty{}

func init() {
	f := CInfa安法.(*安法InfaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "842",
		Title:     "infa",
		TitleName: "安法",
		TitleCode: "c_infa",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾宰穆尔Azemmour = BAzemmour艾宰穆尔
	f.艾宰穆尔Azemmour.SetParent(f)
	
	f.拜赖希德Berrechid = BBerrechid拜赖希德
	f.拜赖希德Berrechid.SetParent(f)
	
	f.安法Infa = BInfa安法
	f.安法Infa.SetParent(f)
	
	f.胡里卜盖Khouribga = BKhouribga胡里卜盖
	f.胡里卜盖Khouribga.SetParent(f)
	
	f.乌勒马斯Oulmes = BOulmes乌勒马斯
	f.乌勒马斯Oulmes.SetParent(f)
	
	f.拉巴特Rabat = BRabat拉巴特
	f.拉巴特Rabat.SetParent(f)
	
	f.塞拉Sale = BSale塞拉
	f.塞拉Sale.SetParent(f)
	
	f.提夫莱特Tiflet = BTiflet提夫莱特
	f.提夫莱特Tiflet.SetParent(f)
	
}
