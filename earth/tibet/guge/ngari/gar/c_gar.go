package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GarCounty interface {
    feud.County
    BChakgang甲岗() 	feud.Barony
    BGar噶尔() 	feud.Barony
    BKunsa昆莎() 	feud.Barony
    BKupa过巴() 	feud.Barony
    BLangchu朗久() 	feud.Barony
    BRabang热帮() 	feud.Barony
    BSenggezangbo森格藏布() 	feud.Barony
}

type 噶尔GarCounty struct {
	feud.BaseCounty
	甲岗Chakgang 	feud.Barony
	噶尔Gar 	feud.Barony
	昆莎Kunsa 	feud.Barony
	过巴Kupa 	feud.Barony
	朗久Langchu 	feud.Barony
	热帮Rabang 	feud.Barony
	森格藏布Senggezangbo 	feud.Barony
}

func (c *噶尔GarCounty) BChakgang甲岗() feud.Barony {
	return c.甲岗Chakgang
}
    
func (c *噶尔GarCounty) BGar噶尔() feud.Barony {
	return c.噶尔Gar
}
    
func (c *噶尔GarCounty) BKunsa昆莎() feud.Barony {
	return c.昆莎Kunsa
}
    
func (c *噶尔GarCounty) BKupa过巴() feud.Barony {
	return c.过巴Kupa
}
    
func (c *噶尔GarCounty) BLangchu朗久() feud.Barony {
	return c.朗久Langchu
}
    
func (c *噶尔GarCounty) BRabang热帮() feud.Barony {
	return c.热帮Rabang
}
    
func (c *噶尔GarCounty) BSenggezangbo森格藏布() feud.Barony {
	return c.森格藏布Senggezangbo
}
    
var CGar噶尔 GarCounty = &噶尔GarCounty{}

func init() {
	f := CGar噶尔.(*噶尔GarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1487",
		Title:     "gar",
		TitleName: "噶尔",
		TitleCode: "c_gar",
		Baronies:  map[string]feud.Barony{},
	}

	f.甲岗Chakgang = BChakgang甲岗
	f.甲岗Chakgang.SetParent(f)
	
	f.噶尔Gar = BGar噶尔
	f.噶尔Gar.SetParent(f)
	
	f.昆莎Kunsa = BKunsa昆莎
	f.昆莎Kunsa.SetParent(f)
	
	f.过巴Kupa = BKupa过巴
	f.过巴Kupa.SetParent(f)
	
	f.朗久Langchu = BLangchu朗久
	f.朗久Langchu.SetParent(f)
	
	f.热帮Rabang = BRabang热帮
	f.热帮Rabang.SetParent(f)
	
	f.森格藏布Senggezangbo = BSenggezangbo森格藏布
	f.森格藏布Senggezangbo.SetParent(f)
	
}
