package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VaudCounty interface {
    feud.County
    BChateaudoex代堡() 	feud.Barony
    BChillon希永() 	feud.Barony
    BEchallens埃沙朗() 	feud.Barony
    BGrandson格朗松() 	feud.Barony
    BGruyere格吕耶尔() 	feud.Barony
    BLausanne洛桑() 	feud.Barony
    BOrbe奥尔布() 	feud.Barony
    BOuchy乌希() 	feud.Barony
    BPayerne帕耶讷() 	feud.Barony
    BRomainmotier罗曼莫捷() 	feud.Barony
    BYverdon伊韦尔东() 	feud.Barony
}

type 沃VaudCounty struct {
	feud.BaseCounty
	代堡Chateaudoex 	feud.Barony
	希永Chillon 	feud.Barony
	埃沙朗Echallens 	feud.Barony
	格朗松Grandson 	feud.Barony
	格吕耶尔Gruyere 	feud.Barony
	洛桑Lausanne 	feud.Barony
	奥尔布Orbe 	feud.Barony
	乌希Ouchy 	feud.Barony
	帕耶讷Payerne 	feud.Barony
	罗曼莫捷Romainmotier 	feud.Barony
	伊韦尔东Yverdon 	feud.Barony
}

func (c *沃VaudCounty) BChateaudoex代堡() feud.Barony {
	return c.代堡Chateaudoex
}
    
func (c *沃VaudCounty) BChillon希永() feud.Barony {
	return c.希永Chillon
}
    
func (c *沃VaudCounty) BEchallens埃沙朗() feud.Barony {
	return c.埃沙朗Echallens
}
    
func (c *沃VaudCounty) BGrandson格朗松() feud.Barony {
	return c.格朗松Grandson
}
    
func (c *沃VaudCounty) BGruyere格吕耶尔() feud.Barony {
	return c.格吕耶尔Gruyere
}
    
func (c *沃VaudCounty) BLausanne洛桑() feud.Barony {
	return c.洛桑Lausanne
}
    
func (c *沃VaudCounty) BOrbe奥尔布() feud.Barony {
	return c.奥尔布Orbe
}
    
func (c *沃VaudCounty) BOuchy乌希() feud.Barony {
	return c.乌希Ouchy
}
    
func (c *沃VaudCounty) BPayerne帕耶讷() feud.Barony {
	return c.帕耶讷Payerne
}
    
func (c *沃VaudCounty) BRomainmotier罗曼莫捷() feud.Barony {
	return c.罗曼莫捷Romainmotier
}
    
func (c *沃VaudCounty) BYverdon伊韦尔东() feud.Barony {
	return c.伊韦尔东Yverdon
}
    
var CVaud沃 VaudCounty = &沃VaudCounty{}

func init() {
	f := CVaud沃.(*沃VaudCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1611",
		Title:     "vaud",
		TitleName: "沃",
		TitleCode: "c_vaud",
		Baronies:  map[string]feud.Barony{},
	}

	f.代堡Chateaudoex = BChateaudoex代堡
	f.代堡Chateaudoex.SetParent(f)
	
	f.希永Chillon = BChillon希永
	f.希永Chillon.SetParent(f)
	
	f.埃沙朗Echallens = BEchallens埃沙朗
	f.埃沙朗Echallens.SetParent(f)
	
	f.格朗松Grandson = BGrandson格朗松
	f.格朗松Grandson.SetParent(f)
	
	f.格吕耶尔Gruyere = BGruyere格吕耶尔
	f.格吕耶尔Gruyere.SetParent(f)
	
	f.洛桑Lausanne = BLausanne洛桑
	f.洛桑Lausanne.SetParent(f)
	
	f.奥尔布Orbe = BOrbe奥尔布
	f.奥尔布Orbe.SetParent(f)
	
	f.乌希Ouchy = BOuchy乌希
	f.乌希Ouchy.SetParent(f)
	
	f.帕耶讷Payerne = BPayerne帕耶讷
	f.帕耶讷Payerne.SetParent(f)
	
	f.罗曼莫捷Romainmotier = BRomainmotier罗曼莫捷
	f.罗曼莫捷Romainmotier.SetParent(f)
	
	f.伊韦尔东Yverdon = BYverdon伊韦尔东
	f.伊韦尔东Yverdon.SetParent(f)
	
}
