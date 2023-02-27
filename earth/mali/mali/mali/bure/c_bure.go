package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BureCounty interface {
    feud.County
    BBoeganiola博埃加尼奥拉() 	feud.Barony
    BBure布尔() 	feud.Barony
    BGaladiou加拉迪乌() 	feud.Barony
    BKamina卡米纳() 	feud.Barony
    BKankan康康() 	feud.Barony
    BKouroussa库鲁萨() 	feud.Barony
    BSiguiri锡吉里() 	feud.Barony
}

type 布尔BureCounty struct {
	feud.BaseCounty
	博埃加尼奥拉Boeganiola 	feud.Barony
	布尔Bure 	feud.Barony
	加拉迪乌Galadiou 	feud.Barony
	卡米纳Kamina 	feud.Barony
	康康Kankan 	feud.Barony
	库鲁萨Kouroussa 	feud.Barony
	锡吉里Siguiri 	feud.Barony
}

func (c *布尔BureCounty) BBoeganiola博埃加尼奥拉() feud.Barony {
	return c.博埃加尼奥拉Boeganiola
}
    
func (c *布尔BureCounty) BBure布尔() feud.Barony {
	return c.布尔Bure
}
    
func (c *布尔BureCounty) BGaladiou加拉迪乌() feud.Barony {
	return c.加拉迪乌Galadiou
}
    
func (c *布尔BureCounty) BKamina卡米纳() feud.Barony {
	return c.卡米纳Kamina
}
    
func (c *布尔BureCounty) BKankan康康() feud.Barony {
	return c.康康Kankan
}
    
func (c *布尔BureCounty) BKouroussa库鲁萨() feud.Barony {
	return c.库鲁萨Kouroussa
}
    
func (c *布尔BureCounty) BSiguiri锡吉里() feud.Barony {
	return c.锡吉里Siguiri
}
    
var CBure布尔 BureCounty = &布尔BureCounty{}

func init() {
	f := CBure布尔.(*布尔BureCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1761",
		Title:     "bure",
		TitleName: "布尔",
		TitleCode: "c_bure",
		Baronies:  map[string]feud.Barony{},
	}

	f.博埃加尼奥拉Boeganiola = BBoeganiola博埃加尼奥拉
	f.博埃加尼奥拉Boeganiola.SetParent(f)
	
	f.布尔Bure = BBure布尔
	f.布尔Bure.SetParent(f)
	
	f.加拉迪乌Galadiou = BGaladiou加拉迪乌
	f.加拉迪乌Galadiou.SetParent(f)
	
	f.卡米纳Kamina = BKamina卡米纳
	f.卡米纳Kamina.SetParent(f)
	
	f.康康Kankan = BKankan康康
	f.康康Kankan.SetParent(f)
	
	f.库鲁萨Kouroussa = BKouroussa库鲁萨
	f.库鲁萨Kouroussa.SetParent(f)
	
	f.锡吉里Siguiri = BSiguiri锡吉里
	f.锡吉里Siguiri.SetParent(f)
	
}
