package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TabrizCounty interface {
    feud.County
    BAhar阿哈尔() 	feud.Barony
    BBabak巴巴克() 	feud.Barony
    BDihnakhirjan迪纳希尔詹() 	feud.Barony
    BSarab萨拉卜() 	feud.Barony
    BShabestar沙贝斯塔尔() 	feud.Barony
    BTabriz大不里士() 	feud.Barony
    BZahhak扎哈克() 	feud.Barony
}

type 大不里士TabrizCounty struct {
	feud.BaseCounty
	阿哈尔Ahar 	feud.Barony
	巴巴克Babak 	feud.Barony
	迪纳希尔詹Dihnakhirjan 	feud.Barony
	萨拉卜Sarab 	feud.Barony
	沙贝斯塔尔Shabestar 	feud.Barony
	大不里士Tabriz 	feud.Barony
	扎哈克Zahhak 	feud.Barony
}

func (c *大不里士TabrizCounty) BAhar阿哈尔() feud.Barony {
	return c.阿哈尔Ahar
}
    
func (c *大不里士TabrizCounty) BBabak巴巴克() feud.Barony {
	return c.巴巴克Babak
}
    
func (c *大不里士TabrizCounty) BDihnakhirjan迪纳希尔詹() feud.Barony {
	return c.迪纳希尔詹Dihnakhirjan
}
    
func (c *大不里士TabrizCounty) BSarab萨拉卜() feud.Barony {
	return c.萨拉卜Sarab
}
    
func (c *大不里士TabrizCounty) BShabestar沙贝斯塔尔() feud.Barony {
	return c.沙贝斯塔尔Shabestar
}
    
func (c *大不里士TabrizCounty) BTabriz大不里士() feud.Barony {
	return c.大不里士Tabriz
}
    
func (c *大不里士TabrizCounty) BZahhak扎哈克() feud.Barony {
	return c.扎哈克Zahhak
}
    
var CTabriz大不里士 TabrizCounty = &大不里士TabrizCounty{}

func init() {
	f := CTabriz大不里士.(*大不里士TabrizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "667",
		Title:     "tabriz",
		TitleName: "大不里士",
		TitleCode: "c_tabriz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿哈尔Ahar = BAhar阿哈尔
	f.阿哈尔Ahar.SetParent(f)
	
	f.巴巴克Babak = BBabak巴巴克
	f.巴巴克Babak.SetParent(f)
	
	f.迪纳希尔詹Dihnakhirjan = BDihnakhirjan迪纳希尔詹
	f.迪纳希尔詹Dihnakhirjan.SetParent(f)
	
	f.萨拉卜Sarab = BSarab萨拉卜
	f.萨拉卜Sarab.SetParent(f)
	
	f.沙贝斯塔尔Shabestar = BShabestar沙贝斯塔尔
	f.沙贝斯塔尔Shabestar.SetParent(f)
	
	f.大不里士Tabriz = BTabriz大不里士
	f.大不里士Tabriz.SetParent(f)
	
	f.扎哈克Zahhak = BZahhak扎哈克
	f.扎哈克Zahhak.SetParent(f)
	
}
