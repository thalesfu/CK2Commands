package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MangyulCounty interface {
    feud.County
    BDram樟木() 	feud.Barony
    BDzongka宗嘎() 	feud.Barony
    BGyagya加加() 	feud.Barony
    BGyirong济咙() 	feud.Barony
    BNyalam尼牙拉木() 	feud.Barony
    BShelkar协格尔() 	feud.Barony
    BTingri定日() 	feud.Barony
}

type 芒隅MangyulCounty struct {
	feud.BaseCounty
	樟木Dram 	feud.Barony
	宗嘎Dzongka 	feud.Barony
	加加Gyagya 	feud.Barony
	济咙Gyirong 	feud.Barony
	尼牙拉木Nyalam 	feud.Barony
	协格尔Shelkar 	feud.Barony
	定日Tingri 	feud.Barony
}

func (c *芒隅MangyulCounty) BDram樟木() feud.Barony {
	return c.樟木Dram
}
    
func (c *芒隅MangyulCounty) BDzongka宗嘎() feud.Barony {
	return c.宗嘎Dzongka
}
    
func (c *芒隅MangyulCounty) BGyagya加加() feud.Barony {
	return c.加加Gyagya
}
    
func (c *芒隅MangyulCounty) BGyirong济咙() feud.Barony {
	return c.济咙Gyirong
}
    
func (c *芒隅MangyulCounty) BNyalam尼牙拉木() feud.Barony {
	return c.尼牙拉木Nyalam
}
    
func (c *芒隅MangyulCounty) BShelkar协格尔() feud.Barony {
	return c.协格尔Shelkar
}
    
func (c *芒隅MangyulCounty) BTingri定日() feud.Barony {
	return c.定日Tingri
}
    
var CMangyul芒隅 MangyulCounty = &芒隅MangyulCounty{}

func init() {
	f := CMangyul芒隅.(*芒隅MangyulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1476",
		Title:     "mangyul",
		TitleName: "芒隅",
		TitleCode: "c_mangyul",
		Baronies:  map[string]feud.Barony{},
	}

	f.樟木Dram = BDram樟木
	f.樟木Dram.SetParent(f)
	
	f.宗嘎Dzongka = BDzongka宗嘎
	f.宗嘎Dzongka.SetParent(f)
	
	f.加加Gyagya = BGyagya加加
	f.加加Gyagya.SetParent(f)
	
	f.济咙Gyirong = BGyirong济咙
	f.济咙Gyirong.SetParent(f)
	
	f.尼牙拉木Nyalam = BNyalam尼牙拉木
	f.尼牙拉木Nyalam.SetParent(f)
	
	f.协格尔Shelkar = BShelkar协格尔
	f.协格尔Shelkar.SetParent(f)
	
	f.定日Tingri = BTingri定日
	f.定日Tingri.SetParent(f)
	
}
