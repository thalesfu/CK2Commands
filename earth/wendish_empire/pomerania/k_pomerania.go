package pomerania

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/brandenburg"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/lausitz"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/mecklemburg"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/meissen"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pomeralia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pommerania"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/rugen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PomeraniaKingdom interface {
    feud.Kingdom
    DBrandenburg勃兰登堡() 	brandenburg.BrandenburgDuke
    DLausitz劳西茨() 	lausitz.LausitzDuke
    DMecklemburg梅克伦堡() 	mecklemburg.MecklemburgDuke
    DMeissen迈森() 	meissen.MeissenDuke
    DPomeralia波米拉利亚() 	pomeralia.PomeraliaDuke
    DPommerania波美拉尼亚() 	pommerania.PommeraniaDuke
    DRugen吕根() 	rugen.RugenDuke
}

type 波美拉尼亚PomeraniaKingdom struct {
	feud.BaseKingdom
	勃兰登堡Brandenburg 	brandenburg.BrandenburgDuke
	劳西茨Lausitz 	lausitz.LausitzDuke
	梅克伦堡Mecklemburg 	mecklemburg.MecklemburgDuke
	迈森Meissen 	meissen.MeissenDuke
	波米拉利亚Pomeralia 	pomeralia.PomeraliaDuke
	波美拉尼亚Pommerania 	pommerania.PommeraniaDuke
	吕根Rugen 	rugen.RugenDuke
}

func (k *波美拉尼亚PomeraniaKingdom) DBrandenburg勃兰登堡() brandenburg.BrandenburgDuke {
	return k.勃兰登堡Brandenburg
}
    
func (k *波美拉尼亚PomeraniaKingdom) DLausitz劳西茨() lausitz.LausitzDuke {
	return k.劳西茨Lausitz
}
    
func (k *波美拉尼亚PomeraniaKingdom) DMecklemburg梅克伦堡() mecklemburg.MecklemburgDuke {
	return k.梅克伦堡Mecklemburg
}
    
func (k *波美拉尼亚PomeraniaKingdom) DMeissen迈森() meissen.MeissenDuke {
	return k.迈森Meissen
}
    
func (k *波美拉尼亚PomeraniaKingdom) DPomeralia波米拉利亚() pomeralia.PomeraliaDuke {
	return k.波米拉利亚Pomeralia
}
    
func (k *波美拉尼亚PomeraniaKingdom) DPommerania波美拉尼亚() pommerania.PommeraniaDuke {
	return k.波美拉尼亚Pommerania
}
    
func (k *波美拉尼亚PomeraniaKingdom) DRugen吕根() rugen.RugenDuke {
	return k.吕根Rugen
}
    
var KPomerania波美拉尼亚 PomeraniaKingdom = &波美拉尼亚PomeraniaKingdom{}

func init() {
	f := KPomerania波美拉尼亚.(*波美拉尼亚PomeraniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "pomerania",
		TitleName: "波美拉尼亚",
		TitleCode: "k_pomerania",
		Dukes:  map[string]feud.Duke{},
	}

	f.勃兰登堡Brandenburg = brandenburg.DBrandenburg勃兰登堡
	f.勃兰登堡Brandenburg.SetParent(f)
	
	f.劳西茨Lausitz = lausitz.DLausitz劳西茨
	f.劳西茨Lausitz.SetParent(f)
	
	f.梅克伦堡Mecklemburg = mecklemburg.DMecklemburg梅克伦堡
	f.梅克伦堡Mecklemburg.SetParent(f)
	
	f.迈森Meissen = meissen.DMeissen迈森
	f.迈森Meissen.SetParent(f)
	
	f.波米拉利亚Pomeralia = pomeralia.DPomeralia波米拉利亚
	f.波米拉利亚Pomeralia.SetParent(f)
	
	f.波美拉尼亚Pommerania = pommerania.DPommerania波美拉尼亚
	f.波美拉尼亚Pommerania.SetParent(f)
	
	f.吕根Rugen = rugen.DRugen吕根
	f.吕根Rugen.SetParent(f)
	
}
