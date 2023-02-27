package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IonaCounty interface {
    feud.County
    BArinagour阿里纳古尔() 	feud.Barony
    BDunyveg丹尼维哥堡() 	feud.Barony
    BFinlaggan芬拉甘() 	feud.Barony
    BIona艾奥纳() 	feud.Barony
    BLaggan拉根() 	feud.Barony
    BScarinish斯卡里尼什() 	feud.Barony
    BTobermory托伯莫里() 	feud.Barony
}

type 艾奥纳IonaCounty struct {
	feud.BaseCounty
	阿里纳古尔Arinagour 	feud.Barony
	丹尼维哥堡Dunyveg 	feud.Barony
	芬拉甘Finlaggan 	feud.Barony
	艾奥纳Iona 	feud.Barony
	拉根Laggan 	feud.Barony
	斯卡里尼什Scarinish 	feud.Barony
	托伯莫里Tobermory 	feud.Barony
}

func (c *艾奥纳IonaCounty) BArinagour阿里纳古尔() feud.Barony {
	return c.阿里纳古尔Arinagour
}
    
func (c *艾奥纳IonaCounty) BDunyveg丹尼维哥堡() feud.Barony {
	return c.丹尼维哥堡Dunyveg
}
    
func (c *艾奥纳IonaCounty) BFinlaggan芬拉甘() feud.Barony {
	return c.芬拉甘Finlaggan
}
    
func (c *艾奥纳IonaCounty) BIona艾奥纳() feud.Barony {
	return c.艾奥纳Iona
}
    
func (c *艾奥纳IonaCounty) BLaggan拉根() feud.Barony {
	return c.拉根Laggan
}
    
func (c *艾奥纳IonaCounty) BScarinish斯卡里尼什() feud.Barony {
	return c.斯卡里尼什Scarinish
}
    
func (c *艾奥纳IonaCounty) BTobermory托伯莫里() feud.Barony {
	return c.托伯莫里Tobermory
}
    
var CIona艾奥纳 IonaCounty = &艾奥纳IonaCounty{}

func init() {
	f := CIona艾奥纳.(*艾奥纳IonaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1939",
		Title:     "iona",
		TitleName: "艾奥纳",
		TitleCode: "c_iona",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿里纳古尔Arinagour = BArinagour阿里纳古尔
	f.阿里纳古尔Arinagour.SetParent(f)
	
	f.丹尼维哥堡Dunyveg = BDunyveg丹尼维哥堡
	f.丹尼维哥堡Dunyveg.SetParent(f)
	
	f.芬拉甘Finlaggan = BFinlaggan芬拉甘
	f.芬拉甘Finlaggan.SetParent(f)
	
	f.艾奥纳Iona = BIona艾奥纳
	f.艾奥纳Iona.SetParent(f)
	
	f.拉根Laggan = BLaggan拉根
	f.拉根Laggan.SetParent(f)
	
	f.斯卡里尼什Scarinish = BScarinish斯卡里尼什
	f.斯卡里尼什Scarinish.SetParent(f)
	
	f.托伯莫里Tobermory = BTobermory托伯莫里
	f.托伯莫里Tobermory.SetParent(f)
	
}
