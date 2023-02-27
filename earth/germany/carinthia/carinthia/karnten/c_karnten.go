package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarntenCounty interface {
    feud.County
    BGreifenburg格赖芬堡() 	feud.Barony
    BKlagenfurt克拉根福特() 	feud.Barony
    BLienz利恩茨() 	feud.Barony
    BSpittl斯皮特尔() 	feud.Barony
    BTreffen特雷芬() 	feud.Barony
    BVillach菲拉赫() 	feud.Barony
    BVorelach费拉赫() 	feud.Barony
}

type 菲拉赫KarntenCounty struct {
	feud.BaseCounty
	格赖芬堡Greifenburg 	feud.Barony
	克拉根福特Klagenfurt 	feud.Barony
	利恩茨Lienz 	feud.Barony
	斯皮特尔Spittl 	feud.Barony
	特雷芬Treffen 	feud.Barony
	菲拉赫Villach 	feud.Barony
	费拉赫Vorelach 	feud.Barony
}

func (c *菲拉赫KarntenCounty) BGreifenburg格赖芬堡() feud.Barony {
	return c.格赖芬堡Greifenburg
}
    
func (c *菲拉赫KarntenCounty) BKlagenfurt克拉根福特() feud.Barony {
	return c.克拉根福特Klagenfurt
}
    
func (c *菲拉赫KarntenCounty) BLienz利恩茨() feud.Barony {
	return c.利恩茨Lienz
}
    
func (c *菲拉赫KarntenCounty) BSpittl斯皮特尔() feud.Barony {
	return c.斯皮特尔Spittl
}
    
func (c *菲拉赫KarntenCounty) BTreffen特雷芬() feud.Barony {
	return c.特雷芬Treffen
}
    
func (c *菲拉赫KarntenCounty) BVillach菲拉赫() feud.Barony {
	return c.菲拉赫Villach
}
    
func (c *菲拉赫KarntenCounty) BVorelach费拉赫() feud.Barony {
	return c.费拉赫Vorelach
}
    
var CKarnten菲拉赫 KarntenCounty = &菲拉赫KarntenCounty{}

func init() {
	f := CKarnten菲拉赫.(*菲拉赫KarntenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "456",
		Title:     "karnten",
		TitleName: "菲拉赫",
		TitleCode: "c_karnten",
		Baronies:  map[string]feud.Barony{},
	}

	f.格赖芬堡Greifenburg = BGreifenburg格赖芬堡
	f.格赖芬堡Greifenburg.SetParent(f)
	
	f.克拉根福特Klagenfurt = BKlagenfurt克拉根福特
	f.克拉根福特Klagenfurt.SetParent(f)
	
	f.利恩茨Lienz = BLienz利恩茨
	f.利恩茨Lienz.SetParent(f)
	
	f.斯皮特尔Spittl = BSpittl斯皮特尔
	f.斯皮特尔Spittl.SetParent(f)
	
	f.特雷芬Treffen = BTreffen特雷芬
	f.特雷芬Treffen.SetParent(f)
	
	f.菲拉赫Villach = BVillach菲拉赫
	f.菲拉赫Villach.SetParent(f)
	
	f.费拉赫Vorelach = BVorelach费拉赫
	f.费拉赫Vorelach.SetParent(f)
	
}
