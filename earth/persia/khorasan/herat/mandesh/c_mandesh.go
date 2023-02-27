package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MandeshCounty interface {
    feud.County
    BAdraskan阿德拉斯坎() 	feud.Barony
    BBaluci俾路支() 	feud.Barony
    BChaghcharan查赫查兰() 	feud.Barony
    BGozareh古扎雷() 	feud.Barony
    BKwajaha科瓦贾哈() 	feud.Barony
    BSaghar萨格哈尔() 	feud.Barony
    BTaywara塔瓦拉() 	feud.Barony
}

type 曼德什MandeshCounty struct {
	feud.BaseCounty
	阿德拉斯坎Adraskan 	feud.Barony
	俾路支Baluci 	feud.Barony
	查赫查兰Chaghcharan 	feud.Barony
	古扎雷Gozareh 	feud.Barony
	科瓦贾哈Kwajaha 	feud.Barony
	萨格哈尔Saghar 	feud.Barony
	塔瓦拉Taywara 	feud.Barony
}

func (c *曼德什MandeshCounty) BAdraskan阿德拉斯坎() feud.Barony {
	return c.阿德拉斯坎Adraskan
}
    
func (c *曼德什MandeshCounty) BBaluci俾路支() feud.Barony {
	return c.俾路支Baluci
}
    
func (c *曼德什MandeshCounty) BChaghcharan查赫查兰() feud.Barony {
	return c.查赫查兰Chaghcharan
}
    
func (c *曼德什MandeshCounty) BGozareh古扎雷() feud.Barony {
	return c.古扎雷Gozareh
}
    
func (c *曼德什MandeshCounty) BKwajaha科瓦贾哈() feud.Barony {
	return c.科瓦贾哈Kwajaha
}
    
func (c *曼德什MandeshCounty) BSaghar萨格哈尔() feud.Barony {
	return c.萨格哈尔Saghar
}
    
func (c *曼德什MandeshCounty) BTaywara塔瓦拉() feud.Barony {
	return c.塔瓦拉Taywara
}
    
var CMandesh曼德什 MandeshCounty = &曼德什MandeshCounty{}

func init() {
	f := CMandesh曼德什.(*曼德什MandeshCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "907",
		Title:     "mandesh",
		TitleName: "曼德什",
		TitleCode: "c_mandesh",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德拉斯坎Adraskan = BAdraskan阿德拉斯坎
	f.阿德拉斯坎Adraskan.SetParent(f)
	
	f.俾路支Baluci = BBaluci俾路支
	f.俾路支Baluci.SetParent(f)
	
	f.查赫查兰Chaghcharan = BChaghcharan查赫查兰
	f.查赫查兰Chaghcharan.SetParent(f)
	
	f.古扎雷Gozareh = BGozareh古扎雷
	f.古扎雷Gozareh.SetParent(f)
	
	f.科瓦贾哈Kwajaha = BKwajaha科瓦贾哈
	f.科瓦贾哈Kwajaha.SetParent(f)
	
	f.萨格哈尔Saghar = BSaghar萨格哈尔
	f.萨格哈尔Saghar.SetParent(f)
	
	f.塔瓦拉Taywara = BTaywara塔瓦拉
	f.塔瓦拉Taywara.SetParent(f)
	
}
