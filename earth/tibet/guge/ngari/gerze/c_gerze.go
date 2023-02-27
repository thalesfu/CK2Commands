package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GerzeCounty interface {
    feud.County
    BGerze改则() 	feud.Barony
    BGotsang古昌() 	feud.Barony
    BJakar_b吓嘎() 	feud.Barony
    BLangkor拉果() 	feud.Barony
    BNangangonna那咔沃那() 	feud.Barony
    BRimar日玛() 	feud.Barony
    BTongtso洞措() 	feud.Barony
}

type 改则GerzeCounty struct {
	feud.BaseCounty
	改则Gerze 	feud.Barony
	古昌Gotsang 	feud.Barony
	吓嘎Jakar_b 	feud.Barony
	拉果Langkor 	feud.Barony
	那咔沃那Nangangonna 	feud.Barony
	日玛Rimar 	feud.Barony
	洞措Tongtso 	feud.Barony
}

func (c *改则GerzeCounty) BGerze改则() feud.Barony {
	return c.改则Gerze
}
    
func (c *改则GerzeCounty) BGotsang古昌() feud.Barony {
	return c.古昌Gotsang
}
    
func (c *改则GerzeCounty) BJakar_b吓嘎() feud.Barony {
	return c.吓嘎Jakar_b
}
    
func (c *改则GerzeCounty) BLangkor拉果() feud.Barony {
	return c.拉果Langkor
}
    
func (c *改则GerzeCounty) BNangangonna那咔沃那() feud.Barony {
	return c.那咔沃那Nangangonna
}
    
func (c *改则GerzeCounty) BRimar日玛() feud.Barony {
	return c.日玛Rimar
}
    
func (c *改则GerzeCounty) BTongtso洞措() feud.Barony {
	return c.洞措Tongtso
}
    
var CGerze改则 GerzeCounty = &改则GerzeCounty{}

func init() {
	f := CGerze改则.(*改则GerzeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1494",
		Title:     "gerze",
		TitleName: "改则",
		TitleCode: "c_gerze",
		Baronies:  map[string]feud.Barony{},
	}

	f.改则Gerze = BGerze改则
	f.改则Gerze.SetParent(f)
	
	f.古昌Gotsang = BGotsang古昌
	f.古昌Gotsang.SetParent(f)
	
	f.吓嘎Jakar_b = BJakar_b吓嘎
	f.吓嘎Jakar_b.SetParent(f)
	
	f.拉果Langkor = BLangkor拉果
	f.拉果Langkor.SetParent(f)
	
	f.那咔沃那Nangangonna = BNangangonna那咔沃那
	f.那咔沃那Nangangonna.SetParent(f)
	
	f.日玛Rimar = BRimar日玛
	f.日玛Rimar.SetParent(f)
	
	f.洞措Tongtso = BTongtso洞措
	f.洞措Tongtso.SetParent(f)
	
}
