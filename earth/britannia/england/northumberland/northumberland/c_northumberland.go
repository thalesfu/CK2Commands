package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorthumberlandCounty interface {
    feud.County
    BAlnwick阿尼克() 	feud.Barony
    BBamburgh班堡() 	feud.Barony
    BHexham赫克瑟姆() 	feud.Barony
    BMitford米特福德() 	feud.Barony
    BMorpeth莫珀斯() 	feud.Barony
    BNewcastle纽卡斯尔() 	feud.Barony
    BSt_aidans圣阿丹() 	feud.Barony
}

type 诺森伯兰NorthumberlandCounty struct {
	feud.BaseCounty
	阿尼克Alnwick 	feud.Barony
	班堡Bamburgh 	feud.Barony
	赫克瑟姆Hexham 	feud.Barony
	米特福德Mitford 	feud.Barony
	莫珀斯Morpeth 	feud.Barony
	纽卡斯尔Newcastle 	feud.Barony
	圣阿丹St_aidans 	feud.Barony
}

func (c *诺森伯兰NorthumberlandCounty) BAlnwick阿尼克() feud.Barony {
	return c.阿尼克Alnwick
}
    
func (c *诺森伯兰NorthumberlandCounty) BBamburgh班堡() feud.Barony {
	return c.班堡Bamburgh
}
    
func (c *诺森伯兰NorthumberlandCounty) BHexham赫克瑟姆() feud.Barony {
	return c.赫克瑟姆Hexham
}
    
func (c *诺森伯兰NorthumberlandCounty) BMitford米特福德() feud.Barony {
	return c.米特福德Mitford
}
    
func (c *诺森伯兰NorthumberlandCounty) BMorpeth莫珀斯() feud.Barony {
	return c.莫珀斯Morpeth
}
    
func (c *诺森伯兰NorthumberlandCounty) BNewcastle纽卡斯尔() feud.Barony {
	return c.纽卡斯尔Newcastle
}
    
func (c *诺森伯兰NorthumberlandCounty) BSt_aidans圣阿丹() feud.Barony {
	return c.圣阿丹St_aidans
}
    
var CNorthumberland诺森伯兰 NorthumberlandCounty = &诺森伯兰NorthumberlandCounty{}

func init() {
	f := CNorthumberland诺森伯兰.(*诺森伯兰NorthumberlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "52",
		Title:     "northumberland",
		TitleName: "诺森伯兰",
		TitleCode: "c_northumberland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尼克Alnwick = BAlnwick阿尼克
	f.阿尼克Alnwick.SetParent(f)
	
	f.班堡Bamburgh = BBamburgh班堡
	f.班堡Bamburgh.SetParent(f)
	
	f.赫克瑟姆Hexham = BHexham赫克瑟姆
	f.赫克瑟姆Hexham.SetParent(f)
	
	f.米特福德Mitford = BMitford米特福德
	f.米特福德Mitford.SetParent(f)
	
	f.莫珀斯Morpeth = BMorpeth莫珀斯
	f.莫珀斯Morpeth.SetParent(f)
	
	f.纽卡斯尔Newcastle = BNewcastle纽卡斯尔
	f.纽卡斯尔Newcastle.SetParent(f)
	
	f.圣阿丹St_aidans = BSt_aidans圣阿丹
	f.圣阿丹St_aidans.SetParent(f)
	
}
