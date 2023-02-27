package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HanyanCounty interface {
    feud.County
    BAlancha埃安查() 	feud.Barony
    BAlkasdir卡斯迪尔() 	feud.Barony
    BBougtob布格图卜() 	feud.Barony
    BHanyan汗燕() 	feud.Barony
    BMagoura马古拉() 	feud.Barony
    BNaama纳马() 	feud.Barony
    BRaselma拉斯马() 	feud.Barony
    BSaida塞伊达() 	feud.Barony
}

type 亚拉拉HanyanCounty struct {
	feud.BaseCounty
	埃安查Alancha 	feud.Barony
	卡斯迪尔Alkasdir 	feud.Barony
	布格图卜Bougtob 	feud.Barony
	汗燕Hanyan 	feud.Barony
	马古拉Magoura 	feud.Barony
	纳马Naama 	feud.Barony
	拉斯马Raselma 	feud.Barony
	塞伊达Saida 	feud.Barony
}

func (c *亚拉拉HanyanCounty) BAlancha埃安查() feud.Barony {
	return c.埃安查Alancha
}
    
func (c *亚拉拉HanyanCounty) BAlkasdir卡斯迪尔() feud.Barony {
	return c.卡斯迪尔Alkasdir
}
    
func (c *亚拉拉HanyanCounty) BBougtob布格图卜() feud.Barony {
	return c.布格图卜Bougtob
}
    
func (c *亚拉拉HanyanCounty) BHanyan汗燕() feud.Barony {
	return c.汗燕Hanyan
}
    
func (c *亚拉拉HanyanCounty) BMagoura马古拉() feud.Barony {
	return c.马古拉Magoura
}
    
func (c *亚拉拉HanyanCounty) BNaama纳马() feud.Barony {
	return c.纳马Naama
}
    
func (c *亚拉拉HanyanCounty) BRaselma拉斯马() feud.Barony {
	return c.拉斯马Raselma
}
    
func (c *亚拉拉HanyanCounty) BSaida塞伊达() feud.Barony {
	return c.塞伊达Saida
}
    
var CHanyan亚拉拉 HanyanCounty = &亚拉拉HanyanCounty{}

func init() {
	f := CHanyan亚拉拉.(*亚拉拉HanyanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "835",
		Title:     "hanyan",
		TitleName: "亚拉拉",
		TitleCode: "c_hanyan",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃安查Alancha = BAlancha埃安查
	f.埃安查Alancha.SetParent(f)
	
	f.卡斯迪尔Alkasdir = BAlkasdir卡斯迪尔
	f.卡斯迪尔Alkasdir.SetParent(f)
	
	f.布格图卜Bougtob = BBougtob布格图卜
	f.布格图卜Bougtob.SetParent(f)
	
	f.汗燕Hanyan = BHanyan汗燕
	f.汗燕Hanyan.SetParent(f)
	
	f.马古拉Magoura = BMagoura马古拉
	f.马古拉Magoura.SetParent(f)
	
	f.纳马Naama = BNaama纳马
	f.纳马Naama.SetParent(f)
	
	f.拉斯马Raselma = BRaselma拉斯马
	f.拉斯马Raselma.SetParent(f)
	
	f.塞伊达Saida = BSaida塞伊达
	f.塞伊达Saida.SetParent(f)
	
}
