package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SzekelyfoldCounty interface {
    feud.County
    BAranyos奥劳纽什塞克() 	feud.Barony
    BCsik奇克() 	feud.Barony
    BHaromszek哈罗姆塞克() 	feud.Barony
    BKezdi凯兹迪() 	feud.Barony
    BMaros毛罗什() 	feud.Barony
    BMarosvasarhely毛罗什瓦沙尔海伊() 	feud.Barony
    BSepsiszentgyorgy谢普希圣哲尔吉() 	feud.Barony
    BSzekelyudvarhely塞凯尤德沃尔海伊() 	feud.Barony
}

type 塞凯伊弗尔德SzekelyfoldCounty struct {
	feud.BaseCounty
	奥劳纽什塞克Aranyos 	feud.Barony
	奇克Csik 	feud.Barony
	哈罗姆塞克Haromszek 	feud.Barony
	凯兹迪Kezdi 	feud.Barony
	毛罗什Maros 	feud.Barony
	毛罗什瓦沙尔海伊Marosvasarhely 	feud.Barony
	谢普希圣哲尔吉Sepsiszentgyorgy 	feud.Barony
	塞凯尤德沃尔海伊Szekelyudvarhely 	feud.Barony
}

func (c *塞凯伊弗尔德SzekelyfoldCounty) BAranyos奥劳纽什塞克() feud.Barony {
	return c.奥劳纽什塞克Aranyos
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BCsik奇克() feud.Barony {
	return c.奇克Csik
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BHaromszek哈罗姆塞克() feud.Barony {
	return c.哈罗姆塞克Haromszek
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BKezdi凯兹迪() feud.Barony {
	return c.凯兹迪Kezdi
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BMaros毛罗什() feud.Barony {
	return c.毛罗什Maros
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BMarosvasarhely毛罗什瓦沙尔海伊() feud.Barony {
	return c.毛罗什瓦沙尔海伊Marosvasarhely
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BSepsiszentgyorgy谢普希圣哲尔吉() feud.Barony {
	return c.谢普希圣哲尔吉Sepsiszentgyorgy
}
    
func (c *塞凯伊弗尔德SzekelyfoldCounty) BSzekelyudvarhely塞凯尤德沃尔海伊() feud.Barony {
	return c.塞凯尤德沃尔海伊Szekelyudvarhely
}
    
var CSzekelyfold塞凯伊弗尔德 SzekelyfoldCounty = &塞凯伊弗尔德SzekelyfoldCounty{}

func init() {
	f := CSzekelyfold塞凯伊弗尔德.(*塞凯伊弗尔德SzekelyfoldCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "540",
		Title:     "szekelyfold",
		TitleName: "塞凯伊弗尔德",
		TitleCode: "c_szekelyfold",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥劳纽什塞克Aranyos = BAranyos奥劳纽什塞克
	f.奥劳纽什塞克Aranyos.SetParent(f)
	
	f.奇克Csik = BCsik奇克
	f.奇克Csik.SetParent(f)
	
	f.哈罗姆塞克Haromszek = BHaromszek哈罗姆塞克
	f.哈罗姆塞克Haromszek.SetParent(f)
	
	f.凯兹迪Kezdi = BKezdi凯兹迪
	f.凯兹迪Kezdi.SetParent(f)
	
	f.毛罗什Maros = BMaros毛罗什
	f.毛罗什Maros.SetParent(f)
	
	f.毛罗什瓦沙尔海伊Marosvasarhely = BMarosvasarhely毛罗什瓦沙尔海伊
	f.毛罗什瓦沙尔海伊Marosvasarhely.SetParent(f)
	
	f.谢普希圣哲尔吉Sepsiszentgyorgy = BSepsiszentgyorgy谢普希圣哲尔吉
	f.谢普希圣哲尔吉Sepsiszentgyorgy.SetParent(f)
	
	f.塞凯尤德沃尔海伊Szekelyudvarhely = BSzekelyudvarhely塞凯尤德沃尔海伊
	f.塞凯尤德沃尔海伊Szekelyudvarhely.SetParent(f)
	
}
