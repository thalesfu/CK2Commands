package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VychegdaCounty interface {
    feud.County
    BAnyb阿内布() 	feud.Barony
    BKulom库洛姆() 	feud.Barony
    BPomozdino波莫兹季诺() 	feud.Barony
    BPozheg波热格() 	feud.Barony
    BRuch鲁奇() 	feud.Barony
    BVol_dino沃利季诺() 	feud.Barony
    BVychegda维切格达() 	feud.Barony
}

type 维切格达VychegdaCounty struct {
	feud.BaseCounty
	阿内布Anyb 	feud.Barony
	库洛姆Kulom 	feud.Barony
	波莫兹季诺Pomozdino 	feud.Barony
	波热格Pozheg 	feud.Barony
	鲁奇Ruch 	feud.Barony
	沃利季诺Vol_dino 	feud.Barony
	维切格达Vychegda 	feud.Barony
}

func (c *维切格达VychegdaCounty) BAnyb阿内布() feud.Barony {
	return c.阿内布Anyb
}
    
func (c *维切格达VychegdaCounty) BKulom库洛姆() feud.Barony {
	return c.库洛姆Kulom
}
    
func (c *维切格达VychegdaCounty) BPomozdino波莫兹季诺() feud.Barony {
	return c.波莫兹季诺Pomozdino
}
    
func (c *维切格达VychegdaCounty) BPozheg波热格() feud.Barony {
	return c.波热格Pozheg
}
    
func (c *维切格达VychegdaCounty) BRuch鲁奇() feud.Barony {
	return c.鲁奇Ruch
}
    
func (c *维切格达VychegdaCounty) BVol_dino沃利季诺() feud.Barony {
	return c.沃利季诺Vol_dino
}
    
func (c *维切格达VychegdaCounty) BVychegda维切格达() feud.Barony {
	return c.维切格达Vychegda
}
    
var CVychegda维切格达 VychegdaCounty = &维切格达VychegdaCounty{}

func init() {
	f := CVychegda维切格达.(*维切格达VychegdaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1838",
		Title:     "vychegda",
		TitleName: "维切格达",
		TitleCode: "c_vychegda",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿内布Anyb = BAnyb阿内布
	f.阿内布Anyb.SetParent(f)
	
	f.库洛姆Kulom = BKulom库洛姆
	f.库洛姆Kulom.SetParent(f)
	
	f.波莫兹季诺Pomozdino = BPomozdino波莫兹季诺
	f.波莫兹季诺Pomozdino.SetParent(f)
	
	f.波热格Pozheg = BPozheg波热格
	f.波热格Pozheg.SetParent(f)
	
	f.鲁奇Ruch = BRuch鲁奇
	f.鲁奇Ruch.SetParent(f)
	
	f.沃利季诺Vol_dino = BVol_dino沃利季诺
	f.沃利季诺Vol_dino.SetParent(f)
	
	f.维切格达Vychegda = BVychegda维切格达
	f.维切格达Vychegda.SetParent(f)
	
}
