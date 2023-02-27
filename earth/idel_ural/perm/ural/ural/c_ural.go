package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UralCounty interface {
    feud.County
    BAsha阿沙() 	feud.Barony
    BChebarkul切巴尔库尔() 	feud.Barony
    BKaslinsky卡斯林斯基() 	feud.Barony
    BKyshtym克什特姆() 	feud.Barony
    BSatka萨特卡() 	feud.Barony
    BUral乌拉尔() 	feud.Barony
    BUstkatav乌斯季_卡塔夫() 	feud.Barony
}

type 乌拉尔UralCounty struct {
	feud.BaseCounty
	阿沙Asha 	feud.Barony
	切巴尔库尔Chebarkul 	feud.Barony
	卡斯林斯基Kaslinsky 	feud.Barony
	克什特姆Kyshtym 	feud.Barony
	萨特卡Satka 	feud.Barony
	乌拉尔Ural 	feud.Barony
	乌斯季_卡塔夫Ustkatav 	feud.Barony
}

func (c *乌拉尔UralCounty) BAsha阿沙() feud.Barony {
	return c.阿沙Asha
}
    
func (c *乌拉尔UralCounty) BChebarkul切巴尔库尔() feud.Barony {
	return c.切巴尔库尔Chebarkul
}
    
func (c *乌拉尔UralCounty) BKaslinsky卡斯林斯基() feud.Barony {
	return c.卡斯林斯基Kaslinsky
}
    
func (c *乌拉尔UralCounty) BKyshtym克什特姆() feud.Barony {
	return c.克什特姆Kyshtym
}
    
func (c *乌拉尔UralCounty) BSatka萨特卡() feud.Barony {
	return c.萨特卡Satka
}
    
func (c *乌拉尔UralCounty) BUral乌拉尔() feud.Barony {
	return c.乌拉尔Ural
}
    
func (c *乌拉尔UralCounty) BUstkatav乌斯季_卡塔夫() feud.Barony {
	return c.乌斯季_卡塔夫Ustkatav
}
    
var CUral乌拉尔 UralCounty = &乌拉尔UralCounty{}

func init() {
	f := CUral乌拉尔.(*乌拉尔UralCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "889",
		Title:     "ural",
		TitleName: "乌拉尔",
		TitleCode: "c_ural",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿沙Asha = BAsha阿沙
	f.阿沙Asha.SetParent(f)
	
	f.切巴尔库尔Chebarkul = BChebarkul切巴尔库尔
	f.切巴尔库尔Chebarkul.SetParent(f)
	
	f.卡斯林斯基Kaslinsky = BKaslinsky卡斯林斯基
	f.卡斯林斯基Kaslinsky.SetParent(f)
	
	f.克什特姆Kyshtym = BKyshtym克什特姆
	f.克什特姆Kyshtym.SetParent(f)
	
	f.萨特卡Satka = BSatka萨特卡
	f.萨特卡Satka.SetParent(f)
	
	f.乌拉尔Ural = BUral乌拉尔
	f.乌拉尔Ural.SetParent(f)
	
	f.乌斯季_卡塔夫Ustkatav = BUstkatav乌斯季_卡塔夫
	f.乌斯季_卡塔夫Ustkatav.SetParent(f)
	
}
