package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PerugiaCounty interface {
    feud.County
    BAssisi阿西西() 	feud.Barony
    BDeruta德鲁塔() 	feud.Barony
    BFratta弗拉塔() 	feud.Barony
    BGubbio古比奥() 	feud.Barony
    BMarsciano马尔夏诺() 	feud.Barony
    BMontone蒙托内() 	feud.Barony
    BPerugia佩鲁贾() 	feud.Barony
}

type 佩鲁贾PerugiaCounty struct {
	feud.BaseCounty
	阿西西Assisi 	feud.Barony
	德鲁塔Deruta 	feud.Barony
	弗拉塔Fratta 	feud.Barony
	古比奥Gubbio 	feud.Barony
	马尔夏诺Marsciano 	feud.Barony
	蒙托内Montone 	feud.Barony
	佩鲁贾Perugia 	feud.Barony
}

func (c *佩鲁贾PerugiaCounty) BAssisi阿西西() feud.Barony {
	return c.阿西西Assisi
}
    
func (c *佩鲁贾PerugiaCounty) BDeruta德鲁塔() feud.Barony {
	return c.德鲁塔Deruta
}
    
func (c *佩鲁贾PerugiaCounty) BFratta弗拉塔() feud.Barony {
	return c.弗拉塔Fratta
}
    
func (c *佩鲁贾PerugiaCounty) BGubbio古比奥() feud.Barony {
	return c.古比奥Gubbio
}
    
func (c *佩鲁贾PerugiaCounty) BMarsciano马尔夏诺() feud.Barony {
	return c.马尔夏诺Marsciano
}
    
func (c *佩鲁贾PerugiaCounty) BMontone蒙托内() feud.Barony {
	return c.蒙托内Montone
}
    
func (c *佩鲁贾PerugiaCounty) BPerugia佩鲁贾() feud.Barony {
	return c.佩鲁贾Perugia
}
    
var CPerugia佩鲁贾 PerugiaCounty = &佩鲁贾PerugiaCounty{}

func init() {
	f := CPerugia佩鲁贾.(*佩鲁贾PerugiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1708",
		Title:     "perugia",
		TitleName: "佩鲁贾",
		TitleCode: "c_perugia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿西西Assisi = BAssisi阿西西
	f.阿西西Assisi.SetParent(f)
	
	f.德鲁塔Deruta = BDeruta德鲁塔
	f.德鲁塔Deruta.SetParent(f)
	
	f.弗拉塔Fratta = BFratta弗拉塔
	f.弗拉塔Fratta.SetParent(f)
	
	f.古比奥Gubbio = BGubbio古比奥
	f.古比奥Gubbio.SetParent(f)
	
	f.马尔夏诺Marsciano = BMarsciano马尔夏诺
	f.马尔夏诺Marsciano.SetParent(f)
	
	f.蒙托内Montone = BMontone蒙托内
	f.蒙托内Montone.SetParent(f)
	
	f.佩鲁贾Perugia = BPerugia佩鲁贾
	f.佩鲁贾Perugia.SetParent(f)
	
}
