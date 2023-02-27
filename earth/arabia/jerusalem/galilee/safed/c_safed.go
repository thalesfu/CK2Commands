package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SafedCounty interface {
    feud.County
    BBanyas巴尼亚斯() 	feud.Barony
    BBelvoir贝尔沃() 	feud.Barony
    BChastelet沙特莱() 	feud.Barony
    BQatsrin卡茨林() 	feud.Barony
    BSafed采法特() 	feud.Barony
    BSubeiba宁录堡() 	feud.Barony
    BToron多隆() 	feud.Barony
}

type 采法特SafedCounty struct {
	feud.BaseCounty
	巴尼亚斯Banyas 	feud.Barony
	贝尔沃Belvoir 	feud.Barony
	沙特莱Chastelet 	feud.Barony
	卡茨林Qatsrin 	feud.Barony
	采法特Safed 	feud.Barony
	宁录堡Subeiba 	feud.Barony
	多隆Toron 	feud.Barony
}

func (c *采法特SafedCounty) BBanyas巴尼亚斯() feud.Barony {
	return c.巴尼亚斯Banyas
}
    
func (c *采法特SafedCounty) BBelvoir贝尔沃() feud.Barony {
	return c.贝尔沃Belvoir
}
    
func (c *采法特SafedCounty) BChastelet沙特莱() feud.Barony {
	return c.沙特莱Chastelet
}
    
func (c *采法特SafedCounty) BQatsrin卡茨林() feud.Barony {
	return c.卡茨林Qatsrin
}
    
func (c *采法特SafedCounty) BSafed采法特() feud.Barony {
	return c.采法特Safed
}
    
func (c *采法特SafedCounty) BSubeiba宁录堡() feud.Barony {
	return c.宁录堡Subeiba
}
    
func (c *采法特SafedCounty) BToron多隆() feud.Barony {
	return c.多隆Toron
}
    
var CSafed采法特 SafedCounty = &采法特SafedCounty{}

func init() {
	f := CSafed采法特.(*采法特SafedCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "769",
		Title:     "safed",
		TitleName: "采法特",
		TitleCode: "c_safed",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尼亚斯Banyas = BBanyas巴尼亚斯
	f.巴尼亚斯Banyas.SetParent(f)
	
	f.贝尔沃Belvoir = BBelvoir贝尔沃
	f.贝尔沃Belvoir.SetParent(f)
	
	f.沙特莱Chastelet = BChastelet沙特莱
	f.沙特莱Chastelet.SetParent(f)
	
	f.卡茨林Qatsrin = BQatsrin卡茨林
	f.卡茨林Qatsrin.SetParent(f)
	
	f.采法特Safed = BSafed采法特
	f.采法特Safed.SetParent(f)
	
	f.宁录堡Subeiba = BSubeiba宁录堡
	f.宁录堡Subeiba.SetParent(f)
	
	f.多隆Toron = BToron多隆
	f.多隆Toron.SetParent(f)
	
}
