package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZakynthosCounty interface {
    feud.County
    BAlikanas阿利卡纳斯() 	feud.Barony
    BKatastari卡塔斯塔里() 	feud.Barony
    BKiliomenos基利奥梅诺斯() 	feud.Barony
    BNavagio纳瓦吉奥() 	feud.Barony
    BStrofadia斯特罗法迪亚() 	feud.Barony
    BVolimes沃利梅斯() 	feud.Barony
    BZakynthos扎金索斯() 	feud.Barony
}

type 扎金索斯ZakynthosCounty struct {
	feud.BaseCounty
	阿利卡纳斯Alikanas 	feud.Barony
	卡塔斯塔里Katastari 	feud.Barony
	基利奥梅诺斯Kiliomenos 	feud.Barony
	纳瓦吉奥Navagio 	feud.Barony
	斯特罗法迪亚Strofadia 	feud.Barony
	沃利梅斯Volimes 	feud.Barony
	扎金索斯Zakynthos 	feud.Barony
}

func (c *扎金索斯ZakynthosCounty) BAlikanas阿利卡纳斯() feud.Barony {
	return c.阿利卡纳斯Alikanas
}
    
func (c *扎金索斯ZakynthosCounty) BKatastari卡塔斯塔里() feud.Barony {
	return c.卡塔斯塔里Katastari
}
    
func (c *扎金索斯ZakynthosCounty) BKiliomenos基利奥梅诺斯() feud.Barony {
	return c.基利奥梅诺斯Kiliomenos
}
    
func (c *扎金索斯ZakynthosCounty) BNavagio纳瓦吉奥() feud.Barony {
	return c.纳瓦吉奥Navagio
}
    
func (c *扎金索斯ZakynthosCounty) BStrofadia斯特罗法迪亚() feud.Barony {
	return c.斯特罗法迪亚Strofadia
}
    
func (c *扎金索斯ZakynthosCounty) BVolimes沃利梅斯() feud.Barony {
	return c.沃利梅斯Volimes
}
    
func (c *扎金索斯ZakynthosCounty) BZakynthos扎金索斯() feud.Barony {
	return c.扎金索斯Zakynthos
}
    
var CZakynthos扎金索斯 ZakynthosCounty = &扎金索斯ZakynthosCounty{}

func init() {
	f := CZakynthos扎金索斯.(*扎金索斯ZakynthosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1889",
		Title:     "zakynthos",
		TitleName: "扎金索斯",
		TitleCode: "c_zakynthos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利卡纳斯Alikanas = BAlikanas阿利卡纳斯
	f.阿利卡纳斯Alikanas.SetParent(f)
	
	f.卡塔斯塔里Katastari = BKatastari卡塔斯塔里
	f.卡塔斯塔里Katastari.SetParent(f)
	
	f.基利奥梅诺斯Kiliomenos = BKiliomenos基利奥梅诺斯
	f.基利奥梅诺斯Kiliomenos.SetParent(f)
	
	f.纳瓦吉奥Navagio = BNavagio纳瓦吉奥
	f.纳瓦吉奥Navagio.SetParent(f)
	
	f.斯特罗法迪亚Strofadia = BStrofadia斯特罗法迪亚
	f.斯特罗法迪亚Strofadia.SetParent(f)
	
	f.沃利梅斯Volimes = BVolimes沃利梅斯
	f.沃利梅斯Volimes.SetParent(f)
	
	f.扎金索斯Zakynthos = BZakynthos扎金索斯
	f.扎金索斯Zakynthos.SetParent(f)
	
}
