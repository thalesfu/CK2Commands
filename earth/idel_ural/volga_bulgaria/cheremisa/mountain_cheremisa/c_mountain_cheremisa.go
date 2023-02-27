package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Mountain_cheremisaCounty interface {
    feud.County
    BBarysh巴雷什() 	feud.Barony
    BButyrskaya布特尔斯卡亚() 	feud.Barony
    BKanadey卡纳杰伊() 	feud.Barony
    BMelekess梅列克斯() 	feud.Barony
    BSimbirsk辛比尔斯克() 	feud.Barony
    BStanichnaya斯塔尼奇纳亚() 	feud.Barony
    BVybornaya维博尔纳亚() 	feud.Barony
}

type 切列米莎Mountain_cheremisaCounty struct {
	feud.BaseCounty
	巴雷什Barysh 	feud.Barony
	布特尔斯卡亚Butyrskaya 	feud.Barony
	卡纳杰伊Kanadey 	feud.Barony
	梅列克斯Melekess 	feud.Barony
	辛比尔斯克Simbirsk 	feud.Barony
	斯塔尼奇纳亚Stanichnaya 	feud.Barony
	维博尔纳亚Vybornaya 	feud.Barony
}

func (c *切列米莎Mountain_cheremisaCounty) BBarysh巴雷什() feud.Barony {
	return c.巴雷什Barysh
}
    
func (c *切列米莎Mountain_cheremisaCounty) BButyrskaya布特尔斯卡亚() feud.Barony {
	return c.布特尔斯卡亚Butyrskaya
}
    
func (c *切列米莎Mountain_cheremisaCounty) BKanadey卡纳杰伊() feud.Barony {
	return c.卡纳杰伊Kanadey
}
    
func (c *切列米莎Mountain_cheremisaCounty) BMelekess梅列克斯() feud.Barony {
	return c.梅列克斯Melekess
}
    
func (c *切列米莎Mountain_cheremisaCounty) BSimbirsk辛比尔斯克() feud.Barony {
	return c.辛比尔斯克Simbirsk
}
    
func (c *切列米莎Mountain_cheremisaCounty) BStanichnaya斯塔尼奇纳亚() feud.Barony {
	return c.斯塔尼奇纳亚Stanichnaya
}
    
func (c *切列米莎Mountain_cheremisaCounty) BVybornaya维博尔纳亚() feud.Barony {
	return c.维博尔纳亚Vybornaya
}
    
var CMountain_cheremisa切列米莎 Mountain_cheremisaCounty = &切列米莎Mountain_cheremisaCounty{}

func init() {
	f := CMountain_cheremisa切列米莎.(*切列米莎Mountain_cheremisaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "591",
		Title:     "mountain_cheremisa",
		TitleName: "切列米莎",
		TitleCode: "c_mountain_cheremisa",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴雷什Barysh = BBarysh巴雷什
	f.巴雷什Barysh.SetParent(f)
	
	f.布特尔斯卡亚Butyrskaya = BButyrskaya布特尔斯卡亚
	f.布特尔斯卡亚Butyrskaya.SetParent(f)
	
	f.卡纳杰伊Kanadey = BKanadey卡纳杰伊
	f.卡纳杰伊Kanadey.SetParent(f)
	
	f.梅列克斯Melekess = BMelekess梅列克斯
	f.梅列克斯Melekess.SetParent(f)
	
	f.辛比尔斯克Simbirsk = BSimbirsk辛比尔斯克
	f.辛比尔斯克Simbirsk.SetParent(f)
	
	f.斯塔尼奇纳亚Stanichnaya = BStanichnaya斯塔尼奇纳亚
	f.斯塔尼奇纳亚Stanichnaya.SetParent(f)
	
	f.维博尔纳亚Vybornaya = BVybornaya维博尔纳亚
	f.维博尔纳亚Vybornaya.SetParent(f)
	
}
