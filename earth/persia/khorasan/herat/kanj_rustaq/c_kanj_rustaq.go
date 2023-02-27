package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kanj_rustaqCounty interface {
    feud.County
    BAufah奥法() 	feud.Barony
    BBaghsur巴格苏尔() 	feud.Barony
    BDiza迪扎() 	feud.Barony
    BFenesk费内斯克() 	feud.Barony
    BKusak库萨克() 	feud.Barony
    BNab纳布() 	feud.Barony
    BTajarg塔贾格() 	feud.Barony
}

type 坎吉鲁斯塔克Kanj_rustaqCounty struct {
	feud.BaseCounty
	奥法Aufah 	feud.Barony
	巴格苏尔Baghsur 	feud.Barony
	迪扎Diza 	feud.Barony
	费内斯克Fenesk 	feud.Barony
	库萨克Kusak 	feud.Barony
	纳布Nab 	feud.Barony
	塔贾格Tajarg 	feud.Barony
}

func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BAufah奥法() feud.Barony {
	return c.奥法Aufah
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BBaghsur巴格苏尔() feud.Barony {
	return c.巴格苏尔Baghsur
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BDiza迪扎() feud.Barony {
	return c.迪扎Diza
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BFenesk费内斯克() feud.Barony {
	return c.费内斯克Fenesk
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BKusak库萨克() feud.Barony {
	return c.库萨克Kusak
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BNab纳布() feud.Barony {
	return c.纳布Nab
}
    
func (c *坎吉鲁斯塔克Kanj_rustaqCounty) BTajarg塔贾格() feud.Barony {
	return c.塔贾格Tajarg
}
    
var CKanj_rustaq坎吉鲁斯塔克 Kanj_rustaqCounty = &坎吉鲁斯塔克Kanj_rustaqCounty{}

func init() {
	f := CKanj_rustaq坎吉鲁斯塔克.(*坎吉鲁斯塔克Kanj_rustaqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1539",
		Title:     "kanj_rustaq",
		TitleName: "坎吉鲁斯塔克",
		TitleCode: "c_kanj_rustaq",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥法Aufah = BAufah奥法
	f.奥法Aufah.SetParent(f)
	
	f.巴格苏尔Baghsur = BBaghsur巴格苏尔
	f.巴格苏尔Baghsur.SetParent(f)
	
	f.迪扎Diza = BDiza迪扎
	f.迪扎Diza.SetParent(f)
	
	f.费内斯克Fenesk = BFenesk费内斯克
	f.费内斯克Fenesk.SetParent(f)
	
	f.库萨克Kusak = BKusak库萨克
	f.库萨克Kusak.SetParent(f)
	
	f.纳布Nab = BNab纳布
	f.纳布Nab.SetParent(f)
	
	f.塔贾格Tajarg = BTajarg塔贾格
	f.塔贾格Tajarg.SetParent(f)
	
}
