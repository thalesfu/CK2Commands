package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CoimbraCounty interface {
    feud.County
    BAveiro阿威罗() 	feud.Barony
    BCantanhede坎塔涅迪() 	feud.Barony
    BCoimbra科英布拉() 	feud.Barony
    BCondeixa孔代沙() 	feud.Barony
    BMontereal蒙蒂雷亚尔() 	feud.Barony
    BPedondo雷东杜() 	feud.Barony
    BPenela佩内拉() 	feud.Barony
    BViseu维塞乌() 	feud.Barony
}

type 科英布拉CoimbraCounty struct {
	feud.BaseCounty
	阿威罗Aveiro 	feud.Barony
	坎塔涅迪Cantanhede 	feud.Barony
	科英布拉Coimbra 	feud.Barony
	孔代沙Condeixa 	feud.Barony
	蒙蒂雷亚尔Montereal 	feud.Barony
	雷东杜Pedondo 	feud.Barony
	佩内拉Penela 	feud.Barony
	维塞乌Viseu 	feud.Barony
}

func (c *科英布拉CoimbraCounty) BAveiro阿威罗() feud.Barony {
	return c.阿威罗Aveiro
}
    
func (c *科英布拉CoimbraCounty) BCantanhede坎塔涅迪() feud.Barony {
	return c.坎塔涅迪Cantanhede
}
    
func (c *科英布拉CoimbraCounty) BCoimbra科英布拉() feud.Barony {
	return c.科英布拉Coimbra
}
    
func (c *科英布拉CoimbraCounty) BCondeixa孔代沙() feud.Barony {
	return c.孔代沙Condeixa
}
    
func (c *科英布拉CoimbraCounty) BMontereal蒙蒂雷亚尔() feud.Barony {
	return c.蒙蒂雷亚尔Montereal
}
    
func (c *科英布拉CoimbraCounty) BPedondo雷东杜() feud.Barony {
	return c.雷东杜Pedondo
}
    
func (c *科英布拉CoimbraCounty) BPenela佩内拉() feud.Barony {
	return c.佩内拉Penela
}
    
func (c *科英布拉CoimbraCounty) BViseu维塞乌() feud.Barony {
	return c.维塞乌Viseu
}
    
var CCoimbra科英布拉 CoimbraCounty = &科英布拉CoimbraCounty{}

func init() {
	f := CCoimbra科英布拉.(*科英布拉CoimbraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "159",
		Title:     "coimbra",
		TitleName: "科英布拉",
		TitleCode: "c_coimbra",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿威罗Aveiro = BAveiro阿威罗
	f.阿威罗Aveiro.SetParent(f)
	
	f.坎塔涅迪Cantanhede = BCantanhede坎塔涅迪
	f.坎塔涅迪Cantanhede.SetParent(f)
	
	f.科英布拉Coimbra = BCoimbra科英布拉
	f.科英布拉Coimbra.SetParent(f)
	
	f.孔代沙Condeixa = BCondeixa孔代沙
	f.孔代沙Condeixa.SetParent(f)
	
	f.蒙蒂雷亚尔Montereal = BMontereal蒙蒂雷亚尔
	f.蒙蒂雷亚尔Montereal.SetParent(f)
	
	f.雷东杜Pedondo = BPedondo雷东杜
	f.雷东杜Pedondo.SetParent(f)
	
	f.佩内拉Penela = BPenela佩内拉
	f.佩内拉Penela.SetParent(f)
	
	f.维塞乌Viseu = BViseu维塞乌
	f.维塞乌Viseu.SetParent(f)
	
}
