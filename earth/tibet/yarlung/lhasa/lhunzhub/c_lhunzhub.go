package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhunzhubCounty interface {
    feud.County
    BCodoi春堆() 	feud.Barony
    BKarze卡孜() 	feud.Barony
    BLhunzhub林周() 	feud.Barony
    BNgarnang阿朗() 	feud.Barony
    BPoindo旁多() 	feud.Barony
    BReting热振() 	feud.Barony
    BSumcheng松盘() 	feud.Barony
}

type 林周LhunzhubCounty struct {
	feud.BaseCounty
	春堆Codoi 	feud.Barony
	卡孜Karze 	feud.Barony
	林周Lhunzhub 	feud.Barony
	阿朗Ngarnang 	feud.Barony
	旁多Poindo 	feud.Barony
	热振Reting 	feud.Barony
	松盘Sumcheng 	feud.Barony
}

func (c *林周LhunzhubCounty) BCodoi春堆() feud.Barony {
	return c.春堆Codoi
}
    
func (c *林周LhunzhubCounty) BKarze卡孜() feud.Barony {
	return c.卡孜Karze
}
    
func (c *林周LhunzhubCounty) BLhunzhub林周() feud.Barony {
	return c.林周Lhunzhub
}
    
func (c *林周LhunzhubCounty) BNgarnang阿朗() feud.Barony {
	return c.阿朗Ngarnang
}
    
func (c *林周LhunzhubCounty) BPoindo旁多() feud.Barony {
	return c.旁多Poindo
}
    
func (c *林周LhunzhubCounty) BReting热振() feud.Barony {
	return c.热振Reting
}
    
func (c *林周LhunzhubCounty) BSumcheng松盘() feud.Barony {
	return c.松盘Sumcheng
}
    
var CLhunzhub林周 LhunzhubCounty = &林周LhunzhubCounty{}

func init() {
	f := CLhunzhub林周.(*林周LhunzhubCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1556",
		Title:     "lhunzhub",
		TitleName: "林周",
		TitleCode: "c_lhunzhub",
		Baronies:  map[string]feud.Barony{},
	}

	f.春堆Codoi = BCodoi春堆
	f.春堆Codoi.SetParent(f)
	
	f.卡孜Karze = BKarze卡孜
	f.卡孜Karze.SetParent(f)
	
	f.林周Lhunzhub = BLhunzhub林周
	f.林周Lhunzhub.SetParent(f)
	
	f.阿朗Ngarnang = BNgarnang阿朗
	f.阿朗Ngarnang.SetParent(f)
	
	f.旁多Poindo = BPoindo旁多
	f.旁多Poindo.SetParent(f)
	
	f.热振Reting = BReting热振
	f.热振Reting.SetParent(f)
	
	f.松盘Sumcheng = BSumcheng松盘
	f.松盘Sumcheng.SetParent(f)
	
}
