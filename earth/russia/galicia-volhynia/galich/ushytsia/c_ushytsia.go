package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UshytsiaCounty interface {
    feud.County
    BDunaivtsi杜纳伊夫齐() 	feud.Barony
    BKamianets卡缅涅茨() 	feud.Barony
    BLitnivtsi利特尼夫齐() 	feud.Barony
    BMurovani穆罗瓦尼() 	feud.Barony
    BUshiyin乌希因() 	feud.Barony
    BUshytsia乌希察() 	feud.Barony
    BVelykyi_zhvanchyk大日万奇克() 	feud.Barony
}

type 乌希察UshytsiaCounty struct {
	feud.BaseCounty
	杜纳伊夫齐Dunaivtsi 	feud.Barony
	卡缅涅茨Kamianets 	feud.Barony
	利特尼夫齐Litnivtsi 	feud.Barony
	穆罗瓦尼Murovani 	feud.Barony
	乌希因Ushiyin 	feud.Barony
	乌希察Ushytsia 	feud.Barony
	大日万奇克Velykyi_zhvanchyk 	feud.Barony
}

func (c *乌希察UshytsiaCounty) BDunaivtsi杜纳伊夫齐() feud.Barony {
	return c.杜纳伊夫齐Dunaivtsi
}
    
func (c *乌希察UshytsiaCounty) BKamianets卡缅涅茨() feud.Barony {
	return c.卡缅涅茨Kamianets
}
    
func (c *乌希察UshytsiaCounty) BLitnivtsi利特尼夫齐() feud.Barony {
	return c.利特尼夫齐Litnivtsi
}
    
func (c *乌希察UshytsiaCounty) BMurovani穆罗瓦尼() feud.Barony {
	return c.穆罗瓦尼Murovani
}
    
func (c *乌希察UshytsiaCounty) BUshiyin乌希因() feud.Barony {
	return c.乌希因Ushiyin
}
    
func (c *乌希察UshytsiaCounty) BUshytsia乌希察() feud.Barony {
	return c.乌希察Ushytsia
}
    
func (c *乌希察UshytsiaCounty) BVelykyi_zhvanchyk大日万奇克() feud.Barony {
	return c.大日万奇克Velykyi_zhvanchyk
}
    
var CUshytsia乌希察 UshytsiaCounty = &乌希察UshytsiaCounty{}

func init() {
	f := CUshytsia乌希察.(*乌希察UshytsiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1637",
		Title:     "ushytsia",
		TitleName: "乌希察",
		TitleCode: "c_ushytsia",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜纳伊夫齐Dunaivtsi = BDunaivtsi杜纳伊夫齐
	f.杜纳伊夫齐Dunaivtsi.SetParent(f)
	
	f.卡缅涅茨Kamianets = BKamianets卡缅涅茨
	f.卡缅涅茨Kamianets.SetParent(f)
	
	f.利特尼夫齐Litnivtsi = BLitnivtsi利特尼夫齐
	f.利特尼夫齐Litnivtsi.SetParent(f)
	
	f.穆罗瓦尼Murovani = BMurovani穆罗瓦尼
	f.穆罗瓦尼Murovani.SetParent(f)
	
	f.乌希因Ushiyin = BUshiyin乌希因
	f.乌希因Ushiyin.SetParent(f)
	
	f.乌希察Ushytsia = BUshytsia乌希察
	f.乌希察Ushytsia.SetParent(f)
	
	f.大日万奇克Velykyi_zhvanchyk = BVelykyi_zhvanchyk大日万奇克
	f.大日万奇克Velykyi_zhvanchyk.SetParent(f)
	
}
