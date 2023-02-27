package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Dhu_zabiCounty interface {
    feud.County
    BDhuzabi阿布扎比() 	feud.Barony
    BDibba迪巴() 	feud.Barony
    BDubai迪拜() 	feud.Barony
    BJulfar佐尔法() 	feud.Barony
    BKhorfakkan豪尔费坎() 	feud.Barony
    BQutuf喀图尔() 	feud.Barony
    BSharjah沙迦() 	feud.Barony
    BSohar苏哈尔() 	feud.Barony
}

type 佐尔法Dhu_zabiCounty struct {
	feud.BaseCounty
	阿布扎比Dhuzabi 	feud.Barony
	迪巴Dibba 	feud.Barony
	迪拜Dubai 	feud.Barony
	佐尔法Julfar 	feud.Barony
	豪尔费坎Khorfakkan 	feud.Barony
	喀图尔Qutuf 	feud.Barony
	沙迦Sharjah 	feud.Barony
	苏哈尔Sohar 	feud.Barony
}

func (c *佐尔法Dhu_zabiCounty) BDhuzabi阿布扎比() feud.Barony {
	return c.阿布扎比Dhuzabi
}
    
func (c *佐尔法Dhu_zabiCounty) BDibba迪巴() feud.Barony {
	return c.迪巴Dibba
}
    
func (c *佐尔法Dhu_zabiCounty) BDubai迪拜() feud.Barony {
	return c.迪拜Dubai
}
    
func (c *佐尔法Dhu_zabiCounty) BJulfar佐尔法() feud.Barony {
	return c.佐尔法Julfar
}
    
func (c *佐尔法Dhu_zabiCounty) BKhorfakkan豪尔费坎() feud.Barony {
	return c.豪尔费坎Khorfakkan
}
    
func (c *佐尔法Dhu_zabiCounty) BQutuf喀图尔() feud.Barony {
	return c.喀图尔Qutuf
}
    
func (c *佐尔法Dhu_zabiCounty) BSharjah沙迦() feud.Barony {
	return c.沙迦Sharjah
}
    
func (c *佐尔法Dhu_zabiCounty) BSohar苏哈尔() feud.Barony {
	return c.苏哈尔Sohar
}
    
var CDhu_zabi佐尔法 Dhu_zabiCounty = &佐尔法Dhu_zabiCounty{}

func init() {
	f := CDhu_zabi佐尔法.(*佐尔法Dhu_zabiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "870",
		Title:     "dhu_zabi",
		TitleName: "佐尔法",
		TitleCode: "c_dhu_zabi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布扎比Dhuzabi = BDhuzabi阿布扎比
	f.阿布扎比Dhuzabi.SetParent(f)
	
	f.迪巴Dibba = BDibba迪巴
	f.迪巴Dibba.SetParent(f)
	
	f.迪拜Dubai = BDubai迪拜
	f.迪拜Dubai.SetParent(f)
	
	f.佐尔法Julfar = BJulfar佐尔法
	f.佐尔法Julfar.SetParent(f)
	
	f.豪尔费坎Khorfakkan = BKhorfakkan豪尔费坎
	f.豪尔费坎Khorfakkan.SetParent(f)
	
	f.喀图尔Qutuf = BQutuf喀图尔
	f.喀图尔Qutuf.SetParent(f)
	
	f.沙迦Sharjah = BSharjah沙迦
	f.沙迦Sharjah.SetParent(f)
	
	f.苏哈尔Sohar = BSohar苏哈尔
	f.苏哈尔Sohar.SetParent(f)
	
}
