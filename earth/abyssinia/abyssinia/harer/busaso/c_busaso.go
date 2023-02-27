package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BusasoCounty interface {
    feud.County
    BBandaraassim班达卡西姆() 	feud.Barony
    BBuraan班拉安() 	feud.Barony
    BErigavo埃里加博() 	feud.Barony
    BGaroowe加罗韦() 	feud.Barony
    BHadaftimo哈达费莫() 	feud.Barony
    BLasanod拉斯阿诺得() 	feud.Barony
    BMosylon莫塞隆() 	feud.Barony
    BQardho加尔多() 	feud.Barony
}

type 博萨索BusasoCounty struct {
	feud.BaseCounty
	班达卡西姆Bandaraassim 	feud.Barony
	班拉安Buraan 	feud.Barony
	埃里加博Erigavo 	feud.Barony
	加罗韦Garoowe 	feud.Barony
	哈达费莫Hadaftimo 	feud.Barony
	拉斯阿诺得Lasanod 	feud.Barony
	莫塞隆Mosylon 	feud.Barony
	加尔多Qardho 	feud.Barony
}

func (c *博萨索BusasoCounty) BBandaraassim班达卡西姆() feud.Barony {
	return c.班达卡西姆Bandaraassim
}
    
func (c *博萨索BusasoCounty) BBuraan班拉安() feud.Barony {
	return c.班拉安Buraan
}
    
func (c *博萨索BusasoCounty) BErigavo埃里加博() feud.Barony {
	return c.埃里加博Erigavo
}
    
func (c *博萨索BusasoCounty) BGaroowe加罗韦() feud.Barony {
	return c.加罗韦Garoowe
}
    
func (c *博萨索BusasoCounty) BHadaftimo哈达费莫() feud.Barony {
	return c.哈达费莫Hadaftimo
}
    
func (c *博萨索BusasoCounty) BLasanod拉斯阿诺得() feud.Barony {
	return c.拉斯阿诺得Lasanod
}
    
func (c *博萨索BusasoCounty) BMosylon莫塞隆() feud.Barony {
	return c.莫塞隆Mosylon
}
    
func (c *博萨索BusasoCounty) BQardho加尔多() feud.Barony {
	return c.加尔多Qardho
}
    
var CBusaso博萨索 BusasoCounty = &博萨索BusasoCounty{}

func init() {
	f := CBusaso博萨索.(*博萨索BusasoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "871",
		Title:     "busaso",
		TitleName: "博萨索",
		TitleCode: "c_busaso",
		Baronies:  map[string]feud.Barony{},
	}

	f.班达卡西姆Bandaraassim = BBandaraassim班达卡西姆
	f.班达卡西姆Bandaraassim.SetParent(f)
	
	f.班拉安Buraan = BBuraan班拉安
	f.班拉安Buraan.SetParent(f)
	
	f.埃里加博Erigavo = BErigavo埃里加博
	f.埃里加博Erigavo.SetParent(f)
	
	f.加罗韦Garoowe = BGaroowe加罗韦
	f.加罗韦Garoowe.SetParent(f)
	
	f.哈达费莫Hadaftimo = BHadaftimo哈达费莫
	f.哈达费莫Hadaftimo.SetParent(f)
	
	f.拉斯阿诺得Lasanod = BLasanod拉斯阿诺得
	f.拉斯阿诺得Lasanod.SetParent(f)
	
	f.莫塞隆Mosylon = BMosylon莫塞隆
	f.莫塞隆Mosylon.SetParent(f)
	
	f.加尔多Qardho = BQardho加尔多
	f.加尔多Qardho.SetParent(f)
	
}
