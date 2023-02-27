package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GwentCounty interface {
    feud.County
    BAbergavenny阿伯加文尼() 	feud.Barony
    BCaerleon卡利恩() 	feud.Barony
    BCaerwent凯尔文特() 	feud.Barony
    BChepstow切普斯托() 	feud.Barony
    BMonmouth蒙茅斯() 	feud.Barony
    BNewport纽波特() 	feud.Barony
    BTintern延特恩() 	feud.Barony
}

type 格温特GwentCounty struct {
	feud.BaseCounty
	阿伯加文尼Abergavenny 	feud.Barony
	卡利恩Caerleon 	feud.Barony
	凯尔文特Caerwent 	feud.Barony
	切普斯托Chepstow 	feud.Barony
	蒙茅斯Monmouth 	feud.Barony
	纽波特Newport 	feud.Barony
	延特恩Tintern 	feud.Barony
}

func (c *格温特GwentCounty) BAbergavenny阿伯加文尼() feud.Barony {
	return c.阿伯加文尼Abergavenny
}
    
func (c *格温特GwentCounty) BCaerleon卡利恩() feud.Barony {
	return c.卡利恩Caerleon
}
    
func (c *格温特GwentCounty) BCaerwent凯尔文特() feud.Barony {
	return c.凯尔文特Caerwent
}
    
func (c *格温特GwentCounty) BChepstow切普斯托() feud.Barony {
	return c.切普斯托Chepstow
}
    
func (c *格温特GwentCounty) BMonmouth蒙茅斯() feud.Barony {
	return c.蒙茅斯Monmouth
}
    
func (c *格温特GwentCounty) BNewport纽波特() feud.Barony {
	return c.纽波特Newport
}
    
func (c *格温特GwentCounty) BTintern延特恩() feud.Barony {
	return c.延特恩Tintern
}
    
var CGwent格温特 GwentCounty = &格温特GwentCounty{}

func init() {
	f := CGwent格温特.(*格温特GwentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "20",
		Title:     "gwent",
		TitleName: "格温特",
		TitleCode: "c_gwent",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯加文尼Abergavenny = BAbergavenny阿伯加文尼
	f.阿伯加文尼Abergavenny.SetParent(f)
	
	f.卡利恩Caerleon = BCaerleon卡利恩
	f.卡利恩Caerleon.SetParent(f)
	
	f.凯尔文特Caerwent = BCaerwent凯尔文特
	f.凯尔文特Caerwent.SetParent(f)
	
	f.切普斯托Chepstow = BChepstow切普斯托
	f.切普斯托Chepstow.SetParent(f)
	
	f.蒙茅斯Monmouth = BMonmouth蒙茅斯
	f.蒙茅斯Monmouth.SetParent(f)
	
	f.纽波特Newport = BNewport纽波特
	f.纽波特Newport.SetParent(f)
	
	f.延特恩Tintern = BTintern延特恩
	f.延特恩Tintern.SetParent(f)
	
}
