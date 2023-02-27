package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AargauCounty interface {
    feud.County
    BAarau_augstgau阿劳() 	feud.Barony
    BBasel巴塞尔() 	feud.Barony
    BDelemont德莱蒙() 	feud.Barony
    BHabsburg哈布斯堡() 	feud.Barony
    BLaufenberg劳芬贝格() 	feud.Barony
    BMoutier穆捷() 	feud.Barony
    BPorrentruy波朗特吕() 	feud.Barony
    BSolothurn索洛图恩() 	feud.Barony
    BSt_ursanne圣于尔萨讷() 	feud.Barony
}

type 巴塞尔AargauCounty struct {
	feud.BaseCounty
	阿劳Aarau_augstgau 	feud.Barony
	巴塞尔Basel 	feud.Barony
	德莱蒙Delemont 	feud.Barony
	哈布斯堡Habsburg 	feud.Barony
	劳芬贝格Laufenberg 	feud.Barony
	穆捷Moutier 	feud.Barony
	波朗特吕Porrentruy 	feud.Barony
	索洛图恩Solothurn 	feud.Barony
	圣于尔萨讷St_ursanne 	feud.Barony
}

func (c *巴塞尔AargauCounty) BAarau_augstgau阿劳() feud.Barony {
	return c.阿劳Aarau_augstgau
}
    
func (c *巴塞尔AargauCounty) BBasel巴塞尔() feud.Barony {
	return c.巴塞尔Basel
}
    
func (c *巴塞尔AargauCounty) BDelemont德莱蒙() feud.Barony {
	return c.德莱蒙Delemont
}
    
func (c *巴塞尔AargauCounty) BHabsburg哈布斯堡() feud.Barony {
	return c.哈布斯堡Habsburg
}
    
func (c *巴塞尔AargauCounty) BLaufenberg劳芬贝格() feud.Barony {
	return c.劳芬贝格Laufenberg
}
    
func (c *巴塞尔AargauCounty) BMoutier穆捷() feud.Barony {
	return c.穆捷Moutier
}
    
func (c *巴塞尔AargauCounty) BPorrentruy波朗特吕() feud.Barony {
	return c.波朗特吕Porrentruy
}
    
func (c *巴塞尔AargauCounty) BSolothurn索洛图恩() feud.Barony {
	return c.索洛图恩Solothurn
}
    
func (c *巴塞尔AargauCounty) BSt_ursanne圣于尔萨讷() feud.Barony {
	return c.圣于尔萨讷St_ursanne
}
    
var CAargau巴塞尔 AargauCounty = &巴塞尔AargauCounty{}

func init() {
	f := CAargau巴塞尔.(*巴塞尔AargauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "242",
		Title:     "aargau",
		TitleName: "巴塞尔",
		TitleCode: "c_aargau",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿劳Aarau_augstgau = BAarau_augstgau阿劳
	f.阿劳Aarau_augstgau.SetParent(f)
	
	f.巴塞尔Basel = BBasel巴塞尔
	f.巴塞尔Basel.SetParent(f)
	
	f.德莱蒙Delemont = BDelemont德莱蒙
	f.德莱蒙Delemont.SetParent(f)
	
	f.哈布斯堡Habsburg = BHabsburg哈布斯堡
	f.哈布斯堡Habsburg.SetParent(f)
	
	f.劳芬贝格Laufenberg = BLaufenberg劳芬贝格
	f.劳芬贝格Laufenberg.SetParent(f)
	
	f.穆捷Moutier = BMoutier穆捷
	f.穆捷Moutier.SetParent(f)
	
	f.波朗特吕Porrentruy = BPorrentruy波朗特吕
	f.波朗特吕Porrentruy.SetParent(f)
	
	f.索洛图恩Solothurn = BSolothurn索洛图恩
	f.索洛图恩Solothurn.SetParent(f)
	
	f.圣于尔萨讷St_ursanne = BSt_ursanne圣于尔萨讷
	f.圣于尔萨讷St_ursanne.SetParent(f)
	
}
