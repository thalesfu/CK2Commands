package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SrirangapatnaCounty interface {
    feud.County
    BBelapura贝拉普拉() 	feud.Barony
    BMysore摩索罗() 	feud.Barony
    BNanjarayapattana难加罗亚帕提塔纳() 	feud.Barony
    BPanasoge波那索戈() 	feud.Barony
    BPeriyapatna佩里亚帕提纳() 	feud.Barony
    BShravanabelagola室罗伐拿白泽() 	feud.Barony
    BSrirangapatna室利浪伽钵多那() 	feud.Barony
}

type 室利浪伽钵多那SrirangapatnaCounty struct {
	feud.BaseCounty
	贝拉普拉Belapura 	feud.Barony
	摩索罗Mysore 	feud.Barony
	难加罗亚帕提塔纳Nanjarayapattana 	feud.Barony
	波那索戈Panasoge 	feud.Barony
	佩里亚帕提纳Periyapatna 	feud.Barony
	室罗伐拿白泽Shravanabelagola 	feud.Barony
	室利浪伽钵多那Srirangapatna 	feud.Barony
}

func (c *室利浪伽钵多那SrirangapatnaCounty) BBelapura贝拉普拉() feud.Barony {
	return c.贝拉普拉Belapura
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BMysore摩索罗() feud.Barony {
	return c.摩索罗Mysore
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BNanjarayapattana难加罗亚帕提塔纳() feud.Barony {
	return c.难加罗亚帕提塔纳Nanjarayapattana
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BPanasoge波那索戈() feud.Barony {
	return c.波那索戈Panasoge
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BPeriyapatna佩里亚帕提纳() feud.Barony {
	return c.佩里亚帕提纳Periyapatna
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BShravanabelagola室罗伐拿白泽() feud.Barony {
	return c.室罗伐拿白泽Shravanabelagola
}
    
func (c *室利浪伽钵多那SrirangapatnaCounty) BSrirangapatna室利浪伽钵多那() feud.Barony {
	return c.室利浪伽钵多那Srirangapatna
}
    
var CSrirangapatna室利浪伽钵多那 SrirangapatnaCounty = &室利浪伽钵多那SrirangapatnaCounty{}

func init() {
	f := CSrirangapatna室利浪伽钵多那.(*室利浪伽钵多那SrirangapatnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1196",
		Title:     "srirangapatna",
		TitleName: "室利浪伽钵多那",
		TitleCode: "c_srirangapatna",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝拉普拉Belapura = BBelapura贝拉普拉
	f.贝拉普拉Belapura.SetParent(f)
	
	f.摩索罗Mysore = BMysore摩索罗
	f.摩索罗Mysore.SetParent(f)
	
	f.难加罗亚帕提塔纳Nanjarayapattana = BNanjarayapattana难加罗亚帕提塔纳
	f.难加罗亚帕提塔纳Nanjarayapattana.SetParent(f)
	
	f.波那索戈Panasoge = BPanasoge波那索戈
	f.波那索戈Panasoge.SetParent(f)
	
	f.佩里亚帕提纳Periyapatna = BPeriyapatna佩里亚帕提纳
	f.佩里亚帕提纳Periyapatna.SetParent(f)
	
	f.室罗伐拿白泽Shravanabelagola = BShravanabelagola室罗伐拿白泽
	f.室罗伐拿白泽Shravanabelagola.SetParent(f)
	
	f.室利浪伽钵多那Srirangapatna = BSrirangapatna室利浪伽钵多那
	f.室利浪伽钵多那Srirangapatna.SetParent(f)
	
}
