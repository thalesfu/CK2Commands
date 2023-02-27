package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WestmeathCounty interface {
    feud.County
    BBirr伯尔() 	feud.Barony
    BCenel_fiachach基内尔菲亚哈赫() 	feud.Barony
    BFir_cell菲尔凯尔() 	feud.Barony
    BFrewin弗雷温() 	feud.Barony
    BTethba泰斯韦() 	feud.Barony
    BUisnech伊什纳赫() 	feud.Barony
    BWestmeath韦斯特米斯() 	feud.Barony
}

type 韦斯特米斯WestmeathCounty struct {
	feud.BaseCounty
	伯尔Birr 	feud.Barony
	基内尔菲亚哈赫Cenel_fiachach 	feud.Barony
	菲尔凯尔Fir_cell 	feud.Barony
	弗雷温Frewin 	feud.Barony
	泰斯韦Tethba 	feud.Barony
	伊什纳赫Uisnech 	feud.Barony
	韦斯特米斯Westmeath 	feud.Barony
}

func (c *韦斯特米斯WestmeathCounty) BBirr伯尔() feud.Barony {
	return c.伯尔Birr
}
    
func (c *韦斯特米斯WestmeathCounty) BCenel_fiachach基内尔菲亚哈赫() feud.Barony {
	return c.基内尔菲亚哈赫Cenel_fiachach
}
    
func (c *韦斯特米斯WestmeathCounty) BFir_cell菲尔凯尔() feud.Barony {
	return c.菲尔凯尔Fir_cell
}
    
func (c *韦斯特米斯WestmeathCounty) BFrewin弗雷温() feud.Barony {
	return c.弗雷温Frewin
}
    
func (c *韦斯特米斯WestmeathCounty) BTethba泰斯韦() feud.Barony {
	return c.泰斯韦Tethba
}
    
func (c *韦斯特米斯WestmeathCounty) BUisnech伊什纳赫() feud.Barony {
	return c.伊什纳赫Uisnech
}
    
func (c *韦斯特米斯WestmeathCounty) BWestmeath韦斯特米斯() feud.Barony {
	return c.韦斯特米斯Westmeath
}
    
var CWestmeath韦斯特米斯 WestmeathCounty = &韦斯特米斯WestmeathCounty{}

func init() {
	f := CWestmeath韦斯特米斯.(*韦斯特米斯WestmeathCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1952",
		Title:     "westmeath",
		TitleName: "韦斯特米斯",
		TitleCode: "c_westmeath",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯尔Birr = BBirr伯尔
	f.伯尔Birr.SetParent(f)
	
	f.基内尔菲亚哈赫Cenel_fiachach = BCenel_fiachach基内尔菲亚哈赫
	f.基内尔菲亚哈赫Cenel_fiachach.SetParent(f)
	
	f.菲尔凯尔Fir_cell = BFir_cell菲尔凯尔
	f.菲尔凯尔Fir_cell.SetParent(f)
	
	f.弗雷温Frewin = BFrewin弗雷温
	f.弗雷温Frewin.SetParent(f)
	
	f.泰斯韦Tethba = BTethba泰斯韦
	f.泰斯韦Tethba.SetParent(f)
	
	f.伊什纳赫Uisnech = BUisnech伊什纳赫
	f.伊什纳赫Uisnech.SetParent(f)
	
	f.韦斯特米斯Westmeath = BWestmeath韦斯特米斯
	f.韦斯特米斯Westmeath.SetParent(f)
	
}
