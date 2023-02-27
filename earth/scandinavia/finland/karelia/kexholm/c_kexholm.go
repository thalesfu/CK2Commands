package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KexholmCounty interface {
    feud.County
    BAntrea安特雷亚() 	feud.Barony
    BJekaborg耶卡伯格() 	feud.Barony
    BKakisalmi凯基萨尔米() 	feud.Barony
    BKoivisto科伊维斯托() 	feud.Barony
    BRaivola赖沃拉() 	feud.Barony
    BTaipale泰帕莱() 	feud.Barony
    BTerijoki泰里约基() 	feud.Barony
    BToksovo托克索沃() 	feud.Barony
}

type 科克斯霍姆KexholmCounty struct {
	feud.BaseCounty
	安特雷亚Antrea 	feud.Barony
	耶卡伯格Jekaborg 	feud.Barony
	凯基萨尔米Kakisalmi 	feud.Barony
	科伊维斯托Koivisto 	feud.Barony
	赖沃拉Raivola 	feud.Barony
	泰帕莱Taipale 	feud.Barony
	泰里约基Terijoki 	feud.Barony
	托克索沃Toksovo 	feud.Barony
}

func (c *科克斯霍姆KexholmCounty) BAntrea安特雷亚() feud.Barony {
	return c.安特雷亚Antrea
}
    
func (c *科克斯霍姆KexholmCounty) BJekaborg耶卡伯格() feud.Barony {
	return c.耶卡伯格Jekaborg
}
    
func (c *科克斯霍姆KexholmCounty) BKakisalmi凯基萨尔米() feud.Barony {
	return c.凯基萨尔米Kakisalmi
}
    
func (c *科克斯霍姆KexholmCounty) BKoivisto科伊维斯托() feud.Barony {
	return c.科伊维斯托Koivisto
}
    
func (c *科克斯霍姆KexholmCounty) BRaivola赖沃拉() feud.Barony {
	return c.赖沃拉Raivola
}
    
func (c *科克斯霍姆KexholmCounty) BTaipale泰帕莱() feud.Barony {
	return c.泰帕莱Taipale
}
    
func (c *科克斯霍姆KexholmCounty) BTerijoki泰里约基() feud.Barony {
	return c.泰里约基Terijoki
}
    
func (c *科克斯霍姆KexholmCounty) BToksovo托克索沃() feud.Barony {
	return c.托克索沃Toksovo
}
    
var CKexholm科克斯霍姆 KexholmCounty = &科克斯霍姆KexholmCounty{}

func init() {
	f := CKexholm科克斯霍姆.(*科克斯霍姆KexholmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "392",
		Title:     "kexholm",
		TitleName: "科克斯霍姆",
		TitleCode: "c_kexholm",
		Baronies:  map[string]feud.Barony{},
	}

	f.安特雷亚Antrea = BAntrea安特雷亚
	f.安特雷亚Antrea.SetParent(f)
	
	f.耶卡伯格Jekaborg = BJekaborg耶卡伯格
	f.耶卡伯格Jekaborg.SetParent(f)
	
	f.凯基萨尔米Kakisalmi = BKakisalmi凯基萨尔米
	f.凯基萨尔米Kakisalmi.SetParent(f)
	
	f.科伊维斯托Koivisto = BKoivisto科伊维斯托
	f.科伊维斯托Koivisto.SetParent(f)
	
	f.赖沃拉Raivola = BRaivola赖沃拉
	f.赖沃拉Raivola.SetParent(f)
	
	f.泰帕莱Taipale = BTaipale泰帕莱
	f.泰帕莱Taipale.SetParent(f)
	
	f.泰里约基Terijoki = BTerijoki泰里约基
	f.泰里约基Terijoki.SetParent(f)
	
	f.托克索沃Toksovo = BToksovo托克索沃
	f.托克索沃Toksovo.SetParent(f)
	
}
