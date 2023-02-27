package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AoudaghostCounty interface {
    feud.County
    BAoudaghost奥达戈斯特() 	feud.Barony
    BHodh霍德() 	feud.Barony
    BIchga伊希加() 	feud.Barony
    BKedama凯达马() 	feud.Barony
    BKsarelbarka巴尔卡堡() 	feud.Barony
    BMoudjeria穆杰里亚() 	feud.Barony
    BSarel萨雷尔() 	feud.Barony
}

type 奥达戈斯特AoudaghostCounty struct {
	feud.BaseCounty
	奥达戈斯特Aoudaghost 	feud.Barony
	霍德Hodh 	feud.Barony
	伊希加Ichga 	feud.Barony
	凯达马Kedama 	feud.Barony
	巴尔卡堡Ksarelbarka 	feud.Barony
	穆杰里亚Moudjeria 	feud.Barony
	萨雷尔Sarel 	feud.Barony
}

func (c *奥达戈斯特AoudaghostCounty) BAoudaghost奥达戈斯特() feud.Barony {
	return c.奥达戈斯特Aoudaghost
}
    
func (c *奥达戈斯特AoudaghostCounty) BHodh霍德() feud.Barony {
	return c.霍德Hodh
}
    
func (c *奥达戈斯特AoudaghostCounty) BIchga伊希加() feud.Barony {
	return c.伊希加Ichga
}
    
func (c *奥达戈斯特AoudaghostCounty) BKedama凯达马() feud.Barony {
	return c.凯达马Kedama
}
    
func (c *奥达戈斯特AoudaghostCounty) BKsarelbarka巴尔卡堡() feud.Barony {
	return c.巴尔卡堡Ksarelbarka
}
    
func (c *奥达戈斯特AoudaghostCounty) BMoudjeria穆杰里亚() feud.Barony {
	return c.穆杰里亚Moudjeria
}
    
func (c *奥达戈斯特AoudaghostCounty) BSarel萨雷尔() feud.Barony {
	return c.萨雷尔Sarel
}
    
var CAoudaghost奥达戈斯特 AoudaghostCounty = &奥达戈斯特AoudaghostCounty{}

func init() {
	f := CAoudaghost奥达戈斯特.(*奥达戈斯特AoudaghostCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "912",
		Title:     "aoudaghost",
		TitleName: "奥达戈斯特",
		TitleCode: "c_aoudaghost",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥达戈斯特Aoudaghost = BAoudaghost奥达戈斯特
	f.奥达戈斯特Aoudaghost.SetParent(f)
	
	f.霍德Hodh = BHodh霍德
	f.霍德Hodh.SetParent(f)
	
	f.伊希加Ichga = BIchga伊希加
	f.伊希加Ichga.SetParent(f)
	
	f.凯达马Kedama = BKedama凯达马
	f.凯达马Kedama.SetParent(f)
	
	f.巴尔卡堡Ksarelbarka = BKsarelbarka巴尔卡堡
	f.巴尔卡堡Ksarelbarka.SetParent(f)
	
	f.穆杰里亚Moudjeria = BMoudjeria穆杰里亚
	f.穆杰里亚Moudjeria.SetParent(f)
	
	f.萨雷尔Sarel = BSarel萨雷尔
	f.萨雷尔Sarel.SetParent(f)
	
}
