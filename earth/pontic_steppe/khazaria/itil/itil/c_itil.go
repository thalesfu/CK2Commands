package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ItilCounty interface {
    feud.County
    BAlga阿尔加() 	feud.Barony
    BItil阿得() 	feud.Barony
    BKamyzyak卡梅贾克() 	feud.Barony
    BKhaganbaligh可汗八里() 	feud.Barony
    BKharabali哈拉巴利() 	feud.Barony
    BTumak图马克() 	feud.Barony
    BXacitarxan哈吉塔尔汉() 	feud.Barony
}

type 阿得ItilCounty struct {
	feud.BaseCounty
	阿尔加Alga 	feud.Barony
	阿得Itil 	feud.Barony
	卡梅贾克Kamyzyak 	feud.Barony
	可汗八里Khaganbaligh 	feud.Barony
	哈拉巴利Kharabali 	feud.Barony
	图马克Tumak 	feud.Barony
	哈吉塔尔汉Xacitarxan 	feud.Barony
}

func (c *阿得ItilCounty) BAlga阿尔加() feud.Barony {
	return c.阿尔加Alga
}
    
func (c *阿得ItilCounty) BItil阿得() feud.Barony {
	return c.阿得Itil
}
    
func (c *阿得ItilCounty) BKamyzyak卡梅贾克() feud.Barony {
	return c.卡梅贾克Kamyzyak
}
    
func (c *阿得ItilCounty) BKhaganbaligh可汗八里() feud.Barony {
	return c.可汗八里Khaganbaligh
}
    
func (c *阿得ItilCounty) BKharabali哈拉巴利() feud.Barony {
	return c.哈拉巴利Kharabali
}
    
func (c *阿得ItilCounty) BTumak图马克() feud.Barony {
	return c.图马克Tumak
}
    
func (c *阿得ItilCounty) BXacitarxan哈吉塔尔汉() feud.Barony {
	return c.哈吉塔尔汉Xacitarxan
}
    
var CItil阿得 ItilCounty = &阿得ItilCounty{}

func init() {
	f := CItil阿得.(*阿得ItilCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "620",
		Title:     "itil",
		TitleName: "阿得",
		TitleCode: "c_itil",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔加Alga = BAlga阿尔加
	f.阿尔加Alga.SetParent(f)
	
	f.阿得Itil = BItil阿得
	f.阿得Itil.SetParent(f)
	
	f.卡梅贾克Kamyzyak = BKamyzyak卡梅贾克
	f.卡梅贾克Kamyzyak.SetParent(f)
	
	f.可汗八里Khaganbaligh = BKhaganbaligh可汗八里
	f.可汗八里Khaganbaligh.SetParent(f)
	
	f.哈拉巴利Kharabali = BKharabali哈拉巴利
	f.哈拉巴利Kharabali.SetParent(f)
	
	f.图马克Tumak = BTumak图马克
	f.图马克Tumak.SetParent(f)
	
	f.哈吉塔尔汉Xacitarxan = BXacitarxan哈吉塔尔汉
	f.哈吉塔尔汉Xacitarxan.SetParent(f)
	
}
