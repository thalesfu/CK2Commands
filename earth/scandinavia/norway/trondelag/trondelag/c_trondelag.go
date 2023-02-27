package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrondelagCounty interface {
    feud.County
    BAustratt奥斯特洛特() 	feud.Barony
    BHaltalen霍尔托伦() 	feud.Barony
    BMaere梅勒() 	feud.Barony
    BNidaros尼达罗斯() 	feud.Barony
    BOra厄拉() 	feud.Barony
    BSteinvikholm斯泰因维克霍姆() 	feud.Barony
    BSverresborg斯韦勒斯堡() 	feud.Barony
    BTrondheim特隆赫姆() 	feud.Barony
}

type 尼达罗斯TrondelagCounty struct {
	feud.BaseCounty
	奥斯特洛特Austratt 	feud.Barony
	霍尔托伦Haltalen 	feud.Barony
	梅勒Maere 	feud.Barony
	尼达罗斯Nidaros 	feud.Barony
	厄拉Ora 	feud.Barony
	斯泰因维克霍姆Steinvikholm 	feud.Barony
	斯韦勒斯堡Sverresborg 	feud.Barony
	特隆赫姆Trondheim 	feud.Barony
}

func (c *尼达罗斯TrondelagCounty) BAustratt奥斯特洛特() feud.Barony {
	return c.奥斯特洛特Austratt
}
    
func (c *尼达罗斯TrondelagCounty) BHaltalen霍尔托伦() feud.Barony {
	return c.霍尔托伦Haltalen
}
    
func (c *尼达罗斯TrondelagCounty) BMaere梅勒() feud.Barony {
	return c.梅勒Maere
}
    
func (c *尼达罗斯TrondelagCounty) BNidaros尼达罗斯() feud.Barony {
	return c.尼达罗斯Nidaros
}
    
func (c *尼达罗斯TrondelagCounty) BOra厄拉() feud.Barony {
	return c.厄拉Ora
}
    
func (c *尼达罗斯TrondelagCounty) BSteinvikholm斯泰因维克霍姆() feud.Barony {
	return c.斯泰因维克霍姆Steinvikholm
}
    
func (c *尼达罗斯TrondelagCounty) BSverresborg斯韦勒斯堡() feud.Barony {
	return c.斯韦勒斯堡Sverresborg
}
    
func (c *尼达罗斯TrondelagCounty) BTrondheim特隆赫姆() feud.Barony {
	return c.特隆赫姆Trondheim
}
    
var CTrondelag尼达罗斯 TrondelagCounty = &尼达罗斯TrondelagCounty{}

func init() {
	f := CTrondelag尼达罗斯.(*尼达罗斯TrondelagCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "275",
		Title:     "trondelag",
		TitleName: "尼达罗斯",
		TitleCode: "c_trondelag",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥斯特洛特Austratt = BAustratt奥斯特洛特
	f.奥斯特洛特Austratt.SetParent(f)
	
	f.霍尔托伦Haltalen = BHaltalen霍尔托伦
	f.霍尔托伦Haltalen.SetParent(f)
	
	f.梅勒Maere = BMaere梅勒
	f.梅勒Maere.SetParent(f)
	
	f.尼达罗斯Nidaros = BNidaros尼达罗斯
	f.尼达罗斯Nidaros.SetParent(f)
	
	f.厄拉Ora = BOra厄拉
	f.厄拉Ora.SetParent(f)
	
	f.斯泰因维克霍姆Steinvikholm = BSteinvikholm斯泰因维克霍姆
	f.斯泰因维克霍姆Steinvikholm.SetParent(f)
	
	f.斯韦勒斯堡Sverresborg = BSverresborg斯韦勒斯堡
	f.斯韦勒斯堡Sverresborg.SetParent(f)
	
	f.特隆赫姆Trondheim = BTrondheim特隆赫姆
	f.特隆赫姆Trondheim.SetParent(f)
	
}
