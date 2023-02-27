package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KurdistanCounty interface {
    feud.County
    BAraden阿拉德() 	feud.Barony
    BBebadi贝巴迪() 	feud.Barony
    BDawodiya达沃蒂亚() 	feud.Barony
    BDehi代希() 	feud.Barony
    BDuhok杜胡克() 	feud.Barony
    BHarmashi赫尔马斯() 	feud.Barony
    BMarqayoma马尔喀马() 	feud.Barony
    BSarsink萨拉斯尼克() 	feud.Barony
}

type 库尔德斯坦KurdistanCounty struct {
	feud.BaseCounty
	阿拉德Araden 	feud.Barony
	贝巴迪Bebadi 	feud.Barony
	达沃蒂亚Dawodiya 	feud.Barony
	代希Dehi 	feud.Barony
	杜胡克Duhok 	feud.Barony
	赫尔马斯Harmashi 	feud.Barony
	马尔喀马Marqayoma 	feud.Barony
	萨拉斯尼克Sarsink 	feud.Barony
}

func (c *库尔德斯坦KurdistanCounty) BAraden阿拉德() feud.Barony {
	return c.阿拉德Araden
}
    
func (c *库尔德斯坦KurdistanCounty) BBebadi贝巴迪() feud.Barony {
	return c.贝巴迪Bebadi
}
    
func (c *库尔德斯坦KurdistanCounty) BDawodiya达沃蒂亚() feud.Barony {
	return c.达沃蒂亚Dawodiya
}
    
func (c *库尔德斯坦KurdistanCounty) BDehi代希() feud.Barony {
	return c.代希Dehi
}
    
func (c *库尔德斯坦KurdistanCounty) BDuhok杜胡克() feud.Barony {
	return c.杜胡克Duhok
}
    
func (c *库尔德斯坦KurdistanCounty) BHarmashi赫尔马斯() feud.Barony {
	return c.赫尔马斯Harmashi
}
    
func (c *库尔德斯坦KurdistanCounty) BMarqayoma马尔喀马() feud.Barony {
	return c.马尔喀马Marqayoma
}
    
func (c *库尔德斯坦KurdistanCounty) BSarsink萨拉斯尼克() feud.Barony {
	return c.萨拉斯尼克Sarsink
}
    
var CKurdistan库尔德斯坦 KurdistanCounty = &库尔德斯坦KurdistanCounty{}

func init() {
	f := CKurdistan库尔德斯坦.(*库尔德斯坦KurdistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "686",
		Title:     "kurdistan",
		TitleName: "库尔德斯坦",
		TitleCode: "c_kurdistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉德Araden = BAraden阿拉德
	f.阿拉德Araden.SetParent(f)
	
	f.贝巴迪Bebadi = BBebadi贝巴迪
	f.贝巴迪Bebadi.SetParent(f)
	
	f.达沃蒂亚Dawodiya = BDawodiya达沃蒂亚
	f.达沃蒂亚Dawodiya.SetParent(f)
	
	f.代希Dehi = BDehi代希
	f.代希Dehi.SetParent(f)
	
	f.杜胡克Duhok = BDuhok杜胡克
	f.杜胡克Duhok.SetParent(f)
	
	f.赫尔马斯Harmashi = BHarmashi赫尔马斯
	f.赫尔马斯Harmashi.SetParent(f)
	
	f.马尔喀马Marqayoma = BMarqayoma马尔喀马
	f.马尔喀马Marqayoma.SetParent(f)
	
	f.萨拉斯尼克Sarsink = BSarsink萨拉斯尼克
	f.萨拉斯尼克Sarsink.SetParent(f)
	
}
