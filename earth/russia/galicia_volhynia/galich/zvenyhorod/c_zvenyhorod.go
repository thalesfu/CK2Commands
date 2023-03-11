package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZvenyhorodCounty interface {
    feud.County
    BButsk布茨克() 	feud.Barony
    BHorodok戈罗多克() 	feud.Barony
    BLvov利沃夫() 	feud.Barony
    BRogozhyn罗戈任() 	feud.Barony
    BSambir桑博尔() 	feud.Barony
    BUdech乌杰奇() 	feud.Barony
    BZvenyhorod兹韦尼戈罗德() 	feud.Barony
}

type 兹韦尼戈罗德ZvenyhorodCounty struct {
	feud.BaseCounty
	布茨克Butsk 	feud.Barony
	戈罗多克Horodok 	feud.Barony
	利沃夫Lvov 	feud.Barony
	罗戈任Rogozhyn 	feud.Barony
	桑博尔Sambir 	feud.Barony
	乌杰奇Udech 	feud.Barony
	兹韦尼戈罗德Zvenyhorod 	feud.Barony
}

func (c *兹韦尼戈罗德ZvenyhorodCounty) BButsk布茨克() feud.Barony {
	return c.布茨克Butsk
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BHorodok戈罗多克() feud.Barony {
	return c.戈罗多克Horodok
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BLvov利沃夫() feud.Barony {
	return c.利沃夫Lvov
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BRogozhyn罗戈任() feud.Barony {
	return c.罗戈任Rogozhyn
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BSambir桑博尔() feud.Barony {
	return c.桑博尔Sambir
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BUdech乌杰奇() feud.Barony {
	return c.乌杰奇Udech
}
    
func (c *兹韦尼戈罗德ZvenyhorodCounty) BZvenyhorod兹韦尼戈罗德() feud.Barony {
	return c.兹韦尼戈罗德Zvenyhorod
}
    
var CZvenyhorod兹韦尼戈罗德 ZvenyhorodCounty = &兹韦尼戈罗德ZvenyhorodCounty{}

func init() {
	f := CZvenyhorod兹韦尼戈罗德.(*兹韦尼戈罗德ZvenyhorodCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1636",
		Title:     "zvenyhorod",
		TitleName: "兹韦尼戈罗德",
		TitleCode: "c_zvenyhorod",
		Baronies:  map[string]feud.Barony{},
	}

	f.布茨克Butsk = BButsk布茨克
	f.布茨克Butsk.SetParent(f)
	
	f.戈罗多克Horodok = BHorodok戈罗多克
	f.戈罗多克Horodok.SetParent(f)
	
	f.利沃夫Lvov = BLvov利沃夫
	f.利沃夫Lvov.SetParent(f)
	
	f.罗戈任Rogozhyn = BRogozhyn罗戈任
	f.罗戈任Rogozhyn.SetParent(f)
	
	f.桑博尔Sambir = BSambir桑博尔
	f.桑博尔Sambir.SetParent(f)
	
	f.乌杰奇Udech = BUdech乌杰奇
	f.乌杰奇Udech.SetParent(f)
	
	f.兹韦尼戈罗德Zvenyhorod = BZvenyhorod兹韦尼戈罗德
	f.兹韦尼戈罗德Zvenyhorod.SetParent(f)
	
}
