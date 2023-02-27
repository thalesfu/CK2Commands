package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Grassland_cheremisaCounty interface {
    feud.County
    BAznaqay阿兹纳卡伊() 	feud.Barony
    BBawli巴夫雷() 	feud.Barony
    BBua布阿() 	feud.Barony
    BCherm切尔姆() 	feud.Barony
    BCistay奇斯塔伊() 	feud.Barony
    BYarcalli亚尔夏勒() 	feud.Barony
    BZay扎伊() 	feud.Barony
}

type 克尔热涅茨Grassland_cheremisaCounty struct {
	feud.BaseCounty
	阿兹纳卡伊Aznaqay 	feud.Barony
	巴夫雷Bawli 	feud.Barony
	布阿Bua 	feud.Barony
	切尔姆Cherm 	feud.Barony
	奇斯塔伊Cistay 	feud.Barony
	亚尔夏勒Yarcalli 	feud.Barony
	扎伊Zay 	feud.Barony
}

func (c *克尔热涅茨Grassland_cheremisaCounty) BAznaqay阿兹纳卡伊() feud.Barony {
	return c.阿兹纳卡伊Aznaqay
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BBawli巴夫雷() feud.Barony {
	return c.巴夫雷Bawli
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BBua布阿() feud.Barony {
	return c.布阿Bua
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BCherm切尔姆() feud.Barony {
	return c.切尔姆Cherm
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BCistay奇斯塔伊() feud.Barony {
	return c.奇斯塔伊Cistay
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BYarcalli亚尔夏勒() feud.Barony {
	return c.亚尔夏勒Yarcalli
}
    
func (c *克尔热涅茨Grassland_cheremisaCounty) BZay扎伊() feud.Barony {
	return c.扎伊Zay
}
    
var CGrassland_cheremisa克尔热涅茨 Grassland_cheremisaCounty = &克尔热涅茨Grassland_cheremisaCounty{}

func init() {
	f := CGrassland_cheremisa克尔热涅茨.(*克尔热涅茨Grassland_cheremisaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "589",
		Title:     "grassland_cheremisa",
		TitleName: "克尔热涅茨",
		TitleCode: "c_grassland_cheremisa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿兹纳卡伊Aznaqay = BAznaqay阿兹纳卡伊
	f.阿兹纳卡伊Aznaqay.SetParent(f)
	
	f.巴夫雷Bawli = BBawli巴夫雷
	f.巴夫雷Bawli.SetParent(f)
	
	f.布阿Bua = BBua布阿
	f.布阿Bua.SetParent(f)
	
	f.切尔姆Cherm = BCherm切尔姆
	f.切尔姆Cherm.SetParent(f)
	
	f.奇斯塔伊Cistay = BCistay奇斯塔伊
	f.奇斯塔伊Cistay.SetParent(f)
	
	f.亚尔夏勒Yarcalli = BYarcalli亚尔夏勒
	f.亚尔夏勒Yarcalli.SetParent(f)
	
	f.扎伊Zay = BZay扎伊
	f.扎伊Zay.SetParent(f)
	
}
