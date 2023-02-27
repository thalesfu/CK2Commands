package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarrickCounty interface {
    feud.County
    BBallantrae巴伦特雷() 	feud.Barony
    BCrossraguel科罗斯拉结尔() 	feud.Barony
    BCulzean克伦() 	feud.Barony
    BDunure达纽尔() 	feud.Barony
    BGreenan格林南() 	feud.Barony
    BLoch_doon杜恩湖() 	feud.Barony
    BMaybole梅博尔() 	feud.Barony
    BTurnberry特恩贝里() 	feud.Barony
}

type 卡里克CarrickCounty struct {
	feud.BaseCounty
	巴伦特雷Ballantrae 	feud.Barony
	科罗斯拉结尔Crossraguel 	feud.Barony
	克伦Culzean 	feud.Barony
	达纽尔Dunure 	feud.Barony
	格林南Greenan 	feud.Barony
	杜恩湖Loch_doon 	feud.Barony
	梅博尔Maybole 	feud.Barony
	特恩贝里Turnberry 	feud.Barony
}

func (c *卡里克CarrickCounty) BBallantrae巴伦特雷() feud.Barony {
	return c.巴伦特雷Ballantrae
}
    
func (c *卡里克CarrickCounty) BCrossraguel科罗斯拉结尔() feud.Barony {
	return c.科罗斯拉结尔Crossraguel
}
    
func (c *卡里克CarrickCounty) BCulzean克伦() feud.Barony {
	return c.克伦Culzean
}
    
func (c *卡里克CarrickCounty) BDunure达纽尔() feud.Barony {
	return c.达纽尔Dunure
}
    
func (c *卡里克CarrickCounty) BGreenan格林南() feud.Barony {
	return c.格林南Greenan
}
    
func (c *卡里克CarrickCounty) BLoch_doon杜恩湖() feud.Barony {
	return c.杜恩湖Loch_doon
}
    
func (c *卡里克CarrickCounty) BMaybole梅博尔() feud.Barony {
	return c.梅博尔Maybole
}
    
func (c *卡里克CarrickCounty) BTurnberry特恩贝里() feud.Barony {
	return c.特恩贝里Turnberry
}
    
var CCarrick卡里克 CarrickCounty = &卡里克CarrickCounty{}

func init() {
	f := CCarrick卡里克.(*卡里克CarrickCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "49",
		Title:     "carrick",
		TitleName: "卡里克",
		TitleCode: "c_carrick",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴伦特雷Ballantrae = BBallantrae巴伦特雷
	f.巴伦特雷Ballantrae.SetParent(f)
	
	f.科罗斯拉结尔Crossraguel = BCrossraguel科罗斯拉结尔
	f.科罗斯拉结尔Crossraguel.SetParent(f)
	
	f.克伦Culzean = BCulzean克伦
	f.克伦Culzean.SetParent(f)
	
	f.达纽尔Dunure = BDunure达纽尔
	f.达纽尔Dunure.SetParent(f)
	
	f.格林南Greenan = BGreenan格林南
	f.格林南Greenan.SetParent(f)
	
	f.杜恩湖Loch_doon = BLoch_doon杜恩湖
	f.杜恩湖Loch_doon.SetParent(f)
	
	f.梅博尔Maybole = BMaybole梅博尔
	f.梅博尔Maybole.SetParent(f)
	
	f.特恩贝里Turnberry = BTurnberry特恩贝里
	f.特恩贝里Turnberry.SetParent(f)
	
}
