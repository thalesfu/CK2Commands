package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KlarjetiCounty interface {
    feud.County
    BAncha安查() 	feud.Barony
    BArtaani阿尔达汉() 	feud.Barony
    BArtanuji阿尔达努奇() 	feud.Barony
    BKhertvisi赫尔特维西() 	feud.Barony
    BKlarjeti克拉尔哲季() 	feud.Barony
    BTbeti特贝蒂() 	feud.Barony
    BTseptha采普萨() 	feud.Barony
}

type 克拉尔哲季KlarjetiCounty struct {
	feud.BaseCounty
	安查Ancha 	feud.Barony
	阿尔达汉Artaani 	feud.Barony
	阿尔达努奇Artanuji 	feud.Barony
	赫尔特维西Khertvisi 	feud.Barony
	克拉尔哲季Klarjeti 	feud.Barony
	特贝蒂Tbeti 	feud.Barony
	采普萨Tseptha 	feud.Barony
}

func (c *克拉尔哲季KlarjetiCounty) BAncha安查() feud.Barony {
	return c.安查Ancha
}
    
func (c *克拉尔哲季KlarjetiCounty) BArtaani阿尔达汉() feud.Barony {
	return c.阿尔达汉Artaani
}
    
func (c *克拉尔哲季KlarjetiCounty) BArtanuji阿尔达努奇() feud.Barony {
	return c.阿尔达努奇Artanuji
}
    
func (c *克拉尔哲季KlarjetiCounty) BKhertvisi赫尔特维西() feud.Barony {
	return c.赫尔特维西Khertvisi
}
    
func (c *克拉尔哲季KlarjetiCounty) BKlarjeti克拉尔哲季() feud.Barony {
	return c.克拉尔哲季Klarjeti
}
    
func (c *克拉尔哲季KlarjetiCounty) BTbeti特贝蒂() feud.Barony {
	return c.特贝蒂Tbeti
}
    
func (c *克拉尔哲季KlarjetiCounty) BTseptha采普萨() feud.Barony {
	return c.采普萨Tseptha
}
    
var CKlarjeti克拉尔哲季 KlarjetiCounty = &克拉尔哲季KlarjetiCounty{}

func init() {
	f := CKlarjeti克拉尔哲季.(*克拉尔哲季KlarjetiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1807",
		Title:     "klarjeti",
		TitleName: "克拉尔哲季",
		TitleCode: "c_klarjeti",
		Baronies:  map[string]feud.Barony{},
	}

	f.安查Ancha = BAncha安查
	f.安查Ancha.SetParent(f)
	
	f.阿尔达汉Artaani = BArtaani阿尔达汉
	f.阿尔达汉Artaani.SetParent(f)
	
	f.阿尔达努奇Artanuji = BArtanuji阿尔达努奇
	f.阿尔达努奇Artanuji.SetParent(f)
	
	f.赫尔特维西Khertvisi = BKhertvisi赫尔特维西
	f.赫尔特维西Khertvisi.SetParent(f)
	
	f.克拉尔哲季Klarjeti = BKlarjeti克拉尔哲季
	f.克拉尔哲季Klarjeti.SetParent(f)
	
	f.特贝蒂Tbeti = BTbeti特贝蒂
	f.特贝蒂Tbeti.SetParent(f)
	
	f.采普萨Tseptha = BTseptha采普萨
	f.采普萨Tseptha.SetParent(f)
	
}
