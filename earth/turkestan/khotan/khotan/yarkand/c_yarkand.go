package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YarkandCounty interface {
    feud.County
    BBaizhangzi柏杖子() 	feud.Barony
    BGujiazhai顾家寨() 	feud.Barony
    BGuma固玛() 	feud.Barony
    BKuqa苦恰() 	feud.Barony
    BMazar麻札() 	feud.Barony
    BYarkand鸦儿看() 	feud.Barony
    BYecheng叶城() 	feud.Barony
}

type 鸦儿看YarkandCounty struct {
	feud.BaseCounty
	柏杖子Baizhangzi 	feud.Barony
	顾家寨Gujiazhai 	feud.Barony
	固玛Guma 	feud.Barony
	苦恰Kuqa 	feud.Barony
	麻札Mazar 	feud.Barony
	鸦儿看Yarkand 	feud.Barony
	叶城Yecheng 	feud.Barony
}

func (c *鸦儿看YarkandCounty) BBaizhangzi柏杖子() feud.Barony {
	return c.柏杖子Baizhangzi
}
    
func (c *鸦儿看YarkandCounty) BGujiazhai顾家寨() feud.Barony {
	return c.顾家寨Gujiazhai
}
    
func (c *鸦儿看YarkandCounty) BGuma固玛() feud.Barony {
	return c.固玛Guma
}
    
func (c *鸦儿看YarkandCounty) BKuqa苦恰() feud.Barony {
	return c.苦恰Kuqa
}
    
func (c *鸦儿看YarkandCounty) BMazar麻札() feud.Barony {
	return c.麻札Mazar
}
    
func (c *鸦儿看YarkandCounty) BYarkand鸦儿看() feud.Barony {
	return c.鸦儿看Yarkand
}
    
func (c *鸦儿看YarkandCounty) BYecheng叶城() feud.Barony {
	return c.叶城Yecheng
}
    
var CYarkand鸦儿看 YarkandCounty = &鸦儿看YarkandCounty{}

func init() {
	f := CYarkand鸦儿看.(*鸦儿看YarkandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1438",
		Title:     "yarkand",
		TitleName: "鸦儿看",
		TitleCode: "c_yarkand",
		Baronies:  map[string]feud.Barony{},
	}

	f.柏杖子Baizhangzi = BBaizhangzi柏杖子
	f.柏杖子Baizhangzi.SetParent(f)
	
	f.顾家寨Gujiazhai = BGujiazhai顾家寨
	f.顾家寨Gujiazhai.SetParent(f)
	
	f.固玛Guma = BGuma固玛
	f.固玛Guma.SetParent(f)
	
	f.苦恰Kuqa = BKuqa苦恰
	f.苦恰Kuqa.SetParent(f)
	
	f.麻札Mazar = BMazar麻札
	f.麻札Mazar.SetParent(f)
	
	f.鸦儿看Yarkand = BYarkand鸦儿看
	f.鸦儿看Yarkand.SetParent(f)
	
	f.叶城Yecheng = BYecheng叶城
	f.叶城Yecheng.SetParent(f)
	
}
