package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ScaloviaCounty interface {
    feud.County
    BJomsberg约姆斯贝格() 	feud.Barony
    BJurgaiten尤尔盖滕() 	feud.Barony
    BRagnit拉格尼特() 	feud.Barony
    BRuss鲁斯() 	feud.Barony
    BSkomanten斯科曼滕() 	feud.Barony
    BSplitter斯普利特() 	feud.Barony
    BStrewa斯特雷瓦() 	feud.Barony
    BTilgit蒂尔吉特() 	feud.Barony
}

type 斯卡洛维亚ScaloviaCounty struct {
	feud.BaseCounty
	约姆斯贝格Jomsberg 	feud.Barony
	尤尔盖滕Jurgaiten 	feud.Barony
	拉格尼特Ragnit 	feud.Barony
	鲁斯Russ 	feud.Barony
	斯科曼滕Skomanten 	feud.Barony
	斯普利特Splitter 	feud.Barony
	斯特雷瓦Strewa 	feud.Barony
	蒂尔吉特Tilgit 	feud.Barony
}

func (c *斯卡洛维亚ScaloviaCounty) BJomsberg约姆斯贝格() feud.Barony {
	return c.约姆斯贝格Jomsberg
}
    
func (c *斯卡洛维亚ScaloviaCounty) BJurgaiten尤尔盖滕() feud.Barony {
	return c.尤尔盖滕Jurgaiten
}
    
func (c *斯卡洛维亚ScaloviaCounty) BRagnit拉格尼特() feud.Barony {
	return c.拉格尼特Ragnit
}
    
func (c *斯卡洛维亚ScaloviaCounty) BRuss鲁斯() feud.Barony {
	return c.鲁斯Russ
}
    
func (c *斯卡洛维亚ScaloviaCounty) BSkomanten斯科曼滕() feud.Barony {
	return c.斯科曼滕Skomanten
}
    
func (c *斯卡洛维亚ScaloviaCounty) BSplitter斯普利特() feud.Barony {
	return c.斯普利特Splitter
}
    
func (c *斯卡洛维亚ScaloviaCounty) BStrewa斯特雷瓦() feud.Barony {
	return c.斯特雷瓦Strewa
}
    
func (c *斯卡洛维亚ScaloviaCounty) BTilgit蒂尔吉特() feud.Barony {
	return c.蒂尔吉特Tilgit
}
    
var CScalovia斯卡洛维亚 ScaloviaCounty = &斯卡洛维亚ScaloviaCounty{}

func init() {
	f := CScalovia斯卡洛维亚.(*斯卡洛维亚ScaloviaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "422",
		Title:     "scalovia",
		TitleName: "斯卡洛维亚",
		TitleCode: "c_scalovia",
		Baronies:  map[string]feud.Barony{},
	}

	f.约姆斯贝格Jomsberg = BJomsberg约姆斯贝格
	f.约姆斯贝格Jomsberg.SetParent(f)
	
	f.尤尔盖滕Jurgaiten = BJurgaiten尤尔盖滕
	f.尤尔盖滕Jurgaiten.SetParent(f)
	
	f.拉格尼特Ragnit = BRagnit拉格尼特
	f.拉格尼特Ragnit.SetParent(f)
	
	f.鲁斯Russ = BRuss鲁斯
	f.鲁斯Russ.SetParent(f)
	
	f.斯科曼滕Skomanten = BSkomanten斯科曼滕
	f.斯科曼滕Skomanten.SetParent(f)
	
	f.斯普利特Splitter = BSplitter斯普利特
	f.斯普利特Splitter.SetParent(f)
	
	f.斯特雷瓦Strewa = BStrewa斯特雷瓦
	f.斯特雷瓦Strewa.SetParent(f)
	
	f.蒂尔吉特Tilgit = BTilgit蒂尔吉特
	f.蒂尔吉特Tilgit.SetParent(f)
	
}
