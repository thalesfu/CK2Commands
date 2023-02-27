package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YoredaleCounty interface {
    feud.County
    BBainbridge班布里奇() 	feud.Barony
    BBolton博尔顿() 	feud.Barony
    BIlkley伊尔克利() 	feud.Barony
    BMiddleham米德尔赫姆() 	feud.Barony
    BSettle塞特尔() 	feud.Barony
    BSkipton斯基普顿() 	feud.Barony
    BYoredale约代尔() 	feud.Barony
}

type 约代尔YoredaleCounty struct {
	feud.BaseCounty
	班布里奇Bainbridge 	feud.Barony
	博尔顿Bolton 	feud.Barony
	伊尔克利Ilkley 	feud.Barony
	米德尔赫姆Middleham 	feud.Barony
	塞特尔Settle 	feud.Barony
	斯基普顿Skipton 	feud.Barony
	约代尔Yoredale 	feud.Barony
}

func (c *约代尔YoredaleCounty) BBainbridge班布里奇() feud.Barony {
	return c.班布里奇Bainbridge
}
    
func (c *约代尔YoredaleCounty) BBolton博尔顿() feud.Barony {
	return c.博尔顿Bolton
}
    
func (c *约代尔YoredaleCounty) BIlkley伊尔克利() feud.Barony {
	return c.伊尔克利Ilkley
}
    
func (c *约代尔YoredaleCounty) BMiddleham米德尔赫姆() feud.Barony {
	return c.米德尔赫姆Middleham
}
    
func (c *约代尔YoredaleCounty) BSettle塞特尔() feud.Barony {
	return c.塞特尔Settle
}
    
func (c *约代尔YoredaleCounty) BSkipton斯基普顿() feud.Barony {
	return c.斯基普顿Skipton
}
    
func (c *约代尔YoredaleCounty) BYoredale约代尔() feud.Barony {
	return c.约代尔Yoredale
}
    
var CYoredale约代尔 YoredaleCounty = &约代尔YoredaleCounty{}

func init() {
	f := CYoredale约代尔.(*约代尔YoredaleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1943",
		Title:     "yoredale",
		TitleName: "约代尔",
		TitleCode: "c_yoredale",
		Baronies:  map[string]feud.Barony{},
	}

	f.班布里奇Bainbridge = BBainbridge班布里奇
	f.班布里奇Bainbridge.SetParent(f)
	
	f.博尔顿Bolton = BBolton博尔顿
	f.博尔顿Bolton.SetParent(f)
	
	f.伊尔克利Ilkley = BIlkley伊尔克利
	f.伊尔克利Ilkley.SetParent(f)
	
	f.米德尔赫姆Middleham = BMiddleham米德尔赫姆
	f.米德尔赫姆Middleham.SetParent(f)
	
	f.塞特尔Settle = BSettle塞特尔
	f.塞特尔Settle.SetParent(f)
	
	f.斯基普顿Skipton = BSkipton斯基普顿
	f.斯基普顿Skipton.SetParent(f)
	
	f.约代尔Yoredale = BYoredale约代尔
	f.约代尔Yoredale.SetParent(f)
	
}
