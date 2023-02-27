package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PereyaslavlCounty interface {
    feud.County
    BBaryshivka巴雷希夫卡() 	feud.Barony
    BBerezan别列赞() 	feud.Barony
    BBoryspil鲍里斯波尔() 	feud.Barony
    BKozelets科泽列齐() 	feud.Barony
    BPereyaslavl佩列亚斯拉夫尔() 	feud.Barony
    BVelyka_dymerka大季梅尔卡() 	feud.Barony
    BYahotyn亚戈京() 	feud.Barony
}

type 佩列亚斯拉夫尔PereyaslavlCounty struct {
	feud.BaseCounty
	巴雷希夫卡Baryshivka 	feud.Barony
	别列赞Berezan 	feud.Barony
	鲍里斯波尔Boryspil 	feud.Barony
	科泽列齐Kozelets 	feud.Barony
	佩列亚斯拉夫尔Pereyaslavl 	feud.Barony
	大季梅尔卡Velyka_dymerka 	feud.Barony
	亚戈京Yahotyn 	feud.Barony
}

func (c *佩列亚斯拉夫尔PereyaslavlCounty) BBaryshivka巴雷希夫卡() feud.Barony {
	return c.巴雷希夫卡Baryshivka
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BBerezan别列赞() feud.Barony {
	return c.别列赞Berezan
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BBoryspil鲍里斯波尔() feud.Barony {
	return c.鲍里斯波尔Boryspil
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BKozelets科泽列齐() feud.Barony {
	return c.科泽列齐Kozelets
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BPereyaslavl佩列亚斯拉夫尔() feud.Barony {
	return c.佩列亚斯拉夫尔Pereyaslavl
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BVelyka_dymerka大季梅尔卡() feud.Barony {
	return c.大季梅尔卡Velyka_dymerka
}
    
func (c *佩列亚斯拉夫尔PereyaslavlCounty) BYahotyn亚戈京() feud.Barony {
	return c.亚戈京Yahotyn
}
    
var CPereyaslavl佩列亚斯拉夫尔 PereyaslavlCounty = &佩列亚斯拉夫尔PereyaslavlCounty{}

func init() {
	f := CPereyaslavl佩列亚斯拉夫尔.(*佩列亚斯拉夫尔PereyaslavlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "555",
		Title:     "pereyaslavl",
		TitleName: "佩列亚斯拉夫尔",
		TitleCode: "c_pereyaslavl",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴雷希夫卡Baryshivka = BBaryshivka巴雷希夫卡
	f.巴雷希夫卡Baryshivka.SetParent(f)
	
	f.别列赞Berezan = BBerezan别列赞
	f.别列赞Berezan.SetParent(f)
	
	f.鲍里斯波尔Boryspil = BBoryspil鲍里斯波尔
	f.鲍里斯波尔Boryspil.SetParent(f)
	
	f.科泽列齐Kozelets = BKozelets科泽列齐
	f.科泽列齐Kozelets.SetParent(f)
	
	f.佩列亚斯拉夫尔Pereyaslavl = BPereyaslavl佩列亚斯拉夫尔
	f.佩列亚斯拉夫尔Pereyaslavl.SetParent(f)
	
	f.大季梅尔卡Velyka_dymerka = BVelyka_dymerka大季梅尔卡
	f.大季梅尔卡Velyka_dymerka.SetParent(f)
	
	f.亚戈京Yahotyn = BYahotyn亚戈京
	f.亚戈京Yahotyn.SetParent(f)
	
}
