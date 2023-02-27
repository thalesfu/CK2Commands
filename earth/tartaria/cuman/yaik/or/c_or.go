package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrCounty interface {
    feud.County
    BAschelisay阿谢利塞() 	feud.Barony
    BIstemis伊斯捷米斯() 	feud.Barony
    BMamyt马梅特() 	feud.Barony
    BMaytobe迈托别() 	feud.Barony
    BOr奥尔() 	feud.Barony
    BOrsk_or奥尔斯克() 	feud.Barony
    BStepnoe斯捷普诺耶() 	feud.Barony
}

type 奥尔OrCounty struct {
	feud.BaseCounty
	阿谢利塞Aschelisay 	feud.Barony
	伊斯捷米斯Istemis 	feud.Barony
	马梅特Mamyt 	feud.Barony
	迈托别Maytobe 	feud.Barony
	奥尔Or 	feud.Barony
	奥尔斯克Orsk_or 	feud.Barony
	斯捷普诺耶Stepnoe 	feud.Barony
}

func (c *奥尔OrCounty) BAschelisay阿谢利塞() feud.Barony {
	return c.阿谢利塞Aschelisay
}
    
func (c *奥尔OrCounty) BIstemis伊斯捷米斯() feud.Barony {
	return c.伊斯捷米斯Istemis
}
    
func (c *奥尔OrCounty) BMamyt马梅特() feud.Barony {
	return c.马梅特Mamyt
}
    
func (c *奥尔OrCounty) BMaytobe迈托别() feud.Barony {
	return c.迈托别Maytobe
}
    
func (c *奥尔OrCounty) BOr奥尔() feud.Barony {
	return c.奥尔Or
}
    
func (c *奥尔OrCounty) BOrsk_or奥尔斯克() feud.Barony {
	return c.奥尔斯克Orsk_or
}
    
func (c *奥尔OrCounty) BStepnoe斯捷普诺耶() feud.Barony {
	return c.斯捷普诺耶Stepnoe
}
    
var COr奥尔 OrCounty = &奥尔OrCounty{}

func init() {
	f := COr奥尔.(*奥尔OrCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1850",
		Title:     "or",
		TitleName: "奥尔",
		TitleCode: "c_or",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿谢利塞Aschelisay = BAschelisay阿谢利塞
	f.阿谢利塞Aschelisay.SetParent(f)
	
	f.伊斯捷米斯Istemis = BIstemis伊斯捷米斯
	f.伊斯捷米斯Istemis.SetParent(f)
	
	f.马梅特Mamyt = BMamyt马梅特
	f.马梅特Mamyt.SetParent(f)
	
	f.迈托别Maytobe = BMaytobe迈托别
	f.迈托别Maytobe.SetParent(f)
	
	f.奥尔Or = BOr奥尔
	f.奥尔Or.SetParent(f)
	
	f.奥尔斯克Orsk_or = BOrsk_or奥尔斯克
	f.奥尔斯克Orsk_or.SetParent(f)
	
	f.斯捷普诺耶Stepnoe = BStepnoe斯捷普诺耶
	f.斯捷普诺耶Stepnoe.SetParent(f)
	
}
