package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamohCounty interface {
    feud.County
    BDamoh佗牟() 	feud.Barony
    BDhamoni佗牟尼() 	feud.Barony
    BNarsinghgarh那罗僧诃堡() 	feud.Barony
    BNimora尼牟罗() 	feud.Barony
    BNohta奴诃多() 	feud.Barony
    BPanwel般韦尔() 	feud.Barony
    BSingrampur僧伽罗摩城() 	feud.Barony
}

type 佗牟DamohCounty struct {
	feud.BaseCounty
	佗牟Damoh 	feud.Barony
	佗牟尼Dhamoni 	feud.Barony
	那罗僧诃堡Narsinghgarh 	feud.Barony
	尼牟罗Nimora 	feud.Barony
	奴诃多Nohta 	feud.Barony
	般韦尔Panwel 	feud.Barony
	僧伽罗摩城Singrampur 	feud.Barony
}

func (c *佗牟DamohCounty) BDamoh佗牟() feud.Barony {
	return c.佗牟Damoh
}
    
func (c *佗牟DamohCounty) BDhamoni佗牟尼() feud.Barony {
	return c.佗牟尼Dhamoni
}
    
func (c *佗牟DamohCounty) BNarsinghgarh那罗僧诃堡() feud.Barony {
	return c.那罗僧诃堡Narsinghgarh
}
    
func (c *佗牟DamohCounty) BNimora尼牟罗() feud.Barony {
	return c.尼牟罗Nimora
}
    
func (c *佗牟DamohCounty) BNohta奴诃多() feud.Barony {
	return c.奴诃多Nohta
}
    
func (c *佗牟DamohCounty) BPanwel般韦尔() feud.Barony {
	return c.般韦尔Panwel
}
    
func (c *佗牟DamohCounty) BSingrampur僧伽罗摩城() feud.Barony {
	return c.僧伽罗摩城Singrampur
}
    
var CDamoh佗牟 DamohCounty = &佗牟DamohCounty{}

func init() {
	f := CDamoh佗牟.(*佗牟DamohCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1274",
		Title:     "damoh",
		TitleName: "佗牟",
		TitleCode: "c_damoh",
		Baronies:  map[string]feud.Barony{},
	}

	f.佗牟Damoh = BDamoh佗牟
	f.佗牟Damoh.SetParent(f)
	
	f.佗牟尼Dhamoni = BDhamoni佗牟尼
	f.佗牟尼Dhamoni.SetParent(f)
	
	f.那罗僧诃堡Narsinghgarh = BNarsinghgarh那罗僧诃堡
	f.那罗僧诃堡Narsinghgarh.SetParent(f)
	
	f.尼牟罗Nimora = BNimora尼牟罗
	f.尼牟罗Nimora.SetParent(f)
	
	f.奴诃多Nohta = BNohta奴诃多
	f.奴诃多Nohta.SetParent(f)
	
	f.般韦尔Panwel = BPanwel般韦尔
	f.般韦尔Panwel.SetParent(f)
	
	f.僧伽罗摩城Singrampur = BSingrampur僧伽罗摩城
	f.僧伽罗摩城Singrampur.SetParent(f)
	
}
