package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BirjandCounty interface {
    feud.County
    BBirjand比尔詹德() 	feud.Barony
    BBoshruyeh博什鲁耶() 	feud.Barony
    BFurg富尔格() 	feud.Barony
    BNehbandan内赫班丹() 	feud.Barony
    BPaeenshahr帕恩_沙赫尔() 	feud.Barony
    BQaen加延() 	feud.Barony
    BSarayan萨拉扬() 	feud.Barony
    BToon顿恩() 	feud.Barony
}

type 比尔詹德BirjandCounty struct {
	feud.BaseCounty
	比尔詹德Birjand 	feud.Barony
	博什鲁耶Boshruyeh 	feud.Barony
	富尔格Furg 	feud.Barony
	内赫班丹Nehbandan 	feud.Barony
	帕恩_沙赫尔Paeenshahr 	feud.Barony
	加延Qaen 	feud.Barony
	萨拉扬Sarayan 	feud.Barony
	顿恩Toon 	feud.Barony
}

func (c *比尔詹德BirjandCounty) BBirjand比尔詹德() feud.Barony {
	return c.比尔詹德Birjand
}
    
func (c *比尔詹德BirjandCounty) BBoshruyeh博什鲁耶() feud.Barony {
	return c.博什鲁耶Boshruyeh
}
    
func (c *比尔詹德BirjandCounty) BFurg富尔格() feud.Barony {
	return c.富尔格Furg
}
    
func (c *比尔詹德BirjandCounty) BNehbandan内赫班丹() feud.Barony {
	return c.内赫班丹Nehbandan
}
    
func (c *比尔詹德BirjandCounty) BPaeenshahr帕恩_沙赫尔() feud.Barony {
	return c.帕恩_沙赫尔Paeenshahr
}
    
func (c *比尔詹德BirjandCounty) BQaen加延() feud.Barony {
	return c.加延Qaen
}
    
func (c *比尔詹德BirjandCounty) BSarayan萨拉扬() feud.Barony {
	return c.萨拉扬Sarayan
}
    
func (c *比尔詹德BirjandCounty) BToon顿恩() feud.Barony {
	return c.顿恩Toon
}
    
var CBirjand比尔詹德 BirjandCounty = &比尔詹德BirjandCounty{}

func init() {
	f := CBirjand比尔詹德.(*比尔詹德BirjandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "906",
		Title:     "birjand",
		TitleName: "比尔詹德",
		TitleCode: "c_birjand",
		Baronies:  map[string]feud.Barony{},
	}

	f.比尔詹德Birjand = BBirjand比尔詹德
	f.比尔詹德Birjand.SetParent(f)
	
	f.博什鲁耶Boshruyeh = BBoshruyeh博什鲁耶
	f.博什鲁耶Boshruyeh.SetParent(f)
	
	f.富尔格Furg = BFurg富尔格
	f.富尔格Furg.SetParent(f)
	
	f.内赫班丹Nehbandan = BNehbandan内赫班丹
	f.内赫班丹Nehbandan.SetParent(f)
	
	f.帕恩_沙赫尔Paeenshahr = BPaeenshahr帕恩_沙赫尔
	f.帕恩_沙赫尔Paeenshahr.SetParent(f)
	
	f.加延Qaen = BQaen加延
	f.加延Qaen.SetParent(f)
	
	f.萨拉扬Sarayan = BSarayan萨拉扬
	f.萨拉扬Sarayan.SetParent(f)
	
	f.顿恩Toon = BToon顿恩
	f.顿恩Toon.SetParent(f)
	
}
