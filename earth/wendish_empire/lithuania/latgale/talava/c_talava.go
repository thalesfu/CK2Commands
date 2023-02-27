package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TalavaCounty interface {
    feud.County
    BAluksne阿卢克斯内() 	feud.Barony
    BGulbene古尔贝内() 	feud.Barony
    BImera伊梅拉() 	feud.Barony
    BMetsene梅特塞内() 	feud.Barony
    BSatekle萨特克莱() 	feud.Barony
    BTrikata特里卡塔() 	feud.Barony
    BVijciems维伊齐埃姆斯() 	feud.Barony
}

type 塔拉瓦TalavaCounty struct {
	feud.BaseCounty
	阿卢克斯内Aluksne 	feud.Barony
	古尔贝内Gulbene 	feud.Barony
	伊梅拉Imera 	feud.Barony
	梅特塞内Metsene 	feud.Barony
	萨特克莱Satekle 	feud.Barony
	特里卡塔Trikata 	feud.Barony
	维伊齐埃姆斯Vijciems 	feud.Barony
}

func (c *塔拉瓦TalavaCounty) BAluksne阿卢克斯内() feud.Barony {
	return c.阿卢克斯内Aluksne
}
    
func (c *塔拉瓦TalavaCounty) BGulbene古尔贝内() feud.Barony {
	return c.古尔贝内Gulbene
}
    
func (c *塔拉瓦TalavaCounty) BImera伊梅拉() feud.Barony {
	return c.伊梅拉Imera
}
    
func (c *塔拉瓦TalavaCounty) BMetsene梅特塞内() feud.Barony {
	return c.梅特塞内Metsene
}
    
func (c *塔拉瓦TalavaCounty) BSatekle萨特克莱() feud.Barony {
	return c.萨特克莱Satekle
}
    
func (c *塔拉瓦TalavaCounty) BTrikata特里卡塔() feud.Barony {
	return c.特里卡塔Trikata
}
    
func (c *塔拉瓦TalavaCounty) BVijciems维伊齐埃姆斯() feud.Barony {
	return c.维伊齐埃姆斯Vijciems
}
    
var CTalava塔拉瓦 TalavaCounty = &塔拉瓦TalavaCounty{}

func init() {
	f := CTalava塔拉瓦.(*塔拉瓦TalavaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1595",
		Title:     "talava",
		TitleName: "塔拉瓦",
		TitleCode: "c_talava",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卢克斯内Aluksne = BAluksne阿卢克斯内
	f.阿卢克斯内Aluksne.SetParent(f)
	
	f.古尔贝内Gulbene = BGulbene古尔贝内
	f.古尔贝内Gulbene.SetParent(f)
	
	f.伊梅拉Imera = BImera伊梅拉
	f.伊梅拉Imera.SetParent(f)
	
	f.梅特塞内Metsene = BMetsene梅特塞内
	f.梅特塞内Metsene.SetParent(f)
	
	f.萨特克莱Satekle = BSatekle萨特克莱
	f.萨特克莱Satekle.SetParent(f)
	
	f.特里卡塔Trikata = BTrikata特里卡塔
	f.特里卡塔Trikata.SetParent(f)
	
	f.维伊齐埃姆斯Vijciems = BVijciems维伊齐埃姆斯
	f.维伊齐埃姆斯Vijciems.SetParent(f)
	
}
