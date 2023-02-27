package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Aj_bogdCounty interface {
    feud.County
    BAj_bogd阿主孛黑答() 	feud.Barony
    BBalikun巴里坤() 	feud.Barony
    BChangcun常村() 	feud.Barony
    BErdene额尔德尼() 	feud.Barony
    BGurvantes古尔班特斯() 	feud.Barony
    BNanlizhuang南里庄() 	feud.Barony
    BQian_pulei前蒲类() 	feud.Barony
}

type 阿主孛黑答Aj_bogdCounty struct {
	feud.BaseCounty
	阿主孛黑答Aj_bogd 	feud.Barony
	巴里坤Balikun 	feud.Barony
	常村Changcun 	feud.Barony
	额尔德尼Erdene 	feud.Barony
	古尔班特斯Gurvantes 	feud.Barony
	南里庄Nanlizhuang 	feud.Barony
	前蒲类Qian_pulei 	feud.Barony
}

func (c *阿主孛黑答Aj_bogdCounty) BAj_bogd阿主孛黑答() feud.Barony {
	return c.阿主孛黑答Aj_bogd
}
    
func (c *阿主孛黑答Aj_bogdCounty) BBalikun巴里坤() feud.Barony {
	return c.巴里坤Balikun
}
    
func (c *阿主孛黑答Aj_bogdCounty) BChangcun常村() feud.Barony {
	return c.常村Changcun
}
    
func (c *阿主孛黑答Aj_bogdCounty) BErdene额尔德尼() feud.Barony {
	return c.额尔德尼Erdene
}
    
func (c *阿主孛黑答Aj_bogdCounty) BGurvantes古尔班特斯() feud.Barony {
	return c.古尔班特斯Gurvantes
}
    
func (c *阿主孛黑答Aj_bogdCounty) BNanlizhuang南里庄() feud.Barony {
	return c.南里庄Nanlizhuang
}
    
func (c *阿主孛黑答Aj_bogdCounty) BQian_pulei前蒲类() feud.Barony {
	return c.前蒲类Qian_pulei
}
    
var CAj_bogd阿主孛黑答 Aj_bogdCounty = &阿主孛黑答Aj_bogdCounty{}

func init() {
	f := CAj_bogd阿主孛黑答.(*阿主孛黑答Aj_bogdCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1453",
		Title:     "aj_bogd",
		TitleName: "阿主孛黑答",
		TitleCode: "c_aj_bogd",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿主孛黑答Aj_bogd = BAj_bogd阿主孛黑答
	f.阿主孛黑答Aj_bogd.SetParent(f)
	
	f.巴里坤Balikun = BBalikun巴里坤
	f.巴里坤Balikun.SetParent(f)
	
	f.常村Changcun = BChangcun常村
	f.常村Changcun.SetParent(f)
	
	f.额尔德尼Erdene = BErdene额尔德尼
	f.额尔德尼Erdene.SetParent(f)
	
	f.古尔班特斯Gurvantes = BGurvantes古尔班特斯
	f.古尔班特斯Gurvantes.SetParent(f)
	
	f.南里庄Nanlizhuang = BNanlizhuang南里庄
	f.南里庄Nanlizhuang.SetParent(f)
	
	f.前蒲类Qian_pulei = BQian_pulei前蒲类
	f.前蒲类Qian_pulei.SetParent(f)
	
}
