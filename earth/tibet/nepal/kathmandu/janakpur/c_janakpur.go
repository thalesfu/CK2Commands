package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JanakpurCounty interface {
    feud.County
    BBharatpur婆罗多补罗() 	feud.Barony
    BBirgunj毗罗健阇() 	feud.Barony
    BDevghat提婆伽吒() 	feud.Barony
    BHetauda黑道达() 	feud.Barony
    BJanakpur阇那迦补罗() 	feud.Barony
    BKalaiya格莱亚() 	feud.Barony
    BKamalamai卡马拉迈() 	feud.Barony
}

type 阇那迦补罗JanakpurCounty struct {
	feud.BaseCounty
	婆罗多补罗Bharatpur 	feud.Barony
	毗罗健阇Birgunj 	feud.Barony
	提婆伽吒Devghat 	feud.Barony
	黑道达Hetauda 	feud.Barony
	阇那迦补罗Janakpur 	feud.Barony
	格莱亚Kalaiya 	feud.Barony
	卡马拉迈Kamalamai 	feud.Barony
}

func (c *阇那迦补罗JanakpurCounty) BBharatpur婆罗多补罗() feud.Barony {
	return c.婆罗多补罗Bharatpur
}
    
func (c *阇那迦补罗JanakpurCounty) BBirgunj毗罗健阇() feud.Barony {
	return c.毗罗健阇Birgunj
}
    
func (c *阇那迦补罗JanakpurCounty) BDevghat提婆伽吒() feud.Barony {
	return c.提婆伽吒Devghat
}
    
func (c *阇那迦补罗JanakpurCounty) BHetauda黑道达() feud.Barony {
	return c.黑道达Hetauda
}
    
func (c *阇那迦补罗JanakpurCounty) BJanakpur阇那迦补罗() feud.Barony {
	return c.阇那迦补罗Janakpur
}
    
func (c *阇那迦补罗JanakpurCounty) BKalaiya格莱亚() feud.Barony {
	return c.格莱亚Kalaiya
}
    
func (c *阇那迦补罗JanakpurCounty) BKamalamai卡马拉迈() feud.Barony {
	return c.卡马拉迈Kamalamai
}
    
var CJanakpur阇那迦补罗 JanakpurCounty = &阇那迦补罗JanakpurCounty{}

func init() {
	f := CJanakpur阇那迦补罗.(*阇那迦补罗JanakpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1477",
		Title:     "janakpur",
		TitleName: "阇那迦补罗",
		TitleCode: "c_janakpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗多补罗Bharatpur = BBharatpur婆罗多补罗
	f.婆罗多补罗Bharatpur.SetParent(f)
	
	f.毗罗健阇Birgunj = BBirgunj毗罗健阇
	f.毗罗健阇Birgunj.SetParent(f)
	
	f.提婆伽吒Devghat = BDevghat提婆伽吒
	f.提婆伽吒Devghat.SetParent(f)
	
	f.黑道达Hetauda = BHetauda黑道达
	f.黑道达Hetauda.SetParent(f)
	
	f.阇那迦补罗Janakpur = BJanakpur阇那迦补罗
	f.阇那迦补罗Janakpur.SetParent(f)
	
	f.格莱亚Kalaiya = BKalaiya格莱亚
	f.格莱亚Kalaiya.SetParent(f)
	
	f.卡马拉迈Kamalamai = BKamalamai卡马拉迈
	f.卡马拉迈Kamalamai.SetParent(f)
	
}
