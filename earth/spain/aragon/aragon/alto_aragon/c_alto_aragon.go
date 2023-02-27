package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Alto_aragonCounty interface {
    feud.County
    BAlmudevar阿尔穆德瓦尔() 	feud.Barony
    BAlquezar阿尔克萨尔() 	feud.Barony
    BAyerbe阿耶韦() 	feud.Barony
    BBarbastro巴瓦斯特罗() 	feud.Barony
    BHuesca韦斯卡() 	feud.Barony
    BJaca哈卡() 	feud.Barony
    BLoarre洛阿雷() 	feud.Barony
    BUncastillo温卡斯蒂略() 	feud.Barony
}

type 上阿拉贡Alto_aragonCounty struct {
	feud.BaseCounty
	阿尔穆德瓦尔Almudevar 	feud.Barony
	阿尔克萨尔Alquezar 	feud.Barony
	阿耶韦Ayerbe 	feud.Barony
	巴瓦斯特罗Barbastro 	feud.Barony
	韦斯卡Huesca 	feud.Barony
	哈卡Jaca 	feud.Barony
	洛阿雷Loarre 	feud.Barony
	温卡斯蒂略Uncastillo 	feud.Barony
}

func (c *上阿拉贡Alto_aragonCounty) BAlmudevar阿尔穆德瓦尔() feud.Barony {
	return c.阿尔穆德瓦尔Almudevar
}
    
func (c *上阿拉贡Alto_aragonCounty) BAlquezar阿尔克萨尔() feud.Barony {
	return c.阿尔克萨尔Alquezar
}
    
func (c *上阿拉贡Alto_aragonCounty) BAyerbe阿耶韦() feud.Barony {
	return c.阿耶韦Ayerbe
}
    
func (c *上阿拉贡Alto_aragonCounty) BBarbastro巴瓦斯特罗() feud.Barony {
	return c.巴瓦斯特罗Barbastro
}
    
func (c *上阿拉贡Alto_aragonCounty) BHuesca韦斯卡() feud.Barony {
	return c.韦斯卡Huesca
}
    
func (c *上阿拉贡Alto_aragonCounty) BJaca哈卡() feud.Barony {
	return c.哈卡Jaca
}
    
func (c *上阿拉贡Alto_aragonCounty) BLoarre洛阿雷() feud.Barony {
	return c.洛阿雷Loarre
}
    
func (c *上阿拉贡Alto_aragonCounty) BUncastillo温卡斯蒂略() feud.Barony {
	return c.温卡斯蒂略Uncastillo
}
    
var CAlto_aragon上阿拉贡 Alto_aragonCounty = &上阿拉贡Alto_aragonCounty{}

func init() {
	f := CAlto_aragon上阿拉贡.(*上阿拉贡Alto_aragonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "207",
		Title:     "alto_aragon",
		TitleName: "上阿拉贡",
		TitleCode: "c_alto_aragon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔穆德瓦尔Almudevar = BAlmudevar阿尔穆德瓦尔
	f.阿尔穆德瓦尔Almudevar.SetParent(f)
	
	f.阿尔克萨尔Alquezar = BAlquezar阿尔克萨尔
	f.阿尔克萨尔Alquezar.SetParent(f)
	
	f.阿耶韦Ayerbe = BAyerbe阿耶韦
	f.阿耶韦Ayerbe.SetParent(f)
	
	f.巴瓦斯特罗Barbastro = BBarbastro巴瓦斯特罗
	f.巴瓦斯特罗Barbastro.SetParent(f)
	
	f.韦斯卡Huesca = BHuesca韦斯卡
	f.韦斯卡Huesca.SetParent(f)
	
	f.哈卡Jaca = BJaca哈卡
	f.哈卡Jaca.SetParent(f)
	
	f.洛阿雷Loarre = BLoarre洛阿雷
	f.洛阿雷Loarre.SetParent(f)
	
	f.温卡斯蒂略Uncastillo = BUncastillo温卡斯蒂略
	f.温卡斯蒂略Uncastillo.SetParent(f)
	
}
