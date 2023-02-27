package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MonyulCounty interface {
    feud.County
    BBameng贝门() 	feud.Barony
    BBomdila邦迪拉() 	feud.Barony
    BItanagar伊塔那噶() 	feud.Barony
    BJoram焦鲁姆() 	feud.Barony
    BKoloriang克罗里安() 	feud.Barony
    BSepla塞巴() 	feud.Barony
    BTawang达旺() 	feud.Barony
}

type 门隅MonyulCounty struct {
	feud.BaseCounty
	贝门Bameng 	feud.Barony
	邦迪拉Bomdila 	feud.Barony
	伊塔那噶Itanagar 	feud.Barony
	焦鲁姆Joram 	feud.Barony
	克罗里安Koloriang 	feud.Barony
	塞巴Sepla 	feud.Barony
	达旺Tawang 	feud.Barony
}

func (c *门隅MonyulCounty) BBameng贝门() feud.Barony {
	return c.贝门Bameng
}
    
func (c *门隅MonyulCounty) BBomdila邦迪拉() feud.Barony {
	return c.邦迪拉Bomdila
}
    
func (c *门隅MonyulCounty) BItanagar伊塔那噶() feud.Barony {
	return c.伊塔那噶Itanagar
}
    
func (c *门隅MonyulCounty) BJoram焦鲁姆() feud.Barony {
	return c.焦鲁姆Joram
}
    
func (c *门隅MonyulCounty) BKoloriang克罗里安() feud.Barony {
	return c.克罗里安Koloriang
}
    
func (c *门隅MonyulCounty) BSepla塞巴() feud.Barony {
	return c.塞巴Sepla
}
    
func (c *门隅MonyulCounty) BTawang达旺() feud.Barony {
	return c.达旺Tawang
}
    
var CMonyul门隅 MonyulCounty = &门隅MonyulCounty{}

func init() {
	f := CMonyul门隅.(*门隅MonyulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1483",
		Title:     "monyul",
		TitleName: "门隅",
		TitleCode: "c_monyul",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝门Bameng = BBameng贝门
	f.贝门Bameng.SetParent(f)
	
	f.邦迪拉Bomdila = BBomdila邦迪拉
	f.邦迪拉Bomdila.SetParent(f)
	
	f.伊塔那噶Itanagar = BItanagar伊塔那噶
	f.伊塔那噶Itanagar.SetParent(f)
	
	f.焦鲁姆Joram = BJoram焦鲁姆
	f.焦鲁姆Joram.SetParent(f)
	
	f.克罗里安Koloriang = BKoloriang克罗里安
	f.克罗里安Koloriang.SetParent(f)
	
	f.塞巴Sepla = BSepla塞巴
	f.塞巴Sepla.SetParent(f)
	
	f.达旺Tawang = BTawang达旺
	f.达旺Tawang.SetParent(f)
	
}
