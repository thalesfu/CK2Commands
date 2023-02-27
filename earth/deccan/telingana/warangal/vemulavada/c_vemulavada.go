package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VemulavadaCounty interface {
    feud.County
    BBalijipeta婆梨恃卑吒() 	feud.Barony
    BIndur因杜尔() 	feud.Barony
    BKaleshwaram卡勒什瓦拉姆() 	feud.Barony
    BKaulas伽郁罗湿() 	feud.Barony
    BKonni拘尼() 	feud.Barony
    BPotali补多利() 	feud.Barony
    BVemulavada吠牟罗婆荼() 	feud.Barony
}

type 吠牟罗婆荼VemulavadaCounty struct {
	feud.BaseCounty
	婆梨恃卑吒Balijipeta 	feud.Barony
	因杜尔Indur 	feud.Barony
	卡勒什瓦拉姆Kaleshwaram 	feud.Barony
	伽郁罗湿Kaulas 	feud.Barony
	拘尼Konni 	feud.Barony
	补多利Potali 	feud.Barony
	吠牟罗婆荼Vemulavada 	feud.Barony
}

func (c *吠牟罗婆荼VemulavadaCounty) BBalijipeta婆梨恃卑吒() feud.Barony {
	return c.婆梨恃卑吒Balijipeta
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BIndur因杜尔() feud.Barony {
	return c.因杜尔Indur
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BKaleshwaram卡勒什瓦拉姆() feud.Barony {
	return c.卡勒什瓦拉姆Kaleshwaram
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BKaulas伽郁罗湿() feud.Barony {
	return c.伽郁罗湿Kaulas
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BKonni拘尼() feud.Barony {
	return c.拘尼Konni
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BPotali补多利() feud.Barony {
	return c.补多利Potali
}
    
func (c *吠牟罗婆荼VemulavadaCounty) BVemulavada吠牟罗婆荼() feud.Barony {
	return c.吠牟罗婆荼Vemulavada
}
    
var CVemulavada吠牟罗婆荼 VemulavadaCounty = &吠牟罗婆荼VemulavadaCounty{}

func init() {
	f := CVemulavada吠牟罗婆荼.(*吠牟罗婆荼VemulavadaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1255",
		Title:     "vemulavada",
		TitleName: "吠牟罗婆荼",
		TitleCode: "c_vemulavada",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆梨恃卑吒Balijipeta = BBalijipeta婆梨恃卑吒
	f.婆梨恃卑吒Balijipeta.SetParent(f)
	
	f.因杜尔Indur = BIndur因杜尔
	f.因杜尔Indur.SetParent(f)
	
	f.卡勒什瓦拉姆Kaleshwaram = BKaleshwaram卡勒什瓦拉姆
	f.卡勒什瓦拉姆Kaleshwaram.SetParent(f)
	
	f.伽郁罗湿Kaulas = BKaulas伽郁罗湿
	f.伽郁罗湿Kaulas.SetParent(f)
	
	f.拘尼Konni = BKonni拘尼
	f.拘尼Konni.SetParent(f)
	
	f.补多利Potali = BPotali补多利
	f.补多利Potali.SetParent(f)
	
	f.吠牟罗婆荼Vemulavada = BVemulavada吠牟罗婆荼
	f.吠牟罗婆荼Vemulavada.SetParent(f)
	
}
