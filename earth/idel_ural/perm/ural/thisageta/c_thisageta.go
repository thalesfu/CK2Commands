package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThisagetaCounty interface {
    feud.County
    BChudova丘多瓦() 	feud.Barony
    BPashiya帕希亚() 	feud.Barony
    BPromysla普罗梅斯拉() 	feud.Barony
    BSarany萨拉内() 	feud.Barony
    BSkal_nyy斯卡利内() 	feud.Barony
    BThisageta杜撒该塔伊() 	feud.Barony
    BVizhay维扎伊() 	feud.Barony
}

type 杜撒该塔伊ThisagetaCounty struct {
	feud.BaseCounty
	丘多瓦Chudova 	feud.Barony
	帕希亚Pashiya 	feud.Barony
	普罗梅斯拉Promysla 	feud.Barony
	萨拉内Sarany 	feud.Barony
	斯卡利内Skal_nyy 	feud.Barony
	杜撒该塔伊Thisageta 	feud.Barony
	维扎伊Vizhay 	feud.Barony
}

func (c *杜撒该塔伊ThisagetaCounty) BChudova丘多瓦() feud.Barony {
	return c.丘多瓦Chudova
}
    
func (c *杜撒该塔伊ThisagetaCounty) BPashiya帕希亚() feud.Barony {
	return c.帕希亚Pashiya
}
    
func (c *杜撒该塔伊ThisagetaCounty) BPromysla普罗梅斯拉() feud.Barony {
	return c.普罗梅斯拉Promysla
}
    
func (c *杜撒该塔伊ThisagetaCounty) BSarany萨拉内() feud.Barony {
	return c.萨拉内Sarany
}
    
func (c *杜撒该塔伊ThisagetaCounty) BSkal_nyy斯卡利内() feud.Barony {
	return c.斯卡利内Skal_nyy
}
    
func (c *杜撒该塔伊ThisagetaCounty) BThisageta杜撒该塔伊() feud.Barony {
	return c.杜撒该塔伊Thisageta
}
    
func (c *杜撒该塔伊ThisagetaCounty) BVizhay维扎伊() feud.Barony {
	return c.维扎伊Vizhay
}
    
var CThisageta杜撒该塔伊 ThisagetaCounty = &杜撒该塔伊ThisagetaCounty{}

func init() {
	f := CThisageta杜撒该塔伊.(*杜撒该塔伊ThisagetaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1843",
		Title:     "thisageta",
		TitleName: "杜撒该塔伊",
		TitleCode: "c_thisageta",
		Baronies:  map[string]feud.Barony{},
	}

	f.丘多瓦Chudova = BChudova丘多瓦
	f.丘多瓦Chudova.SetParent(f)
	
	f.帕希亚Pashiya = BPashiya帕希亚
	f.帕希亚Pashiya.SetParent(f)
	
	f.普罗梅斯拉Promysla = BPromysla普罗梅斯拉
	f.普罗梅斯拉Promysla.SetParent(f)
	
	f.萨拉内Sarany = BSarany萨拉内
	f.萨拉内Sarany.SetParent(f)
	
	f.斯卡利内Skal_nyy = BSkal_nyy斯卡利内
	f.斯卡利内Skal_nyy.SetParent(f)
	
	f.杜撒该塔伊Thisageta = BThisageta杜撒该塔伊
	f.杜撒该塔伊Thisageta.SetParent(f)
	
	f.维扎伊Vizhay = BVizhay维扎伊
	f.维扎伊Vizhay.SetParent(f)
	
}
