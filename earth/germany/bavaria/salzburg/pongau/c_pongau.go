package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PongauCounty interface {
    feud.County
    BHoven霍文() 	feud.Barony
    BMittersill米特西尔() 	feud.Barony
    BPongau蓬高() 	feud.Barony
    BSalvet萨尔韦特() 	feud.Barony
    BSt_maximillian圣马克西米利安() 	feud.Barony
    BTaemswich塔姆斯韦格() 	feud.Barony
    BZelle采尔() 	feud.Barony
}

type 蓬高PongauCounty struct {
	feud.BaseCounty
	霍文Hoven 	feud.Barony
	米特西尔Mittersill 	feud.Barony
	蓬高Pongau 	feud.Barony
	萨尔韦特Salvet 	feud.Barony
	圣马克西米利安St_maximillian 	feud.Barony
	塔姆斯韦格Taemswich 	feud.Barony
	采尔Zelle 	feud.Barony
}

func (c *蓬高PongauCounty) BHoven霍文() feud.Barony {
	return c.霍文Hoven
}
    
func (c *蓬高PongauCounty) BMittersill米特西尔() feud.Barony {
	return c.米特西尔Mittersill
}
    
func (c *蓬高PongauCounty) BPongau蓬高() feud.Barony {
	return c.蓬高Pongau
}
    
func (c *蓬高PongauCounty) BSalvet萨尔韦特() feud.Barony {
	return c.萨尔韦特Salvet
}
    
func (c *蓬高PongauCounty) BSt_maximillian圣马克西米利安() feud.Barony {
	return c.圣马克西米利安St_maximillian
}
    
func (c *蓬高PongauCounty) BTaemswich塔姆斯韦格() feud.Barony {
	return c.塔姆斯韦格Taemswich
}
    
func (c *蓬高PongauCounty) BZelle采尔() feud.Barony {
	return c.采尔Zelle
}
    
var CPongau蓬高 PongauCounty = &蓬高PongauCounty{}

func init() {
	f := CPongau蓬高.(*蓬高PongauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1689",
		Title:     "pongau",
		TitleName: "蓬高",
		TitleCode: "c_pongau",
		Baronies:  map[string]feud.Barony{},
	}

	f.霍文Hoven = BHoven霍文
	f.霍文Hoven.SetParent(f)
	
	f.米特西尔Mittersill = BMittersill米特西尔
	f.米特西尔Mittersill.SetParent(f)
	
	f.蓬高Pongau = BPongau蓬高
	f.蓬高Pongau.SetParent(f)
	
	f.萨尔韦特Salvet = BSalvet萨尔韦特
	f.萨尔韦特Salvet.SetParent(f)
	
	f.圣马克西米利安St_maximillian = BSt_maximillian圣马克西米利安
	f.圣马克西米利安St_maximillian.SetParent(f)
	
	f.塔姆斯韦格Taemswich = BTaemswich塔姆斯韦格
	f.塔姆斯韦格Taemswich.SetParent(f)
	
	f.采尔Zelle = BZelle采尔
	f.采尔Zelle.SetParent(f)
	
}
