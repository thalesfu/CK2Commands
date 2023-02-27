package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GegyaiCounty interface {
    feud.County
    BArkog阿尔过() 	feud.Barony
    BChuktsang却藏() 	feud.Barony
    BGegyai革吉() 	feud.Barony
    BKelle吉勒() 	feud.Barony
    BSekhadrig色喀执() 	feud.Barony
    BYakra亚热() 	feud.Barony
    BZhungpa雄巴() 	feud.Barony
}

type 革吉GegyaiCounty struct {
	feud.BaseCounty
	阿尔过Arkog 	feud.Barony
	却藏Chuktsang 	feud.Barony
	革吉Gegyai 	feud.Barony
	吉勒Kelle 	feud.Barony
	色喀执Sekhadrig 	feud.Barony
	亚热Yakra 	feud.Barony
	雄巴Zhungpa 	feud.Barony
}

func (c *革吉GegyaiCounty) BArkog阿尔过() feud.Barony {
	return c.阿尔过Arkog
}
    
func (c *革吉GegyaiCounty) BChuktsang却藏() feud.Barony {
	return c.却藏Chuktsang
}
    
func (c *革吉GegyaiCounty) BGegyai革吉() feud.Barony {
	return c.革吉Gegyai
}
    
func (c *革吉GegyaiCounty) BKelle吉勒() feud.Barony {
	return c.吉勒Kelle
}
    
func (c *革吉GegyaiCounty) BSekhadrig色喀执() feud.Barony {
	return c.色喀执Sekhadrig
}
    
func (c *革吉GegyaiCounty) BYakra亚热() feud.Barony {
	return c.亚热Yakra
}
    
func (c *革吉GegyaiCounty) BZhungpa雄巴() feud.Barony {
	return c.雄巴Zhungpa
}
    
var CGegyai革吉 GegyaiCounty = &革吉GegyaiCounty{}

func init() {
	f := CGegyai革吉.(*革吉GegyaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1563",
		Title:     "gegyai",
		TitleName: "革吉",
		TitleCode: "c_gegyai",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔过Arkog = BArkog阿尔过
	f.阿尔过Arkog.SetParent(f)
	
	f.却藏Chuktsang = BChuktsang却藏
	f.却藏Chuktsang.SetParent(f)
	
	f.革吉Gegyai = BGegyai革吉
	f.革吉Gegyai.SetParent(f)
	
	f.吉勒Kelle = BKelle吉勒
	f.吉勒Kelle.SetParent(f)
	
	f.色喀执Sekhadrig = BSekhadrig色喀执
	f.色喀执Sekhadrig.SetParent(f)
	
	f.亚热Yakra = BYakra亚热
	f.亚热Yakra.SetParent(f)
	
	f.雄巴Zhungpa = BZhungpa雄巴
	f.雄巴Zhungpa.SetParent(f)
	
}
