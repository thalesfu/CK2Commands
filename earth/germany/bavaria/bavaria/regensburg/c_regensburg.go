package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RegensburgCounty interface {
    feud.County
    BIngolstadt因戈尔施塔特() 	feud.Barony
    BKelheim凯尔海姆() 	feud.Barony
    BLandshut兰茨胡特() 	feud.Barony
    BLerchenfeld莱兴费尔德() 	feud.Barony
    BRegensburg雷根斯堡() 	feud.Barony
    BRohrbach罗尔巴赫() 	feud.Barony
    BStraubing施特劳宾() 	feud.Barony
}

type 雷根斯堡RegensburgCounty struct {
	feud.BaseCounty
	因戈尔施塔特Ingolstadt 	feud.Barony
	凯尔海姆Kelheim 	feud.Barony
	兰茨胡特Landshut 	feud.Barony
	莱兴费尔德Lerchenfeld 	feud.Barony
	雷根斯堡Regensburg 	feud.Barony
	罗尔巴赫Rohrbach 	feud.Barony
	施特劳宾Straubing 	feud.Barony
}

func (c *雷根斯堡RegensburgCounty) BIngolstadt因戈尔施塔特() feud.Barony {
	return c.因戈尔施塔特Ingolstadt
}
    
func (c *雷根斯堡RegensburgCounty) BKelheim凯尔海姆() feud.Barony {
	return c.凯尔海姆Kelheim
}
    
func (c *雷根斯堡RegensburgCounty) BLandshut兰茨胡特() feud.Barony {
	return c.兰茨胡特Landshut
}
    
func (c *雷根斯堡RegensburgCounty) BLerchenfeld莱兴费尔德() feud.Barony {
	return c.莱兴费尔德Lerchenfeld
}
    
func (c *雷根斯堡RegensburgCounty) BRegensburg雷根斯堡() feud.Barony {
	return c.雷根斯堡Regensburg
}
    
func (c *雷根斯堡RegensburgCounty) BRohrbach罗尔巴赫() feud.Barony {
	return c.罗尔巴赫Rohrbach
}
    
func (c *雷根斯堡RegensburgCounty) BStraubing施特劳宾() feud.Barony {
	return c.施特劳宾Straubing
}
    
var CRegensburg雷根斯堡 RegensburgCounty = &雷根斯堡RegensburgCounty{}

func init() {
	f := CRegensburg雷根斯堡.(*雷根斯堡RegensburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1700",
		Title:     "regensburg",
		TitleName: "雷根斯堡",
		TitleCode: "c_regensburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.因戈尔施塔特Ingolstadt = BIngolstadt因戈尔施塔特
	f.因戈尔施塔特Ingolstadt.SetParent(f)
	
	f.凯尔海姆Kelheim = BKelheim凯尔海姆
	f.凯尔海姆Kelheim.SetParent(f)
	
	f.兰茨胡特Landshut = BLandshut兰茨胡特
	f.兰茨胡特Landshut.SetParent(f)
	
	f.莱兴费尔德Lerchenfeld = BLerchenfeld莱兴费尔德
	f.莱兴费尔德Lerchenfeld.SetParent(f)
	
	f.雷根斯堡Regensburg = BRegensburg雷根斯堡
	f.雷根斯堡Regensburg.SetParent(f)
	
	f.罗尔巴赫Rohrbach = BRohrbach罗尔巴赫
	f.罗尔巴赫Rohrbach.SetParent(f)
	
	f.施特劳宾Straubing = BStraubing施特劳宾
	f.施特劳宾Straubing.SetParent(f)
	
}
