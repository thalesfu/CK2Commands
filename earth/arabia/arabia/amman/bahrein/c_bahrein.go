package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BahreinCounty interface {
    feud.County
    BAl_khuwayr胡韦尔() 	feud.Barony
    BAl_masqar马斯加尔() 	feud.Barony
    BAl_naman纳曼() 	feud.Barony
    BBatn_ardasir巴特恩阿尔达希尔() 	feud.Barony
    BMurwab穆尔瓦比() 	feud.Barony
    BTuraynah图赖纳() 	feud.Barony
    BZubarah祖巴拉() 	feud.Barony
}

type 巴林BahreinCounty struct {
	feud.BaseCounty
	胡韦尔Al_khuwayr 	feud.Barony
	马斯加尔Al_masqar 	feud.Barony
	纳曼Al_naman 	feud.Barony
	巴特恩阿尔达希尔Batn_ardasir 	feud.Barony
	穆尔瓦比Murwab 	feud.Barony
	图赖纳Turaynah 	feud.Barony
	祖巴拉Zubarah 	feud.Barony
}

func (c *巴林BahreinCounty) BAl_khuwayr胡韦尔() feud.Barony {
	return c.胡韦尔Al_khuwayr
}
    
func (c *巴林BahreinCounty) BAl_masqar马斯加尔() feud.Barony {
	return c.马斯加尔Al_masqar
}
    
func (c *巴林BahreinCounty) BAl_naman纳曼() feud.Barony {
	return c.纳曼Al_naman
}
    
func (c *巴林BahreinCounty) BBatn_ardasir巴特恩阿尔达希尔() feud.Barony {
	return c.巴特恩阿尔达希尔Batn_ardasir
}
    
func (c *巴林BahreinCounty) BMurwab穆尔瓦比() feud.Barony {
	return c.穆尔瓦比Murwab
}
    
func (c *巴林BahreinCounty) BTuraynah图赖纳() feud.Barony {
	return c.图赖纳Turaynah
}
    
func (c *巴林BahreinCounty) BZubarah祖巴拉() feud.Barony {
	return c.祖巴拉Zubarah
}
    
var CBahrein巴林 BahreinCounty = &巴林BahreinCounty{}

func init() {
	f := CBahrein巴林.(*巴林BahreinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "653",
		Title:     "bahrein",
		TitleName: "巴林",
		TitleCode: "c_bahrein",
		Baronies:  map[string]feud.Barony{},
	}

	f.胡韦尔Al_khuwayr = BAl_khuwayr胡韦尔
	f.胡韦尔Al_khuwayr.SetParent(f)
	
	f.马斯加尔Al_masqar = BAl_masqar马斯加尔
	f.马斯加尔Al_masqar.SetParent(f)
	
	f.纳曼Al_naman = BAl_naman纳曼
	f.纳曼Al_naman.SetParent(f)
	
	f.巴特恩阿尔达希尔Batn_ardasir = BBatn_ardasir巴特恩阿尔达希尔
	f.巴特恩阿尔达希尔Batn_ardasir.SetParent(f)
	
	f.穆尔瓦比Murwab = BMurwab穆尔瓦比
	f.穆尔瓦比Murwab.SetParent(f)
	
	f.图赖纳Turaynah = BTuraynah图赖纳
	f.图赖纳Turaynah.SetParent(f)
	
	f.祖巴拉Zubarah = BZubarah祖巴拉
	f.祖巴拉Zubarah.SetParent(f)
	
}
