package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuchaCounty interface {
    feud.County
    BBay拜城() 	feud.Barony
    BKizil克孜尔() 	feud.Barony
    BKucha苦叉() 	feud.Barony
    BKumtura库木吐喇() 	feud.Barony
    BSubashi苏巴什() 	feud.Barony
    BXayar沙雅() 	feud.Barony
    BZhangjiawan张家湾() 	feud.Barony
}

type 苦叉KuchaCounty struct {
	feud.BaseCounty
	拜城Bay 	feud.Barony
	克孜尔Kizil 	feud.Barony
	苦叉Kucha 	feud.Barony
	库木吐喇Kumtura 	feud.Barony
	苏巴什Subashi 	feud.Barony
	沙雅Xayar 	feud.Barony
	张家湾Zhangjiawan 	feud.Barony
}

func (c *苦叉KuchaCounty) BBay拜城() feud.Barony {
	return c.拜城Bay
}
    
func (c *苦叉KuchaCounty) BKizil克孜尔() feud.Barony {
	return c.克孜尔Kizil
}
    
func (c *苦叉KuchaCounty) BKucha苦叉() feud.Barony {
	return c.苦叉Kucha
}
    
func (c *苦叉KuchaCounty) BKumtura库木吐喇() feud.Barony {
	return c.库木吐喇Kumtura
}
    
func (c *苦叉KuchaCounty) BSubashi苏巴什() feud.Barony {
	return c.苏巴什Subashi
}
    
func (c *苦叉KuchaCounty) BXayar沙雅() feud.Barony {
	return c.沙雅Xayar
}
    
func (c *苦叉KuchaCounty) BZhangjiawan张家湾() feud.Barony {
	return c.张家湾Zhangjiawan
}
    
var CKucha苦叉 KuchaCounty = &苦叉KuchaCounty{}

func init() {
	f := CKucha苦叉.(*苦叉KuchaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1444",
		Title:     "kucha",
		TitleName: "苦叉",
		TitleCode: "c_kucha",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜城Bay = BBay拜城
	f.拜城Bay.SetParent(f)
	
	f.克孜尔Kizil = BKizil克孜尔
	f.克孜尔Kizil.SetParent(f)
	
	f.苦叉Kucha = BKucha苦叉
	f.苦叉Kucha.SetParent(f)
	
	f.库木吐喇Kumtura = BKumtura库木吐喇
	f.库木吐喇Kumtura.SetParent(f)
	
	f.苏巴什Subashi = BSubashi苏巴什
	f.苏巴什Subashi.SetParent(f)
	
	f.沙雅Xayar = BXayar沙雅
	f.沙雅Xayar.SetParent(f)
	
	f.张家湾Zhangjiawan = BZhangjiawan张家湾
	f.张家湾Zhangjiawan.SetParent(f)
	
}
