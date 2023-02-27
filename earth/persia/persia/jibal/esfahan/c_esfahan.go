package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EsfahanCounty interface {
    feud.County
    BAbyaneh阿卜亚内赫() 	feud.Barony
    BArdestan阿尔德斯坦() 	feud.Barony
    BEsfahan伊斯法罕() 	feud.Barony
    BKashan卡尚() 	feud.Barony
    BKhansar汉萨尔() 	feud.Barony
    BQomsheh恰什哈() 	feud.Barony
    BSedeh塞代() 	feud.Barony
    BZarrinshahr扎林沙赫尔() 	feud.Barony
}

type 伊斯法罕EsfahanCounty struct {
	feud.BaseCounty
	阿卜亚内赫Abyaneh 	feud.Barony
	阿尔德斯坦Ardestan 	feud.Barony
	伊斯法罕Esfahan 	feud.Barony
	卡尚Kashan 	feud.Barony
	汉萨尔Khansar 	feud.Barony
	恰什哈Qomsheh 	feud.Barony
	塞代Sedeh 	feud.Barony
	扎林沙赫尔Zarrinshahr 	feud.Barony
}

func (c *伊斯法罕EsfahanCounty) BAbyaneh阿卜亚内赫() feud.Barony {
	return c.阿卜亚内赫Abyaneh
}
    
func (c *伊斯法罕EsfahanCounty) BArdestan阿尔德斯坦() feud.Barony {
	return c.阿尔德斯坦Ardestan
}
    
func (c *伊斯法罕EsfahanCounty) BEsfahan伊斯法罕() feud.Barony {
	return c.伊斯法罕Esfahan
}
    
func (c *伊斯法罕EsfahanCounty) BKashan卡尚() feud.Barony {
	return c.卡尚Kashan
}
    
func (c *伊斯法罕EsfahanCounty) BKhansar汉萨尔() feud.Barony {
	return c.汉萨尔Khansar
}
    
func (c *伊斯法罕EsfahanCounty) BQomsheh恰什哈() feud.Barony {
	return c.恰什哈Qomsheh
}
    
func (c *伊斯法罕EsfahanCounty) BSedeh塞代() feud.Barony {
	return c.塞代Sedeh
}
    
func (c *伊斯法罕EsfahanCounty) BZarrinshahr扎林沙赫尔() feud.Barony {
	return c.扎林沙赫尔Zarrinshahr
}
    
var CEsfahan伊斯法罕 EsfahanCounty = &伊斯法罕EsfahanCounty{}

func init() {
	f := CEsfahan伊斯法罕.(*伊斯法罕EsfahanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "646",
		Title:     "esfahan",
		TitleName: "伊斯法罕",
		TitleCode: "c_esfahan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜亚内赫Abyaneh = BAbyaneh阿卜亚内赫
	f.阿卜亚内赫Abyaneh.SetParent(f)
	
	f.阿尔德斯坦Ardestan = BArdestan阿尔德斯坦
	f.阿尔德斯坦Ardestan.SetParent(f)
	
	f.伊斯法罕Esfahan = BEsfahan伊斯法罕
	f.伊斯法罕Esfahan.SetParent(f)
	
	f.卡尚Kashan = BKashan卡尚
	f.卡尚Kashan.SetParent(f)
	
	f.汉萨尔Khansar = BKhansar汉萨尔
	f.汉萨尔Khansar.SetParent(f)
	
	f.恰什哈Qomsheh = BQomsheh恰什哈
	f.恰什哈Qomsheh.SetParent(f)
	
	f.塞代Sedeh = BSedeh塞代
	f.塞代Sedeh.SetParent(f)
	
	f.扎林沙赫尔Zarrinshahr = BZarrinshahr扎林沙赫尔
	f.扎林沙赫尔Zarrinshahr.SetParent(f)
	
}
