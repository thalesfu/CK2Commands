package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BraganzaCounty interface {
    feud.County
    BAzinhoso阿齐尼奥苏() 	feud.Barony
    BBraganza布拉干萨() 	feud.Barony
    BCastelomelhor梅略尔堡() 	feud.Barony
    BCastelorodrigo罗德里戈堡() 	feud.Barony
    BChaves沙维什() 	feud.Barony
    BMogadouro莫加多鲁() 	feud.Barony
    BTorredemoncorvo托里迪蒙科尔武() 	feud.Barony
    BVilareal雷阿尔城() 	feud.Barony
}

type 布拉干萨BraganzaCounty struct {
	feud.BaseCounty
	阿齐尼奥苏Azinhoso 	feud.Barony
	布拉干萨Braganza 	feud.Barony
	梅略尔堡Castelomelhor 	feud.Barony
	罗德里戈堡Castelorodrigo 	feud.Barony
	沙维什Chaves 	feud.Barony
	莫加多鲁Mogadouro 	feud.Barony
	托里迪蒙科尔武Torredemoncorvo 	feud.Barony
	雷阿尔城Vilareal 	feud.Barony
}

func (c *布拉干萨BraganzaCounty) BAzinhoso阿齐尼奥苏() feud.Barony {
	return c.阿齐尼奥苏Azinhoso
}
    
func (c *布拉干萨BraganzaCounty) BBraganza布拉干萨() feud.Barony {
	return c.布拉干萨Braganza
}
    
func (c *布拉干萨BraganzaCounty) BCastelomelhor梅略尔堡() feud.Barony {
	return c.梅略尔堡Castelomelhor
}
    
func (c *布拉干萨BraganzaCounty) BCastelorodrigo罗德里戈堡() feud.Barony {
	return c.罗德里戈堡Castelorodrigo
}
    
func (c *布拉干萨BraganzaCounty) BChaves沙维什() feud.Barony {
	return c.沙维什Chaves
}
    
func (c *布拉干萨BraganzaCounty) BMogadouro莫加多鲁() feud.Barony {
	return c.莫加多鲁Mogadouro
}
    
func (c *布拉干萨BraganzaCounty) BTorredemoncorvo托里迪蒙科尔武() feud.Barony {
	return c.托里迪蒙科尔武Torredemoncorvo
}
    
func (c *布拉干萨BraganzaCounty) BVilareal雷阿尔城() feud.Barony {
	return c.雷阿尔城Vilareal
}
    
var CBraganza布拉干萨 BraganzaCounty = &布拉干萨BraganzaCounty{}

func init() {
	f := CBraganza布拉干萨.(*布拉干萨BraganzaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "188",
		Title:     "braganza",
		TitleName: "布拉干萨",
		TitleCode: "c_braganza",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿齐尼奥苏Azinhoso = BAzinhoso阿齐尼奥苏
	f.阿齐尼奥苏Azinhoso.SetParent(f)
	
	f.布拉干萨Braganza = BBraganza布拉干萨
	f.布拉干萨Braganza.SetParent(f)
	
	f.梅略尔堡Castelomelhor = BCastelomelhor梅略尔堡
	f.梅略尔堡Castelomelhor.SetParent(f)
	
	f.罗德里戈堡Castelorodrigo = BCastelorodrigo罗德里戈堡
	f.罗德里戈堡Castelorodrigo.SetParent(f)
	
	f.沙维什Chaves = BChaves沙维什
	f.沙维什Chaves.SetParent(f)
	
	f.莫加多鲁Mogadouro = BMogadouro莫加多鲁
	f.莫加多鲁Mogadouro.SetParent(f)
	
	f.托里迪蒙科尔武Torredemoncorvo = BTorredemoncorvo托里迪蒙科尔武
	f.托里迪蒙科尔武Torredemoncorvo.SetParent(f)
	
	f.雷阿尔城Vilareal = BVilareal雷阿尔城
	f.雷阿尔城Vilareal.SetParent(f)
	
}
