package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BerberaCounty interface {
    feud.County
    BBerbera拨拔力() 	feud.Barony
    BBurco布尔奥() 	feud.Barony
    BDaarbuduq达尔布都奇() 	feud.Barony
    BGabiley盖比莱() 	feud.Barony
    BMandera曼德拉() 	feud.Barony
    BMaydh迈德() 	feud.Barony
    BShiikh谢赫() 	feud.Barony
}

type 拨拔力BerberaCounty struct {
	feud.BaseCounty
	拨拔力Berbera 	feud.Barony
	布尔奥Burco 	feud.Barony
	达尔布都奇Daarbuduq 	feud.Barony
	盖比莱Gabiley 	feud.Barony
	曼德拉Mandera 	feud.Barony
	迈德Maydh 	feud.Barony
	谢赫Shiikh 	feud.Barony
}

func (c *拨拔力BerberaCounty) BBerbera拨拔力() feud.Barony {
	return c.拨拔力Berbera
}
    
func (c *拨拔力BerberaCounty) BBurco布尔奥() feud.Barony {
	return c.布尔奥Burco
}
    
func (c *拨拔力BerberaCounty) BDaarbuduq达尔布都奇() feud.Barony {
	return c.达尔布都奇Daarbuduq
}
    
func (c *拨拔力BerberaCounty) BGabiley盖比莱() feud.Barony {
	return c.盖比莱Gabiley
}
    
func (c *拨拔力BerberaCounty) BMandera曼德拉() feud.Barony {
	return c.曼德拉Mandera
}
    
func (c *拨拔力BerberaCounty) BMaydh迈德() feud.Barony {
	return c.迈德Maydh
}
    
func (c *拨拔力BerberaCounty) BShiikh谢赫() feud.Barony {
	return c.谢赫Shiikh
}
    
var CBerbera拨拔力 BerberaCounty = &拨拔力BerberaCounty{}

func init() {
	f := CBerbera拨拔力.(*拨拔力BerberaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "872",
		Title:     "berbera",
		TitleName: "拨拔力",
		TitleCode: "c_berbera",
		Baronies:  map[string]feud.Barony{},
	}

	f.拨拔力Berbera = BBerbera拨拔力
	f.拨拔力Berbera.SetParent(f)
	
	f.布尔奥Burco = BBurco布尔奥
	f.布尔奥Burco.SetParent(f)
	
	f.达尔布都奇Daarbuduq = BDaarbuduq达尔布都奇
	f.达尔布都奇Daarbuduq.SetParent(f)
	
	f.盖比莱Gabiley = BGabiley盖比莱
	f.盖比莱Gabiley.SetParent(f)
	
	f.曼德拉Mandera = BMandera曼德拉
	f.曼德拉Mandera.SetParent(f)
	
	f.迈德Maydh = BMaydh迈德
	f.迈德Maydh.SetParent(f)
	
	f.谢赫Shiikh = BShiikh谢赫
	f.谢赫Shiikh.SetParent(f)
	
}
