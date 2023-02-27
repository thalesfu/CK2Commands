package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Zanjan_abharCounty interface {
    feud.County
    BAba阿巴() 	feud.Barony
    BAbhar阿卜哈尔() 	feud.Barony
    BMozaffari莫扎法里() 	feud.Barony
    BQuzlucheh古兹卢切() 	feud.Barony
    BSuhraward苏赫拉瓦德() 	feud.Barony
    BWasamkuh瓦桑库赫() 	feud.Barony
    BZahanj扎汗杰() 	feud.Barony
}

type 赞詹_阿卜哈尔Zanjan_abharCounty struct {
	feud.BaseCounty
	阿巴Aba 	feud.Barony
	阿卜哈尔Abhar 	feud.Barony
	莫扎法里Mozaffari 	feud.Barony
	古兹卢切Quzlucheh 	feud.Barony
	苏赫拉瓦德Suhraward 	feud.Barony
	瓦桑库赫Wasamkuh 	feud.Barony
	扎汗杰Zahanj 	feud.Barony
}

func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BAba阿巴() feud.Barony {
	return c.阿巴Aba
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BAbhar阿卜哈尔() feud.Barony {
	return c.阿卜哈尔Abhar
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BMozaffari莫扎法里() feud.Barony {
	return c.莫扎法里Mozaffari
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BQuzlucheh古兹卢切() feud.Barony {
	return c.古兹卢切Quzlucheh
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BSuhraward苏赫拉瓦德() feud.Barony {
	return c.苏赫拉瓦德Suhraward
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BWasamkuh瓦桑库赫() feud.Barony {
	return c.瓦桑库赫Wasamkuh
}
    
func (c *赞詹_阿卜哈尔Zanjan_abharCounty) BZahanj扎汗杰() feud.Barony {
	return c.扎汗杰Zahanj
}
    
var CZanjan_abhar赞詹_阿卜哈尔 Zanjan_abharCounty = &赞詹_阿卜哈尔Zanjan_abharCounty{}

func init() {
	f := CZanjan_abhar赞詹_阿卜哈尔.(*赞詹_阿卜哈尔Zanjan_abharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "715",
		Title:     "zanjan_abhar",
		TitleName: "赞詹_阿卜哈尔",
		TitleCode: "c_zanjan_abhar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴Aba = BAba阿巴
	f.阿巴Aba.SetParent(f)
	
	f.阿卜哈尔Abhar = BAbhar阿卜哈尔
	f.阿卜哈尔Abhar.SetParent(f)
	
	f.莫扎法里Mozaffari = BMozaffari莫扎法里
	f.莫扎法里Mozaffari.SetParent(f)
	
	f.古兹卢切Quzlucheh = BQuzlucheh古兹卢切
	f.古兹卢切Quzlucheh.SetParent(f)
	
	f.苏赫拉瓦德Suhraward = BSuhraward苏赫拉瓦德
	f.苏赫拉瓦德Suhraward.SetParent(f)
	
	f.瓦桑库赫Wasamkuh = BWasamkuh瓦桑库赫
	f.瓦桑库赫Wasamkuh.SetParent(f)
	
	f.扎汗杰Zahanj = BZahanj扎汗杰
	f.扎汗杰Zahanj.SetParent(f)
	
}
