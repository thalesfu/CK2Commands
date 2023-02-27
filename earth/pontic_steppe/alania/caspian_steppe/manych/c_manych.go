package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ManychCounty interface {
    feud.County
    BKarabulak卡拉布拉克() 	feud.Barony
    BKunay库奈() 	feud.Barony
    BMakhach马哈奇() 	feud.Barony
    BMakhmudmekteb马赫穆德_梅克捷布() 	feud.Barony
    BMalgobekbalka马尔戈别克_巴尔卡() 	feud.Barony
    BMozdok莫兹多克() 	feud.Barony
    BNartkala纳尔特卡拉() 	feud.Barony
}

type 马内奇ManychCounty struct {
	feud.BaseCounty
	卡拉布拉克Karabulak 	feud.Barony
	库奈Kunay 	feud.Barony
	马哈奇Makhach 	feud.Barony
	马赫穆德_梅克捷布Makhmudmekteb 	feud.Barony
	马尔戈别克_巴尔卡Malgobekbalka 	feud.Barony
	莫兹多克Mozdok 	feud.Barony
	纳尔特卡拉Nartkala 	feud.Barony
}

func (c *马内奇ManychCounty) BKarabulak卡拉布拉克() feud.Barony {
	return c.卡拉布拉克Karabulak
}
    
func (c *马内奇ManychCounty) BKunay库奈() feud.Barony {
	return c.库奈Kunay
}
    
func (c *马内奇ManychCounty) BMakhach马哈奇() feud.Barony {
	return c.马哈奇Makhach
}
    
func (c *马内奇ManychCounty) BMakhmudmekteb马赫穆德_梅克捷布() feud.Barony {
	return c.马赫穆德_梅克捷布Makhmudmekteb
}
    
func (c *马内奇ManychCounty) BMalgobekbalka马尔戈别克_巴尔卡() feud.Barony {
	return c.马尔戈别克_巴尔卡Malgobekbalka
}
    
func (c *马内奇ManychCounty) BMozdok莫兹多克() feud.Barony {
	return c.莫兹多克Mozdok
}
    
func (c *马内奇ManychCounty) BNartkala纳尔特卡拉() feud.Barony {
	return c.纳尔特卡拉Nartkala
}
    
var CManych马内奇 ManychCounty = &马内奇ManychCounty{}

func init() {
	f := CManych马内奇.(*马内奇ManychCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "605",
		Title:     "manych",
		TitleName: "马内奇",
		TitleCode: "c_manych",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉布拉克Karabulak = BKarabulak卡拉布拉克
	f.卡拉布拉克Karabulak.SetParent(f)
	
	f.库奈Kunay = BKunay库奈
	f.库奈Kunay.SetParent(f)
	
	f.马哈奇Makhach = BMakhach马哈奇
	f.马哈奇Makhach.SetParent(f)
	
	f.马赫穆德_梅克捷布Makhmudmekteb = BMakhmudmekteb马赫穆德_梅克捷布
	f.马赫穆德_梅克捷布Makhmudmekteb.SetParent(f)
	
	f.马尔戈别克_巴尔卡Malgobekbalka = BMalgobekbalka马尔戈别克_巴尔卡
	f.马尔戈别克_巴尔卡Malgobekbalka.SetParent(f)
	
	f.莫兹多克Mozdok = BMozdok莫兹多克
	f.莫兹多克Mozdok.SetParent(f)
	
	f.纳尔特卡拉Nartkala = BNartkala纳尔特卡拉
	f.纳尔特卡拉Nartkala.SetParent(f)
	
}
