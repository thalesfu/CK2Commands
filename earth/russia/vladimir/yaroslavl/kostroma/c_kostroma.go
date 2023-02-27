package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KostromaCounty interface {
    feud.County
    BApraksino阿普拉克西诺() 	feud.Barony
    BKosmynino科斯梅尼诺() 	feud.Barony
    BKostroma科斯特罗马() 	feud.Barony
    BNekrasoskoye涅克拉索斯科耶() 	feud.Barony
    BNerektha涅列赫塔() 	feud.Barony
    BPlyos普廖斯() 	feud.Barony
    BSudislavl苏季斯拉夫尔() 	feud.Barony
    BZavolzhsk扎沃尔日斯克() 	feud.Barony
}

type 科斯特罗马KostromaCounty struct {
	feud.BaseCounty
	阿普拉克西诺Apraksino 	feud.Barony
	科斯梅尼诺Kosmynino 	feud.Barony
	科斯特罗马Kostroma 	feud.Barony
	涅克拉索斯科耶Nekrasoskoye 	feud.Barony
	涅列赫塔Nerektha 	feud.Barony
	普廖斯Plyos 	feud.Barony
	苏季斯拉夫尔Sudislavl 	feud.Barony
	扎沃尔日斯克Zavolzhsk 	feud.Barony
}

func (c *科斯特罗马KostromaCounty) BApraksino阿普拉克西诺() feud.Barony {
	return c.阿普拉克西诺Apraksino
}
    
func (c *科斯特罗马KostromaCounty) BKosmynino科斯梅尼诺() feud.Barony {
	return c.科斯梅尼诺Kosmynino
}
    
func (c *科斯特罗马KostromaCounty) BKostroma科斯特罗马() feud.Barony {
	return c.科斯特罗马Kostroma
}
    
func (c *科斯特罗马KostromaCounty) BNekrasoskoye涅克拉索斯科耶() feud.Barony {
	return c.涅克拉索斯科耶Nekrasoskoye
}
    
func (c *科斯特罗马KostromaCounty) BNerektha涅列赫塔() feud.Barony {
	return c.涅列赫塔Nerektha
}
    
func (c *科斯特罗马KostromaCounty) BPlyos普廖斯() feud.Barony {
	return c.普廖斯Plyos
}
    
func (c *科斯特罗马KostromaCounty) BSudislavl苏季斯拉夫尔() feud.Barony {
	return c.苏季斯拉夫尔Sudislavl
}
    
func (c *科斯特罗马KostromaCounty) BZavolzhsk扎沃尔日斯克() feud.Barony {
	return c.扎沃尔日斯克Zavolzhsk
}
    
var CKostroma科斯特罗马 KostromaCounty = &科斯特罗马KostromaCounty{}

func init() {
	f := CKostroma科斯特罗马.(*科斯特罗马KostromaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "407",
		Title:     "kostroma",
		TitleName: "科斯特罗马",
		TitleCode: "c_kostroma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普拉克西诺Apraksino = BApraksino阿普拉克西诺
	f.阿普拉克西诺Apraksino.SetParent(f)
	
	f.科斯梅尼诺Kosmynino = BKosmynino科斯梅尼诺
	f.科斯梅尼诺Kosmynino.SetParent(f)
	
	f.科斯特罗马Kostroma = BKostroma科斯特罗马
	f.科斯特罗马Kostroma.SetParent(f)
	
	f.涅克拉索斯科耶Nekrasoskoye = BNekrasoskoye涅克拉索斯科耶
	f.涅克拉索斯科耶Nekrasoskoye.SetParent(f)
	
	f.涅列赫塔Nerektha = BNerektha涅列赫塔
	f.涅列赫塔Nerektha.SetParent(f)
	
	f.普廖斯Plyos = BPlyos普廖斯
	f.普廖斯Plyos.SetParent(f)
	
	f.苏季斯拉夫尔Sudislavl = BSudislavl苏季斯拉夫尔
	f.苏季斯拉夫尔Sudislavl.SetParent(f)
	
	f.扎沃尔日斯克Zavolzhsk = BZavolzhsk扎沃尔日斯克
	f.扎沃尔日斯克Zavolzhsk.SetParent(f)
	
}
