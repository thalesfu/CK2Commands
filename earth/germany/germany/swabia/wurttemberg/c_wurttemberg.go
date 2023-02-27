package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WurttembergCounty interface {
    feud.County
    BAsperg阿斯佩格() 	feud.Barony
    BEsslingen埃斯林根() 	feud.Barony
    BGmund格明德() 	feud.Barony
    BHeilbronn海尔布隆() 	feud.Barony
    BReutlingen罗伊特林根() 	feud.Barony
    BStaufen施陶芬() 	feud.Barony
    BStuttgart斯图加特() 	feud.Barony
    BWaiblingen魏布林根() 	feud.Barony
}

type 符滕堡WurttembergCounty struct {
	feud.BaseCounty
	阿斯佩格Asperg 	feud.Barony
	埃斯林根Esslingen 	feud.Barony
	格明德Gmund 	feud.Barony
	海尔布隆Heilbronn 	feud.Barony
	罗伊特林根Reutlingen 	feud.Barony
	施陶芬Staufen 	feud.Barony
	斯图加特Stuttgart 	feud.Barony
	魏布林根Waiblingen 	feud.Barony
}

func (c *符滕堡WurttembergCounty) BAsperg阿斯佩格() feud.Barony {
	return c.阿斯佩格Asperg
}
    
func (c *符滕堡WurttembergCounty) BEsslingen埃斯林根() feud.Barony {
	return c.埃斯林根Esslingen
}
    
func (c *符滕堡WurttembergCounty) BGmund格明德() feud.Barony {
	return c.格明德Gmund
}
    
func (c *符滕堡WurttembergCounty) BHeilbronn海尔布隆() feud.Barony {
	return c.海尔布隆Heilbronn
}
    
func (c *符滕堡WurttembergCounty) BReutlingen罗伊特林根() feud.Barony {
	return c.罗伊特林根Reutlingen
}
    
func (c *符滕堡WurttembergCounty) BStaufen施陶芬() feud.Barony {
	return c.施陶芬Staufen
}
    
func (c *符滕堡WurttembergCounty) BStuttgart斯图加特() feud.Barony {
	return c.斯图加特Stuttgart
}
    
func (c *符滕堡WurttembergCounty) BWaiblingen魏布林根() feud.Barony {
	return c.魏布林根Waiblingen
}
    
var CWurttemberg符滕堡 WurttembergCounty = &符滕堡WurttembergCounty{}

func init() {
	f := CWurttemberg符滕堡.(*符滕堡WurttembergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "253",
		Title:     "wurttemberg",
		TitleName: "符滕堡",
		TitleCode: "c_wurttemberg",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯佩格Asperg = BAsperg阿斯佩格
	f.阿斯佩格Asperg.SetParent(f)
	
	f.埃斯林根Esslingen = BEsslingen埃斯林根
	f.埃斯林根Esslingen.SetParent(f)
	
	f.格明德Gmund = BGmund格明德
	f.格明德Gmund.SetParent(f)
	
	f.海尔布隆Heilbronn = BHeilbronn海尔布隆
	f.海尔布隆Heilbronn.SetParent(f)
	
	f.罗伊特林根Reutlingen = BReutlingen罗伊特林根
	f.罗伊特林根Reutlingen.SetParent(f)
	
	f.施陶芬Staufen = BStaufen施陶芬
	f.施陶芬Staufen.SetParent(f)
	
	f.斯图加特Stuttgart = BStuttgart斯图加特
	f.斯图加特Stuttgart.SetParent(f)
	
	f.魏布林根Waiblingen = BWaiblingen魏布林根
	f.魏布林根Waiblingen.SetParent(f)
	
}
