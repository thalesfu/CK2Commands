package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DhamalpurCounty interface {
    feud.County
    BDhamalpur陀摩罗补罗() 	feud.Barony
    BDhrol陀罗() 	feud.Barony
    BHurshi呼尸() 	feud.Barony
    BJade阇提() 	feud.Barony
    BKhasta迦多() 	feud.Barony
    BLakhota罗诃陀() 	feud.Barony
    BMorvi摩尔比() 	feud.Barony
}

type 陀摩罗补罗DhamalpurCounty struct {
	feud.BaseCounty
	陀摩罗补罗Dhamalpur 	feud.Barony
	陀罗Dhrol 	feud.Barony
	呼尸Hurshi 	feud.Barony
	阇提Jade 	feud.Barony
	迦多Khasta 	feud.Barony
	罗诃陀Lakhota 	feud.Barony
	摩尔比Morvi 	feud.Barony
}

func (c *陀摩罗补罗DhamalpurCounty) BDhamalpur陀摩罗补罗() feud.Barony {
	return c.陀摩罗补罗Dhamalpur
}
    
func (c *陀摩罗补罗DhamalpurCounty) BDhrol陀罗() feud.Barony {
	return c.陀罗Dhrol
}
    
func (c *陀摩罗补罗DhamalpurCounty) BHurshi呼尸() feud.Barony {
	return c.呼尸Hurshi
}
    
func (c *陀摩罗补罗DhamalpurCounty) BJade阇提() feud.Barony {
	return c.阇提Jade
}
    
func (c *陀摩罗补罗DhamalpurCounty) BKhasta迦多() feud.Barony {
	return c.迦多Khasta
}
    
func (c *陀摩罗补罗DhamalpurCounty) BLakhota罗诃陀() feud.Barony {
	return c.罗诃陀Lakhota
}
    
func (c *陀摩罗补罗DhamalpurCounty) BMorvi摩尔比() feud.Barony {
	return c.摩尔比Morvi
}
    
var CDhamalpur陀摩罗补罗 DhamalpurCounty = &陀摩罗补罗DhamalpurCounty{}

func init() {
	f := CDhamalpur陀摩罗补罗.(*陀摩罗补罗DhamalpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1268",
		Title:     "dhamalpur",
		TitleName: "陀摩罗补罗",
		TitleCode: "c_dhamalpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.陀摩罗补罗Dhamalpur = BDhamalpur陀摩罗补罗
	f.陀摩罗补罗Dhamalpur.SetParent(f)
	
	f.陀罗Dhrol = BDhrol陀罗
	f.陀罗Dhrol.SetParent(f)
	
	f.呼尸Hurshi = BHurshi呼尸
	f.呼尸Hurshi.SetParent(f)
	
	f.阇提Jade = BJade阇提
	f.阇提Jade.SetParent(f)
	
	f.迦多Khasta = BKhasta迦多
	f.迦多Khasta.SetParent(f)
	
	f.罗诃陀Lakhota = BLakhota罗诃陀
	f.罗诃陀Lakhota.SetParent(f)
	
	f.摩尔比Morvi = BMorvi摩尔比
	f.摩尔比Morvi.SetParent(f)
	
}
