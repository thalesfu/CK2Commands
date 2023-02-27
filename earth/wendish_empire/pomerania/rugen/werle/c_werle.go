package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WerleCounty interface {
    feud.County
    BDemmin代明() 	feud.Barony
    BFriedland弗里德兰() 	feud.Barony
    BGreifswald格赖夫斯瓦尔德() 	feud.Barony
    BStralsund施特拉尔松德() 	feud.Barony
    BTemplin滕普林() 	feud.Barony
    BTreptow特雷普托() 	feud.Barony
    BTribsees特里布塞斯() 	feud.Barony
    BWaren瓦伦() 	feud.Barony
}

type 韦尔勒WerleCounty struct {
	feud.BaseCounty
	代明Demmin 	feud.Barony
	弗里德兰Friedland 	feud.Barony
	格赖夫斯瓦尔德Greifswald 	feud.Barony
	施特拉尔松德Stralsund 	feud.Barony
	滕普林Templin 	feud.Barony
	特雷普托Treptow 	feud.Barony
	特里布塞斯Tribsees 	feud.Barony
	瓦伦Waren 	feud.Barony
}

func (c *韦尔勒WerleCounty) BDemmin代明() feud.Barony {
	return c.代明Demmin
}
    
func (c *韦尔勒WerleCounty) BFriedland弗里德兰() feud.Barony {
	return c.弗里德兰Friedland
}
    
func (c *韦尔勒WerleCounty) BGreifswald格赖夫斯瓦尔德() feud.Barony {
	return c.格赖夫斯瓦尔德Greifswald
}
    
func (c *韦尔勒WerleCounty) BStralsund施特拉尔松德() feud.Barony {
	return c.施特拉尔松德Stralsund
}
    
func (c *韦尔勒WerleCounty) BTemplin滕普林() feud.Barony {
	return c.滕普林Templin
}
    
func (c *韦尔勒WerleCounty) BTreptow特雷普托() feud.Barony {
	return c.特雷普托Treptow
}
    
func (c *韦尔勒WerleCounty) BTribsees特里布塞斯() feud.Barony {
	return c.特里布塞斯Tribsees
}
    
func (c *韦尔勒WerleCounty) BWaren瓦伦() feud.Barony {
	return c.瓦伦Waren
}
    
var CWerle韦尔勒 WerleCounty = &韦尔勒WerleCounty{}

func init() {
	f := CWerle韦尔勒.(*韦尔勒WerleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "307",
		Title:     "werle",
		TitleName: "韦尔勒",
		TitleCode: "c_werle",
		Baronies:  map[string]feud.Barony{},
	}

	f.代明Demmin = BDemmin代明
	f.代明Demmin.SetParent(f)
	
	f.弗里德兰Friedland = BFriedland弗里德兰
	f.弗里德兰Friedland.SetParent(f)
	
	f.格赖夫斯瓦尔德Greifswald = BGreifswald格赖夫斯瓦尔德
	f.格赖夫斯瓦尔德Greifswald.SetParent(f)
	
	f.施特拉尔松德Stralsund = BStralsund施特拉尔松德
	f.施特拉尔松德Stralsund.SetParent(f)
	
	f.滕普林Templin = BTemplin滕普林
	f.滕普林Templin.SetParent(f)
	
	f.特雷普托Treptow = BTreptow特雷普托
	f.特雷普托Treptow.SetParent(f)
	
	f.特里布塞斯Tribsees = BTribsees特里布塞斯
	f.特里布塞斯Tribsees.SetParent(f)
	
	f.瓦伦Waren = BWaren瓦伦
	f.瓦伦Waren.SetParent(f)
	
}
