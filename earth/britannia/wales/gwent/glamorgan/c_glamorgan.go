package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GlamorganCounty interface {
    feud.County
    BCaerphilly卡菲利() 	feud.Barony
    BCardiff加的夫() 	feud.Barony
    BLlandaff兰达夫() 	feud.Barony
    BMargam马格姆() 	feud.Barony
    BNeath尼思() 	feud.Barony
    BOgmore奥格莫尔() 	feud.Barony
    BSt_donats圣多纳茨() 	feud.Barony
}

type 格拉摩根GlamorganCounty struct {
	feud.BaseCounty
	卡菲利Caerphilly 	feud.Barony
	加的夫Cardiff 	feud.Barony
	兰达夫Llandaff 	feud.Barony
	马格姆Margam 	feud.Barony
	尼思Neath 	feud.Barony
	奥格莫尔Ogmore 	feud.Barony
	圣多纳茨St_donats 	feud.Barony
}

func (c *格拉摩根GlamorganCounty) BCaerphilly卡菲利() feud.Barony {
	return c.卡菲利Caerphilly
}
    
func (c *格拉摩根GlamorganCounty) BCardiff加的夫() feud.Barony {
	return c.加的夫Cardiff
}
    
func (c *格拉摩根GlamorganCounty) BLlandaff兰达夫() feud.Barony {
	return c.兰达夫Llandaff
}
    
func (c *格拉摩根GlamorganCounty) BMargam马格姆() feud.Barony {
	return c.马格姆Margam
}
    
func (c *格拉摩根GlamorganCounty) BNeath尼思() feud.Barony {
	return c.尼思Neath
}
    
func (c *格拉摩根GlamorganCounty) BOgmore奥格莫尔() feud.Barony {
	return c.奥格莫尔Ogmore
}
    
func (c *格拉摩根GlamorganCounty) BSt_donats圣多纳茨() feud.Barony {
	return c.圣多纳茨St_donats
}
    
var CGlamorgan格拉摩根 GlamorganCounty = &格拉摩根GlamorganCounty{}

func init() {
	f := CGlamorgan格拉摩根.(*格拉摩根GlamorganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "19",
		Title:     "glamorgan",
		TitleName: "格拉摩根",
		TitleCode: "c_glamorgan",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡菲利Caerphilly = BCaerphilly卡菲利
	f.卡菲利Caerphilly.SetParent(f)
	
	f.加的夫Cardiff = BCardiff加的夫
	f.加的夫Cardiff.SetParent(f)
	
	f.兰达夫Llandaff = BLlandaff兰达夫
	f.兰达夫Llandaff.SetParent(f)
	
	f.马格姆Margam = BMargam马格姆
	f.马格姆Margam.SetParent(f)
	
	f.尼思Neath = BNeath尼思
	f.尼思Neath.SetParent(f)
	
	f.奥格莫尔Ogmore = BOgmore奥格莫尔
	f.奥格莫尔Ogmore.SetParent(f)
	
	f.圣多纳茨St_donats = BSt_donats圣多纳茨
	f.圣多纳茨St_donats.SetParent(f)
	
}
