package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MakariCounty interface {
    feud.County
    BChilaro奇拉罗() 	feud.Barony
    BDiama迪亚马() 	feud.Barony
    BGambila甘比拉() 	feud.Barony
    BGegu盖古() 	feud.Barony
    BGidan_hunu吉丹胡努() 	feud.Barony
    BKussiri库西里() 	feud.Barony
    BMakari马卡里() 	feud.Barony
}

type 马卡里MakariCounty struct {
	feud.BaseCounty
	奇拉罗Chilaro 	feud.Barony
	迪亚马Diama 	feud.Barony
	甘比拉Gambila 	feud.Barony
	盖古Gegu 	feud.Barony
	吉丹胡努Gidan_hunu 	feud.Barony
	库西里Kussiri 	feud.Barony
	马卡里Makari 	feud.Barony
}

func (c *马卡里MakariCounty) BChilaro奇拉罗() feud.Barony {
	return c.奇拉罗Chilaro
}
    
func (c *马卡里MakariCounty) BDiama迪亚马() feud.Barony {
	return c.迪亚马Diama
}
    
func (c *马卡里MakariCounty) BGambila甘比拉() feud.Barony {
	return c.甘比拉Gambila
}
    
func (c *马卡里MakariCounty) BGegu盖古() feud.Barony {
	return c.盖古Gegu
}
    
func (c *马卡里MakariCounty) BGidan_hunu吉丹胡努() feud.Barony {
	return c.吉丹胡努Gidan_hunu
}
    
func (c *马卡里MakariCounty) BKussiri库西里() feud.Barony {
	return c.库西里Kussiri
}
    
func (c *马卡里MakariCounty) BMakari马卡里() feud.Barony {
	return c.马卡里Makari
}
    
var CMakari马卡里 MakariCounty = &马卡里MakariCounty{}

func init() {
	f := CMakari马卡里.(*马卡里MakariCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1741",
		Title:     "makari",
		TitleName: "马卡里",
		TitleCode: "c_makari",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇拉罗Chilaro = BChilaro奇拉罗
	f.奇拉罗Chilaro.SetParent(f)
	
	f.迪亚马Diama = BDiama迪亚马
	f.迪亚马Diama.SetParent(f)
	
	f.甘比拉Gambila = BGambila甘比拉
	f.甘比拉Gambila.SetParent(f)
	
	f.盖古Gegu = BGegu盖古
	f.盖古Gegu.SetParent(f)
	
	f.吉丹胡努Gidan_hunu = BGidan_hunu吉丹胡努
	f.吉丹胡努Gidan_hunu.SetParent(f)
	
	f.库西里Kussiri = BKussiri库西里
	f.库西里Kussiri.SetParent(f)
	
	f.马卡里Makari = BMakari马卡里
	f.马卡里Makari.SetParent(f)
	
}
