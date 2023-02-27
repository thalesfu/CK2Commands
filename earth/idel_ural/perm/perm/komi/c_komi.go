package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KomiCounty interface {
    feud.County
    BAykino艾基诺() 	feud.Barony
    BInta因塔() 	feud.Barony
    BKhalmeryu哈利梅尔_尤() 	feud.Barony
    BKharp哈尔普() 	feud.Barony
    BUstkolom乌斯季_科洛姆() 	feud.Barony
    BUsttsilma乌斯季_齐利马() 	feud.Barony
    BVorgashor沃尔加绍尔() 	feud.Barony
    BVorkuta沃尔库塔() 	feud.Barony
}

type 皮利瓦KomiCounty struct {
	feud.BaseCounty
	艾基诺Aykino 	feud.Barony
	因塔Inta 	feud.Barony
	哈利梅尔_尤Khalmeryu 	feud.Barony
	哈尔普Kharp 	feud.Barony
	乌斯季_科洛姆Ustkolom 	feud.Barony
	乌斯季_齐利马Usttsilma 	feud.Barony
	沃尔加绍尔Vorgashor 	feud.Barony
	沃尔库塔Vorkuta 	feud.Barony
}

func (c *皮利瓦KomiCounty) BAykino艾基诺() feud.Barony {
	return c.艾基诺Aykino
}
    
func (c *皮利瓦KomiCounty) BInta因塔() feud.Barony {
	return c.因塔Inta
}
    
func (c *皮利瓦KomiCounty) BKhalmeryu哈利梅尔_尤() feud.Barony {
	return c.哈利梅尔_尤Khalmeryu
}
    
func (c *皮利瓦KomiCounty) BKharp哈尔普() feud.Barony {
	return c.哈尔普Kharp
}
    
func (c *皮利瓦KomiCounty) BUstkolom乌斯季_科洛姆() feud.Barony {
	return c.乌斯季_科洛姆Ustkolom
}
    
func (c *皮利瓦KomiCounty) BUsttsilma乌斯季_齐利马() feud.Barony {
	return c.乌斯季_齐利马Usttsilma
}
    
func (c *皮利瓦KomiCounty) BVorgashor沃尔加绍尔() feud.Barony {
	return c.沃尔加绍尔Vorgashor
}
    
func (c *皮利瓦KomiCounty) BVorkuta沃尔库塔() feud.Barony {
	return c.沃尔库塔Vorkuta
}
    
var CKomi皮利瓦 KomiCounty = &皮利瓦KomiCounty{}

func init() {
	f := CKomi皮利瓦.(*皮利瓦KomiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "888",
		Title:     "komi",
		TitleName: "皮利瓦",
		TitleCode: "c_komi",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾基诺Aykino = BAykino艾基诺
	f.艾基诺Aykino.SetParent(f)
	
	f.因塔Inta = BInta因塔
	f.因塔Inta.SetParent(f)
	
	f.哈利梅尔_尤Khalmeryu = BKhalmeryu哈利梅尔_尤
	f.哈利梅尔_尤Khalmeryu.SetParent(f)
	
	f.哈尔普Kharp = BKharp哈尔普
	f.哈尔普Kharp.SetParent(f)
	
	f.乌斯季_科洛姆Ustkolom = BUstkolom乌斯季_科洛姆
	f.乌斯季_科洛姆Ustkolom.SetParent(f)
	
	f.乌斯季_齐利马Usttsilma = BUsttsilma乌斯季_齐利马
	f.乌斯季_齐利马Usttsilma.SetParent(f)
	
	f.沃尔加绍尔Vorgashor = BVorgashor沃尔加绍尔
	f.沃尔加绍尔Vorgashor.SetParent(f)
	
	f.沃尔库塔Vorkuta = BVorkuta沃尔库塔
	f.沃尔库塔Vorkuta.SetParent(f)
	
}
