package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AralCounty interface {
    feud.County
    BAszczeatrik阿谢_阿特里克() 	feud.Barony
    BDawletgirei达夫列特_吉雷() 	feud.Barony
    BKarakul卡拉库尔() 	feud.Barony
    BKassarma卡萨尔马() 	feud.Barony
    BKaszkarata卡什卡拉塔() 	feud.Barony
    BKosbulak科斯_布拉克() 	feud.Barony
    BKuljandi库良季() 	feud.Barony
    BSokyrbulak索克尔_布拉克() 	feud.Barony
}

type 库坦布拉克AralCounty struct {
	feud.BaseCounty
	阿谢_阿特里克Aszczeatrik 	feud.Barony
	达夫列特_吉雷Dawletgirei 	feud.Barony
	卡拉库尔Karakul 	feud.Barony
	卡萨尔马Kassarma 	feud.Barony
	卡什卡拉塔Kaszkarata 	feud.Barony
	科斯_布拉克Kosbulak 	feud.Barony
	库良季Kuljandi 	feud.Barony
	索克尔_布拉克Sokyrbulak 	feud.Barony
}

func (c *库坦布拉克AralCounty) BAszczeatrik阿谢_阿特里克() feud.Barony {
	return c.阿谢_阿特里克Aszczeatrik
}
    
func (c *库坦布拉克AralCounty) BDawletgirei达夫列特_吉雷() feud.Barony {
	return c.达夫列特_吉雷Dawletgirei
}
    
func (c *库坦布拉克AralCounty) BKarakul卡拉库尔() feud.Barony {
	return c.卡拉库尔Karakul
}
    
func (c *库坦布拉克AralCounty) BKassarma卡萨尔马() feud.Barony {
	return c.卡萨尔马Kassarma
}
    
func (c *库坦布拉克AralCounty) BKaszkarata卡什卡拉塔() feud.Barony {
	return c.卡什卡拉塔Kaszkarata
}
    
func (c *库坦布拉克AralCounty) BKosbulak科斯_布拉克() feud.Barony {
	return c.科斯_布拉克Kosbulak
}
    
func (c *库坦布拉克AralCounty) BKuljandi库良季() feud.Barony {
	return c.库良季Kuljandi
}
    
func (c *库坦布拉克AralCounty) BSokyrbulak索克尔_布拉克() feud.Barony {
	return c.索克尔_布拉克Sokyrbulak
}
    
var CAral库坦布拉克 AralCounty = &库坦布拉克AralCounty{}

func init() {
	f := CAral库坦布拉克.(*库坦布拉克AralCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "623",
		Title:     "aral",
		TitleName: "库坦布拉克",
		TitleCode: "c_aral",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿谢_阿特里克Aszczeatrik = BAszczeatrik阿谢_阿特里克
	f.阿谢_阿特里克Aszczeatrik.SetParent(f)
	
	f.达夫列特_吉雷Dawletgirei = BDawletgirei达夫列特_吉雷
	f.达夫列特_吉雷Dawletgirei.SetParent(f)
	
	f.卡拉库尔Karakul = BKarakul卡拉库尔
	f.卡拉库尔Karakul.SetParent(f)
	
	f.卡萨尔马Kassarma = BKassarma卡萨尔马
	f.卡萨尔马Kassarma.SetParent(f)
	
	f.卡什卡拉塔Kaszkarata = BKaszkarata卡什卡拉塔
	f.卡什卡拉塔Kaszkarata.SetParent(f)
	
	f.科斯_布拉克Kosbulak = BKosbulak科斯_布拉克
	f.科斯_布拉克Kosbulak.SetParent(f)
	
	f.库良季Kuljandi = BKuljandi库良季
	f.库良季Kuljandi.SetParent(f)
	
	f.索克尔_布拉克Sokyrbulak = BSokyrbulak索克尔_布拉克
	f.索克尔_布拉克Sokyrbulak.SetParent(f)
	
}
