package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_mafraqCounty interface {
    feud.County
    BBuwayda布未达() 	feud.Barony
    BElemtaih艾勒穆塔() 	feud.Barony
    BJarash杰拉什() 	feud.Barony
    BMafraq马弗拉克() 	feud.Barony
    BNasib奈西卜() 	feud.Barony
    BRamtah赖姆萨() 	feud.Barony
    BTaebah塔埃巴() 	feud.Barony
    BThughra特胡哈拉() 	feud.Barony
}

type 马弗拉克Al_mafraqCounty struct {
	feud.BaseCounty
	布未达Buwayda 	feud.Barony
	艾勒穆塔Elemtaih 	feud.Barony
	杰拉什Jarash 	feud.Barony
	马弗拉克Mafraq 	feud.Barony
	奈西卜Nasib 	feud.Barony
	赖姆萨Ramtah 	feud.Barony
	塔埃巴Taebah 	feud.Barony
	特胡哈拉Thughra 	feud.Barony
}

func (c *马弗拉克Al_mafraqCounty) BBuwayda布未达() feud.Barony {
	return c.布未达Buwayda
}
    
func (c *马弗拉克Al_mafraqCounty) BElemtaih艾勒穆塔() feud.Barony {
	return c.艾勒穆塔Elemtaih
}
    
func (c *马弗拉克Al_mafraqCounty) BJarash杰拉什() feud.Barony {
	return c.杰拉什Jarash
}
    
func (c *马弗拉克Al_mafraqCounty) BMafraq马弗拉克() feud.Barony {
	return c.马弗拉克Mafraq
}
    
func (c *马弗拉克Al_mafraqCounty) BNasib奈西卜() feud.Barony {
	return c.奈西卜Nasib
}
    
func (c *马弗拉克Al_mafraqCounty) BRamtah赖姆萨() feud.Barony {
	return c.赖姆萨Ramtah
}
    
func (c *马弗拉克Al_mafraqCounty) BTaebah塔埃巴() feud.Barony {
	return c.塔埃巴Taebah
}
    
func (c *马弗拉克Al_mafraqCounty) BThughra特胡哈拉() feud.Barony {
	return c.特胡哈拉Thughra
}
    
var CAl_mafraq马弗拉克 Al_mafraqCounty = &马弗拉克Al_mafraqCounty{}

func init() {
	f := CAl_mafraq马弗拉克.(*马弗拉克Al_mafraqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "726",
		Title:     "al_mafraq",
		TitleName: "马弗拉克",
		TitleCode: "c_al_mafraq",
		Baronies:  map[string]feud.Barony{},
	}

	f.布未达Buwayda = BBuwayda布未达
	f.布未达Buwayda.SetParent(f)
	
	f.艾勒穆塔Elemtaih = BElemtaih艾勒穆塔
	f.艾勒穆塔Elemtaih.SetParent(f)
	
	f.杰拉什Jarash = BJarash杰拉什
	f.杰拉什Jarash.SetParent(f)
	
	f.马弗拉克Mafraq = BMafraq马弗拉克
	f.马弗拉克Mafraq.SetParent(f)
	
	f.奈西卜Nasib = BNasib奈西卜
	f.奈西卜Nasib.SetParent(f)
	
	f.赖姆萨Ramtah = BRamtah赖姆萨
	f.赖姆萨Ramtah.SetParent(f)
	
	f.塔埃巴Taebah = BTaebah塔埃巴
	f.塔埃巴Taebah.SetParent(f)
	
	f.特胡哈拉Thughra = BThughra特胡哈拉
	f.特胡哈拉Thughra.SetParent(f)
	
}
