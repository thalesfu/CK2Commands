package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NarkeCounty interface {
    feud.County
    BAskersund阿斯克松德() 	feud.Barony
    BGoksholm约克斯霍尔姆() 	feud.Barony
    BHallsberg哈尔斯贝里() 	feud.Barony
    BKumla库姆拉() 	feud.Barony
    BMosas穆索斯() 	feud.Barony
    BNora努拉() 	feud.Barony
    BOrebro厄勒布鲁() 	feud.Barony
    BRiseberga里瑟贝里亚() 	feud.Barony
}

type 内尔克NarkeCounty struct {
	feud.BaseCounty
	阿斯克松德Askersund 	feud.Barony
	约克斯霍尔姆Goksholm 	feud.Barony
	哈尔斯贝里Hallsberg 	feud.Barony
	库姆拉Kumla 	feud.Barony
	穆索斯Mosas 	feud.Barony
	努拉Nora 	feud.Barony
	厄勒布鲁Orebro 	feud.Barony
	里瑟贝里亚Riseberga 	feud.Barony
}

func (c *内尔克NarkeCounty) BAskersund阿斯克松德() feud.Barony {
	return c.阿斯克松德Askersund
}
    
func (c *内尔克NarkeCounty) BGoksholm约克斯霍尔姆() feud.Barony {
	return c.约克斯霍尔姆Goksholm
}
    
func (c *内尔克NarkeCounty) BHallsberg哈尔斯贝里() feud.Barony {
	return c.哈尔斯贝里Hallsberg
}
    
func (c *内尔克NarkeCounty) BKumla库姆拉() feud.Barony {
	return c.库姆拉Kumla
}
    
func (c *内尔克NarkeCounty) BMosas穆索斯() feud.Barony {
	return c.穆索斯Mosas
}
    
func (c *内尔克NarkeCounty) BNora努拉() feud.Barony {
	return c.努拉Nora
}
    
func (c *内尔克NarkeCounty) BOrebro厄勒布鲁() feud.Barony {
	return c.厄勒布鲁Orebro
}
    
func (c *内尔克NarkeCounty) BRiseberga里瑟贝里亚() feud.Barony {
	return c.里瑟贝里亚Riseberga
}
    
var CNarke内尔克 NarkeCounty = &内尔克NarkeCounty{}

func init() {
	f := CNarke内尔克.(*内尔克NarkeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "294",
		Title:     "narke",
		TitleName: "内尔克",
		TitleCode: "c_narke",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯克松德Askersund = BAskersund阿斯克松德
	f.阿斯克松德Askersund.SetParent(f)
	
	f.约克斯霍尔姆Goksholm = BGoksholm约克斯霍尔姆
	f.约克斯霍尔姆Goksholm.SetParent(f)
	
	f.哈尔斯贝里Hallsberg = BHallsberg哈尔斯贝里
	f.哈尔斯贝里Hallsberg.SetParent(f)
	
	f.库姆拉Kumla = BKumla库姆拉
	f.库姆拉Kumla.SetParent(f)
	
	f.穆索斯Mosas = BMosas穆索斯
	f.穆索斯Mosas.SetParent(f)
	
	f.努拉Nora = BNora努拉
	f.努拉Nora.SetParent(f)
	
	f.厄勒布鲁Orebro = BOrebro厄勒布鲁
	f.厄勒布鲁Orebro.SetParent(f)
	
	f.里瑟贝里亚Riseberga = BRiseberga里瑟贝里亚
	f.里瑟贝里亚Riseberga.SetParent(f)
	
}
