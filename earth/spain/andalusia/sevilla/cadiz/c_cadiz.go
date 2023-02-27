package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CadizCounty interface {
    feud.County
    BAlcaladelosgazules加苏莱斯堡() 	feud.Barony
    BArcos阿尔科斯() 	feud.Barony
    BCadiz加的斯() 	feud.Barony
    BJerez赫雷斯() 	feud.Barony
    BMedinasidonia锡多尼亚城() 	feud.Barony
    BSanfernando圣费尔南多() 	feud.Barony
    BSanjosedelvalle圣何塞斯德尔巴列() 	feud.Barony
    BSanlucadebarrameda桑卢卡尔德巴拉梅达() 	feud.Barony
}

type 加的斯CadizCounty struct {
	feud.BaseCounty
	加苏莱斯堡Alcaladelosgazules 	feud.Barony
	阿尔科斯Arcos 	feud.Barony
	加的斯Cadiz 	feud.Barony
	赫雷斯Jerez 	feud.Barony
	锡多尼亚城Medinasidonia 	feud.Barony
	圣费尔南多Sanfernando 	feud.Barony
	圣何塞斯德尔巴列Sanjosedelvalle 	feud.Barony
	桑卢卡尔德巴拉梅达Sanlucadebarrameda 	feud.Barony
}

func (c *加的斯CadizCounty) BAlcaladelosgazules加苏莱斯堡() feud.Barony {
	return c.加苏莱斯堡Alcaladelosgazules
}
    
func (c *加的斯CadizCounty) BArcos阿尔科斯() feud.Barony {
	return c.阿尔科斯Arcos
}
    
func (c *加的斯CadizCounty) BCadiz加的斯() feud.Barony {
	return c.加的斯Cadiz
}
    
func (c *加的斯CadizCounty) BJerez赫雷斯() feud.Barony {
	return c.赫雷斯Jerez
}
    
func (c *加的斯CadizCounty) BMedinasidonia锡多尼亚城() feud.Barony {
	return c.锡多尼亚城Medinasidonia
}
    
func (c *加的斯CadizCounty) BSanfernando圣费尔南多() feud.Barony {
	return c.圣费尔南多Sanfernando
}
    
func (c *加的斯CadizCounty) BSanjosedelvalle圣何塞斯德尔巴列() feud.Barony {
	return c.圣何塞斯德尔巴列Sanjosedelvalle
}
    
func (c *加的斯CadizCounty) BSanlucadebarrameda桑卢卡尔德巴拉梅达() feud.Barony {
	return c.桑卢卡尔德巴拉梅达Sanlucadebarrameda
}
    
var CCadiz加的斯 CadizCounty = &加的斯CadizCounty{}

func init() {
	f := CCadiz加的斯.(*加的斯CadizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "165",
		Title:     "cadiz",
		TitleName: "加的斯",
		TitleCode: "c_cadiz",
		Baronies:  map[string]feud.Barony{},
	}

	f.加苏莱斯堡Alcaladelosgazules = BAlcaladelosgazules加苏莱斯堡
	f.加苏莱斯堡Alcaladelosgazules.SetParent(f)
	
	f.阿尔科斯Arcos = BArcos阿尔科斯
	f.阿尔科斯Arcos.SetParent(f)
	
	f.加的斯Cadiz = BCadiz加的斯
	f.加的斯Cadiz.SetParent(f)
	
	f.赫雷斯Jerez = BJerez赫雷斯
	f.赫雷斯Jerez.SetParent(f)
	
	f.锡多尼亚城Medinasidonia = BMedinasidonia锡多尼亚城
	f.锡多尼亚城Medinasidonia.SetParent(f)
	
	f.圣费尔南多Sanfernando = BSanfernando圣费尔南多
	f.圣费尔南多Sanfernando.SetParent(f)
	
	f.圣何塞斯德尔巴列Sanjosedelvalle = BSanjosedelvalle圣何塞斯德尔巴列
	f.圣何塞斯德尔巴列Sanjosedelvalle.SetParent(f)
	
	f.桑卢卡尔德巴拉梅达Sanlucadebarrameda = BSanlucadebarrameda桑卢卡尔德巴拉梅达
	f.桑卢卡尔德巴拉梅达Sanlucadebarrameda.SetParent(f)
	
}
