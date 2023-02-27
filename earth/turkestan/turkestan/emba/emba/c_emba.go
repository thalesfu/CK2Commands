package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EmbaCounty interface {
    feud.County
    BAkbulak阿克布拉克() 	feud.Barony
    BAkkube阿库别() 	feud.Barony
    BAzgyl阿兹格尔() 	feud.Barony
    BBakachi巴卡奇() 	feud.Barony
    BBesbay别斯拜() 	feud.Barony
    BEmba恩巴() 	feud.Barony
    BMukur穆库尔() 	feud.Barony
}

type 恩巴EmbaCounty struct {
	feud.BaseCounty
	阿克布拉克Akbulak 	feud.Barony
	阿库别Akkube 	feud.Barony
	阿兹格尔Azgyl 	feud.Barony
	巴卡奇Bakachi 	feud.Barony
	别斯拜Besbay 	feud.Barony
	恩巴Emba 	feud.Barony
	穆库尔Mukur 	feud.Barony
}

func (c *恩巴EmbaCounty) BAkbulak阿克布拉克() feud.Barony {
	return c.阿克布拉克Akbulak
}
    
func (c *恩巴EmbaCounty) BAkkube阿库别() feud.Barony {
	return c.阿库别Akkube
}
    
func (c *恩巴EmbaCounty) BAzgyl阿兹格尔() feud.Barony {
	return c.阿兹格尔Azgyl
}
    
func (c *恩巴EmbaCounty) BBakachi巴卡奇() feud.Barony {
	return c.巴卡奇Bakachi
}
    
func (c *恩巴EmbaCounty) BBesbay别斯拜() feud.Barony {
	return c.别斯拜Besbay
}
    
func (c *恩巴EmbaCounty) BEmba恩巴() feud.Barony {
	return c.恩巴Emba
}
    
func (c *恩巴EmbaCounty) BMukur穆库尔() feud.Barony {
	return c.穆库尔Mukur
}
    
var CEmba恩巴 EmbaCounty = &恩巴EmbaCounty{}

func init() {
	f := CEmba恩巴.(*恩巴EmbaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1793",
		Title:     "emba",
		TitleName: "恩巴",
		TitleCode: "c_emba",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克布拉克Akbulak = BAkbulak阿克布拉克
	f.阿克布拉克Akbulak.SetParent(f)
	
	f.阿库别Akkube = BAkkube阿库别
	f.阿库别Akkube.SetParent(f)
	
	f.阿兹格尔Azgyl = BAzgyl阿兹格尔
	f.阿兹格尔Azgyl.SetParent(f)
	
	f.巴卡奇Bakachi = BBakachi巴卡奇
	f.巴卡奇Bakachi.SetParent(f)
	
	f.别斯拜Besbay = BBesbay别斯拜
	f.别斯拜Besbay.SetParent(f)
	
	f.恩巴Emba = BEmba恩巴
	f.恩巴Emba.SetParent(f)
	
	f.穆库尔Mukur = BMukur穆库尔
	f.穆库尔Mukur.SetParent(f)
	
}
