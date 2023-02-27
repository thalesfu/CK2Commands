package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaninCounty interface {
    feud.County
    BChesha乔沙() 	feud.Barony
    BChizha奇扎() 	feud.Barony
    BKanin卡宁() 	feud.Barony
    BKiya_kanin基亚() 	feud.Barony
    BMesna梅斯纳() 	feud.Barony
    BOzero奥泽罗() 	feud.Barony
    BShoyna绍伊纳() 	feud.Barony
}

type 卡宁KaninCounty struct {
	feud.BaseCounty
	乔沙Chesha 	feud.Barony
	奇扎Chizha 	feud.Barony
	卡宁Kanin 	feud.Barony
	基亚Kiya_kanin 	feud.Barony
	梅斯纳Mesna 	feud.Barony
	奥泽罗Ozero 	feud.Barony
	绍伊纳Shoyna 	feud.Barony
}

func (c *卡宁KaninCounty) BChesha乔沙() feud.Barony {
	return c.乔沙Chesha
}
    
func (c *卡宁KaninCounty) BChizha奇扎() feud.Barony {
	return c.奇扎Chizha
}
    
func (c *卡宁KaninCounty) BKanin卡宁() feud.Barony {
	return c.卡宁Kanin
}
    
func (c *卡宁KaninCounty) BKiya_kanin基亚() feud.Barony {
	return c.基亚Kiya_kanin
}
    
func (c *卡宁KaninCounty) BMesna梅斯纳() feud.Barony {
	return c.梅斯纳Mesna
}
    
func (c *卡宁KaninCounty) BOzero奥泽罗() feud.Barony {
	return c.奥泽罗Ozero
}
    
func (c *卡宁KaninCounty) BShoyna绍伊纳() feud.Barony {
	return c.绍伊纳Shoyna
}
    
var CKanin卡宁 KaninCounty = &卡宁KaninCounty{}

func init() {
	f := CKanin卡宁.(*卡宁KaninCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1820",
		Title:     "kanin",
		TitleName: "卡宁",
		TitleCode: "c_kanin",
		Baronies:  map[string]feud.Barony{},
	}

	f.乔沙Chesha = BChesha乔沙
	f.乔沙Chesha.SetParent(f)
	
	f.奇扎Chizha = BChizha奇扎
	f.奇扎Chizha.SetParent(f)
	
	f.卡宁Kanin = BKanin卡宁
	f.卡宁Kanin.SetParent(f)
	
	f.基亚Kiya_kanin = BKiya_kanin基亚
	f.基亚Kiya_kanin.SetParent(f)
	
	f.梅斯纳Mesna = BMesna梅斯纳
	f.梅斯纳Mesna.SetParent(f)
	
	f.奥泽罗Ozero = BOzero奥泽罗
	f.奥泽罗Ozero.SetParent(f)
	
	f.绍伊纳Shoyna = BShoyna绍伊纳
	f.绍伊纳Shoyna.SetParent(f)
	
}
