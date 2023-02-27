package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DipalpurCounty interface {
    feud.County
    BAbohar阿博赫尔() 	feud.Barony
    BAjodhan阿殊檀() 	feud.Barony
    BChichawatni奇查瓦尼() 	feud.Barony
    BDipalpur提波罗补罗() 	feud.Barony
    BJongbandha树俱般陀() 	feud.Barony
    BPancapura般遮补罗() 	feud.Barony
    BSatghara萨特迦拉() 	feud.Barony
}

type 提波罗补罗DipalpurCounty struct {
	feud.BaseCounty
	阿博赫尔Abohar 	feud.Barony
	阿殊檀Ajodhan 	feud.Barony
	奇查瓦尼Chichawatni 	feud.Barony
	提波罗补罗Dipalpur 	feud.Barony
	树俱般陀Jongbandha 	feud.Barony
	般遮补罗Pancapura 	feud.Barony
	萨特迦拉Satghara 	feud.Barony
}

func (c *提波罗补罗DipalpurCounty) BAbohar阿博赫尔() feud.Barony {
	return c.阿博赫尔Abohar
}
    
func (c *提波罗补罗DipalpurCounty) BAjodhan阿殊檀() feud.Barony {
	return c.阿殊檀Ajodhan
}
    
func (c *提波罗补罗DipalpurCounty) BChichawatni奇查瓦尼() feud.Barony {
	return c.奇查瓦尼Chichawatni
}
    
func (c *提波罗补罗DipalpurCounty) BDipalpur提波罗补罗() feud.Barony {
	return c.提波罗补罗Dipalpur
}
    
func (c *提波罗补罗DipalpurCounty) BJongbandha树俱般陀() feud.Barony {
	return c.树俱般陀Jongbandha
}
    
func (c *提波罗补罗DipalpurCounty) BPancapura般遮补罗() feud.Barony {
	return c.般遮补罗Pancapura
}
    
func (c *提波罗补罗DipalpurCounty) BSatghara萨特迦拉() feud.Barony {
	return c.萨特迦拉Satghara
}
    
var CDipalpur提波罗补罗 DipalpurCounty = &提波罗补罗DipalpurCounty{}

func init() {
	f := CDipalpur提波罗补罗.(*提波罗补罗DipalpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1363",
		Title:     "dipalpur",
		TitleName: "提波罗补罗",
		TitleCode: "c_dipalpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿博赫尔Abohar = BAbohar阿博赫尔
	f.阿博赫尔Abohar.SetParent(f)
	
	f.阿殊檀Ajodhan = BAjodhan阿殊檀
	f.阿殊檀Ajodhan.SetParent(f)
	
	f.奇查瓦尼Chichawatni = BChichawatni奇查瓦尼
	f.奇查瓦尼Chichawatni.SetParent(f)
	
	f.提波罗补罗Dipalpur = BDipalpur提波罗补罗
	f.提波罗补罗Dipalpur.SetParent(f)
	
	f.树俱般陀Jongbandha = BJongbandha树俱般陀
	f.树俱般陀Jongbandha.SetParent(f)
	
	f.般遮补罗Pancapura = BPancapura般遮补罗
	f.般遮补罗Pancapura.SetParent(f)
	
	f.萨特迦拉Satghara = BSatghara萨特迦拉
	f.萨特迦拉Satghara.SetParent(f)
	
}
