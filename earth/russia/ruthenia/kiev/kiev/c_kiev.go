package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KievCounty interface {
    feud.County
    BFastiv法斯托夫() 	feud.Barony
    BKaniv卡尼夫() 	feud.Barony
    BKiev基辅() 	feud.Barony
    BVasykliv瓦西里基夫() 	feud.Barony
    BVyshhorod维什戈罗德() 	feud.Barony
    BYuriev尤里耶夫() 	feud.Barony
    BZhytomyr日托米尔() 	feud.Barony
}

type 基辅KievCounty struct {
	feud.BaseCounty
	法斯托夫Fastiv 	feud.Barony
	卡尼夫Kaniv 	feud.Barony
	基辅Kiev 	feud.Barony
	瓦西里基夫Vasykliv 	feud.Barony
	维什戈罗德Vyshhorod 	feud.Barony
	尤里耶夫Yuriev 	feud.Barony
	日托米尔Zhytomyr 	feud.Barony
}

func (c *基辅KievCounty) BFastiv法斯托夫() feud.Barony {
	return c.法斯托夫Fastiv
}
    
func (c *基辅KievCounty) BKaniv卡尼夫() feud.Barony {
	return c.卡尼夫Kaniv
}
    
func (c *基辅KievCounty) BKiev基辅() feud.Barony {
	return c.基辅Kiev
}
    
func (c *基辅KievCounty) BVasykliv瓦西里基夫() feud.Barony {
	return c.瓦西里基夫Vasykliv
}
    
func (c *基辅KievCounty) BVyshhorod维什戈罗德() feud.Barony {
	return c.维什戈罗德Vyshhorod
}
    
func (c *基辅KievCounty) BYuriev尤里耶夫() feud.Barony {
	return c.尤里耶夫Yuriev
}
    
func (c *基辅KievCounty) BZhytomyr日托米尔() feud.Barony {
	return c.日托米尔Zhytomyr
}
    
var CKiev基辅 KievCounty = &基辅KievCounty{}

func init() {
	f := CKiev基辅.(*基辅KievCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "547",
		Title:     "kiev",
		TitleName: "基辅",
		TitleCode: "c_kiev",
		Baronies:  map[string]feud.Barony{},
	}

	f.法斯托夫Fastiv = BFastiv法斯托夫
	f.法斯托夫Fastiv.SetParent(f)
	
	f.卡尼夫Kaniv = BKaniv卡尼夫
	f.卡尼夫Kaniv.SetParent(f)
	
	f.基辅Kiev = BKiev基辅
	f.基辅Kiev.SetParent(f)
	
	f.瓦西里基夫Vasykliv = BVasykliv瓦西里基夫
	f.瓦西里基夫Vasykliv.SetParent(f)
	
	f.维什戈罗德Vyshhorod = BVyshhorod维什戈罗德
	f.维什戈罗德Vyshhorod.SetParent(f)
	
	f.尤里耶夫Yuriev = BYuriev尤里耶夫
	f.尤里耶夫Yuriev.SetParent(f)
	
	f.日托米尔Zhytomyr = BZhytomyr日托米尔
	f.日托米尔Zhytomyr.SetParent(f)
	
}
