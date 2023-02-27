package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaltaCounty interface {
    feud.County
    BBirzebbuga奇塔代拉() 	feud.Barony
    BGozo海堡() 	feud.Barony
    BMarsascala博尔戈() 	feud.Barony
    BMdina姆迪纳() 	feud.Barony
    BMgarr祖里科() 	feud.Barony
    BSangjilan纳斯卡里奥() 	feud.Barony
    BSanpawl圣保罗() 	feud.Barony
    BSliema拉巴特() 	feud.Barony
}

type 马耳他MaltaCounty struct {
	feud.BaseCounty
	奇塔代拉Birzebbuga 	feud.Barony
	海堡Gozo 	feud.Barony
	博尔戈Marsascala 	feud.Barony
	姆迪纳Mdina 	feud.Barony
	祖里科Mgarr 	feud.Barony
	纳斯卡里奥Sangjilan 	feud.Barony
	圣保罗Sanpawl 	feud.Barony
	拉巴特Sliema 	feud.Barony
}

func (c *马耳他MaltaCounty) BBirzebbuga奇塔代拉() feud.Barony {
	return c.奇塔代拉Birzebbuga
}
    
func (c *马耳他MaltaCounty) BGozo海堡() feud.Barony {
	return c.海堡Gozo
}
    
func (c *马耳他MaltaCounty) BMarsascala博尔戈() feud.Barony {
	return c.博尔戈Marsascala
}
    
func (c *马耳他MaltaCounty) BMdina姆迪纳() feud.Barony {
	return c.姆迪纳Mdina
}
    
func (c *马耳他MaltaCounty) BMgarr祖里科() feud.Barony {
	return c.祖里科Mgarr
}
    
func (c *马耳他MaltaCounty) BSangjilan纳斯卡里奥() feud.Barony {
	return c.纳斯卡里奥Sangjilan
}
    
func (c *马耳他MaltaCounty) BSanpawl圣保罗() feud.Barony {
	return c.圣保罗Sanpawl
}
    
func (c *马耳他MaltaCounty) BSliema拉巴特() feud.Barony {
	return c.拉巴特Sliema
}
    
var CMalta马耳他 MaltaCounty = &马耳他MaltaCounty{}

func init() {
	f := CMalta马耳他.(*马耳他MaltaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "812",
		Title:     "malta",
		TitleName: "马耳他",
		TitleCode: "c_malta",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇塔代拉Birzebbuga = BBirzebbuga奇塔代拉
	f.奇塔代拉Birzebbuga.SetParent(f)
	
	f.海堡Gozo = BGozo海堡
	f.海堡Gozo.SetParent(f)
	
	f.博尔戈Marsascala = BMarsascala博尔戈
	f.博尔戈Marsascala.SetParent(f)
	
	f.姆迪纳Mdina = BMdina姆迪纳
	f.姆迪纳Mdina.SetParent(f)
	
	f.祖里科Mgarr = BMgarr祖里科
	f.祖里科Mgarr.SetParent(f)
	
	f.纳斯卡里奥Sangjilan = BSangjilan纳斯卡里奥
	f.纳斯卡里奥Sangjilan.SetParent(f)
	
	f.圣保罗Sanpawl = BSanpawl圣保罗
	f.圣保罗Sanpawl.SetParent(f)
	
	f.拉巴特Sliema = BSliema拉巴特
	f.拉巴特Sliema.SetParent(f)
	
}
