package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArboreaCounty interface {
    feud.County
    BCabras卡布拉斯() 	feud.Barony
    BFordongianus福尔东贾努斯() 	feud.Barony
    BOristano奥里斯塔诺() 	feud.Barony
    BPabillonis帕比洛尼斯() 	feud.Barony
    BSanluri圣卢里() 	feud.Barony
    BSanta_giusta圣朱斯塔() 	feud.Barony
    BSorgono索尔戈诺() 	feud.Barony
    BTharros塔罗斯() 	feud.Barony
}

type 阿尔博雷亚ArboreaCounty struct {
	feud.BaseCounty
	卡布拉斯Cabras 	feud.Barony
	福尔东贾努斯Fordongianus 	feud.Barony
	奥里斯塔诺Oristano 	feud.Barony
	帕比洛尼斯Pabillonis 	feud.Barony
	圣卢里Sanluri 	feud.Barony
	圣朱斯塔Santa_giusta 	feud.Barony
	索尔戈诺Sorgono 	feud.Barony
	塔罗斯Tharros 	feud.Barony
}

func (c *阿尔博雷亚ArboreaCounty) BCabras卡布拉斯() feud.Barony {
	return c.卡布拉斯Cabras
}
    
func (c *阿尔博雷亚ArboreaCounty) BFordongianus福尔东贾努斯() feud.Barony {
	return c.福尔东贾努斯Fordongianus
}
    
func (c *阿尔博雷亚ArboreaCounty) BOristano奥里斯塔诺() feud.Barony {
	return c.奥里斯塔诺Oristano
}
    
func (c *阿尔博雷亚ArboreaCounty) BPabillonis帕比洛尼斯() feud.Barony {
	return c.帕比洛尼斯Pabillonis
}
    
func (c *阿尔博雷亚ArboreaCounty) BSanluri圣卢里() feud.Barony {
	return c.圣卢里Sanluri
}
    
func (c *阿尔博雷亚ArboreaCounty) BSanta_giusta圣朱斯塔() feud.Barony {
	return c.圣朱斯塔Santa_giusta
}
    
func (c *阿尔博雷亚ArboreaCounty) BSorgono索尔戈诺() feud.Barony {
	return c.索尔戈诺Sorgono
}
    
func (c *阿尔博雷亚ArboreaCounty) BTharros塔罗斯() feud.Barony {
	return c.塔罗斯Tharros
}
    
var CArborea阿尔博雷亚 ArboreaCounty = &阿尔博雷亚ArboreaCounty{}

func init() {
	f := CArborea阿尔博雷亚.(*阿尔博雷亚ArboreaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "325",
		Title:     "arborea",
		TitleName: "阿尔博雷亚",
		TitleCode: "c_arborea",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡布拉斯Cabras = BCabras卡布拉斯
	f.卡布拉斯Cabras.SetParent(f)
	
	f.福尔东贾努斯Fordongianus = BFordongianus福尔东贾努斯
	f.福尔东贾努斯Fordongianus.SetParent(f)
	
	f.奥里斯塔诺Oristano = BOristano奥里斯塔诺
	f.奥里斯塔诺Oristano.SetParent(f)
	
	f.帕比洛尼斯Pabillonis = BPabillonis帕比洛尼斯
	f.帕比洛尼斯Pabillonis.SetParent(f)
	
	f.圣卢里Sanluri = BSanluri圣卢里
	f.圣卢里Sanluri.SetParent(f)
	
	f.圣朱斯塔Santa_giusta = BSanta_giusta圣朱斯塔
	f.圣朱斯塔Santa_giusta.SetParent(f)
	
	f.索尔戈诺Sorgono = BSorgono索尔戈诺
	f.索尔戈诺Sorgono.SetParent(f)
	
	f.塔罗斯Tharros = BTharros塔罗斯
	f.塔罗斯Tharros.SetParent(f)
	
}
