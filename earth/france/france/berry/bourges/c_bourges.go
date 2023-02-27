package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BourgesCounty interface {
    feud.County
    BBourges布尔日() 	feud.Barony
    BChateauroux沙托鲁() 	feud.Barony
    BDeols代奥勒() 	feud.Barony
    BIssoudun伊苏丹() 	feud.Barony
    BMehun默安() 	feud.Barony
    BStsatur圣萨图尔() 	feud.Barony
    BVierzon维耶尔宗() 	feud.Barony
}

type 布尔日BourgesCounty struct {
	feud.BaseCounty
	布尔日Bourges 	feud.Barony
	沙托鲁Chateauroux 	feud.Barony
	代奥勒Deols 	feud.Barony
	伊苏丹Issoudun 	feud.Barony
	默安Mehun 	feud.Barony
	圣萨图尔Stsatur 	feud.Barony
	维耶尔宗Vierzon 	feud.Barony
}

func (c *布尔日BourgesCounty) BBourges布尔日() feud.Barony {
	return c.布尔日Bourges
}
    
func (c *布尔日BourgesCounty) BChateauroux沙托鲁() feud.Barony {
	return c.沙托鲁Chateauroux
}
    
func (c *布尔日BourgesCounty) BDeols代奥勒() feud.Barony {
	return c.代奥勒Deols
}
    
func (c *布尔日BourgesCounty) BIssoudun伊苏丹() feud.Barony {
	return c.伊苏丹Issoudun
}
    
func (c *布尔日BourgesCounty) BMehun默安() feud.Barony {
	return c.默安Mehun
}
    
func (c *布尔日BourgesCounty) BStsatur圣萨图尔() feud.Barony {
	return c.圣萨图尔Stsatur
}
    
func (c *布尔日BourgesCounty) BVierzon维耶尔宗() feud.Barony {
	return c.维耶尔宗Vierzon
}
    
var CBourges布尔日 BourgesCounty = &布尔日BourgesCounty{}

func init() {
	f := CBourges布尔日.(*布尔日BourgesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "139",
		Title:     "bourges",
		TitleName: "布尔日",
		TitleCode: "c_bourges",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔日Bourges = BBourges布尔日
	f.布尔日Bourges.SetParent(f)
	
	f.沙托鲁Chateauroux = BChateauroux沙托鲁
	f.沙托鲁Chateauroux.SetParent(f)
	
	f.代奥勒Deols = BDeols代奥勒
	f.代奥勒Deols.SetParent(f)
	
	f.伊苏丹Issoudun = BIssoudun伊苏丹
	f.伊苏丹Issoudun.SetParent(f)
	
	f.默安Mehun = BMehun默安
	f.默安Mehun.SetParent(f)
	
	f.圣萨图尔Stsatur = BStsatur圣萨图尔
	f.圣萨图尔Stsatur.SetParent(f)
	
	f.维耶尔宗Vierzon = BVierzon维耶尔宗
	f.维耶尔宗Vierzon.SetParent(f)
	
}
