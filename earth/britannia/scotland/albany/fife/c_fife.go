package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FifeCounty interface {
    feud.County
    BCupar库珀() 	feud.Barony
    BDunfermline邓弗姆林() 	feud.Barony
    BFalkland福克兰() 	feud.Barony
    BKinross金罗斯() 	feud.Barony
    BKirkcaldy柯科迪() 	feud.Barony
    BLeuchars卢赫斯() 	feud.Barony
    BLochore洛霍尔() 	feud.Barony
    BSt_andrews圣安德鲁() 	feud.Barony
}

type 法夫FifeCounty struct {
	feud.BaseCounty
	库珀Cupar 	feud.Barony
	邓弗姆林Dunfermline 	feud.Barony
	福克兰Falkland 	feud.Barony
	金罗斯Kinross 	feud.Barony
	柯科迪Kirkcaldy 	feud.Barony
	卢赫斯Leuchars 	feud.Barony
	洛霍尔Lochore 	feud.Barony
	圣安德鲁St_andrews 	feud.Barony
}

func (c *法夫FifeCounty) BCupar库珀() feud.Barony {
	return c.库珀Cupar
}
    
func (c *法夫FifeCounty) BDunfermline邓弗姆林() feud.Barony {
	return c.邓弗姆林Dunfermline
}
    
func (c *法夫FifeCounty) BFalkland福克兰() feud.Barony {
	return c.福克兰Falkland
}
    
func (c *法夫FifeCounty) BKinross金罗斯() feud.Barony {
	return c.金罗斯Kinross
}
    
func (c *法夫FifeCounty) BKirkcaldy柯科迪() feud.Barony {
	return c.柯科迪Kirkcaldy
}
    
func (c *法夫FifeCounty) BLeuchars卢赫斯() feud.Barony {
	return c.卢赫斯Leuchars
}
    
func (c *法夫FifeCounty) BLochore洛霍尔() feud.Barony {
	return c.洛霍尔Lochore
}
    
func (c *法夫FifeCounty) BSt_andrews圣安德鲁() feud.Barony {
	return c.圣安德鲁St_andrews
}
    
var CFife法夫 FifeCounty = &法夫FifeCounty{}

func init() {
	f := CFife法夫.(*法夫FifeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "46",
		Title:     "fife",
		TitleName: "法夫",
		TitleCode: "c_fife",
		Baronies:  map[string]feud.Barony{},
	}

	f.库珀Cupar = BCupar库珀
	f.库珀Cupar.SetParent(f)
	
	f.邓弗姆林Dunfermline = BDunfermline邓弗姆林
	f.邓弗姆林Dunfermline.SetParent(f)
	
	f.福克兰Falkland = BFalkland福克兰
	f.福克兰Falkland.SetParent(f)
	
	f.金罗斯Kinross = BKinross金罗斯
	f.金罗斯Kinross.SetParent(f)
	
	f.柯科迪Kirkcaldy = BKirkcaldy柯科迪
	f.柯科迪Kirkcaldy.SetParent(f)
	
	f.卢赫斯Leuchars = BLeuchars卢赫斯
	f.卢赫斯Leuchars.SetParent(f)
	
	f.洛霍尔Lochore = BLochore洛霍尔
	f.洛霍尔Lochore.SetParent(f)
	
	f.圣安德鲁St_andrews = BSt_andrews圣安德鲁
	f.圣安德鲁St_andrews.SetParent(f)
	
}
