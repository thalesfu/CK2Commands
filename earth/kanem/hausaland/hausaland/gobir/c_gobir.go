package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GobirCounty interface {
    feud.County
    BBirnin_konni比尔尼恩孔尼() 	feud.Barony
    BDoba多巴() 	feud.Barony
    BEsuk_mesembe埃苏克梅森贝() 	feud.Barony
    BGobir戈比尔() 	feud.Barony
    BIerkpen伊尔克彭() 	feud.Barony
    BKpem佩梅() 	feud.Barony
    BWaja瓦贾() 	feud.Barony
}

type 戈比尔GobirCounty struct {
	feud.BaseCounty
	比尔尼恩孔尼Birnin_konni 	feud.Barony
	多巴Doba 	feud.Barony
	埃苏克梅森贝Esuk_mesembe 	feud.Barony
	戈比尔Gobir 	feud.Barony
	伊尔克彭Ierkpen 	feud.Barony
	佩梅Kpem 	feud.Barony
	瓦贾Waja 	feud.Barony
}

func (c *戈比尔GobirCounty) BBirnin_konni比尔尼恩孔尼() feud.Barony {
	return c.比尔尼恩孔尼Birnin_konni
}
    
func (c *戈比尔GobirCounty) BDoba多巴() feud.Barony {
	return c.多巴Doba
}
    
func (c *戈比尔GobirCounty) BEsuk_mesembe埃苏克梅森贝() feud.Barony {
	return c.埃苏克梅森贝Esuk_mesembe
}
    
func (c *戈比尔GobirCounty) BGobir戈比尔() feud.Barony {
	return c.戈比尔Gobir
}
    
func (c *戈比尔GobirCounty) BIerkpen伊尔克彭() feud.Barony {
	return c.伊尔克彭Ierkpen
}
    
func (c *戈比尔GobirCounty) BKpem佩梅() feud.Barony {
	return c.佩梅Kpem
}
    
func (c *戈比尔GobirCounty) BWaja瓦贾() feud.Barony {
	return c.瓦贾Waja
}
    
var CGobir戈比尔 GobirCounty = &戈比尔GobirCounty{}

func init() {
	f := CGobir戈比尔.(*戈比尔GobirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1748",
		Title:     "gobir",
		TitleName: "戈比尔",
		TitleCode: "c_gobir",
		Baronies:  map[string]feud.Barony{},
	}

	f.比尔尼恩孔尼Birnin_konni = BBirnin_konni比尔尼恩孔尼
	f.比尔尼恩孔尼Birnin_konni.SetParent(f)
	
	f.多巴Doba = BDoba多巴
	f.多巴Doba.SetParent(f)
	
	f.埃苏克梅森贝Esuk_mesembe = BEsuk_mesembe埃苏克梅森贝
	f.埃苏克梅森贝Esuk_mesembe.SetParent(f)
	
	f.戈比尔Gobir = BGobir戈比尔
	f.戈比尔Gobir.SetParent(f)
	
	f.伊尔克彭Ierkpen = BIerkpen伊尔克彭
	f.伊尔克彭Ierkpen.SetParent(f)
	
	f.佩梅Kpem = BKpem佩梅
	f.佩梅Kpem.SetParent(f)
	
	f.瓦贾Waja = BWaja瓦贾
	f.瓦贾Waja.SetParent(f)
	
}
