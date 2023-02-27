package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BejaijaCounty interface {
    feud.County
    BAkaoudj阿克奥吉() 	feud.Barony
    BBejaija贝贾亚() 	feud.Barony
    BBoudaoud布达乌德() 	feud.Barony
    BElbekara贝卡拉() 	feud.Barony
    BEldjabia德贾比亚() 	feud.Barony
    BGantas甘塔斯() 	feud.Barony
    BJijel吉杰尔() 	feud.Barony
    BMila密拉() 	feud.Barony
}

type 贝贾亚BejaijaCounty struct {
	feud.BaseCounty
	阿克奥吉Akaoudj 	feud.Barony
	贝贾亚Bejaija 	feud.Barony
	布达乌德Boudaoud 	feud.Barony
	贝卡拉Elbekara 	feud.Barony
	德贾比亚Eldjabia 	feud.Barony
	甘塔斯Gantas 	feud.Barony
	吉杰尔Jijel 	feud.Barony
	密拉Mila 	feud.Barony
}

func (c *贝贾亚BejaijaCounty) BAkaoudj阿克奥吉() feud.Barony {
	return c.阿克奥吉Akaoudj
}
    
func (c *贝贾亚BejaijaCounty) BBejaija贝贾亚() feud.Barony {
	return c.贝贾亚Bejaija
}
    
func (c *贝贾亚BejaijaCounty) BBoudaoud布达乌德() feud.Barony {
	return c.布达乌德Boudaoud
}
    
func (c *贝贾亚BejaijaCounty) BElbekara贝卡拉() feud.Barony {
	return c.贝卡拉Elbekara
}
    
func (c *贝贾亚BejaijaCounty) BEldjabia德贾比亚() feud.Barony {
	return c.德贾比亚Eldjabia
}
    
func (c *贝贾亚BejaijaCounty) BGantas甘塔斯() feud.Barony {
	return c.甘塔斯Gantas
}
    
func (c *贝贾亚BejaijaCounty) BJijel吉杰尔() feud.Barony {
	return c.吉杰尔Jijel
}
    
func (c *贝贾亚BejaijaCounty) BMila密拉() feud.Barony {
	return c.密拉Mila
}
    
var CBejaija贝贾亚 BejaijaCounty = &贝贾亚BejaijaCounty{}

func init() {
	f := CBejaija贝贾亚.(*贝贾亚BejaijaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "822",
		Title:     "bejaija",
		TitleName: "贝贾亚",
		TitleCode: "c_bejaija",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克奥吉Akaoudj = BAkaoudj阿克奥吉
	f.阿克奥吉Akaoudj.SetParent(f)
	
	f.贝贾亚Bejaija = BBejaija贝贾亚
	f.贝贾亚Bejaija.SetParent(f)
	
	f.布达乌德Boudaoud = BBoudaoud布达乌德
	f.布达乌德Boudaoud.SetParent(f)
	
	f.贝卡拉Elbekara = BElbekara贝卡拉
	f.贝卡拉Elbekara.SetParent(f)
	
	f.德贾比亚Eldjabia = BEldjabia德贾比亚
	f.德贾比亚Eldjabia.SetParent(f)
	
	f.甘塔斯Gantas = BGantas甘塔斯
	f.甘塔斯Gantas.SetParent(f)
	
	f.吉杰尔Jijel = BJijel吉杰尔
	f.吉杰尔Jijel.SetParent(f)
	
	f.密拉Mila = BMila密拉
	f.密拉Mila.SetParent(f)
	
}
