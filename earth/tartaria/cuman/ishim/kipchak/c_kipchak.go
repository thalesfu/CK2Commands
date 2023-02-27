package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KipchakCounty interface {
    feud.County
    BAmankaragaj阿曼卡拉盖() 	feud.Barony
    BKipchak钦察() 	feud.Barony
    BKipgaj基普加伊() 	feud.Barony
    BKusma库斯马() 	feud.Barony
    BKusmuryn库斯穆伦() 	feud.Barony
    BZhailma扎伊尔马() 	feud.Barony
    BZhitikara杰特加拉() 	feud.Barony
}

type 田吉兹KipchakCounty struct {
	feud.BaseCounty
	阿曼卡拉盖Amankaragaj 	feud.Barony
	钦察Kipchak 	feud.Barony
	基普加伊Kipgaj 	feud.Barony
	库斯马Kusma 	feud.Barony
	库斯穆伦Kusmuryn 	feud.Barony
	扎伊尔马Zhailma 	feud.Barony
	杰特加拉Zhitikara 	feud.Barony
}

func (c *田吉兹KipchakCounty) BAmankaragaj阿曼卡拉盖() feud.Barony {
	return c.阿曼卡拉盖Amankaragaj
}
    
func (c *田吉兹KipchakCounty) BKipchak钦察() feud.Barony {
	return c.钦察Kipchak
}
    
func (c *田吉兹KipchakCounty) BKipgaj基普加伊() feud.Barony {
	return c.基普加伊Kipgaj
}
    
func (c *田吉兹KipchakCounty) BKusma库斯马() feud.Barony {
	return c.库斯马Kusma
}
    
func (c *田吉兹KipchakCounty) BKusmuryn库斯穆伦() feud.Barony {
	return c.库斯穆伦Kusmuryn
}
    
func (c *田吉兹KipchakCounty) BZhailma扎伊尔马() feud.Barony {
	return c.扎伊尔马Zhailma
}
    
func (c *田吉兹KipchakCounty) BZhitikara杰特加拉() feud.Barony {
	return c.杰特加拉Zhitikara
}
    
var CKipchak田吉兹 KipchakCounty = &田吉兹KipchakCounty{}

func init() {
	f := CKipchak田吉兹.(*田吉兹KipchakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1430",
		Title:     "kipchak",
		TitleName: "田吉兹",
		TitleCode: "c_kipchak",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿曼卡拉盖Amankaragaj = BAmankaragaj阿曼卡拉盖
	f.阿曼卡拉盖Amankaragaj.SetParent(f)
	
	f.钦察Kipchak = BKipchak钦察
	f.钦察Kipchak.SetParent(f)
	
	f.基普加伊Kipgaj = BKipgaj基普加伊
	f.基普加伊Kipgaj.SetParent(f)
	
	f.库斯马Kusma = BKusma库斯马
	f.库斯马Kusma.SetParent(f)
	
	f.库斯穆伦Kusmuryn = BKusmuryn库斯穆伦
	f.库斯穆伦Kusmuryn.SetParent(f)
	
	f.扎伊尔马Zhailma = BZhailma扎伊尔马
	f.扎伊尔马Zhailma.SetParent(f)
	
	f.杰特加拉Zhitikara = BZhitikara杰特加拉
	f.杰特加拉Zhitikara.SetParent(f)
	
}
