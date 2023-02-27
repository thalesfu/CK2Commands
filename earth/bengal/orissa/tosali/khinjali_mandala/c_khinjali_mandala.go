package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Khinjali_mandalaCounty interface {
    feud.County
    BBaramba巴兰巴() 	feud.Barony
    BKhinjali契那阇梨() 	feud.Barony
    BKujar俱折罗() 	feud.Barony
    BKunian鸠尼案() 	feud.Barony
    BMalther摩谛() 	feud.Barony
    BNayagarh那耶姞利呬() 	feud.Barony
    BNilamadhav尼罗摩陀婆() 	feud.Barony
}

type 契那阇梨曼荼罗Khinjali_mandalaCounty struct {
	feud.BaseCounty
	巴兰巴Baramba 	feud.Barony
	契那阇梨Khinjali 	feud.Barony
	俱折罗Kujar 	feud.Barony
	鸠尼案Kunian 	feud.Barony
	摩谛Malther 	feud.Barony
	那耶姞利呬Nayagarh 	feud.Barony
	尼罗摩陀婆Nilamadhav 	feud.Barony
}

func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BBaramba巴兰巴() feud.Barony {
	return c.巴兰巴Baramba
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BKhinjali契那阇梨() feud.Barony {
	return c.契那阇梨Khinjali
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BKujar俱折罗() feud.Barony {
	return c.俱折罗Kujar
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BKunian鸠尼案() feud.Barony {
	return c.鸠尼案Kunian
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BMalther摩谛() feud.Barony {
	return c.摩谛Malther
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BNayagarh那耶姞利呬() feud.Barony {
	return c.那耶姞利呬Nayagarh
}
    
func (c *契那阇梨曼荼罗Khinjali_mandalaCounty) BNilamadhav尼罗摩陀婆() feud.Barony {
	return c.尼罗摩陀婆Nilamadhav
}
    
var CKhinjali_mandala契那阇梨曼荼罗 Khinjali_mandalaCounty = &契那阇梨曼荼罗Khinjali_mandalaCounty{}

func init() {
	f := CKhinjali_mandala契那阇梨曼荼罗.(*契那阇梨曼荼罗Khinjali_mandalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1229",
		Title:     "khinjali_mandala",
		TitleName: "契那阇梨曼荼罗",
		TitleCode: "c_khinjali_mandala",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴兰巴Baramba = BBaramba巴兰巴
	f.巴兰巴Baramba.SetParent(f)
	
	f.契那阇梨Khinjali = BKhinjali契那阇梨
	f.契那阇梨Khinjali.SetParent(f)
	
	f.俱折罗Kujar = BKujar俱折罗
	f.俱折罗Kujar.SetParent(f)
	
	f.鸠尼案Kunian = BKunian鸠尼案
	f.鸠尼案Kunian.SetParent(f)
	
	f.摩谛Malther = BMalther摩谛
	f.摩谛Malther.SetParent(f)
	
	f.那耶姞利呬Nayagarh = BNayagarh那耶姞利呬
	f.那耶姞利呬Nayagarh.SetParent(f)
	
	f.尼罗摩陀婆Nilamadhav = BNilamadhav尼罗摩陀婆
	f.尼罗摩陀婆Nilamadhav.SetParent(f)
	
}
