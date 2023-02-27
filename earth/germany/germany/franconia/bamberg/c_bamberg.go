package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BambergCounty interface {
    feud.County
    BAnsbach安斯巴赫() 	feud.Barony
    BBabenberg巴本堡() 	feud.Barony
    BBamberg班贝格() 	feud.Barony
    BCadolzburg卡多尔茨堡() 	feud.Barony
    BColmberg科尔姆贝格() 	feud.Barony
    BCrailsheim克赖尔斯海姆() 	feud.Barony
    BRoth罗特() 	feud.Barony
}

type 班贝格BambergCounty struct {
	feud.BaseCounty
	安斯巴赫Ansbach 	feud.Barony
	巴本堡Babenberg 	feud.Barony
	班贝格Bamberg 	feud.Barony
	卡多尔茨堡Cadolzburg 	feud.Barony
	科尔姆贝格Colmberg 	feud.Barony
	克赖尔斯海姆Crailsheim 	feud.Barony
	罗特Roth 	feud.Barony
}

func (c *班贝格BambergCounty) BAnsbach安斯巴赫() feud.Barony {
	return c.安斯巴赫Ansbach
}
    
func (c *班贝格BambergCounty) BBabenberg巴本堡() feud.Barony {
	return c.巴本堡Babenberg
}
    
func (c *班贝格BambergCounty) BBamberg班贝格() feud.Barony {
	return c.班贝格Bamberg
}
    
func (c *班贝格BambergCounty) BCadolzburg卡多尔茨堡() feud.Barony {
	return c.卡多尔茨堡Cadolzburg
}
    
func (c *班贝格BambergCounty) BColmberg科尔姆贝格() feud.Barony {
	return c.科尔姆贝格Colmberg
}
    
func (c *班贝格BambergCounty) BCrailsheim克赖尔斯海姆() feud.Barony {
	return c.克赖尔斯海姆Crailsheim
}
    
func (c *班贝格BambergCounty) BRoth罗特() feud.Barony {
	return c.罗特Roth
}
    
var CBamberg班贝格 BambergCounty = &班贝格BambergCounty{}

func init() {
	f := CBamberg班贝格.(*班贝格BambergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "313",
		Title:     "bamberg",
		TitleName: "班贝格",
		TitleCode: "c_bamberg",
		Baronies:  map[string]feud.Barony{},
	}

	f.安斯巴赫Ansbach = BAnsbach安斯巴赫
	f.安斯巴赫Ansbach.SetParent(f)
	
	f.巴本堡Babenberg = BBabenberg巴本堡
	f.巴本堡Babenberg.SetParent(f)
	
	f.班贝格Bamberg = BBamberg班贝格
	f.班贝格Bamberg.SetParent(f)
	
	f.卡多尔茨堡Cadolzburg = BCadolzburg卡多尔茨堡
	f.卡多尔茨堡Cadolzburg.SetParent(f)
	
	f.科尔姆贝格Colmberg = BColmberg科尔姆贝格
	f.科尔姆贝格Colmberg.SetParent(f)
	
	f.克赖尔斯海姆Crailsheim = BCrailsheim克赖尔斯海姆
	f.克赖尔斯海姆Crailsheim.SetParent(f)
	
	f.罗特Roth = BRoth罗特
	f.罗特Roth.SetParent(f)
	
}
