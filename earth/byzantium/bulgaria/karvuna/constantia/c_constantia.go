package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ConstantiaCounty interface {
    feud.County
    BAdamclisi阿达姆克利西() 	feud.Barony
    BCarachioi卡拉基奥() 	feud.Barony
    BCobadin科巴丁() 	feud.Barony
    BCogealac科贾拉克() 	feud.Barony
    BConstantia康斯坦察() 	feud.Barony
    BMangalia曼加利亚() 	feud.Barony
    BMesgidia梅吉迪亚() 	feud.Barony
    BTopraisar托普拉伊萨尔() 	feud.Barony
}

type 康斯坦察ConstantiaCounty struct {
	feud.BaseCounty
	阿达姆克利西Adamclisi 	feud.Barony
	卡拉基奥Carachioi 	feud.Barony
	科巴丁Cobadin 	feud.Barony
	科贾拉克Cogealac 	feud.Barony
	康斯坦察Constantia 	feud.Barony
	曼加利亚Mangalia 	feud.Barony
	梅吉迪亚Mesgidia 	feud.Barony
	托普拉伊萨尔Topraisar 	feud.Barony
}

func (c *康斯坦察ConstantiaCounty) BAdamclisi阿达姆克利西() feud.Barony {
	return c.阿达姆克利西Adamclisi
}
    
func (c *康斯坦察ConstantiaCounty) BCarachioi卡拉基奥() feud.Barony {
	return c.卡拉基奥Carachioi
}
    
func (c *康斯坦察ConstantiaCounty) BCobadin科巴丁() feud.Barony {
	return c.科巴丁Cobadin
}
    
func (c *康斯坦察ConstantiaCounty) BCogealac科贾拉克() feud.Barony {
	return c.科贾拉克Cogealac
}
    
func (c *康斯坦察ConstantiaCounty) BConstantia康斯坦察() feud.Barony {
	return c.康斯坦察Constantia
}
    
func (c *康斯坦察ConstantiaCounty) BMangalia曼加利亚() feud.Barony {
	return c.曼加利亚Mangalia
}
    
func (c *康斯坦察ConstantiaCounty) BMesgidia梅吉迪亚() feud.Barony {
	return c.梅吉迪亚Mesgidia
}
    
func (c *康斯坦察ConstantiaCounty) BTopraisar托普拉伊萨尔() feud.Barony {
	return c.托普拉伊萨尔Topraisar
}
    
var CConstantia康斯坦察 ConstantiaCounty = &康斯坦察ConstantiaCounty{}

func init() {
	f := CConstantia康斯坦察.(*康斯坦察ConstantiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "510",
		Title:     "constantia",
		TitleName: "康斯坦察",
		TitleCode: "c_constantia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达姆克利西Adamclisi = BAdamclisi阿达姆克利西
	f.阿达姆克利西Adamclisi.SetParent(f)
	
	f.卡拉基奥Carachioi = BCarachioi卡拉基奥
	f.卡拉基奥Carachioi.SetParent(f)
	
	f.科巴丁Cobadin = BCobadin科巴丁
	f.科巴丁Cobadin.SetParent(f)
	
	f.科贾拉克Cogealac = BCogealac科贾拉克
	f.科贾拉克Cogealac.SetParent(f)
	
	f.康斯坦察Constantia = BConstantia康斯坦察
	f.康斯坦察Constantia.SetParent(f)
	
	f.曼加利亚Mangalia = BMangalia曼加利亚
	f.曼加利亚Mangalia.SetParent(f)
	
	f.梅吉迪亚Mesgidia = BMesgidia梅吉迪亚
	f.梅吉迪亚Mesgidia.SetParent(f)
	
	f.托普拉伊萨尔Topraisar = BTopraisar托普拉伊萨尔
	f.托普拉伊萨尔Topraisar.SetParent(f)
	
}
