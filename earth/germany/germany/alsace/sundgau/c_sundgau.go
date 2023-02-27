package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SundgauCounty interface {
    feud.County
    BAltkirch阿尔特基什() 	feud.Barony
    BEnsisheim昂西塞姆() 	feud.Barony
    BFerette费雷特() 	feud.Barony
    BKolmar科尔马() 	feud.Barony
    BLandser朗塞() 	feud.Barony
    BMulhouse米卢斯() 	feud.Barony
    BMurbach米尔巴克() 	feud.Barony
    BThann坦恩() 	feud.Barony
}

type 松德高SundgauCounty struct {
	feud.BaseCounty
	阿尔特基什Altkirch 	feud.Barony
	昂西塞姆Ensisheim 	feud.Barony
	费雷特Ferette 	feud.Barony
	科尔马Kolmar 	feud.Barony
	朗塞Landser 	feud.Barony
	米卢斯Mulhouse 	feud.Barony
	米尔巴克Murbach 	feud.Barony
	坦恩Thann 	feud.Barony
}

func (c *松德高SundgauCounty) BAltkirch阿尔特基什() feud.Barony {
	return c.阿尔特基什Altkirch
}
    
func (c *松德高SundgauCounty) BEnsisheim昂西塞姆() feud.Barony {
	return c.昂西塞姆Ensisheim
}
    
func (c *松德高SundgauCounty) BFerette费雷特() feud.Barony {
	return c.费雷特Ferette
}
    
func (c *松德高SundgauCounty) BKolmar科尔马() feud.Barony {
	return c.科尔马Kolmar
}
    
func (c *松德高SundgauCounty) BLandser朗塞() feud.Barony {
	return c.朗塞Landser
}
    
func (c *松德高SundgauCounty) BMulhouse米卢斯() feud.Barony {
	return c.米卢斯Mulhouse
}
    
func (c *松德高SundgauCounty) BMurbach米尔巴克() feud.Barony {
	return c.米尔巴克Murbach
}
    
func (c *松德高SundgauCounty) BThann坦恩() feud.Barony {
	return c.坦恩Thann
}
    
var CSundgau松德高 SundgauCounty = &松德高SundgauCounty{}

func init() {
	f := CSundgau松德高.(*松德高SundgauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "134",
		Title:     "sundgau",
		TitleName: "松德高",
		TitleCode: "c_sundgau",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔特基什Altkirch = BAltkirch阿尔特基什
	f.阿尔特基什Altkirch.SetParent(f)
	
	f.昂西塞姆Ensisheim = BEnsisheim昂西塞姆
	f.昂西塞姆Ensisheim.SetParent(f)
	
	f.费雷特Ferette = BFerette费雷特
	f.费雷特Ferette.SetParent(f)
	
	f.科尔马Kolmar = BKolmar科尔马
	f.科尔马Kolmar.SetParent(f)
	
	f.朗塞Landser = BLandser朗塞
	f.朗塞Landser.SetParent(f)
	
	f.米卢斯Mulhouse = BMulhouse米卢斯
	f.米卢斯Mulhouse.SetParent(f)
	
	f.米尔巴克Murbach = BMurbach米尔巴克
	f.米尔巴克Murbach.SetParent(f)
	
	f.坦恩Thann = BThann坦恩
	f.坦恩Thann.SetParent(f)
	
}
