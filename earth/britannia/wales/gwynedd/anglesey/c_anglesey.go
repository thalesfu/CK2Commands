package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AngleseyCounty interface {
    feud.County
    BAberffraw阿伯弗劳() 	feud.Barony
    BAmlwch阿姆卢赫() 	feud.Barony
    BBeaumaris博马里斯() 	feud.Barony
    BBenllech本莱赫() 	feud.Barony
    BHolyhead霍利黑德() 	feud.Barony
    BLlanfaes兰法斯() 	feud.Barony
    BLlangefni兰盖夫尼() 	feud.Barony
}

type 安格尔西AngleseyCounty struct {
	feud.BaseCounty
	阿伯弗劳Aberffraw 	feud.Barony
	阿姆卢赫Amlwch 	feud.Barony
	博马里斯Beaumaris 	feud.Barony
	本莱赫Benllech 	feud.Barony
	霍利黑德Holyhead 	feud.Barony
	兰法斯Llanfaes 	feud.Barony
	兰盖夫尼Llangefni 	feud.Barony
}

func (c *安格尔西AngleseyCounty) BAberffraw阿伯弗劳() feud.Barony {
	return c.阿伯弗劳Aberffraw
}
    
func (c *安格尔西AngleseyCounty) BAmlwch阿姆卢赫() feud.Barony {
	return c.阿姆卢赫Amlwch
}
    
func (c *安格尔西AngleseyCounty) BBeaumaris博马里斯() feud.Barony {
	return c.博马里斯Beaumaris
}
    
func (c *安格尔西AngleseyCounty) BBenllech本莱赫() feud.Barony {
	return c.本莱赫Benllech
}
    
func (c *安格尔西AngleseyCounty) BHolyhead霍利黑德() feud.Barony {
	return c.霍利黑德Holyhead
}
    
func (c *安格尔西AngleseyCounty) BLlanfaes兰法斯() feud.Barony {
	return c.兰法斯Llanfaes
}
    
func (c *安格尔西AngleseyCounty) BLlangefni兰盖夫尼() feud.Barony {
	return c.兰盖夫尼Llangefni
}
    
var CAnglesey安格尔西 AngleseyCounty = &安格尔西AngleseyCounty{}

func init() {
	f := CAnglesey安格尔西.(*安格尔西AngleseyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1949",
		Title:     "anglesey",
		TitleName: "安格尔西",
		TitleCode: "c_anglesey",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯弗劳Aberffraw = BAberffraw阿伯弗劳
	f.阿伯弗劳Aberffraw.SetParent(f)
	
	f.阿姆卢赫Amlwch = BAmlwch阿姆卢赫
	f.阿姆卢赫Amlwch.SetParent(f)
	
	f.博马里斯Beaumaris = BBeaumaris博马里斯
	f.博马里斯Beaumaris.SetParent(f)
	
	f.本莱赫Benllech = BBenllech本莱赫
	f.本莱赫Benllech.SetParent(f)
	
	f.霍利黑德Holyhead = BHolyhead霍利黑德
	f.霍利黑德Holyhead.SetParent(f)
	
	f.兰法斯Llanfaes = BLlanfaes兰法斯
	f.兰法斯Llanfaes.SetParent(f)
	
	f.兰盖夫尼Llangefni = BLlangefni兰盖夫尼
	f.兰盖夫尼Llangefni.SetParent(f)
	
}
