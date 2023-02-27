package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BratslavCounty interface {
    feud.County
    BBratslav布拉茨拉夫() 	feud.Barony
    BKryzhopil克雷若皮尔() 	feud.Barony
    BNemyriv涅米罗夫() 	feud.Barony
    BSharhorod沙尔霍罗德() 	feud.Barony
    BTomashpil托马什皮尔() 	feud.Barony
    BYampil扬皮尔() 	feud.Barony
    BZhmerynka日梅林卡() 	feud.Barony
}

type 布拉茨拉夫BratslavCounty struct {
	feud.BaseCounty
	布拉茨拉夫Bratslav 	feud.Barony
	克雷若皮尔Kryzhopil 	feud.Barony
	涅米罗夫Nemyriv 	feud.Barony
	沙尔霍罗德Sharhorod 	feud.Barony
	托马什皮尔Tomashpil 	feud.Barony
	扬皮尔Yampil 	feud.Barony
	日梅林卡Zhmerynka 	feud.Barony
}

func (c *布拉茨拉夫BratslavCounty) BBratslav布拉茨拉夫() feud.Barony {
	return c.布拉茨拉夫Bratslav
}
    
func (c *布拉茨拉夫BratslavCounty) BKryzhopil克雷若皮尔() feud.Barony {
	return c.克雷若皮尔Kryzhopil
}
    
func (c *布拉茨拉夫BratslavCounty) BNemyriv涅米罗夫() feud.Barony {
	return c.涅米罗夫Nemyriv
}
    
func (c *布拉茨拉夫BratslavCounty) BSharhorod沙尔霍罗德() feud.Barony {
	return c.沙尔霍罗德Sharhorod
}
    
func (c *布拉茨拉夫BratslavCounty) BTomashpil托马什皮尔() feud.Barony {
	return c.托马什皮尔Tomashpil
}
    
func (c *布拉茨拉夫BratslavCounty) BYampil扬皮尔() feud.Barony {
	return c.扬皮尔Yampil
}
    
func (c *布拉茨拉夫BratslavCounty) BZhmerynka日梅林卡() feud.Barony {
	return c.日梅林卡Zhmerynka
}
    
var CBratslav布拉茨拉夫 BratslavCounty = &布拉茨拉夫BratslavCounty{}

func init() {
	f := CBratslav布拉茨拉夫.(*布拉茨拉夫BratslavCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1641",
		Title:     "bratslav",
		TitleName: "布拉茨拉夫",
		TitleCode: "c_bratslav",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉茨拉夫Bratslav = BBratslav布拉茨拉夫
	f.布拉茨拉夫Bratslav.SetParent(f)
	
	f.克雷若皮尔Kryzhopil = BKryzhopil克雷若皮尔
	f.克雷若皮尔Kryzhopil.SetParent(f)
	
	f.涅米罗夫Nemyriv = BNemyriv涅米罗夫
	f.涅米罗夫Nemyriv.SetParent(f)
	
	f.沙尔霍罗德Sharhorod = BSharhorod沙尔霍罗德
	f.沙尔霍罗德Sharhorod.SetParent(f)
	
	f.托马什皮尔Tomashpil = BTomashpil托马什皮尔
	f.托马什皮尔Tomashpil.SetParent(f)
	
	f.扬皮尔Yampil = BYampil扬皮尔
	f.扬皮尔Yampil.SetParent(f)
	
	f.日梅林卡Zhmerynka = BZhmerynka日梅林卡
	f.日梅林卡Zhmerynka.SetParent(f)
	
}
