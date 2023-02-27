package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LabourdCounty interface {
    feud.County
    BAire艾尔() 	feud.Barony
    BBayonne巴约讷() 	feud.Barony
    BDax达克斯() 	feud.Barony
    BHasparren阿斯帕伦() 	feud.Barony
    BLabastideclairence克莱朗斯堡() 	feud.Barony
    BSauveterre索沃泰尔() 	feud.Barony
    BSorde索尔德() 	feud.Barony
    BStsever圣瑟韦() 	feud.Barony
}

type 达克斯LabourdCounty struct {
	feud.BaseCounty
	艾尔Aire 	feud.Barony
	巴约讷Bayonne 	feud.Barony
	达克斯Dax 	feud.Barony
	阿斯帕伦Hasparren 	feud.Barony
	克莱朗斯堡Labastideclairence 	feud.Barony
	索沃泰尔Sauveterre 	feud.Barony
	索尔德Sorde 	feud.Barony
	圣瑟韦Stsever 	feud.Barony
}

func (c *达克斯LabourdCounty) BAire艾尔() feud.Barony {
	return c.艾尔Aire
}
    
func (c *达克斯LabourdCounty) BBayonne巴约讷() feud.Barony {
	return c.巴约讷Bayonne
}
    
func (c *达克斯LabourdCounty) BDax达克斯() feud.Barony {
	return c.达克斯Dax
}
    
func (c *达克斯LabourdCounty) BHasparren阿斯帕伦() feud.Barony {
	return c.阿斯帕伦Hasparren
}
    
func (c *达克斯LabourdCounty) BLabastideclairence克莱朗斯堡() feud.Barony {
	return c.克莱朗斯堡Labastideclairence
}
    
func (c *达克斯LabourdCounty) BSauveterre索沃泰尔() feud.Barony {
	return c.索沃泰尔Sauveterre
}
    
func (c *达克斯LabourdCounty) BSorde索尔德() feud.Barony {
	return c.索尔德Sorde
}
    
func (c *达克斯LabourdCounty) BStsever圣瑟韦() feud.Barony {
	return c.圣瑟韦Stsever
}
    
var CLabourd达克斯 LabourdCounty = &达克斯LabourdCounty{}

func init() {
	f := CLabourd达克斯.(*达克斯LabourdCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "151",
		Title:     "labourd",
		TitleName: "达克斯",
		TitleCode: "c_labourd",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾尔Aire = BAire艾尔
	f.艾尔Aire.SetParent(f)
	
	f.巴约讷Bayonne = BBayonne巴约讷
	f.巴约讷Bayonne.SetParent(f)
	
	f.达克斯Dax = BDax达克斯
	f.达克斯Dax.SetParent(f)
	
	f.阿斯帕伦Hasparren = BHasparren阿斯帕伦
	f.阿斯帕伦Hasparren.SetParent(f)
	
	f.克莱朗斯堡Labastideclairence = BLabastideclairence克莱朗斯堡
	f.克莱朗斯堡Labastideclairence.SetParent(f)
	
	f.索沃泰尔Sauveterre = BSauveterre索沃泰尔
	f.索沃泰尔Sauveterre.SetParent(f)
	
	f.索尔德Sorde = BSorde索尔德
	f.索尔德Sorde.SetParent(f)
	
	f.圣瑟韦Stsever = BStsever圣瑟韦
	f.圣瑟韦Stsever.SetParent(f)
	
}
