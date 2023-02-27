package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Beni_yanniCounty interface {
    feud.County
    BAzazga阿扎兹加() 	feud.Barony
    BAzeffoun阿宰丰() 	feud.Barony
    BBourmedes布米尔达斯() 	feud.Barony
    BBujima布济马() 	feud.Barony
    BTafoughalt塔夫亚拉特() 	feud.Barony
    BTaoudouch套多斯() 	feud.Barony
    BThiza希泽() 	feud.Barony
    BTigzirt提格济尔特() 	feud.Barony
}

type 贝尼耶尼Beni_yanniCounty struct {
	feud.BaseCounty
	阿扎兹加Azazga 	feud.Barony
	阿宰丰Azeffoun 	feud.Barony
	布米尔达斯Bourmedes 	feud.Barony
	布济马Bujima 	feud.Barony
	塔夫亚拉特Tafoughalt 	feud.Barony
	套多斯Taoudouch 	feud.Barony
	希泽Thiza 	feud.Barony
	提格济尔特Tigzirt 	feud.Barony
}

func (c *贝尼耶尼Beni_yanniCounty) BAzazga阿扎兹加() feud.Barony {
	return c.阿扎兹加Azazga
}
    
func (c *贝尼耶尼Beni_yanniCounty) BAzeffoun阿宰丰() feud.Barony {
	return c.阿宰丰Azeffoun
}
    
func (c *贝尼耶尼Beni_yanniCounty) BBourmedes布米尔达斯() feud.Barony {
	return c.布米尔达斯Bourmedes
}
    
func (c *贝尼耶尼Beni_yanniCounty) BBujima布济马() feud.Barony {
	return c.布济马Bujima
}
    
func (c *贝尼耶尼Beni_yanniCounty) BTafoughalt塔夫亚拉特() feud.Barony {
	return c.塔夫亚拉特Tafoughalt
}
    
func (c *贝尼耶尼Beni_yanniCounty) BTaoudouch套多斯() feud.Barony {
	return c.套多斯Taoudouch
}
    
func (c *贝尼耶尼Beni_yanniCounty) BThiza希泽() feud.Barony {
	return c.希泽Thiza
}
    
func (c *贝尼耶尼Beni_yanniCounty) BTigzirt提格济尔特() feud.Barony {
	return c.提格济尔特Tigzirt
}
    
var CBeni_yanni贝尼耶尼 Beni_yanniCounty = &贝尼耶尼Beni_yanniCounty{}

func init() {
	f := CBeni_yanni贝尼耶尼.(*贝尼耶尼Beni_yanniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "825",
		Title:     "beni_yanni",
		TitleName: "贝尼耶尼",
		TitleCode: "c_beni_yanni",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿扎兹加Azazga = BAzazga阿扎兹加
	f.阿扎兹加Azazga.SetParent(f)
	
	f.阿宰丰Azeffoun = BAzeffoun阿宰丰
	f.阿宰丰Azeffoun.SetParent(f)
	
	f.布米尔达斯Bourmedes = BBourmedes布米尔达斯
	f.布米尔达斯Bourmedes.SetParent(f)
	
	f.布济马Bujima = BBujima布济马
	f.布济马Bujima.SetParent(f)
	
	f.塔夫亚拉特Tafoughalt = BTafoughalt塔夫亚拉特
	f.塔夫亚拉特Tafoughalt.SetParent(f)
	
	f.套多斯Taoudouch = BTaoudouch套多斯
	f.套多斯Taoudouch.SetParent(f)
	
	f.希泽Thiza = BThiza希泽
	f.希泽Thiza.SetParent(f)
	
	f.提格济尔特Tigzirt = BTigzirt提格济尔特
	f.提格济尔特Tigzirt.SetParent(f)
	
}
