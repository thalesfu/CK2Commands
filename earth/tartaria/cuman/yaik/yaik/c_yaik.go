package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YaikCounty interface {
    feud.County
    BKurgan库尔干() 	feud.Barony
    BKyzalyar克孜勒_亚尔() 	feud.Barony
    BLebyazhe列比亚日耶() 	feud.Barony
    BMakushino马库希诺() 	feud.Barony
    BMishkino米什基诺() 	feud.Barony
    BShumikha舒米哈() 	feud.Barony
    BYurgamysh尤尔加梅什() 	feud.Barony
}

type 亚伊克YaikCounty struct {
	feud.BaseCounty
	库尔干Kurgan 	feud.Barony
	克孜勒_亚尔Kyzalyar 	feud.Barony
	列比亚日耶Lebyazhe 	feud.Barony
	马库希诺Makushino 	feud.Barony
	米什基诺Mishkino 	feud.Barony
	舒米哈Shumikha 	feud.Barony
	尤尔加梅什Yurgamysh 	feud.Barony
}

func (c *亚伊克YaikCounty) BKurgan库尔干() feud.Barony {
	return c.库尔干Kurgan
}
    
func (c *亚伊克YaikCounty) BKyzalyar克孜勒_亚尔() feud.Barony {
	return c.克孜勒_亚尔Kyzalyar
}
    
func (c *亚伊克YaikCounty) BLebyazhe列比亚日耶() feud.Barony {
	return c.列比亚日耶Lebyazhe
}
    
func (c *亚伊克YaikCounty) BMakushino马库希诺() feud.Barony {
	return c.马库希诺Makushino
}
    
func (c *亚伊克YaikCounty) BMishkino米什基诺() feud.Barony {
	return c.米什基诺Mishkino
}
    
func (c *亚伊克YaikCounty) BShumikha舒米哈() feud.Barony {
	return c.舒米哈Shumikha
}
    
func (c *亚伊克YaikCounty) BYurgamysh尤尔加梅什() feud.Barony {
	return c.尤尔加梅什Yurgamysh
}
    
var CYaik亚伊克 YaikCounty = &亚伊克YaikCounty{}

func init() {
	f := CYaik亚伊克.(*亚伊克YaikCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "895",
		Title:     "yaik",
		TitleName: "亚伊克",
		TitleCode: "c_yaik",
		Baronies:  map[string]feud.Barony{},
	}

	f.库尔干Kurgan = BKurgan库尔干
	f.库尔干Kurgan.SetParent(f)
	
	f.克孜勒_亚尔Kyzalyar = BKyzalyar克孜勒_亚尔
	f.克孜勒_亚尔Kyzalyar.SetParent(f)
	
	f.列比亚日耶Lebyazhe = BLebyazhe列比亚日耶
	f.列比亚日耶Lebyazhe.SetParent(f)
	
	f.马库希诺Makushino = BMakushino马库希诺
	f.马库希诺Makushino.SetParent(f)
	
	f.米什基诺Mishkino = BMishkino米什基诺
	f.米什基诺Mishkino.SetParent(f)
	
	f.舒米哈Shumikha = BShumikha舒米哈
	f.舒米哈Shumikha.SetParent(f)
	
	f.尤尔加梅什Yurgamysh = BYurgamysh尤尔加梅什
	f.尤尔加梅什Yurgamysh.SetParent(f)
	
}
