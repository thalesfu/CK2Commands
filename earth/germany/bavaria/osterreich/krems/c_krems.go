package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KremsCounty interface {
    feud.County
    BEggenburg埃根堡() 	feud.Barony
    BGars加尔斯() 	feud.Barony
    BHorn_krems霍恩() 	feud.Barony
    BKrems克雷米斯() 	feud.Barony
    BKuenring屈恩灵() 	feud.Barony
    BRetz雷茨() 	feud.Barony
    BZwettl茨韦特尔() 	feud.Barony
}

type 克雷米斯KremsCounty struct {
	feud.BaseCounty
	埃根堡Eggenburg 	feud.Barony
	加尔斯Gars 	feud.Barony
	霍恩Horn_krems 	feud.Barony
	克雷米斯Krems 	feud.Barony
	屈恩灵Kuenring 	feud.Barony
	雷茨Retz 	feud.Barony
	茨韦特尔Zwettl 	feud.Barony
}

func (c *克雷米斯KremsCounty) BEggenburg埃根堡() feud.Barony {
	return c.埃根堡Eggenburg
}
    
func (c *克雷米斯KremsCounty) BGars加尔斯() feud.Barony {
	return c.加尔斯Gars
}
    
func (c *克雷米斯KremsCounty) BHorn_krems霍恩() feud.Barony {
	return c.霍恩Horn_krems
}
    
func (c *克雷米斯KremsCounty) BKrems克雷米斯() feud.Barony {
	return c.克雷米斯Krems
}
    
func (c *克雷米斯KremsCounty) BKuenring屈恩灵() feud.Barony {
	return c.屈恩灵Kuenring
}
    
func (c *克雷米斯KremsCounty) BRetz雷茨() feud.Barony {
	return c.雷茨Retz
}
    
func (c *克雷米斯KremsCounty) BZwettl茨韦特尔() feud.Barony {
	return c.茨韦特尔Zwettl
}
    
var CKrems克雷米斯 KremsCounty = &克雷米斯KremsCounty{}

func init() {
	f := CKrems克雷米斯.(*克雷米斯KremsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1697",
		Title:     "krems",
		TitleName: "克雷米斯",
		TitleCode: "c_krems",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃根堡Eggenburg = BEggenburg埃根堡
	f.埃根堡Eggenburg.SetParent(f)
	
	f.加尔斯Gars = BGars加尔斯
	f.加尔斯Gars.SetParent(f)
	
	f.霍恩Horn_krems = BHorn_krems霍恩
	f.霍恩Horn_krems.SetParent(f)
	
	f.克雷米斯Krems = BKrems克雷米斯
	f.克雷米斯Krems.SetParent(f)
	
	f.屈恩灵Kuenring = BKuenring屈恩灵
	f.屈恩灵Kuenring.SetParent(f)
	
	f.雷茨Retz = BRetz雷茨
	f.雷茨Retz.SetParent(f)
	
	f.茨韦特尔Zwettl = BZwettl茨韦特尔
	f.茨韦特尔Zwettl.SetParent(f)
	
}
