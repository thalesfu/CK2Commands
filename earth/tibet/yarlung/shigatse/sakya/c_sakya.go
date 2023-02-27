package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SakyaCounty interface {
    feud.County
    BDinggye定结() 	feud.Barony
    BGamba岗巴() 	feud.Barony
    BGeding吉定() 	feud.Barony
    BLungrong龙中() 	feud.Barony
    BPelsakya吉祥萨迦() 	feud.Barony
    BSakya萨迦() 	feud.Barony
    BZhentang陈塘() 	feud.Barony
}

type 萨迦SakyaCounty struct {
	feud.BaseCounty
	定结Dinggye 	feud.Barony
	岗巴Gamba 	feud.Barony
	吉定Geding 	feud.Barony
	龙中Lungrong 	feud.Barony
	吉祥萨迦Pelsakya 	feud.Barony
	萨迦Sakya 	feud.Barony
	陈塘Zhentang 	feud.Barony
}

func (c *萨迦SakyaCounty) BDinggye定结() feud.Barony {
	return c.定结Dinggye
}
    
func (c *萨迦SakyaCounty) BGamba岗巴() feud.Barony {
	return c.岗巴Gamba
}
    
func (c *萨迦SakyaCounty) BGeding吉定() feud.Barony {
	return c.吉定Geding
}
    
func (c *萨迦SakyaCounty) BLungrong龙中() feud.Barony {
	return c.龙中Lungrong
}
    
func (c *萨迦SakyaCounty) BPelsakya吉祥萨迦() feud.Barony {
	return c.吉祥萨迦Pelsakya
}
    
func (c *萨迦SakyaCounty) BSakya萨迦() feud.Barony {
	return c.萨迦Sakya
}
    
func (c *萨迦SakyaCounty) BZhentang陈塘() feud.Barony {
	return c.陈塘Zhentang
}
    
var CSakya萨迦 SakyaCounty = &萨迦SakyaCounty{}

func init() {
	f := CSakya萨迦.(*萨迦SakyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1492",
		Title:     "sakya",
		TitleName: "萨迦",
		TitleCode: "c_sakya",
		Baronies:  map[string]feud.Barony{},
	}

	f.定结Dinggye = BDinggye定结
	f.定结Dinggye.SetParent(f)
	
	f.岗巴Gamba = BGamba岗巴
	f.岗巴Gamba.SetParent(f)
	
	f.吉定Geding = BGeding吉定
	f.吉定Geding.SetParent(f)
	
	f.龙中Lungrong = BLungrong龙中
	f.龙中Lungrong.SetParent(f)
	
	f.吉祥萨迦Pelsakya = BPelsakya吉祥萨迦
	f.吉祥萨迦Pelsakya.SetParent(f)
	
	f.萨迦Sakya = BSakya萨迦
	f.萨迦Sakya.SetParent(f)
	
	f.陈塘Zhentang = BZhentang陈塘
	f.陈塘Zhentang.SetParent(f)
	
}
