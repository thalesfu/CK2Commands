package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SozopolisCounty interface {
    feud.County
    BCadi加多() 	feud.Barony
    BColossae科洛塞() 	feud.Barony
    BDinar第纳尔() 	feud.Barony
    BIsparta伊斯帕尔塔() 	feud.Barony
    BKelainai克莱奈() 	feud.Barony
    BPolidorion波利德良() 	feud.Barony
    BSouzopolis索佐波利斯() 	feud.Barony
}

type 索佐波利斯SozopolisCounty struct {
	feud.BaseCounty
	加多Cadi 	feud.Barony
	科洛塞Colossae 	feud.Barony
	第纳尔Dinar 	feud.Barony
	伊斯帕尔塔Isparta 	feud.Barony
	克莱奈Kelainai 	feud.Barony
	波利德良Polidorion 	feud.Barony
	索佐波利斯Souzopolis 	feud.Barony
}

func (c *索佐波利斯SozopolisCounty) BCadi加多() feud.Barony {
	return c.加多Cadi
}
    
func (c *索佐波利斯SozopolisCounty) BColossae科洛塞() feud.Barony {
	return c.科洛塞Colossae
}
    
func (c *索佐波利斯SozopolisCounty) BDinar第纳尔() feud.Barony {
	return c.第纳尔Dinar
}
    
func (c *索佐波利斯SozopolisCounty) BIsparta伊斯帕尔塔() feud.Barony {
	return c.伊斯帕尔塔Isparta
}
    
func (c *索佐波利斯SozopolisCounty) BKelainai克莱奈() feud.Barony {
	return c.克莱奈Kelainai
}
    
func (c *索佐波利斯SozopolisCounty) BPolidorion波利德良() feud.Barony {
	return c.波利德良Polidorion
}
    
func (c *索佐波利斯SozopolisCounty) BSouzopolis索佐波利斯() feud.Barony {
	return c.索佐波利斯Souzopolis
}
    
var CSozopolis索佐波利斯 SozopolisCounty = &索佐波利斯SozopolisCounty{}

func init() {
	f := CSozopolis索佐波利斯.(*索佐波利斯SozopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "754",
		Title:     "sozopolis",
		TitleName: "索佐波利斯",
		TitleCode: "c_sozopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.加多Cadi = BCadi加多
	f.加多Cadi.SetParent(f)
	
	f.科洛塞Colossae = BColossae科洛塞
	f.科洛塞Colossae.SetParent(f)
	
	f.第纳尔Dinar = BDinar第纳尔
	f.第纳尔Dinar.SetParent(f)
	
	f.伊斯帕尔塔Isparta = BIsparta伊斯帕尔塔
	f.伊斯帕尔塔Isparta.SetParent(f)
	
	f.克莱奈Kelainai = BKelainai克莱奈
	f.克莱奈Kelainai.SetParent(f)
	
	f.波利德良Polidorion = BPolidorion波利德良
	f.波利德良Polidorion.SetParent(f)
	
	f.索佐波利斯Souzopolis = BSouzopolis索佐波利斯
	f.索佐波利斯Souzopolis.SetParent(f)
	
}
