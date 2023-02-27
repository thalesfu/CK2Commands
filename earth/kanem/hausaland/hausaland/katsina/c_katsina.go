package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KatsinaCounty interface {
    feud.County
    BAsani阿沙尼() 	feud.Barony
    BFurana富拉纳() 	feud.Barony
    BGausambe戈桑布() 	feud.Barony
    BKaduna卡杜纳() 	feud.Barony
    BKatsina卡齐纳() 	feud.Barony
    BOgunde奥贡德() 	feud.Barony
    BZaria扎里亚() 	feud.Barony
}

type 卡齐纳KatsinaCounty struct {
	feud.BaseCounty
	阿沙尼Asani 	feud.Barony
	富拉纳Furana 	feud.Barony
	戈桑布Gausambe 	feud.Barony
	卡杜纳Kaduna 	feud.Barony
	卡齐纳Katsina 	feud.Barony
	奥贡德Ogunde 	feud.Barony
	扎里亚Zaria 	feud.Barony
}

func (c *卡齐纳KatsinaCounty) BAsani阿沙尼() feud.Barony {
	return c.阿沙尼Asani
}
    
func (c *卡齐纳KatsinaCounty) BFurana富拉纳() feud.Barony {
	return c.富拉纳Furana
}
    
func (c *卡齐纳KatsinaCounty) BGausambe戈桑布() feud.Barony {
	return c.戈桑布Gausambe
}
    
func (c *卡齐纳KatsinaCounty) BKaduna卡杜纳() feud.Barony {
	return c.卡杜纳Kaduna
}
    
func (c *卡齐纳KatsinaCounty) BKatsina卡齐纳() feud.Barony {
	return c.卡齐纳Katsina
}
    
func (c *卡齐纳KatsinaCounty) BOgunde奥贡德() feud.Barony {
	return c.奥贡德Ogunde
}
    
func (c *卡齐纳KatsinaCounty) BZaria扎里亚() feud.Barony {
	return c.扎里亚Zaria
}
    
var CKatsina卡齐纳 KatsinaCounty = &卡齐纳KatsinaCounty{}

func init() {
	f := CKatsina卡齐纳.(*卡齐纳KatsinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1751",
		Title:     "katsina",
		TitleName: "卡齐纳",
		TitleCode: "c_katsina",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿沙尼Asani = BAsani阿沙尼
	f.阿沙尼Asani.SetParent(f)
	
	f.富拉纳Furana = BFurana富拉纳
	f.富拉纳Furana.SetParent(f)
	
	f.戈桑布Gausambe = BGausambe戈桑布
	f.戈桑布Gausambe.SetParent(f)
	
	f.卡杜纳Kaduna = BKaduna卡杜纳
	f.卡杜纳Kaduna.SetParent(f)
	
	f.卡齐纳Katsina = BKatsina卡齐纳
	f.卡齐纳Katsina.SetParent(f)
	
	f.奥贡德Ogunde = BOgunde奥贡德
	f.奥贡德Ogunde.SetParent(f)
	
	f.扎里亚Zaria = BZaria扎里亚
	f.扎里亚Zaria.SetParent(f)
	
}
