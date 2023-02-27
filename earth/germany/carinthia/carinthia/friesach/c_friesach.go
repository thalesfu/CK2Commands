package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FriesachCounty interface {
    feud.County
    BAlthofen阿尔特霍芬() 	feud.Barony
    BMicheldorf米歇尔多夫() 	feud.Barony
    BPetersberg彼得斯贝格() 	feud.Barony
    BSankt_lambrecht圣兰布雷希特() 	feud.Barony
    BSankt_veit圣法伊特() 	feud.Barony
    BSt_nicholas_friesach圣尼古拉斯() 	feud.Barony
    BWolfsberg沃尔夫斯堡() 	feud.Barony
}

type 弗里萨赫FriesachCounty struct {
	feud.BaseCounty
	阿尔特霍芬Althofen 	feud.Barony
	米歇尔多夫Micheldorf 	feud.Barony
	彼得斯贝格Petersberg 	feud.Barony
	圣兰布雷希特Sankt_lambrecht 	feud.Barony
	圣法伊特Sankt_veit 	feud.Barony
	圣尼古拉斯St_nicholas_friesach 	feud.Barony
	沃尔夫斯堡Wolfsberg 	feud.Barony
}

func (c *弗里萨赫FriesachCounty) BAlthofen阿尔特霍芬() feud.Barony {
	return c.阿尔特霍芬Althofen
}
    
func (c *弗里萨赫FriesachCounty) BMicheldorf米歇尔多夫() feud.Barony {
	return c.米歇尔多夫Micheldorf
}
    
func (c *弗里萨赫FriesachCounty) BPetersberg彼得斯贝格() feud.Barony {
	return c.彼得斯贝格Petersberg
}
    
func (c *弗里萨赫FriesachCounty) BSankt_lambrecht圣兰布雷希特() feud.Barony {
	return c.圣兰布雷希特Sankt_lambrecht
}
    
func (c *弗里萨赫FriesachCounty) BSankt_veit圣法伊特() feud.Barony {
	return c.圣法伊特Sankt_veit
}
    
func (c *弗里萨赫FriesachCounty) BSt_nicholas_friesach圣尼古拉斯() feud.Barony {
	return c.圣尼古拉斯St_nicholas_friesach
}
    
func (c *弗里萨赫FriesachCounty) BWolfsberg沃尔夫斯堡() feud.Barony {
	return c.沃尔夫斯堡Wolfsberg
}
    
var CFriesach弗里萨赫 FriesachCounty = &弗里萨赫FriesachCounty{}

func init() {
	f := CFriesach弗里萨赫.(*弗里萨赫FriesachCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1696",
		Title:     "friesach",
		TitleName: "弗里萨赫",
		TitleCode: "c_friesach",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔特霍芬Althofen = BAlthofen阿尔特霍芬
	f.阿尔特霍芬Althofen.SetParent(f)
	
	f.米歇尔多夫Micheldorf = BMicheldorf米歇尔多夫
	f.米歇尔多夫Micheldorf.SetParent(f)
	
	f.彼得斯贝格Petersberg = BPetersberg彼得斯贝格
	f.彼得斯贝格Petersberg.SetParent(f)
	
	f.圣兰布雷希特Sankt_lambrecht = BSankt_lambrecht圣兰布雷希特
	f.圣兰布雷希特Sankt_lambrecht.SetParent(f)
	
	f.圣法伊特Sankt_veit = BSankt_veit圣法伊特
	f.圣法伊特Sankt_veit.SetParent(f)
	
	f.圣尼古拉斯St_nicholas_friesach = BSt_nicholas_friesach圣尼古拉斯
	f.圣尼古拉斯St_nicholas_friesach.SetParent(f)
	
	f.沃尔夫斯堡Wolfsberg = BWolfsberg沃尔夫斯堡
	f.沃尔夫斯堡Wolfsberg.SetParent(f)
	
}
