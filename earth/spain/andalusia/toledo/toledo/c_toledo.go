package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToledoCounty interface {
    feud.County
    BConsuegra孔苏埃格拉() 	feud.Barony
    BFuensalida丰萨利达() 	feud.Barony
    BIllescas伊列斯卡斯() 	feud.Barony
    BMadrid马德里() 	feud.Barony
    BOrgaz奥尔加斯() 	feud.Barony
    BTalavera塔拉韦拉() 	feud.Barony
    BToledo托雷多() 	feud.Barony
    BTolemora莫拉() 	feud.Barony
}

type 托雷多ToledoCounty struct {
	feud.BaseCounty
	孔苏埃格拉Consuegra 	feud.Barony
	丰萨利达Fuensalida 	feud.Barony
	伊列斯卡斯Illescas 	feud.Barony
	马德里Madrid 	feud.Barony
	奥尔加斯Orgaz 	feud.Barony
	塔拉韦拉Talavera 	feud.Barony
	托雷多Toledo 	feud.Barony
	莫拉Tolemora 	feud.Barony
}

func (c *托雷多ToledoCounty) BConsuegra孔苏埃格拉() feud.Barony {
	return c.孔苏埃格拉Consuegra
}
    
func (c *托雷多ToledoCounty) BFuensalida丰萨利达() feud.Barony {
	return c.丰萨利达Fuensalida
}
    
func (c *托雷多ToledoCounty) BIllescas伊列斯卡斯() feud.Barony {
	return c.伊列斯卡斯Illescas
}
    
func (c *托雷多ToledoCounty) BMadrid马德里() feud.Barony {
	return c.马德里Madrid
}
    
func (c *托雷多ToledoCounty) BOrgaz奥尔加斯() feud.Barony {
	return c.奥尔加斯Orgaz
}
    
func (c *托雷多ToledoCounty) BTalavera塔拉韦拉() feud.Barony {
	return c.塔拉韦拉Talavera
}
    
func (c *托雷多ToledoCounty) BToledo托雷多() feud.Barony {
	return c.托雷多Toledo
}
    
func (c *托雷多ToledoCounty) BTolemora莫拉() feud.Barony {
	return c.莫拉Tolemora
}
    
var CToledo托雷多 ToledoCounty = &托雷多ToledoCounty{}

func init() {
	f := CToledo托雷多.(*托雷多ToledoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "197",
		Title:     "toledo",
		TitleName: "托雷多",
		TitleCode: "c_toledo",
		Baronies:  map[string]feud.Barony{},
	}

	f.孔苏埃格拉Consuegra = BConsuegra孔苏埃格拉
	f.孔苏埃格拉Consuegra.SetParent(f)
	
	f.丰萨利达Fuensalida = BFuensalida丰萨利达
	f.丰萨利达Fuensalida.SetParent(f)
	
	f.伊列斯卡斯Illescas = BIllescas伊列斯卡斯
	f.伊列斯卡斯Illescas.SetParent(f)
	
	f.马德里Madrid = BMadrid马德里
	f.马德里Madrid.SetParent(f)
	
	f.奥尔加斯Orgaz = BOrgaz奥尔加斯
	f.奥尔加斯Orgaz.SetParent(f)
	
	f.塔拉韦拉Talavera = BTalavera塔拉韦拉
	f.塔拉韦拉Talavera.SetParent(f)
	
	f.托雷多Toledo = BToledo托雷多
	f.托雷多Toledo.SetParent(f)
	
	f.莫拉Tolemora = BTolemora莫拉
	f.莫拉Tolemora.SetParent(f)
	
}
