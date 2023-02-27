package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Asturias_de_santillanaCounty interface {
    feud.County
    BCamargo卡马戈() 	feud.Barony
    BCastrourdiales乌迪亚莱斯堡() 	feud.Barony
    BLaredo拉雷多() 	feud.Barony
    BReinosa雷诺萨() 	feud.Barony
    BSantander桑坦德() 	feud.Barony
    BSantillanadelmar滨海桑蒂利亚纳() 	feud.Barony
    BSantona桑托尼亚() 	feud.Barony
    BSanvicente圣维森特() 	feud.Barony
    BSuances苏安塞斯() 	feud.Barony
}

type 阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty struct {
	feud.BaseCounty
	卡马戈Camargo 	feud.Barony
	乌迪亚莱斯堡Castrourdiales 	feud.Barony
	拉雷多Laredo 	feud.Barony
	雷诺萨Reinosa 	feud.Barony
	桑坦德Santander 	feud.Barony
	滨海桑蒂利亚纳Santillanadelmar 	feud.Barony
	桑托尼亚Santona 	feud.Barony
	圣维森特Sanvicente 	feud.Barony
	苏安塞斯Suances 	feud.Barony
}

func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BCamargo卡马戈() feud.Barony {
	return c.卡马戈Camargo
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BCastrourdiales乌迪亚莱斯堡() feud.Barony {
	return c.乌迪亚莱斯堡Castrourdiales
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BLaredo拉雷多() feud.Barony {
	return c.拉雷多Laredo
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BReinosa雷诺萨() feud.Barony {
	return c.雷诺萨Reinosa
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BSantander桑坦德() feud.Barony {
	return c.桑坦德Santander
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BSantillanadelmar滨海桑蒂利亚纳() feud.Barony {
	return c.滨海桑蒂利亚纳Santillanadelmar
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BSantona桑托尼亚() feud.Barony {
	return c.桑托尼亚Santona
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BSanvicente圣维森特() feud.Barony {
	return c.圣维森特Sanvicente
}
    
func (c *阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty) BSuances苏安塞斯() feud.Barony {
	return c.苏安塞斯Suances
}
    
var CAsturias_de_santillana阿斯图里亚斯德桑蒂利亚纳 Asturias_de_santillanaCounty = &阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty{}

func init() {
	f := CAsturias_de_santillana阿斯图里亚斯德桑蒂利亚纳.(*阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "154",
		Title:     "asturias_de_santillana",
		TitleName: "阿斯图里亚斯德桑蒂利亚纳",
		TitleCode: "c_asturias_de_santillana",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡马戈Camargo = BCamargo卡马戈
	f.卡马戈Camargo.SetParent(f)
	
	f.乌迪亚莱斯堡Castrourdiales = BCastrourdiales乌迪亚莱斯堡
	f.乌迪亚莱斯堡Castrourdiales.SetParent(f)
	
	f.拉雷多Laredo = BLaredo拉雷多
	f.拉雷多Laredo.SetParent(f)
	
	f.雷诺萨Reinosa = BReinosa雷诺萨
	f.雷诺萨Reinosa.SetParent(f)
	
	f.桑坦德Santander = BSantander桑坦德
	f.桑坦德Santander.SetParent(f)
	
	f.滨海桑蒂利亚纳Santillanadelmar = BSantillanadelmar滨海桑蒂利亚纳
	f.滨海桑蒂利亚纳Santillanadelmar.SetParent(f)
	
	f.桑托尼亚Santona = BSantona桑托尼亚
	f.桑托尼亚Santona.SetParent(f)
	
	f.圣维森特Sanvicente = BSanvicente圣维森特
	f.圣维森特Sanvicente.SetParent(f)
	
	f.苏安塞斯Suances = BSuances苏安塞斯
	f.苏安塞斯Suances.SetParent(f)
	
}
