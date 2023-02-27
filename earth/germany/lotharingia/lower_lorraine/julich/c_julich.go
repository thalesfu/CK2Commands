package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JulichCounty interface {
    feud.County
    BAachen亚琛() 	feud.Barony
    BDuren迪伦() 	feud.Barony
    BGeldern盖尔登() 	feud.Barony
    BJulich于利希() 	feud.Barony
    BMoers默尔斯() 	feud.Barony
    BMonschau蒙绍() 	feud.Barony
    BPrum普吕姆() 	feud.Barony
    BRoermond鲁尔蒙德() 	feud.Barony
}

type 于利希JulichCounty struct {
	feud.BaseCounty
	亚琛Aachen 	feud.Barony
	迪伦Duren 	feud.Barony
	盖尔登Geldern 	feud.Barony
	于利希Julich 	feud.Barony
	默尔斯Moers 	feud.Barony
	蒙绍Monschau 	feud.Barony
	普吕姆Prum 	feud.Barony
	鲁尔蒙德Roermond 	feud.Barony
}

func (c *于利希JulichCounty) BAachen亚琛() feud.Barony {
	return c.亚琛Aachen
}
    
func (c *于利希JulichCounty) BDuren迪伦() feud.Barony {
	return c.迪伦Duren
}
    
func (c *于利希JulichCounty) BGeldern盖尔登() feud.Barony {
	return c.盖尔登Geldern
}
    
func (c *于利希JulichCounty) BJulich于利希() feud.Barony {
	return c.于利希Julich
}
    
func (c *于利希JulichCounty) BMoers默尔斯() feud.Barony {
	return c.默尔斯Moers
}
    
func (c *于利希JulichCounty) BMonschau蒙绍() feud.Barony {
	return c.蒙绍Monschau
}
    
func (c *于利希JulichCounty) BPrum普吕姆() feud.Barony {
	return c.普吕姆Prum
}
    
func (c *于利希JulichCounty) BRoermond鲁尔蒙德() feud.Barony {
	return c.鲁尔蒙德Roermond
}
    
var CJulich于利希 JulichCounty = &于利希JulichCounty{}

func init() {
	f := CJulich于利希.(*于利希JulichCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "90",
		Title:     "julich",
		TitleName: "于利希",
		TitleCode: "c_julich",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚琛Aachen = BAachen亚琛
	f.亚琛Aachen.SetParent(f)
	
	f.迪伦Duren = BDuren迪伦
	f.迪伦Duren.SetParent(f)
	
	f.盖尔登Geldern = BGeldern盖尔登
	f.盖尔登Geldern.SetParent(f)
	
	f.于利希Julich = BJulich于利希
	f.于利希Julich.SetParent(f)
	
	f.默尔斯Moers = BMoers默尔斯
	f.默尔斯Moers.SetParent(f)
	
	f.蒙绍Monschau = BMonschau蒙绍
	f.蒙绍Monschau.SetParent(f)
	
	f.普吕姆Prum = BPrum普吕姆
	f.普吕姆Prum.SetParent(f)
	
	f.鲁尔蒙德Roermond = BRoermond鲁尔蒙德
	f.鲁尔蒙德Roermond.SetParent(f)
	
}
