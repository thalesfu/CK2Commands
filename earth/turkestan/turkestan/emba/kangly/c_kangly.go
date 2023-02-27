package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanglyCounty interface {
    feud.County
    BKangly康里() 	feud.Barony
    BKarasha卡拉沙() 	feud.Barony
    BKizay基泽乌() 	feud.Barony
    BKoshkar科什卡尔() 	feud.Barony
    BMakat马卡特() 	feud.Barony
    BQulsary库利萨雷() 	feud.Barony
    BTengiz_kangly田吉兹() 	feud.Barony
}

type 康里KanglyCounty struct {
	feud.BaseCounty
	康里Kangly 	feud.Barony
	卡拉沙Karasha 	feud.Barony
	基泽乌Kizay 	feud.Barony
	科什卡尔Koshkar 	feud.Barony
	马卡特Makat 	feud.Barony
	库利萨雷Qulsary 	feud.Barony
	田吉兹Tengiz_kangly 	feud.Barony
}

func (c *康里KanglyCounty) BKangly康里() feud.Barony {
	return c.康里Kangly
}
    
func (c *康里KanglyCounty) BKarasha卡拉沙() feud.Barony {
	return c.卡拉沙Karasha
}
    
func (c *康里KanglyCounty) BKizay基泽乌() feud.Barony {
	return c.基泽乌Kizay
}
    
func (c *康里KanglyCounty) BKoshkar科什卡尔() feud.Barony {
	return c.科什卡尔Koshkar
}
    
func (c *康里KanglyCounty) BMakat马卡特() feud.Barony {
	return c.马卡特Makat
}
    
func (c *康里KanglyCounty) BQulsary库利萨雷() feud.Barony {
	return c.库利萨雷Qulsary
}
    
func (c *康里KanglyCounty) BTengiz_kangly田吉兹() feud.Barony {
	return c.田吉兹Tengiz_kangly
}
    
var CKangly康里 KanglyCounty = &康里KanglyCounty{}

func init() {
	f := CKangly康里.(*康里KanglyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "621",
		Title:     "kangly",
		TitleName: "康里",
		TitleCode: "c_kangly",
		Baronies:  map[string]feud.Barony{},
	}

	f.康里Kangly = BKangly康里
	f.康里Kangly.SetParent(f)
	
	f.卡拉沙Karasha = BKarasha卡拉沙
	f.卡拉沙Karasha.SetParent(f)
	
	f.基泽乌Kizay = BKizay基泽乌
	f.基泽乌Kizay.SetParent(f)
	
	f.科什卡尔Koshkar = BKoshkar科什卡尔
	f.科什卡尔Koshkar.SetParent(f)
	
	f.马卡特Makat = BMakat马卡特
	f.马卡特Makat.SetParent(f)
	
	f.库利萨雷Qulsary = BQulsary库利萨雷
	f.库利萨雷Qulsary.SetParent(f)
	
	f.田吉兹Tengiz_kangly = BTengiz_kangly田吉兹
	f.田吉兹Tengiz_kangly.SetParent(f)
	
}
