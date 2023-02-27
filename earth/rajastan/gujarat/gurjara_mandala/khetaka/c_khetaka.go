package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhetakaCounty interface {
    feud.County
    BAshaval阿舍伐罗() 	feud.Barony
    BBhadran波多兰() 	feud.Barony
    BDholka陀尔迦() 	feud.Barony
    BKhambhat堪婆多() 	feud.Barony
    BKhetaka契吒迦() 	feud.Barony
    BNadiad那地耶陀() 	feud.Barony
    BSanand桑恩德() 	feud.Barony
    BTarapur多罗城() 	feud.Barony
}

type 契吒迦KhetakaCounty struct {
	feud.BaseCounty
	阿舍伐罗Ashaval 	feud.Barony
	波多兰Bhadran 	feud.Barony
	陀尔迦Dholka 	feud.Barony
	堪婆多Khambhat 	feud.Barony
	契吒迦Khetaka 	feud.Barony
	那地耶陀Nadiad 	feud.Barony
	桑恩德Sanand 	feud.Barony
	多罗城Tarapur 	feud.Barony
}

func (c *契吒迦KhetakaCounty) BAshaval阿舍伐罗() feud.Barony {
	return c.阿舍伐罗Ashaval
}
    
func (c *契吒迦KhetakaCounty) BBhadran波多兰() feud.Barony {
	return c.波多兰Bhadran
}
    
func (c *契吒迦KhetakaCounty) BDholka陀尔迦() feud.Barony {
	return c.陀尔迦Dholka
}
    
func (c *契吒迦KhetakaCounty) BKhambhat堪婆多() feud.Barony {
	return c.堪婆多Khambhat
}
    
func (c *契吒迦KhetakaCounty) BKhetaka契吒迦() feud.Barony {
	return c.契吒迦Khetaka
}
    
func (c *契吒迦KhetakaCounty) BNadiad那地耶陀() feud.Barony {
	return c.那地耶陀Nadiad
}
    
func (c *契吒迦KhetakaCounty) BSanand桑恩德() feud.Barony {
	return c.桑恩德Sanand
}
    
func (c *契吒迦KhetakaCounty) BTarapur多罗城() feud.Barony {
	return c.多罗城Tarapur
}
    
var CKhetaka契吒迦 KhetakaCounty = &契吒迦KhetakaCounty{}

func init() {
	f := CKhetaka契吒迦.(*契吒迦KhetakaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1290",
		Title:     "khetaka",
		TitleName: "契吒迦",
		TitleCode: "c_khetaka",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿舍伐罗Ashaval = BAshaval阿舍伐罗
	f.阿舍伐罗Ashaval.SetParent(f)
	
	f.波多兰Bhadran = BBhadran波多兰
	f.波多兰Bhadran.SetParent(f)
	
	f.陀尔迦Dholka = BDholka陀尔迦
	f.陀尔迦Dholka.SetParent(f)
	
	f.堪婆多Khambhat = BKhambhat堪婆多
	f.堪婆多Khambhat.SetParent(f)
	
	f.契吒迦Khetaka = BKhetaka契吒迦
	f.契吒迦Khetaka.SetParent(f)
	
	f.那地耶陀Nadiad = BNadiad那地耶陀
	f.那地耶陀Nadiad.SetParent(f)
	
	f.桑恩德Sanand = BSanand桑恩德
	f.桑恩德Sanand.SetParent(f)
	
	f.多罗城Tarapur = BTarapur多罗城
	f.多罗城Tarapur.SetParent(f)
	
}
