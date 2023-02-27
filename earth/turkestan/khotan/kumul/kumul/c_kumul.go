package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KumulCounty interface {
    feud.County
    BDahe大河() 	feud.Barony
    BKumul哈密() 	feud.Barony
    BNanhu_tarim南湖() 	feud.Barony
    BPiqan辟展() 	feud.Barony
    BToyuq吐峪沟() 	feud.Barony
    BYanghai洋海() 	feud.Barony
    BYiwu伊吾() 	feud.Barony
}

type 哈密KumulCounty struct {
	feud.BaseCounty
	大河Dahe 	feud.Barony
	哈密Kumul 	feud.Barony
	南湖Nanhu_tarim 	feud.Barony
	辟展Piqan 	feud.Barony
	吐峪沟Toyuq 	feud.Barony
	洋海Yanghai 	feud.Barony
	伊吾Yiwu 	feud.Barony
}

func (c *哈密KumulCounty) BDahe大河() feud.Barony {
	return c.大河Dahe
}
    
func (c *哈密KumulCounty) BKumul哈密() feud.Barony {
	return c.哈密Kumul
}
    
func (c *哈密KumulCounty) BNanhu_tarim南湖() feud.Barony {
	return c.南湖Nanhu_tarim
}
    
func (c *哈密KumulCounty) BPiqan辟展() feud.Barony {
	return c.辟展Piqan
}
    
func (c *哈密KumulCounty) BToyuq吐峪沟() feud.Barony {
	return c.吐峪沟Toyuq
}
    
func (c *哈密KumulCounty) BYanghai洋海() feud.Barony {
	return c.洋海Yanghai
}
    
func (c *哈密KumulCounty) BYiwu伊吾() feud.Barony {
	return c.伊吾Yiwu
}
    
var CKumul哈密 KumulCounty = &哈密KumulCounty{}

func init() {
	f := CKumul哈密.(*哈密KumulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1450",
		Title:     "kumul",
		TitleName: "哈密",
		TitleCode: "c_kumul",
		Baronies:  map[string]feud.Barony{},
	}

	f.大河Dahe = BDahe大河
	f.大河Dahe.SetParent(f)
	
	f.哈密Kumul = BKumul哈密
	f.哈密Kumul.SetParent(f)
	
	f.南湖Nanhu_tarim = BNanhu_tarim南湖
	f.南湖Nanhu_tarim.SetParent(f)
	
	f.辟展Piqan = BPiqan辟展
	f.辟展Piqan.SetParent(f)
	
	f.吐峪沟Toyuq = BToyuq吐峪沟
	f.吐峪沟Toyuq.SetParent(f)
	
	f.洋海Yanghai = BYanghai洋海
	f.洋海Yanghai.SetParent(f)
	
	f.伊吾Yiwu = BYiwu伊吾
	f.伊吾Yiwu.SetParent(f)
	
}
