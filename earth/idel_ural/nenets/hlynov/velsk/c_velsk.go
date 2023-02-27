package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VelskCounty interface {
    feud.County
    BKuloy_velsk库洛伊() 	feud.Barony
    BPezhma佩日马() 	feud.Barony
    BShilovskaya希洛夫斯卡亚() 	feud.Barony
    BShunema舒涅马() 	feud.Barony
    BVaga_velsk瓦加() 	feud.Barony
    BVel_velsk韦尔() 	feud.Barony
    BVelsk韦利斯克() 	feud.Barony
}

type 韦利斯克VelskCounty struct {
	feud.BaseCounty
	库洛伊Kuloy_velsk 	feud.Barony
	佩日马Pezhma 	feud.Barony
	希洛夫斯卡亚Shilovskaya 	feud.Barony
	舒涅马Shunema 	feud.Barony
	瓦加Vaga_velsk 	feud.Barony
	韦尔Vel_velsk 	feud.Barony
	韦利斯克Velsk 	feud.Barony
}

func (c *韦利斯克VelskCounty) BKuloy_velsk库洛伊() feud.Barony {
	return c.库洛伊Kuloy_velsk
}
    
func (c *韦利斯克VelskCounty) BPezhma佩日马() feud.Barony {
	return c.佩日马Pezhma
}
    
func (c *韦利斯克VelskCounty) BShilovskaya希洛夫斯卡亚() feud.Barony {
	return c.希洛夫斯卡亚Shilovskaya
}
    
func (c *韦利斯克VelskCounty) BShunema舒涅马() feud.Barony {
	return c.舒涅马Shunema
}
    
func (c *韦利斯克VelskCounty) BVaga_velsk瓦加() feud.Barony {
	return c.瓦加Vaga_velsk
}
    
func (c *韦利斯克VelskCounty) BVel_velsk韦尔() feud.Barony {
	return c.韦尔Vel_velsk
}
    
func (c *韦利斯克VelskCounty) BVelsk韦利斯克() feud.Barony {
	return c.韦利斯克Velsk
}
    
var CVelsk韦利斯克 VelskCounty = &韦利斯克VelskCounty{}

func init() {
	f := CVelsk韦利斯克.(*韦利斯克VelskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1826",
		Title:     "velsk",
		TitleName: "韦利斯克",
		TitleCode: "c_velsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.库洛伊Kuloy_velsk = BKuloy_velsk库洛伊
	f.库洛伊Kuloy_velsk.SetParent(f)
	
	f.佩日马Pezhma = BPezhma佩日马
	f.佩日马Pezhma.SetParent(f)
	
	f.希洛夫斯卡亚Shilovskaya = BShilovskaya希洛夫斯卡亚
	f.希洛夫斯卡亚Shilovskaya.SetParent(f)
	
	f.舒涅马Shunema = BShunema舒涅马
	f.舒涅马Shunema.SetParent(f)
	
	f.瓦加Vaga_velsk = BVaga_velsk瓦加
	f.瓦加Vaga_velsk.SetParent(f)
	
	f.韦尔Vel_velsk = BVel_velsk韦尔
	f.韦尔Vel_velsk.SetParent(f)
	
	f.韦利斯克Velsk = BVelsk韦利斯克
	f.韦利斯克Velsk.SetParent(f)
	
}
