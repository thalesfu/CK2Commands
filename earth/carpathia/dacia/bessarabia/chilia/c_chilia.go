package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChiliaCounty interface {
    feud.County
    BCahul卡胡尔() 	feud.Barony
    BChilia基利亚() 	feud.Barony
    BLisky利斯基() 	feud.Barony
    BOblucita奥布卢奇察() 	feud.Barony
    BPalada帕拉德() 	feud.Barony
    BReni_ro列尼() 	feud.Barony
    BTigheci蒂盖奇() 	feud.Barony
}

type 基利亚ChiliaCounty struct {
	feud.BaseCounty
	卡胡尔Cahul 	feud.Barony
	基利亚Chilia 	feud.Barony
	利斯基Lisky 	feud.Barony
	奥布卢奇察Oblucita 	feud.Barony
	帕拉德Palada 	feud.Barony
	列尼Reni_ro 	feud.Barony
	蒂盖奇Tigheci 	feud.Barony
}

func (c *基利亚ChiliaCounty) BCahul卡胡尔() feud.Barony {
	return c.卡胡尔Cahul
}
    
func (c *基利亚ChiliaCounty) BChilia基利亚() feud.Barony {
	return c.基利亚Chilia
}
    
func (c *基利亚ChiliaCounty) BLisky利斯基() feud.Barony {
	return c.利斯基Lisky
}
    
func (c *基利亚ChiliaCounty) BOblucita奥布卢奇察() feud.Barony {
	return c.奥布卢奇察Oblucita
}
    
func (c *基利亚ChiliaCounty) BPalada帕拉德() feud.Barony {
	return c.帕拉德Palada
}
    
func (c *基利亚ChiliaCounty) BReni_ro列尼() feud.Barony {
	return c.列尼Reni_ro
}
    
func (c *基利亚ChiliaCounty) BTigheci蒂盖奇() feud.Barony {
	return c.蒂盖奇Tigheci
}
    
var CChilia基利亚 ChiliaCounty = &基利亚ChiliaCounty{}

func init() {
	f := CChilia基利亚.(*基利亚ChiliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1640",
		Title:     "chilia",
		TitleName: "基利亚",
		TitleCode: "c_chilia",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡胡尔Cahul = BCahul卡胡尔
	f.卡胡尔Cahul.SetParent(f)
	
	f.基利亚Chilia = BChilia基利亚
	f.基利亚Chilia.SetParent(f)
	
	f.利斯基Lisky = BLisky利斯基
	f.利斯基Lisky.SetParent(f)
	
	f.奥布卢奇察Oblucita = BOblucita奥布卢奇察
	f.奥布卢奇察Oblucita.SetParent(f)
	
	f.帕拉德Palada = BPalada帕拉德
	f.帕拉德Palada.SetParent(f)
	
	f.列尼Reni_ro = BReni_ro列尼
	f.列尼Reni_ro.SetParent(f)
	
	f.蒂盖奇Tigheci = BTigheci蒂盖奇
	f.蒂盖奇Tigheci.SetParent(f)
	
}
