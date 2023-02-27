package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaqsinCounty interface {
    feud.County
    BBundin本丁() 	feud.Barony
    BEl_ton埃尔顿() 	feud.Barony
    BKatrichev卡特里切夫() 	feud.Barony
    BMaksima马克西马() 	feud.Barony
    BSaqsin撒哈辛() 	feud.Barony
    BTsarev察列夫() 	feud.Barony
    BZhyra日拉() 	feud.Barony
}

type 撒哈辛SaqsinCounty struct {
	feud.BaseCounty
	本丁Bundin 	feud.Barony
	埃尔顿El_ton 	feud.Barony
	卡特里切夫Katrichev 	feud.Barony
	马克西马Maksima 	feud.Barony
	撒哈辛Saqsin 	feud.Barony
	察列夫Tsarev 	feud.Barony
	日拉Zhyra 	feud.Barony
}

func (c *撒哈辛SaqsinCounty) BBundin本丁() feud.Barony {
	return c.本丁Bundin
}
    
func (c *撒哈辛SaqsinCounty) BEl_ton埃尔顿() feud.Barony {
	return c.埃尔顿El_ton
}
    
func (c *撒哈辛SaqsinCounty) BKatrichev卡特里切夫() feud.Barony {
	return c.卡特里切夫Katrichev
}
    
func (c *撒哈辛SaqsinCounty) BMaksima马克西马() feud.Barony {
	return c.马克西马Maksima
}
    
func (c *撒哈辛SaqsinCounty) BSaqsin撒哈辛() feud.Barony {
	return c.撒哈辛Saqsin
}
    
func (c *撒哈辛SaqsinCounty) BTsarev察列夫() feud.Barony {
	return c.察列夫Tsarev
}
    
func (c *撒哈辛SaqsinCounty) BZhyra日拉() feud.Barony {
	return c.日拉Zhyra
}
    
var CSaqsin撒哈辛 SaqsinCounty = &撒哈辛SaqsinCounty{}

func init() {
	f := CSaqsin撒哈辛.(*撒哈辛SaqsinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1812",
		Title:     "saqsin",
		TitleName: "撒哈辛",
		TitleCode: "c_saqsin",
		Baronies:  map[string]feud.Barony{},
	}

	f.本丁Bundin = BBundin本丁
	f.本丁Bundin.SetParent(f)
	
	f.埃尔顿El_ton = BEl_ton埃尔顿
	f.埃尔顿El_ton.SetParent(f)
	
	f.卡特里切夫Katrichev = BKatrichev卡特里切夫
	f.卡特里切夫Katrichev.SetParent(f)
	
	f.马克西马Maksima = BMaksima马克西马
	f.马克西马Maksima.SetParent(f)
	
	f.撒哈辛Saqsin = BSaqsin撒哈辛
	f.撒哈辛Saqsin.SetParent(f)
	
	f.察列夫Tsarev = BTsarev察列夫
	f.察列夫Tsarev.SetParent(f)
	
	f.日拉Zhyra = BZhyra日拉
	f.日拉Zhyra.SetParent(f)
	
}
