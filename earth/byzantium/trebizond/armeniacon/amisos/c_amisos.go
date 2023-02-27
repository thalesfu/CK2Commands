package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmisosCounty interface {
    feud.County
    BAmasia阿马希亚() 	feud.Barony
    BAmisos阿米索斯() 	feud.Barony
    BEupatoria叶夫帕托里亚() 	feud.Barony
    BNeokaisarea新凯撒利亚() 	feud.Barony
    BPhadisane法蒂森() 	feud.Barony
    BPhazimonitis法齐莫尼蒂斯() 	feud.Barony
    BThermodon忒耳摩冬() 	feud.Barony
}

type 阿米索斯AmisosCounty struct {
	feud.BaseCounty
	阿马希亚Amasia 	feud.Barony
	阿米索斯Amisos 	feud.Barony
	叶夫帕托里亚Eupatoria 	feud.Barony
	新凯撒利亚Neokaisarea 	feud.Barony
	法蒂森Phadisane 	feud.Barony
	法齐莫尼蒂斯Phazimonitis 	feud.Barony
	忒耳摩冬Thermodon 	feud.Barony
}

func (c *阿米索斯AmisosCounty) BAmasia阿马希亚() feud.Barony {
	return c.阿马希亚Amasia
}
    
func (c *阿米索斯AmisosCounty) BAmisos阿米索斯() feud.Barony {
	return c.阿米索斯Amisos
}
    
func (c *阿米索斯AmisosCounty) BEupatoria叶夫帕托里亚() feud.Barony {
	return c.叶夫帕托里亚Eupatoria
}
    
func (c *阿米索斯AmisosCounty) BNeokaisarea新凯撒利亚() feud.Barony {
	return c.新凯撒利亚Neokaisarea
}
    
func (c *阿米索斯AmisosCounty) BPhadisane法蒂森() feud.Barony {
	return c.法蒂森Phadisane
}
    
func (c *阿米索斯AmisosCounty) BPhazimonitis法齐莫尼蒂斯() feud.Barony {
	return c.法齐莫尼蒂斯Phazimonitis
}
    
func (c *阿米索斯AmisosCounty) BThermodon忒耳摩冬() feud.Barony {
	return c.忒耳摩冬Thermodon
}
    
var CAmisos阿米索斯 AmisosCounty = &阿米索斯AmisosCounty{}

func init() {
	f := CAmisos阿米索斯.(*阿米索斯AmisosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "738",
		Title:     "amisos",
		TitleName: "阿米索斯",
		TitleCode: "c_amisos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马希亚Amasia = BAmasia阿马希亚
	f.阿马希亚Amasia.SetParent(f)
	
	f.阿米索斯Amisos = BAmisos阿米索斯
	f.阿米索斯Amisos.SetParent(f)
	
	f.叶夫帕托里亚Eupatoria = BEupatoria叶夫帕托里亚
	f.叶夫帕托里亚Eupatoria.SetParent(f)
	
	f.新凯撒利亚Neokaisarea = BNeokaisarea新凯撒利亚
	f.新凯撒利亚Neokaisarea.SetParent(f)
	
	f.法蒂森Phadisane = BPhadisane法蒂森
	f.法蒂森Phadisane.SetParent(f)
	
	f.法齐莫尼蒂斯Phazimonitis = BPhazimonitis法齐莫尼蒂斯
	f.法齐莫尼蒂斯Phazimonitis.SetParent(f)
	
	f.忒耳摩冬Thermodon = BThermodon忒耳摩冬
	f.忒耳摩冬Thermodon.SetParent(f)
	
}
