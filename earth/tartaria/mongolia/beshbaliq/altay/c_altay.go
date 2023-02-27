package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AltayCounty interface {
    feud.County
    BAltay阿尔泰() 	feud.Barony
    BBurqin布尔津() 	feud.Barony
    BBurultokay布伦托海() 	feud.Barony
    BJeminay吉木乃() 	feud.Barony
    BKokagax阔克阿尕什() 	feud.Barony
    BKoktokay可可托海() 	feud.Barony
    BQinggil青格里() 	feud.Barony
}

type 兀泷古AltayCounty struct {
	feud.BaseCounty
	阿尔泰Altay 	feud.Barony
	布尔津Burqin 	feud.Barony
	布伦托海Burultokay 	feud.Barony
	吉木乃Jeminay 	feud.Barony
	阔克阿尕什Kokagax 	feud.Barony
	可可托海Koktokay 	feud.Barony
	青格里Qinggil 	feud.Barony
}

func (c *兀泷古AltayCounty) BAltay阿尔泰() feud.Barony {
	return c.阿尔泰Altay
}
    
func (c *兀泷古AltayCounty) BBurqin布尔津() feud.Barony {
	return c.布尔津Burqin
}
    
func (c *兀泷古AltayCounty) BBurultokay布伦托海() feud.Barony {
	return c.布伦托海Burultokay
}
    
func (c *兀泷古AltayCounty) BJeminay吉木乃() feud.Barony {
	return c.吉木乃Jeminay
}
    
func (c *兀泷古AltayCounty) BKokagax阔克阿尕什() feud.Barony {
	return c.阔克阿尕什Kokagax
}
    
func (c *兀泷古AltayCounty) BKoktokay可可托海() feud.Barony {
	return c.可可托海Koktokay
}
    
func (c *兀泷古AltayCounty) BQinggil青格里() feud.Barony {
	return c.青格里Qinggil
}
    
var CAltay兀泷古 AltayCounty = &兀泷古AltayCounty{}

func init() {
	f := CAltay兀泷古.(*兀泷古AltayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1451",
		Title:     "altay",
		TitleName: "兀泷古",
		TitleCode: "c_altay",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔泰Altay = BAltay阿尔泰
	f.阿尔泰Altay.SetParent(f)
	
	f.布尔津Burqin = BBurqin布尔津
	f.布尔津Burqin.SetParent(f)
	
	f.布伦托海Burultokay = BBurultokay布伦托海
	f.布伦托海Burultokay.SetParent(f)
	
	f.吉木乃Jeminay = BJeminay吉木乃
	f.吉木乃Jeminay.SetParent(f)
	
	f.阔克阿尕什Kokagax = BKokagax阔克阿尕什
	f.阔克阿尕什Kokagax.SetParent(f)
	
	f.可可托海Koktokay = BKoktokay可可托海
	f.可可托海Koktokay.SetParent(f)
	
	f.青格里Qinggil = BQinggil青格里
	f.青格里Qinggil.SetParent(f)
	
}
