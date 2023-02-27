package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TuraCounty interface {
    feud.County
    BAke阿克() 	feud.Barony
    BAlmaty阿拉木图() 	feud.Barony
    BBestobe别斯托别() 	feud.Barony
    BKishkenekol基什克涅科尔() 	feud.Barony
    BTeke铁克() 	feud.Barony
    BTura图拉() 	feud.Barony
    BZalivino_tura扎利维诺() 	feud.Barony
}

type 图拉TuraCounty struct {
	feud.BaseCounty
	阿克Ake 	feud.Barony
	阿拉木图Almaty 	feud.Barony
	别斯托别Bestobe 	feud.Barony
	基什克涅科尔Kishkenekol 	feud.Barony
	铁克Teke 	feud.Barony
	图拉Tura 	feud.Barony
	扎利维诺Zalivino_tura 	feud.Barony
}

func (c *图拉TuraCounty) BAke阿克() feud.Barony {
	return c.阿克Ake
}
    
func (c *图拉TuraCounty) BAlmaty阿拉木图() feud.Barony {
	return c.阿拉木图Almaty
}
    
func (c *图拉TuraCounty) BBestobe别斯托别() feud.Barony {
	return c.别斯托别Bestobe
}
    
func (c *图拉TuraCounty) BKishkenekol基什克涅科尔() feud.Barony {
	return c.基什克涅科尔Kishkenekol
}
    
func (c *图拉TuraCounty) BTeke铁克() feud.Barony {
	return c.铁克Teke
}
    
func (c *图拉TuraCounty) BTura图拉() feud.Barony {
	return c.图拉Tura
}
    
func (c *图拉TuraCounty) BZalivino_tura扎利维诺() feud.Barony {
	return c.扎利维诺Zalivino_tura
}
    
var CTura图拉 TuraCounty = &图拉TuraCounty{}

func init() {
	f := CTura图拉.(*图拉TuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1306",
		Title:     "tura",
		TitleName: "图拉",
		TitleCode: "c_tura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克Ake = BAke阿克
	f.阿克Ake.SetParent(f)
	
	f.阿拉木图Almaty = BAlmaty阿拉木图
	f.阿拉木图Almaty.SetParent(f)
	
	f.别斯托别Bestobe = BBestobe别斯托别
	f.别斯托别Bestobe.SetParent(f)
	
	f.基什克涅科尔Kishkenekol = BKishkenekol基什克涅科尔
	f.基什克涅科尔Kishkenekol.SetParent(f)
	
	f.铁克Teke = BTeke铁克
	f.铁克Teke.SetParent(f)
	
	f.图拉Tura = BTura图拉
	f.图拉Tura.SetParent(f)
	
	f.扎利维诺Zalivino_tura = BZalivino_tura扎利维诺
	f.扎利维诺Zalivino_tura.SetParent(f)
	
}
