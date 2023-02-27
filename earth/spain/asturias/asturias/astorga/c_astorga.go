package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AstorgaCounty interface {
    feud.County
    BAstorga阿斯托加() 	feud.Barony
    BBembibre本比夫雷() 	feud.Barony
    BCacabelos卡卡韦洛斯() 	feud.Barony
    BCamponaraya坎波那拉亚() 	feud.Barony
    BFabero法韦罗() 	feud.Barony
    BPonferrada蓬费拉达() 	feud.Barony
    BRibadelago里瓦德拉戈() 	feud.Barony
    BToreno托雷诺() 	feud.Barony
}

type 阿斯托加AstorgaCounty struct {
	feud.BaseCounty
	阿斯托加Astorga 	feud.Barony
	本比夫雷Bembibre 	feud.Barony
	卡卡韦洛斯Cacabelos 	feud.Barony
	坎波那拉亚Camponaraya 	feud.Barony
	法韦罗Fabero 	feud.Barony
	蓬费拉达Ponferrada 	feud.Barony
	里瓦德拉戈Ribadelago 	feud.Barony
	托雷诺Toreno 	feud.Barony
}

func (c *阿斯托加AstorgaCounty) BAstorga阿斯托加() feud.Barony {
	return c.阿斯托加Astorga
}
    
func (c *阿斯托加AstorgaCounty) BBembibre本比夫雷() feud.Barony {
	return c.本比夫雷Bembibre
}
    
func (c *阿斯托加AstorgaCounty) BCacabelos卡卡韦洛斯() feud.Barony {
	return c.卡卡韦洛斯Cacabelos
}
    
func (c *阿斯托加AstorgaCounty) BCamponaraya坎波那拉亚() feud.Barony {
	return c.坎波那拉亚Camponaraya
}
    
func (c *阿斯托加AstorgaCounty) BFabero法韦罗() feud.Barony {
	return c.法韦罗Fabero
}
    
func (c *阿斯托加AstorgaCounty) BPonferrada蓬费拉达() feud.Barony {
	return c.蓬费拉达Ponferrada
}
    
func (c *阿斯托加AstorgaCounty) BRibadelago里瓦德拉戈() feud.Barony {
	return c.里瓦德拉戈Ribadelago
}
    
func (c *阿斯托加AstorgaCounty) BToreno托雷诺() feud.Barony {
	return c.托雷诺Toreno
}
    
var CAstorga阿斯托加 AstorgaCounty = &阿斯托加AstorgaCounty{}

func init() {
	f := CAstorga阿斯托加.(*阿斯托加AstorgaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "189",
		Title:     "astorga",
		TitleName: "阿斯托加",
		TitleCode: "c_astorga",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯托加Astorga = BAstorga阿斯托加
	f.阿斯托加Astorga.SetParent(f)
	
	f.本比夫雷Bembibre = BBembibre本比夫雷
	f.本比夫雷Bembibre.SetParent(f)
	
	f.卡卡韦洛斯Cacabelos = BCacabelos卡卡韦洛斯
	f.卡卡韦洛斯Cacabelos.SetParent(f)
	
	f.坎波那拉亚Camponaraya = BCamponaraya坎波那拉亚
	f.坎波那拉亚Camponaraya.SetParent(f)
	
	f.法韦罗Fabero = BFabero法韦罗
	f.法韦罗Fabero.SetParent(f)
	
	f.蓬费拉达Ponferrada = BPonferrada蓬费拉达
	f.蓬费拉达Ponferrada.SetParent(f)
	
	f.里瓦德拉戈Ribadelago = BRibadelago里瓦德拉戈
	f.里瓦德拉戈Ribadelago.SetParent(f)
	
	f.托雷诺Toreno = BToreno托雷诺
	f.托雷诺Toreno.SetParent(f)
	
}
