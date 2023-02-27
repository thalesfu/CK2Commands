package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MemaCounty interface {
    feud.County
    BBakara巴卡拉() 	feud.Barony
    BDia迪亚() 	feud.Barony
    BKabora卡博拉() 	feud.Barony
    BMacina马西纳() 	feud.Barony
    BMema梅马() 	feud.Barony
    BPerge佩尔热() 	feud.Barony
    BSadiabougouda萨迪亚布古达() 	feud.Barony
}

type 梅马MemaCounty struct {
	feud.BaseCounty
	巴卡拉Bakara 	feud.Barony
	迪亚Dia 	feud.Barony
	卡博拉Kabora 	feud.Barony
	马西纳Macina 	feud.Barony
	梅马Mema 	feud.Barony
	佩尔热Perge 	feud.Barony
	萨迪亚布古达Sadiabougouda 	feud.Barony
}

func (c *梅马MemaCounty) BBakara巴卡拉() feud.Barony {
	return c.巴卡拉Bakara
}
    
func (c *梅马MemaCounty) BDia迪亚() feud.Barony {
	return c.迪亚Dia
}
    
func (c *梅马MemaCounty) BKabora卡博拉() feud.Barony {
	return c.卡博拉Kabora
}
    
func (c *梅马MemaCounty) BMacina马西纳() feud.Barony {
	return c.马西纳Macina
}
    
func (c *梅马MemaCounty) BMema梅马() feud.Barony {
	return c.梅马Mema
}
    
func (c *梅马MemaCounty) BPerge佩尔热() feud.Barony {
	return c.佩尔热Perge
}
    
func (c *梅马MemaCounty) BSadiabougouda萨迪亚布古达() feud.Barony {
	return c.萨迪亚布古达Sadiabougouda
}
    
var CMema梅马 MemaCounty = &梅马MemaCounty{}

func init() {
	f := CMema梅马.(*梅马MemaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1768",
		Title:     "mema",
		TitleName: "梅马",
		TitleCode: "c_mema",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴卡拉Bakara = BBakara巴卡拉
	f.巴卡拉Bakara.SetParent(f)
	
	f.迪亚Dia = BDia迪亚
	f.迪亚Dia.SetParent(f)
	
	f.卡博拉Kabora = BKabora卡博拉
	f.卡博拉Kabora.SetParent(f)
	
	f.马西纳Macina = BMacina马西纳
	f.马西纳Macina.SetParent(f)
	
	f.梅马Mema = BMema梅马
	f.梅马Mema.SetParent(f)
	
	f.佩尔热Perge = BPerge佩尔热
	f.佩尔热Perge.SetParent(f)
	
	f.萨迪亚布古达Sadiabougouda = BSadiabougouda萨迪亚布古达
	f.萨迪亚布古达Sadiabougouda.SetParent(f)
	
}
