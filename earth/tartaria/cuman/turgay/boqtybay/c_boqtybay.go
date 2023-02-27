package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BoqtybayCounty interface {
    feud.County
    BBoqtybay博克特拜() 	feud.Barony
    BGuberlinskiy古别尔林斯基() 	feud.Barony
    BKlochkovo克洛奇科沃() 	feud.Barony
    BPushkino普希金诺() 	feud.Barony
    BTaldykol塔尔德科尔() 	feud.Barony
    BUrkash乌尔卡什() 	feud.Barony
    BZhailma_boqtybay扎伊尔马() 	feud.Barony
}

type 博克特拜BoqtybayCounty struct {
	feud.BaseCounty
	博克特拜Boqtybay 	feud.Barony
	古别尔林斯基Guberlinskiy 	feud.Barony
	克洛奇科沃Klochkovo 	feud.Barony
	普希金诺Pushkino 	feud.Barony
	塔尔德科尔Taldykol 	feud.Barony
	乌尔卡什Urkash 	feud.Barony
	扎伊尔马Zhailma_boqtybay 	feud.Barony
}

func (c *博克特拜BoqtybayCounty) BBoqtybay博克特拜() feud.Barony {
	return c.博克特拜Boqtybay
}
    
func (c *博克特拜BoqtybayCounty) BGuberlinskiy古别尔林斯基() feud.Barony {
	return c.古别尔林斯基Guberlinskiy
}
    
func (c *博克特拜BoqtybayCounty) BKlochkovo克洛奇科沃() feud.Barony {
	return c.克洛奇科沃Klochkovo
}
    
func (c *博克特拜BoqtybayCounty) BPushkino普希金诺() feud.Barony {
	return c.普希金诺Pushkino
}
    
func (c *博克特拜BoqtybayCounty) BTaldykol塔尔德科尔() feud.Barony {
	return c.塔尔德科尔Taldykol
}
    
func (c *博克特拜BoqtybayCounty) BUrkash乌尔卡什() feud.Barony {
	return c.乌尔卡什Urkash
}
    
func (c *博克特拜BoqtybayCounty) BZhailma_boqtybay扎伊尔马() feud.Barony {
	return c.扎伊尔马Zhailma_boqtybay
}
    
var CBoqtybay博克特拜 BoqtybayCounty = &博克特拜BoqtybayCounty{}

func init() {
	f := CBoqtybay博克特拜.(*博克特拜BoqtybayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1870",
		Title:     "boqtybay",
		TitleName: "博克特拜",
		TitleCode: "c_boqtybay",
		Baronies:  map[string]feud.Barony{},
	}

	f.博克特拜Boqtybay = BBoqtybay博克特拜
	f.博克特拜Boqtybay.SetParent(f)
	
	f.古别尔林斯基Guberlinskiy = BGuberlinskiy古别尔林斯基
	f.古别尔林斯基Guberlinskiy.SetParent(f)
	
	f.克洛奇科沃Klochkovo = BKlochkovo克洛奇科沃
	f.克洛奇科沃Klochkovo.SetParent(f)
	
	f.普希金诺Pushkino = BPushkino普希金诺
	f.普希金诺Pushkino.SetParent(f)
	
	f.塔尔德科尔Taldykol = BTaldykol塔尔德科尔
	f.塔尔德科尔Taldykol.SetParent(f)
	
	f.乌尔卡什Urkash = BUrkash乌尔卡什
	f.乌尔卡什Urkash.SetParent(f)
	
	f.扎伊尔马Zhailma_boqtybay = BZhailma_boqtybay扎伊尔马
	f.扎伊尔马Zhailma_boqtybay.SetParent(f)
	
}
