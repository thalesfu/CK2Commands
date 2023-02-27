package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KursCounty interface {
    feud.County
    BAizpute艾兹普泰() 	feud.Barony
    BApule阿普莱() 	feud.Barony
    BGrobin格罗宾() 	feud.Barony
    BImpilte伊皮尔蒂斯() 	feud.Barony
    BKuldiga库尔迪加() 	feud.Barony
    BMedze梅泽() 	feud.Barony
    BVartaja瓦尔塔亚() 	feud.Barony
}

type 库尔泽梅KursCounty struct {
	feud.BaseCounty
	艾兹普泰Aizpute 	feud.Barony
	阿普莱Apule 	feud.Barony
	格罗宾Grobin 	feud.Barony
	伊皮尔蒂斯Impilte 	feud.Barony
	库尔迪加Kuldiga 	feud.Barony
	梅泽Medze 	feud.Barony
	瓦尔塔亚Vartaja 	feud.Barony
}

func (c *库尔泽梅KursCounty) BAizpute艾兹普泰() feud.Barony {
	return c.艾兹普泰Aizpute
}
    
func (c *库尔泽梅KursCounty) BApule阿普莱() feud.Barony {
	return c.阿普莱Apule
}
    
func (c *库尔泽梅KursCounty) BGrobin格罗宾() feud.Barony {
	return c.格罗宾Grobin
}
    
func (c *库尔泽梅KursCounty) BImpilte伊皮尔蒂斯() feud.Barony {
	return c.伊皮尔蒂斯Impilte
}
    
func (c *库尔泽梅KursCounty) BKuldiga库尔迪加() feud.Barony {
	return c.库尔迪加Kuldiga
}
    
func (c *库尔泽梅KursCounty) BMedze梅泽() feud.Barony {
	return c.梅泽Medze
}
    
func (c *库尔泽梅KursCounty) BVartaja瓦尔塔亚() feud.Barony {
	return c.瓦尔塔亚Vartaja
}
    
var CKurs库尔泽梅 KursCounty = &库尔泽梅KursCounty{}

func init() {
	f := CKurs库尔泽梅.(*库尔泽梅KursCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "373",
		Title:     "kurs",
		TitleName: "库尔泽梅",
		TitleCode: "c_kurs",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾兹普泰Aizpute = BAizpute艾兹普泰
	f.艾兹普泰Aizpute.SetParent(f)
	
	f.阿普莱Apule = BApule阿普莱
	f.阿普莱Apule.SetParent(f)
	
	f.格罗宾Grobin = BGrobin格罗宾
	f.格罗宾Grobin.SetParent(f)
	
	f.伊皮尔蒂斯Impilte = BImpilte伊皮尔蒂斯
	f.伊皮尔蒂斯Impilte.SetParent(f)
	
	f.库尔迪加Kuldiga = BKuldiga库尔迪加
	f.库尔迪加Kuldiga.SetParent(f)
	
	f.梅泽Medze = BMedze梅泽
	f.梅泽Medze.SetParent(f)
	
	f.瓦尔塔亚Vartaja = BVartaja瓦尔塔亚
	f.瓦尔塔亚Vartaja.SetParent(f)
	
}
