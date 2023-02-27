package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BumthangCounty interface {
    feud.County
    BBumthang布姆唐() 	feud.Barony
    BJambay强贝() 	feud.Barony
    BKongchogsum贡觉孙() 	feud.Barony
    BSingye辛格() 	feud.Barony
    BTrashigang扎西岗() 	feud.Barony
    BTrongsa通萨() 	feud.Barony
    BZugne祖恩() 	feud.Barony
}

type 布姆唐BumthangCounty struct {
	feud.BaseCounty
	布姆唐Bumthang 	feud.Barony
	强贝Jambay 	feud.Barony
	贡觉孙Kongchogsum 	feud.Barony
	辛格Singye 	feud.Barony
	扎西岗Trashigang 	feud.Barony
	通萨Trongsa 	feud.Barony
	祖恩Zugne 	feud.Barony
}

func (c *布姆唐BumthangCounty) BBumthang布姆唐() feud.Barony {
	return c.布姆唐Bumthang
}
    
func (c *布姆唐BumthangCounty) BJambay强贝() feud.Barony {
	return c.强贝Jambay
}
    
func (c *布姆唐BumthangCounty) BKongchogsum贡觉孙() feud.Barony {
	return c.贡觉孙Kongchogsum
}
    
func (c *布姆唐BumthangCounty) BSingye辛格() feud.Barony {
	return c.辛格Singye
}
    
func (c *布姆唐BumthangCounty) BTrashigang扎西岗() feud.Barony {
	return c.扎西岗Trashigang
}
    
func (c *布姆唐BumthangCounty) BTrongsa通萨() feud.Barony {
	return c.通萨Trongsa
}
    
func (c *布姆唐BumthangCounty) BZugne祖恩() feud.Barony {
	return c.祖恩Zugne
}
    
var CBumthang布姆唐 BumthangCounty = &布姆唐BumthangCounty{}

func init() {
	f := CBumthang布姆唐.(*布姆唐BumthangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1482",
		Title:     "bumthang",
		TitleName: "布姆唐",
		TitleCode: "c_bumthang",
		Baronies:  map[string]feud.Barony{},
	}

	f.布姆唐Bumthang = BBumthang布姆唐
	f.布姆唐Bumthang.SetParent(f)
	
	f.强贝Jambay = BJambay强贝
	f.强贝Jambay.SetParent(f)
	
	f.贡觉孙Kongchogsum = BKongchogsum贡觉孙
	f.贡觉孙Kongchogsum.SetParent(f)
	
	f.辛格Singye = BSingye辛格
	f.辛格Singye.SetParent(f)
	
	f.扎西岗Trashigang = BTrashigang扎西岗
	f.扎西岗Trashigang.SetParent(f)
	
	f.通萨Trongsa = BTrongsa通萨
	f.通萨Trongsa.SetParent(f)
	
	f.祖恩Zugne = BZugne祖恩
	f.祖恩Zugne.SetParent(f)
	
}
