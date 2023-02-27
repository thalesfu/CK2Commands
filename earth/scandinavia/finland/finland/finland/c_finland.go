package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinlandCounty interface {
    feud.County
    BJukarsborg于卡斯伯格() 	feud.Barony
    BKuusisto库西斯托() 	feud.Barony
    BLieto列托() 	feud.Barony
    BNaantali楠塔利() 	feud.Barony
    BRauma劳马() 	feud.Barony
    BRikala里卡拉() 	feud.Barony
    BStenberga斯腾贝里亚() 	feud.Barony
    BTurku图尔库() 	feud.Barony
}

type 索米FinlandCounty struct {
	feud.BaseCounty
	于卡斯伯格Jukarsborg 	feud.Barony
	库西斯托Kuusisto 	feud.Barony
	列托Lieto 	feud.Barony
	楠塔利Naantali 	feud.Barony
	劳马Rauma 	feud.Barony
	里卡拉Rikala 	feud.Barony
	斯腾贝里亚Stenberga 	feud.Barony
	图尔库Turku 	feud.Barony
}

func (c *索米FinlandCounty) BJukarsborg于卡斯伯格() feud.Barony {
	return c.于卡斯伯格Jukarsborg
}
    
func (c *索米FinlandCounty) BKuusisto库西斯托() feud.Barony {
	return c.库西斯托Kuusisto
}
    
func (c *索米FinlandCounty) BLieto列托() feud.Barony {
	return c.列托Lieto
}
    
func (c *索米FinlandCounty) BNaantali楠塔利() feud.Barony {
	return c.楠塔利Naantali
}
    
func (c *索米FinlandCounty) BRauma劳马() feud.Barony {
	return c.劳马Rauma
}
    
func (c *索米FinlandCounty) BRikala里卡拉() feud.Barony {
	return c.里卡拉Rikala
}
    
func (c *索米FinlandCounty) BStenberga斯腾贝里亚() feud.Barony {
	return c.斯腾贝里亚Stenberga
}
    
func (c *索米FinlandCounty) BTurku图尔库() feud.Barony {
	return c.图尔库Turku
}
    
var CFinland索米 FinlandCounty = &索米FinlandCounty{}

func init() {
	f := CFinland索米.(*索米FinlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "382",
		Title:     "finland",
		TitleName: "索米",
		TitleCode: "c_finland",
		Baronies:  map[string]feud.Barony{},
	}

	f.于卡斯伯格Jukarsborg = BJukarsborg于卡斯伯格
	f.于卡斯伯格Jukarsborg.SetParent(f)
	
	f.库西斯托Kuusisto = BKuusisto库西斯托
	f.库西斯托Kuusisto.SetParent(f)
	
	f.列托Lieto = BLieto列托
	f.列托Lieto.SetParent(f)
	
	f.楠塔利Naantali = BNaantali楠塔利
	f.楠塔利Naantali.SetParent(f)
	
	f.劳马Rauma = BRauma劳马
	f.劳马Rauma.SetParent(f)
	
	f.里卡拉Rikala = BRikala里卡拉
	f.里卡拉Rikala.SetParent(f)
	
	f.斯腾贝里亚Stenberga = BStenberga斯腾贝里亚
	f.斯腾贝里亚Stenberga.SetParent(f)
	
	f.图尔库Turku = BTurku图尔库
	f.图尔库Turku.SetParent(f)
	
}
