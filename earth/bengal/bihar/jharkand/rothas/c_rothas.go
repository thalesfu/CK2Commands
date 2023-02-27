package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RothasCounty interface {
    feud.County
    BBetla吠多罗() 	feud.Barony
    BKarangi伽蓝耆() 	feud.Barony
    BKarumbakham伽蓝婆甘() 	feud.Barony
    BPalamau波罗摩郁() 	feud.Barony
    BRohtas卢诃多娑() 	feud.Barony
    BRohtasan卢诃多珊() 	feud.Barony
    BRohtasgarh卢诃多娑姞利呬() 	feud.Barony
}

type 卢诃多娑RothasCounty struct {
	feud.BaseCounty
	吠多罗Betla 	feud.Barony
	伽蓝耆Karangi 	feud.Barony
	伽蓝婆甘Karumbakham 	feud.Barony
	波罗摩郁Palamau 	feud.Barony
	卢诃多娑Rohtas 	feud.Barony
	卢诃多珊Rohtasan 	feud.Barony
	卢诃多娑姞利呬Rohtasgarh 	feud.Barony
}

func (c *卢诃多娑RothasCounty) BBetla吠多罗() feud.Barony {
	return c.吠多罗Betla
}
    
func (c *卢诃多娑RothasCounty) BKarangi伽蓝耆() feud.Barony {
	return c.伽蓝耆Karangi
}
    
func (c *卢诃多娑RothasCounty) BKarumbakham伽蓝婆甘() feud.Barony {
	return c.伽蓝婆甘Karumbakham
}
    
func (c *卢诃多娑RothasCounty) BPalamau波罗摩郁() feud.Barony {
	return c.波罗摩郁Palamau
}
    
func (c *卢诃多娑RothasCounty) BRohtas卢诃多娑() feud.Barony {
	return c.卢诃多娑Rohtas
}
    
func (c *卢诃多娑RothasCounty) BRohtasan卢诃多珊() feud.Barony {
	return c.卢诃多珊Rohtasan
}
    
func (c *卢诃多娑RothasCounty) BRohtasgarh卢诃多娑姞利呬() feud.Barony {
	return c.卢诃多娑姞利呬Rohtasgarh
}
    
var CRothas卢诃多娑 RothasCounty = &卢诃多娑RothasCounty{}

func init() {
	f := CRothas卢诃多娑.(*卢诃多娑RothasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1327",
		Title:     "rothas",
		TitleName: "卢诃多娑",
		TitleCode: "c_rothas",
		Baronies:  map[string]feud.Barony{},
	}

	f.吠多罗Betla = BBetla吠多罗
	f.吠多罗Betla.SetParent(f)
	
	f.伽蓝耆Karangi = BKarangi伽蓝耆
	f.伽蓝耆Karangi.SetParent(f)
	
	f.伽蓝婆甘Karumbakham = BKarumbakham伽蓝婆甘
	f.伽蓝婆甘Karumbakham.SetParent(f)
	
	f.波罗摩郁Palamau = BPalamau波罗摩郁
	f.波罗摩郁Palamau.SetParent(f)
	
	f.卢诃多娑Rohtas = BRohtas卢诃多娑
	f.卢诃多娑Rohtas.SetParent(f)
	
	f.卢诃多珊Rohtasan = BRohtasan卢诃多珊
	f.卢诃多珊Rohtasan.SetParent(f)
	
	f.卢诃多娑姞利呬Rohtasgarh = BRohtasgarh卢诃多娑姞利呬
	f.卢诃多娑姞利呬Rohtasgarh.SetParent(f)
	
}
