package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChervenCounty interface {
    feud.County
    BCherven切尔文() 	feud.Barony
    BHrubeshiv赫鲁别舒夫() 	feud.Barony
    BKholm_ch霍尔姆() 	feud.Barony
    BKrasnystaw克拉斯内斯塔夫() 	feud.Barony
    BSaciaska松夏德卡() 	feud.Barony
    BSuteysk苏捷伊斯克() 	feud.Barony
    BVolyn沃伦() 	feud.Barony
}

type 切尔文ChervenCounty struct {
	feud.BaseCounty
	切尔文Cherven 	feud.Barony
	赫鲁别舒夫Hrubeshiv 	feud.Barony
	霍尔姆Kholm_ch 	feud.Barony
	克拉斯内斯塔夫Krasnystaw 	feud.Barony
	松夏德卡Saciaska 	feud.Barony
	苏捷伊斯克Suteysk 	feud.Barony
	沃伦Volyn 	feud.Barony
}

func (c *切尔文ChervenCounty) BCherven切尔文() feud.Barony {
	return c.切尔文Cherven
}
    
func (c *切尔文ChervenCounty) BHrubeshiv赫鲁别舒夫() feud.Barony {
	return c.赫鲁别舒夫Hrubeshiv
}
    
func (c *切尔文ChervenCounty) BKholm_ch霍尔姆() feud.Barony {
	return c.霍尔姆Kholm_ch
}
    
func (c *切尔文ChervenCounty) BKrasnystaw克拉斯内斯塔夫() feud.Barony {
	return c.克拉斯内斯塔夫Krasnystaw
}
    
func (c *切尔文ChervenCounty) BSaciaska松夏德卡() feud.Barony {
	return c.松夏德卡Saciaska
}
    
func (c *切尔文ChervenCounty) BSuteysk苏捷伊斯克() feud.Barony {
	return c.苏捷伊斯克Suteysk
}
    
func (c *切尔文ChervenCounty) BVolyn沃伦() feud.Barony {
	return c.沃伦Volyn
}
    
var CCherven切尔文 ChervenCounty = &切尔文ChervenCounty{}

func init() {
	f := CCherven切尔文.(*切尔文ChervenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1635",
		Title:     "cherven",
		TitleName: "切尔文",
		TitleCode: "c_cherven",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔文Cherven = BCherven切尔文
	f.切尔文Cherven.SetParent(f)
	
	f.赫鲁别舒夫Hrubeshiv = BHrubeshiv赫鲁别舒夫
	f.赫鲁别舒夫Hrubeshiv.SetParent(f)
	
	f.霍尔姆Kholm_ch = BKholm_ch霍尔姆
	f.霍尔姆Kholm_ch.SetParent(f)
	
	f.克拉斯内斯塔夫Krasnystaw = BKrasnystaw克拉斯内斯塔夫
	f.克拉斯内斯塔夫Krasnystaw.SetParent(f)
	
	f.松夏德卡Saciaska = BSaciaska松夏德卡
	f.松夏德卡Saciaska.SetParent(f)
	
	f.苏捷伊斯克Suteysk = BSuteysk苏捷伊斯克
	f.苏捷伊斯克Suteysk.SetParent(f)
	
	f.沃伦Volyn = BVolyn沃伦
	f.沃伦Volyn.SetParent(f)
	
}
