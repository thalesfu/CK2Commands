package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KartalyCounty interface {
    feud.County
    BKachar卡恰尔() 	feud.Barony
    BKartaly卡尔塔雷_阿亚特() 	feud.Barony
    BKirovka基洛夫卡() 	feud.Barony
    BKostanay库斯塔奈() 	feud.Barony
    BTalapker塔拉普克尔() 	feud.Barony
    BZarechny扎列奇内() 	feud.Barony
    BZhdanovka日丹诺夫卡() 	feud.Barony
}

type 卡尔塔雷_阿亚特KartalyCounty struct {
	feud.BaseCounty
	卡恰尔Kachar 	feud.Barony
	卡尔塔雷_阿亚特Kartaly 	feud.Barony
	基洛夫卡Kirovka 	feud.Barony
	库斯塔奈Kostanay 	feud.Barony
	塔拉普克尔Talapker 	feud.Barony
	扎列奇内Zarechny 	feud.Barony
	日丹诺夫卡Zhdanovka 	feud.Barony
}

func (c *卡尔塔雷_阿亚特KartalyCounty) BKachar卡恰尔() feud.Barony {
	return c.卡恰尔Kachar
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BKartaly卡尔塔雷_阿亚特() feud.Barony {
	return c.卡尔塔雷_阿亚特Kartaly
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BKirovka基洛夫卡() feud.Barony {
	return c.基洛夫卡Kirovka
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BKostanay库斯塔奈() feud.Barony {
	return c.库斯塔奈Kostanay
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BTalapker塔拉普克尔() feud.Barony {
	return c.塔拉普克尔Talapker
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BZarechny扎列奇内() feud.Barony {
	return c.扎列奇内Zarechny
}
    
func (c *卡尔塔雷_阿亚特KartalyCounty) BZhdanovka日丹诺夫卡() feud.Barony {
	return c.日丹诺夫卡Zhdanovka
}
    
var CKartaly卡尔塔雷_阿亚特 KartalyCounty = &卡尔塔雷_阿亚特KartalyCounty{}

func init() {
	f := CKartaly卡尔塔雷_阿亚特.(*卡尔塔雷_阿亚特KartalyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1854",
		Title:     "kartaly",
		TitleName: "卡尔塔雷_阿亚特",
		TitleCode: "c_kartaly",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡恰尔Kachar = BKachar卡恰尔
	f.卡恰尔Kachar.SetParent(f)
	
	f.卡尔塔雷_阿亚特Kartaly = BKartaly卡尔塔雷_阿亚特
	f.卡尔塔雷_阿亚特Kartaly.SetParent(f)
	
	f.基洛夫卡Kirovka = BKirovka基洛夫卡
	f.基洛夫卡Kirovka.SetParent(f)
	
	f.库斯塔奈Kostanay = BKostanay库斯塔奈
	f.库斯塔奈Kostanay.SetParent(f)
	
	f.塔拉普克尔Talapker = BTalapker塔拉普克尔
	f.塔拉普克尔Talapker.SetParent(f)
	
	f.扎列奇内Zarechny = BZarechny扎列奇内
	f.扎列奇内Zarechny.SetParent(f)
	
	f.日丹诺夫卡Zhdanovka = BZhdanovka日丹诺夫卡
	f.日丹诺夫卡Zhdanovka.SetParent(f)
	
}
