package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FachiCounty interface {
    feud.County
    BFachi法希() 	feud.Barony
    BGeger盖格尔() 	feud.Barony
    BGoulo古洛() 	feud.Barony
    BKemkaga康卡加() 	feud.Barony
    BMbarou姆巴鲁() 	feud.Barony
    BMitena米泰纳() 	feud.Barony
    BTiadam蒂阿达姆() 	feud.Barony
}

type 法希FachiCounty struct {
	feud.BaseCounty
	法希Fachi 	feud.Barony
	盖格尔Geger 	feud.Barony
	古洛Goulo 	feud.Barony
	康卡加Kemkaga 	feud.Barony
	姆巴鲁Mbarou 	feud.Barony
	米泰纳Mitena 	feud.Barony
	蒂阿达姆Tiadam 	feud.Barony
}

func (c *法希FachiCounty) BFachi法希() feud.Barony {
	return c.法希Fachi
}
    
func (c *法希FachiCounty) BGeger盖格尔() feud.Barony {
	return c.盖格尔Geger
}
    
func (c *法希FachiCounty) BGoulo古洛() feud.Barony {
	return c.古洛Goulo
}
    
func (c *法希FachiCounty) BKemkaga康卡加() feud.Barony {
	return c.康卡加Kemkaga
}
    
func (c *法希FachiCounty) BMbarou姆巴鲁() feud.Barony {
	return c.姆巴鲁Mbarou
}
    
func (c *法希FachiCounty) BMitena米泰纳() feud.Barony {
	return c.米泰纳Mitena
}
    
func (c *法希FachiCounty) BTiadam蒂阿达姆() feud.Barony {
	return c.蒂阿达姆Tiadam
}
    
var CFachi法希 FachiCounty = &法希FachiCounty{}

func init() {
	f := CFachi法希.(*法希FachiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1737",
		Title:     "fachi",
		TitleName: "法希",
		TitleCode: "c_fachi",
		Baronies:  map[string]feud.Barony{},
	}

	f.法希Fachi = BFachi法希
	f.法希Fachi.SetParent(f)
	
	f.盖格尔Geger = BGeger盖格尔
	f.盖格尔Geger.SetParent(f)
	
	f.古洛Goulo = BGoulo古洛
	f.古洛Goulo.SetParent(f)
	
	f.康卡加Kemkaga = BKemkaga康卡加
	f.康卡加Kemkaga.SetParent(f)
	
	f.姆巴鲁Mbarou = BMbarou姆巴鲁
	f.姆巴鲁Mbarou.SetParent(f)
	
	f.米泰纳Mitena = BMitena米泰纳
	f.米泰纳Mitena.SetParent(f)
	
	f.蒂阿达姆Tiadam = BTiadam蒂阿达姆
	f.蒂阿达姆Tiadam.SetParent(f)
	
}
