package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MemelCounty interface {
    feud.County
    BDreverna德雷韦尔纳() 	feud.Barony
    BGargzdai加尔格兹代() 	feud.Barony
    BJuodkrante尤奥德克兰特() 	feud.Barony
    BKaup考普() 	feud.Barony
    BKretingale克雷廷加莱() 	feud.Barony
    BMemel梅梅尔() 	feud.Barony
    BNida尼达() 	feud.Barony
    BPalanga帕兰加() 	feud.Barony
}

type 梅梅尔MemelCounty struct {
	feud.BaseCounty
	德雷韦尔纳Dreverna 	feud.Barony
	加尔格兹代Gargzdai 	feud.Barony
	尤奥德克兰特Juodkrante 	feud.Barony
	考普Kaup 	feud.Barony
	克雷廷加莱Kretingale 	feud.Barony
	梅梅尔Memel 	feud.Barony
	尼达Nida 	feud.Barony
	帕兰加Palanga 	feud.Barony
}

func (c *梅梅尔MemelCounty) BDreverna德雷韦尔纳() feud.Barony {
	return c.德雷韦尔纳Dreverna
}
    
func (c *梅梅尔MemelCounty) BGargzdai加尔格兹代() feud.Barony {
	return c.加尔格兹代Gargzdai
}
    
func (c *梅梅尔MemelCounty) BJuodkrante尤奥德克兰特() feud.Barony {
	return c.尤奥德克兰特Juodkrante
}
    
func (c *梅梅尔MemelCounty) BKaup考普() feud.Barony {
	return c.考普Kaup
}
    
func (c *梅梅尔MemelCounty) BKretingale克雷廷加莱() feud.Barony {
	return c.克雷廷加莱Kretingale
}
    
func (c *梅梅尔MemelCounty) BMemel梅梅尔() feud.Barony {
	return c.梅梅尔Memel
}
    
func (c *梅梅尔MemelCounty) BNida尼达() feud.Barony {
	return c.尼达Nida
}
    
func (c *梅梅尔MemelCounty) BPalanga帕兰加() feud.Barony {
	return c.帕兰加Palanga
}
    
var CMemel梅梅尔 MemelCounty = &梅梅尔MemelCounty{}

func init() {
	f := CMemel梅梅尔.(*梅梅尔MemelCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "372",
		Title:     "memel",
		TitleName: "梅梅尔",
		TitleCode: "c_memel",
		Baronies:  map[string]feud.Barony{},
	}

	f.德雷韦尔纳Dreverna = BDreverna德雷韦尔纳
	f.德雷韦尔纳Dreverna.SetParent(f)
	
	f.加尔格兹代Gargzdai = BGargzdai加尔格兹代
	f.加尔格兹代Gargzdai.SetParent(f)
	
	f.尤奥德克兰特Juodkrante = BJuodkrante尤奥德克兰特
	f.尤奥德克兰特Juodkrante.SetParent(f)
	
	f.考普Kaup = BKaup考普
	f.考普Kaup.SetParent(f)
	
	f.克雷廷加莱Kretingale = BKretingale克雷廷加莱
	f.克雷廷加莱Kretingale.SetParent(f)
	
	f.梅梅尔Memel = BMemel梅梅尔
	f.梅梅尔Memel.SetParent(f)
	
	f.尼达Nida = BNida尼达
	f.尼达Nida.SetParent(f)
	
	f.帕兰加Palanga = BPalanga帕兰加
	f.帕兰加Palanga.SetParent(f)
	
}
