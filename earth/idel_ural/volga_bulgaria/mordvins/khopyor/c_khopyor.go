package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhopyorCounty interface {
    feud.County
    BBalashov巴拉绍夫() 	feud.Barony
    BBorisoglebsk鲍里索格列布斯克() 	feud.Barony
    BKhopyorsk霍皮奥尔斯克() 	feud.Barony
    BKirsanov基尔萨诺夫() 	feud.Barony
    BTambov坦波夫() 	feud.Barony
    BUryupin乌留平() 	feud.Barony
    BUstmedveditskaya乌斯季_梅德韦季茨卡亚() 	feud.Barony
    BUvarovo乌瓦罗沃() 	feud.Barony
}

type 霍皮奥尔KhopyorCounty struct {
	feud.BaseCounty
	巴拉绍夫Balashov 	feud.Barony
	鲍里索格列布斯克Borisoglebsk 	feud.Barony
	霍皮奥尔斯克Khopyorsk 	feud.Barony
	基尔萨诺夫Kirsanov 	feud.Barony
	坦波夫Tambov 	feud.Barony
	乌留平Uryupin 	feud.Barony
	乌斯季_梅德韦季茨卡亚Ustmedveditskaya 	feud.Barony
	乌瓦罗沃Uvarovo 	feud.Barony
}

func (c *霍皮奥尔KhopyorCounty) BBalashov巴拉绍夫() feud.Barony {
	return c.巴拉绍夫Balashov
}
    
func (c *霍皮奥尔KhopyorCounty) BBorisoglebsk鲍里索格列布斯克() feud.Barony {
	return c.鲍里索格列布斯克Borisoglebsk
}
    
func (c *霍皮奥尔KhopyorCounty) BKhopyorsk霍皮奥尔斯克() feud.Barony {
	return c.霍皮奥尔斯克Khopyorsk
}
    
func (c *霍皮奥尔KhopyorCounty) BKirsanov基尔萨诺夫() feud.Barony {
	return c.基尔萨诺夫Kirsanov
}
    
func (c *霍皮奥尔KhopyorCounty) BTambov坦波夫() feud.Barony {
	return c.坦波夫Tambov
}
    
func (c *霍皮奥尔KhopyorCounty) BUryupin乌留平() feud.Barony {
	return c.乌留平Uryupin
}
    
func (c *霍皮奥尔KhopyorCounty) BUstmedveditskaya乌斯季_梅德韦季茨卡亚() feud.Barony {
	return c.乌斯季_梅德韦季茨卡亚Ustmedveditskaya
}
    
func (c *霍皮奥尔KhopyorCounty) BUvarovo乌瓦罗沃() feud.Barony {
	return c.乌瓦罗沃Uvarovo
}
    
var CKhopyor霍皮奥尔 KhopyorCounty = &霍皮奥尔KhopyorCounty{}

func init() {
	f := CKhopyor霍皮奥尔.(*霍皮奥尔KhopyorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "593",
		Title:     "khopyor",
		TitleName: "霍皮奥尔",
		TitleCode: "c_khopyor",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉绍夫Balashov = BBalashov巴拉绍夫
	f.巴拉绍夫Balashov.SetParent(f)
	
	f.鲍里索格列布斯克Borisoglebsk = BBorisoglebsk鲍里索格列布斯克
	f.鲍里索格列布斯克Borisoglebsk.SetParent(f)
	
	f.霍皮奥尔斯克Khopyorsk = BKhopyorsk霍皮奥尔斯克
	f.霍皮奥尔斯克Khopyorsk.SetParent(f)
	
	f.基尔萨诺夫Kirsanov = BKirsanov基尔萨诺夫
	f.基尔萨诺夫Kirsanov.SetParent(f)
	
	f.坦波夫Tambov = BTambov坦波夫
	f.坦波夫Tambov.SetParent(f)
	
	f.乌留平Uryupin = BUryupin乌留平
	f.乌留平Uryupin.SetParent(f)
	
	f.乌斯季_梅德韦季茨卡亚Ustmedveditskaya = BUstmedveditskaya乌斯季_梅德韦季茨卡亚
	f.乌斯季_梅德韦季茨卡亚Ustmedveditskaya.SetParent(f)
	
	f.乌瓦罗沃Uvarovo = BUvarovo乌瓦罗沃
	f.乌瓦罗沃Uvarovo.SetParent(f)
	
}
