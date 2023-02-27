package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type West_dvinaCounty interface {
    feud.County
    BBalvi巴尔维() 	feud.Barony
    BDaugavpils陶格夫匹尔斯() 	feud.Barony
    BErle埃尔勒() 	feud.Barony
    BKraslava克拉斯拉瓦() 	feud.Barony
    BKreuzburg克罗伊茨堡() 	feud.Barony
    BLudza卢扎() 	feud.Barony
    BRezekne雷泽克内() 	feud.Barony
    BVarakjani瓦拉克利亚尼() 	feud.Barony
    BVilaka维利亚卡() 	feud.Barony
}

type 西德维纳West_dvinaCounty struct {
	feud.BaseCounty
	巴尔维Balvi 	feud.Barony
	陶格夫匹尔斯Daugavpils 	feud.Barony
	埃尔勒Erle 	feud.Barony
	克拉斯拉瓦Kraslava 	feud.Barony
	克罗伊茨堡Kreuzburg 	feud.Barony
	卢扎Ludza 	feud.Barony
	雷泽克内Rezekne 	feud.Barony
	瓦拉克利亚尼Varakjani 	feud.Barony
	维利亚卡Vilaka 	feud.Barony
}

func (c *西德维纳West_dvinaCounty) BBalvi巴尔维() feud.Barony {
	return c.巴尔维Balvi
}
    
func (c *西德维纳West_dvinaCounty) BDaugavpils陶格夫匹尔斯() feud.Barony {
	return c.陶格夫匹尔斯Daugavpils
}
    
func (c *西德维纳West_dvinaCounty) BErle埃尔勒() feud.Barony {
	return c.埃尔勒Erle
}
    
func (c *西德维纳West_dvinaCounty) BKraslava克拉斯拉瓦() feud.Barony {
	return c.克拉斯拉瓦Kraslava
}
    
func (c *西德维纳West_dvinaCounty) BKreuzburg克罗伊茨堡() feud.Barony {
	return c.克罗伊茨堡Kreuzburg
}
    
func (c *西德维纳West_dvinaCounty) BLudza卢扎() feud.Barony {
	return c.卢扎Ludza
}
    
func (c *西德维纳West_dvinaCounty) BRezekne雷泽克内() feud.Barony {
	return c.雷泽克内Rezekne
}
    
func (c *西德维纳West_dvinaCounty) BVarakjani瓦拉克利亚尼() feud.Barony {
	return c.瓦拉克利亚尼Varakjani
}
    
func (c *西德维纳West_dvinaCounty) BVilaka维利亚卡() feud.Barony {
	return c.维利亚卡Vilaka
}
    
var CWest_dvina西德维纳 West_dvinaCounty = &西德维纳West_dvinaCounty{}

func init() {
	f := CWest_dvina西德维纳.(*西德维纳West_dvinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "416",
		Title:     "west_dvina",
		TitleName: "西德维纳",
		TitleCode: "c_west_dvina",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔维Balvi = BBalvi巴尔维
	f.巴尔维Balvi.SetParent(f)
	
	f.陶格夫匹尔斯Daugavpils = BDaugavpils陶格夫匹尔斯
	f.陶格夫匹尔斯Daugavpils.SetParent(f)
	
	f.埃尔勒Erle = BErle埃尔勒
	f.埃尔勒Erle.SetParent(f)
	
	f.克拉斯拉瓦Kraslava = BKraslava克拉斯拉瓦
	f.克拉斯拉瓦Kraslava.SetParent(f)
	
	f.克罗伊茨堡Kreuzburg = BKreuzburg克罗伊茨堡
	f.克罗伊茨堡Kreuzburg.SetParent(f)
	
	f.卢扎Ludza = BLudza卢扎
	f.卢扎Ludza.SetParent(f)
	
	f.雷泽克内Rezekne = BRezekne雷泽克内
	f.雷泽克内Rezekne.SetParent(f)
	
	f.瓦拉克利亚尼Varakjani = BVarakjani瓦拉克利亚尼
	f.瓦拉克利亚尼Varakjani.SetParent(f)
	
	f.维利亚卡Vilaka = BVilaka维利亚卡
	f.维利亚卡Vilaka.SetParent(f)
	
}
