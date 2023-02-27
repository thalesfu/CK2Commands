package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VoinCounty interface {
    feud.County
    BBohodukhivka博戈杜霍夫卡() 	feud.Barony
    BChornobai乔尔诺拜() 	feud.Barony
    BDrabiv德拉比夫() 	feud.Barony
    BIrkliiv伊尔克利伊夫() 	feud.Barony
    BPishchane皮夏涅() 	feud.Barony
    BVoin沃因() 	feud.Barony
    BZolotonosha佐洛托诺沙() 	feud.Barony
}

type 沃因VoinCounty struct {
	feud.BaseCounty
	博戈杜霍夫卡Bohodukhivka 	feud.Barony
	乔尔诺拜Chornobai 	feud.Barony
	德拉比夫Drabiv 	feud.Barony
	伊尔克利伊夫Irkliiv 	feud.Barony
	皮夏涅Pishchane 	feud.Barony
	沃因Voin 	feud.Barony
	佐洛托诺沙Zolotonosha 	feud.Barony
}

func (c *沃因VoinCounty) BBohodukhivka博戈杜霍夫卡() feud.Barony {
	return c.博戈杜霍夫卡Bohodukhivka
}
    
func (c *沃因VoinCounty) BChornobai乔尔诺拜() feud.Barony {
	return c.乔尔诺拜Chornobai
}
    
func (c *沃因VoinCounty) BDrabiv德拉比夫() feud.Barony {
	return c.德拉比夫Drabiv
}
    
func (c *沃因VoinCounty) BIrkliiv伊尔克利伊夫() feud.Barony {
	return c.伊尔克利伊夫Irkliiv
}
    
func (c *沃因VoinCounty) BPishchane皮夏涅() feud.Barony {
	return c.皮夏涅Pishchane
}
    
func (c *沃因VoinCounty) BVoin沃因() feud.Barony {
	return c.沃因Voin
}
    
func (c *沃因VoinCounty) BZolotonosha佐洛托诺沙() feud.Barony {
	return c.佐洛托诺沙Zolotonosha
}
    
var CVoin沃因 VoinCounty = &沃因VoinCounty{}

func init() {
	f := CVoin沃因.(*沃因VoinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1657",
		Title:     "voin",
		TitleName: "沃因",
		TitleCode: "c_voin",
		Baronies:  map[string]feud.Barony{},
	}

	f.博戈杜霍夫卡Bohodukhivka = BBohodukhivka博戈杜霍夫卡
	f.博戈杜霍夫卡Bohodukhivka.SetParent(f)
	
	f.乔尔诺拜Chornobai = BChornobai乔尔诺拜
	f.乔尔诺拜Chornobai.SetParent(f)
	
	f.德拉比夫Drabiv = BDrabiv德拉比夫
	f.德拉比夫Drabiv.SetParent(f)
	
	f.伊尔克利伊夫Irkliiv = BIrkliiv伊尔克利伊夫
	f.伊尔克利伊夫Irkliiv.SetParent(f)
	
	f.皮夏涅Pishchane = BPishchane皮夏涅
	f.皮夏涅Pishchane.SetParent(f)
	
	f.沃因Voin = BVoin沃因
	f.沃因Voin.SetParent(f)
	
	f.佐洛托诺沙Zolotonosha = BZolotonosha佐洛托诺沙
	f.佐洛托诺沙Zolotonosha.SetParent(f)
	
}
