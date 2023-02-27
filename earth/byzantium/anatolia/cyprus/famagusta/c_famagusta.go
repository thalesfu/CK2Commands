package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FamagustaCounty interface {
    feud.County
    BBuffavento布法文托() 	feud.Barony
    BCithium基蒂翁() 	feud.Barony
    BFamagusta法马古斯塔() 	feud.Barony
    BKantara坎塔拉() 	feud.Barony
    BKyrenia凯里尼亚() 	feud.Barony
    BNikosia尼科西亚() 	feud.Barony
    BPeristerona派里斯泰罗纳() 	feud.Barony
    BSthilarion圣伊拉里() 	feud.Barony
}

type 法马古斯塔FamagustaCounty struct {
	feud.BaseCounty
	布法文托Buffavento 	feud.Barony
	基蒂翁Cithium 	feud.Barony
	法马古斯塔Famagusta 	feud.Barony
	坎塔拉Kantara 	feud.Barony
	凯里尼亚Kyrenia 	feud.Barony
	尼科西亚Nikosia 	feud.Barony
	派里斯泰罗纳Peristerona 	feud.Barony
	圣伊拉里Sthilarion 	feud.Barony
}

func (c *法马古斯塔FamagustaCounty) BBuffavento布法文托() feud.Barony {
	return c.布法文托Buffavento
}
    
func (c *法马古斯塔FamagustaCounty) BCithium基蒂翁() feud.Barony {
	return c.基蒂翁Cithium
}
    
func (c *法马古斯塔FamagustaCounty) BFamagusta法马古斯塔() feud.Barony {
	return c.法马古斯塔Famagusta
}
    
func (c *法马古斯塔FamagustaCounty) BKantara坎塔拉() feud.Barony {
	return c.坎塔拉Kantara
}
    
func (c *法马古斯塔FamagustaCounty) BKyrenia凯里尼亚() feud.Barony {
	return c.凯里尼亚Kyrenia
}
    
func (c *法马古斯塔FamagustaCounty) BNikosia尼科西亚() feud.Barony {
	return c.尼科西亚Nikosia
}
    
func (c *法马古斯塔FamagustaCounty) BPeristerona派里斯泰罗纳() feud.Barony {
	return c.派里斯泰罗纳Peristerona
}
    
func (c *法马古斯塔FamagustaCounty) BSthilarion圣伊拉里() feud.Barony {
	return c.圣伊拉里Sthilarion
}
    
var CFamagusta法马古斯塔 FamagustaCounty = &法马古斯塔FamagustaCounty{}

func init() {
	f := CFamagusta法马古斯塔.(*法马古斯塔FamagustaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "757",
		Title:     "famagusta",
		TitleName: "法马古斯塔",
		TitleCode: "c_famagusta",
		Baronies:  map[string]feud.Barony{},
	}

	f.布法文托Buffavento = BBuffavento布法文托
	f.布法文托Buffavento.SetParent(f)
	
	f.基蒂翁Cithium = BCithium基蒂翁
	f.基蒂翁Cithium.SetParent(f)
	
	f.法马古斯塔Famagusta = BFamagusta法马古斯塔
	f.法马古斯塔Famagusta.SetParent(f)
	
	f.坎塔拉Kantara = BKantara坎塔拉
	f.坎塔拉Kantara.SetParent(f)
	
	f.凯里尼亚Kyrenia = BKyrenia凯里尼亚
	f.凯里尼亚Kyrenia.SetParent(f)
	
	f.尼科西亚Nikosia = BNikosia尼科西亚
	f.尼科西亚Nikosia.SetParent(f)
	
	f.派里斯泰罗纳Peristerona = BPeristerona派里斯泰罗纳
	f.派里斯泰罗纳Peristerona.SetParent(f)
	
	f.圣伊拉里Sthilarion = BSthilarion圣伊拉里
	f.圣伊拉里Sthilarion.SetParent(f)
	
}
