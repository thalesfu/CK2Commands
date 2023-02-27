package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RamagiriCounty interface {
    feud.County
    BBanjapalli槃阇波梨() 	feud.Barony
    BBehtar吠诃怛罗() 	feud.Barony
    BBelagula吠罗瞿罗() 	feud.Barony
    BMansar曼瑟尔() 	feud.Barony
    BNandivardhana难提拔檀那() 	feud.Barony
    BPavnar普拉纳() 	feud.Barony
    BRamagiri罗摩耆厘() 	feud.Barony
}

type 罗摩耆厘RamagiriCounty struct {
	feud.BaseCounty
	槃阇波梨Banjapalli 	feud.Barony
	吠诃怛罗Behtar 	feud.Barony
	吠罗瞿罗Belagula 	feud.Barony
	曼瑟尔Mansar 	feud.Barony
	难提拔檀那Nandivardhana 	feud.Barony
	普拉纳Pavnar 	feud.Barony
	罗摩耆厘Ramagiri 	feud.Barony
}

func (c *罗摩耆厘RamagiriCounty) BBanjapalli槃阇波梨() feud.Barony {
	return c.槃阇波梨Banjapalli
}
    
func (c *罗摩耆厘RamagiriCounty) BBehtar吠诃怛罗() feud.Barony {
	return c.吠诃怛罗Behtar
}
    
func (c *罗摩耆厘RamagiriCounty) BBelagula吠罗瞿罗() feud.Barony {
	return c.吠罗瞿罗Belagula
}
    
func (c *罗摩耆厘RamagiriCounty) BMansar曼瑟尔() feud.Barony {
	return c.曼瑟尔Mansar
}
    
func (c *罗摩耆厘RamagiriCounty) BNandivardhana难提拔檀那() feud.Barony {
	return c.难提拔檀那Nandivardhana
}
    
func (c *罗摩耆厘RamagiriCounty) BPavnar普拉纳() feud.Barony {
	return c.普拉纳Pavnar
}
    
func (c *罗摩耆厘RamagiriCounty) BRamagiri罗摩耆厘() feud.Barony {
	return c.罗摩耆厘Ramagiri
}
    
var CRamagiri罗摩耆厘 RamagiriCounty = &罗摩耆厘RamagiriCounty{}

func init() {
	f := CRamagiri罗摩耆厘.(*罗摩耆厘RamagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1159",
		Title:     "ramagiri",
		TitleName: "罗摩耆厘",
		TitleCode: "c_ramagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.槃阇波梨Banjapalli = BBanjapalli槃阇波梨
	f.槃阇波梨Banjapalli.SetParent(f)
	
	f.吠诃怛罗Behtar = BBehtar吠诃怛罗
	f.吠诃怛罗Behtar.SetParent(f)
	
	f.吠罗瞿罗Belagula = BBelagula吠罗瞿罗
	f.吠罗瞿罗Belagula.SetParent(f)
	
	f.曼瑟尔Mansar = BMansar曼瑟尔
	f.曼瑟尔Mansar.SetParent(f)
	
	f.难提拔檀那Nandivardhana = BNandivardhana难提拔檀那
	f.难提拔檀那Nandivardhana.SetParent(f)
	
	f.普拉纳Pavnar = BPavnar普拉纳
	f.普拉纳Pavnar.SetParent(f)
	
	f.罗摩耆厘Ramagiri = BRamagiri罗摩耆厘
	f.罗摩耆厘Ramagiri.SetParent(f)
	
}
