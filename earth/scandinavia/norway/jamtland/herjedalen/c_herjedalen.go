package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HerjedalenCounty interface {
    feud.County
    BHogvalen赫格沃伦() 	feud.Barony
    BLillhardal小海达尔() 	feud.Barony
    BLjusnedal尤斯内达尔() 	feud.Barony
    BSveg斯韦格() 	feud.Barony
    BTannas坦奈斯() 	feud.Barony
    BUlvkalla乌尔夫谢拉() 	feud.Barony
    BVemdalen韦姆达伦() 	feud.Barony
}

type 海里耶达伦HerjedalenCounty struct {
	feud.BaseCounty
	赫格沃伦Hogvalen 	feud.Barony
	小海达尔Lillhardal 	feud.Barony
	尤斯内达尔Ljusnedal 	feud.Barony
	斯韦格Sveg 	feud.Barony
	坦奈斯Tannas 	feud.Barony
	乌尔夫谢拉Ulvkalla 	feud.Barony
	韦姆达伦Vemdalen 	feud.Barony
}

func (c *海里耶达伦HerjedalenCounty) BHogvalen赫格沃伦() feud.Barony {
	return c.赫格沃伦Hogvalen
}
    
func (c *海里耶达伦HerjedalenCounty) BLillhardal小海达尔() feud.Barony {
	return c.小海达尔Lillhardal
}
    
func (c *海里耶达伦HerjedalenCounty) BLjusnedal尤斯内达尔() feud.Barony {
	return c.尤斯内达尔Ljusnedal
}
    
func (c *海里耶达伦HerjedalenCounty) BSveg斯韦格() feud.Barony {
	return c.斯韦格Sveg
}
    
func (c *海里耶达伦HerjedalenCounty) BTannas坦奈斯() feud.Barony {
	return c.坦奈斯Tannas
}
    
func (c *海里耶达伦HerjedalenCounty) BUlvkalla乌尔夫谢拉() feud.Barony {
	return c.乌尔夫谢拉Ulvkalla
}
    
func (c *海里耶达伦HerjedalenCounty) BVemdalen韦姆达伦() feud.Barony {
	return c.韦姆达伦Vemdalen
}
    
var CHerjedalen海里耶达伦 HerjedalenCounty = &海里耶达伦HerjedalenCounty{}

func init() {
	f := CHerjedalen海里耶达伦.(*海里耶达伦HerjedalenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "284",
		Title:     "herjedalen",
		TitleName: "海里耶达伦",
		TitleCode: "c_herjedalen",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫格沃伦Hogvalen = BHogvalen赫格沃伦
	f.赫格沃伦Hogvalen.SetParent(f)
	
	f.小海达尔Lillhardal = BLillhardal小海达尔
	f.小海达尔Lillhardal.SetParent(f)
	
	f.尤斯内达尔Ljusnedal = BLjusnedal尤斯内达尔
	f.尤斯内达尔Ljusnedal.SetParent(f)
	
	f.斯韦格Sveg = BSveg斯韦格
	f.斯韦格Sveg.SetParent(f)
	
	f.坦奈斯Tannas = BTannas坦奈斯
	f.坦奈斯Tannas.SetParent(f)
	
	f.乌尔夫谢拉Ulvkalla = BUlvkalla乌尔夫谢拉
	f.乌尔夫谢拉Ulvkalla.SetParent(f)
	
	f.韦姆达伦Vemdalen = BVemdalen韦姆达伦
	f.韦姆达伦Vemdalen.SetParent(f)
	
}
