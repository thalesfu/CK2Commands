package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BuqtirmaCounty interface {
    feud.County
    BAkku阿库() 	feud.Barony
    BBuqtirma布赫塔尔马() 	feud.Barony
    BKazantay卡赞泰() 	feud.Barony
    BKoyanbay科扬拜() 	feud.Barony
    BShamshi沙姆希() 	feud.Barony
    BShoktal朔克塔尔() 	feud.Barony
    BZhanatan扎纳坦() 	feud.Barony
}

type 布赫塔尔马BuqtirmaCounty struct {
	feud.BaseCounty
	阿库Akku 	feud.Barony
	布赫塔尔马Buqtirma 	feud.Barony
	卡赞泰Kazantay 	feud.Barony
	科扬拜Koyanbay 	feud.Barony
	沙姆希Shamshi 	feud.Barony
	朔克塔尔Shoktal 	feud.Barony
	扎纳坦Zhanatan 	feud.Barony
}

func (c *布赫塔尔马BuqtirmaCounty) BAkku阿库() feud.Barony {
	return c.阿库Akku
}
    
func (c *布赫塔尔马BuqtirmaCounty) BBuqtirma布赫塔尔马() feud.Barony {
	return c.布赫塔尔马Buqtirma
}
    
func (c *布赫塔尔马BuqtirmaCounty) BKazantay卡赞泰() feud.Barony {
	return c.卡赞泰Kazantay
}
    
func (c *布赫塔尔马BuqtirmaCounty) BKoyanbay科扬拜() feud.Barony {
	return c.科扬拜Koyanbay
}
    
func (c *布赫塔尔马BuqtirmaCounty) BShamshi沙姆希() feud.Barony {
	return c.沙姆希Shamshi
}
    
func (c *布赫塔尔马BuqtirmaCounty) BShoktal朔克塔尔() feud.Barony {
	return c.朔克塔尔Shoktal
}
    
func (c *布赫塔尔马BuqtirmaCounty) BZhanatan扎纳坦() feud.Barony {
	return c.扎纳坦Zhanatan
}
    
var CBuqtirma布赫塔尔马 BuqtirmaCounty = &布赫塔尔马BuqtirmaCounty{}

func init() {
	f := CBuqtirma布赫塔尔马.(*布赫塔尔马BuqtirmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1877",
		Title:     "buqtirma",
		TitleName: "布赫塔尔马",
		TitleCode: "c_buqtirma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿库Akku = BAkku阿库
	f.阿库Akku.SetParent(f)
	
	f.布赫塔尔马Buqtirma = BBuqtirma布赫塔尔马
	f.布赫塔尔马Buqtirma.SetParent(f)
	
	f.卡赞泰Kazantay = BKazantay卡赞泰
	f.卡赞泰Kazantay.SetParent(f)
	
	f.科扬拜Koyanbay = BKoyanbay科扬拜
	f.科扬拜Koyanbay.SetParent(f)
	
	f.沙姆希Shamshi = BShamshi沙姆希
	f.沙姆希Shamshi.SetParent(f)
	
	f.朔克塔尔Shoktal = BShoktal朔克塔尔
	f.朔克塔尔Shoktal.SetParent(f)
	
	f.扎纳坦Zhanatan = BZhanatan扎纳坦
	f.扎纳坦Zhanatan.SetParent(f)
	
}
