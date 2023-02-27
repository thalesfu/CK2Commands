package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kodalaka_mandalaCounty interface {
    feud.County
    BBajrakot巴加罗科塔() 	feud.Barony
    BDeogarh提婆姞利呬() 	feud.Barony
    BDhenkanal登卡纳尔() 	feud.Barony
    BJoranda_gadhi基兰达加迪() 	feud.Barony
    BKapilash卡佩拉什() 	feud.Barony
    BKaramul卡拉穆尔() 	feud.Barony
    BKodalaka拘陀罗迦() 	feud.Barony
    BTalcher塔尔切尔() 	feud.Barony
}

type 拘陀罗迦曼荼罗Kodalaka_mandalaCounty struct {
	feud.BaseCounty
	巴加罗科塔Bajrakot 	feud.Barony
	提婆姞利呬Deogarh 	feud.Barony
	登卡纳尔Dhenkanal 	feud.Barony
	基兰达加迪Joranda_gadhi 	feud.Barony
	卡佩拉什Kapilash 	feud.Barony
	卡拉穆尔Karamul 	feud.Barony
	拘陀罗迦Kodalaka 	feud.Barony
	塔尔切尔Talcher 	feud.Barony
}

func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BBajrakot巴加罗科塔() feud.Barony {
	return c.巴加罗科塔Bajrakot
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BDeogarh提婆姞利呬() feud.Barony {
	return c.提婆姞利呬Deogarh
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BDhenkanal登卡纳尔() feud.Barony {
	return c.登卡纳尔Dhenkanal
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BJoranda_gadhi基兰达加迪() feud.Barony {
	return c.基兰达加迪Joranda_gadhi
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BKapilash卡佩拉什() feud.Barony {
	return c.卡佩拉什Kapilash
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BKaramul卡拉穆尔() feud.Barony {
	return c.卡拉穆尔Karamul
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BKodalaka拘陀罗迦() feud.Barony {
	return c.拘陀罗迦Kodalaka
}
    
func (c *拘陀罗迦曼荼罗Kodalaka_mandalaCounty) BTalcher塔尔切尔() feud.Barony {
	return c.塔尔切尔Talcher
}
    
var CKodalaka_mandala拘陀罗迦曼荼罗 Kodalaka_mandalaCounty = &拘陀罗迦曼荼罗Kodalaka_mandalaCounty{}

func init() {
	f := CKodalaka_mandala拘陀罗迦曼荼罗.(*拘陀罗迦曼荼罗Kodalaka_mandalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1156",
		Title:     "kodalaka_mandala",
		TitleName: "拘陀罗迦曼荼罗",
		TitleCode: "c_kodalaka_mandala",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴加罗科塔Bajrakot = BBajrakot巴加罗科塔
	f.巴加罗科塔Bajrakot.SetParent(f)
	
	f.提婆姞利呬Deogarh = BDeogarh提婆姞利呬
	f.提婆姞利呬Deogarh.SetParent(f)
	
	f.登卡纳尔Dhenkanal = BDhenkanal登卡纳尔
	f.登卡纳尔Dhenkanal.SetParent(f)
	
	f.基兰达加迪Joranda_gadhi = BJoranda_gadhi基兰达加迪
	f.基兰达加迪Joranda_gadhi.SetParent(f)
	
	f.卡佩拉什Kapilash = BKapilash卡佩拉什
	f.卡佩拉什Kapilash.SetParent(f)
	
	f.卡拉穆尔Karamul = BKaramul卡拉穆尔
	f.卡拉穆尔Karamul.SetParent(f)
	
	f.拘陀罗迦Kodalaka = BKodalaka拘陀罗迦
	f.拘陀罗迦Kodalaka.SetParent(f)
	
	f.塔尔切尔Talcher = BTalcher塔尔切尔
	f.塔尔切尔Talcher.SetParent(f)
	
}
