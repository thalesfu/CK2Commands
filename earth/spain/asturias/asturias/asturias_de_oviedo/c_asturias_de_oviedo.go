package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Asturias_de_oviedoCounty interface {
    feud.County
    BCangasdelnarcea纳尔塞阿河畔坎加斯() 	feud.Barony
    BCangasdeonis坎加斯德奥尼斯() 	feud.Barony
    BGijon希洪() 	feud.Barony
    BLuarca卢阿尔卡() 	feud.Barony
    BNorena诺雷尼亚() 	feud.Barony
    BOviedo奥维耶多() 	feud.Barony
    BTineo蒂内奥() 	feud.Barony
    BVillaviciosa比利亚维西奥萨() 	feud.Barony
}

type 阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty struct {
	feud.BaseCounty
	纳尔塞阿河畔坎加斯Cangasdelnarcea 	feud.Barony
	坎加斯德奥尼斯Cangasdeonis 	feud.Barony
	希洪Gijon 	feud.Barony
	卢阿尔卡Luarca 	feud.Barony
	诺雷尼亚Norena 	feud.Barony
	奥维耶多Oviedo 	feud.Barony
	蒂内奥Tineo 	feud.Barony
	比利亚维西奥萨Villaviciosa 	feud.Barony
}

func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BCangasdelnarcea纳尔塞阿河畔坎加斯() feud.Barony {
	return c.纳尔塞阿河畔坎加斯Cangasdelnarcea
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BCangasdeonis坎加斯德奥尼斯() feud.Barony {
	return c.坎加斯德奥尼斯Cangasdeonis
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BGijon希洪() feud.Barony {
	return c.希洪Gijon
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BLuarca卢阿尔卡() feud.Barony {
	return c.卢阿尔卡Luarca
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BNorena诺雷尼亚() feud.Barony {
	return c.诺雷尼亚Norena
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BOviedo奥维耶多() feud.Barony {
	return c.奥维耶多Oviedo
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BTineo蒂内奥() feud.Barony {
	return c.蒂内奥Tineo
}
    
func (c *阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty) BVillaviciosa比利亚维西奥萨() feud.Barony {
	return c.比利亚维西奥萨Villaviciosa
}
    
var CAsturias_de_oviedo阿斯图里亚斯德奥维耶多 Asturias_de_oviedoCounty = &阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty{}

func init() {
	f := CAsturias_de_oviedo阿斯图里亚斯德奥维耶多.(*阿斯图里亚斯德奥维耶多Asturias_de_oviedoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "155",
		Title:     "asturias_de_oviedo",
		TitleName: "阿斯图里亚斯德奥维耶多",
		TitleCode: "c_asturias_de_oviedo",
		Baronies:  map[string]feud.Barony{},
	}

	f.纳尔塞阿河畔坎加斯Cangasdelnarcea = BCangasdelnarcea纳尔塞阿河畔坎加斯
	f.纳尔塞阿河畔坎加斯Cangasdelnarcea.SetParent(f)
	
	f.坎加斯德奥尼斯Cangasdeonis = BCangasdeonis坎加斯德奥尼斯
	f.坎加斯德奥尼斯Cangasdeonis.SetParent(f)
	
	f.希洪Gijon = BGijon希洪
	f.希洪Gijon.SetParent(f)
	
	f.卢阿尔卡Luarca = BLuarca卢阿尔卡
	f.卢阿尔卡Luarca.SetParent(f)
	
	f.诺雷尼亚Norena = BNorena诺雷尼亚
	f.诺雷尼亚Norena.SetParent(f)
	
	f.奥维耶多Oviedo = BOviedo奥维耶多
	f.奥维耶多Oviedo.SetParent(f)
	
	f.蒂内奥Tineo = BTineo蒂内奥
	f.蒂内奥Tineo.SetParent(f)
	
	f.比利亚维西奥萨Villaviciosa = BVillaviciosa比利亚维西奥萨
	f.比利亚维西奥萨Villaviciosa.SetParent(f)
	
}
