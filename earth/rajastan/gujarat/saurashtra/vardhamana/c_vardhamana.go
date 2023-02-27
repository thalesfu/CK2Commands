package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VardhamanaCounty interface {
    feud.County
    BAmarvalli阿摩伐梨() 	feud.Barony
    BHalvad哈尔瓦() 	feud.Barony
    BHardi曷罗提() 	feud.Barony
    BHarpura曷罗补罗() 	feud.Barony
    BJakhri阇佉梨() 	feud.Barony
    BRanpur兰补罗() 	feud.Barony
    BVardhamana筏陀摩那() 	feud.Barony
}

type 筏陀摩那VardhamanaCounty struct {
	feud.BaseCounty
	阿摩伐梨Amarvalli 	feud.Barony
	哈尔瓦Halvad 	feud.Barony
	曷罗提Hardi 	feud.Barony
	曷罗补罗Harpura 	feud.Barony
	阇佉梨Jakhri 	feud.Barony
	兰补罗Ranpur 	feud.Barony
	筏陀摩那Vardhamana 	feud.Barony
}

func (c *筏陀摩那VardhamanaCounty) BAmarvalli阿摩伐梨() feud.Barony {
	return c.阿摩伐梨Amarvalli
}
    
func (c *筏陀摩那VardhamanaCounty) BHalvad哈尔瓦() feud.Barony {
	return c.哈尔瓦Halvad
}
    
func (c *筏陀摩那VardhamanaCounty) BHardi曷罗提() feud.Barony {
	return c.曷罗提Hardi
}
    
func (c *筏陀摩那VardhamanaCounty) BHarpura曷罗补罗() feud.Barony {
	return c.曷罗补罗Harpura
}
    
func (c *筏陀摩那VardhamanaCounty) BJakhri阇佉梨() feud.Barony {
	return c.阇佉梨Jakhri
}
    
func (c *筏陀摩那VardhamanaCounty) BRanpur兰补罗() feud.Barony {
	return c.兰补罗Ranpur
}
    
func (c *筏陀摩那VardhamanaCounty) BVardhamana筏陀摩那() feud.Barony {
	return c.筏陀摩那Vardhamana
}
    
var CVardhamana筏陀摩那 VardhamanaCounty = &筏陀摩那VardhamanaCounty{}

func init() {
	f := CVardhamana筏陀摩那.(*筏陀摩那VardhamanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1267",
		Title:     "vardhamana",
		TitleName: "筏陀摩那",
		TitleCode: "c_vardhamana",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿摩伐梨Amarvalli = BAmarvalli阿摩伐梨
	f.阿摩伐梨Amarvalli.SetParent(f)
	
	f.哈尔瓦Halvad = BHalvad哈尔瓦
	f.哈尔瓦Halvad.SetParent(f)
	
	f.曷罗提Hardi = BHardi曷罗提
	f.曷罗提Hardi.SetParent(f)
	
	f.曷罗补罗Harpura = BHarpura曷罗补罗
	f.曷罗补罗Harpura.SetParent(f)
	
	f.阇佉梨Jakhri = BJakhri阇佉梨
	f.阇佉梨Jakhri.SetParent(f)
	
	f.兰补罗Ranpur = BRanpur兰补罗
	f.兰补罗Ranpur.SetParent(f)
	
	f.筏陀摩那Vardhamana = BVardhamana筏陀摩那
	f.筏陀摩那Vardhamana.SetParent(f)
	
}
