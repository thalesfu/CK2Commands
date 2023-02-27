package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RoslavlCounty interface {
    feud.County
    BAleksandr亚历山大() 	feud.Barony
    BChocimsk霍季姆斯克() 	feud.Barony
    BImaslavl伊马斯拉夫尔() 	feud.Barony
    BKankol坎科尔() 	feud.Barony
    BPochinok波奇诺克() 	feud.Barony
    BRoslavl罗斯拉夫尔() 	feud.Barony
    BVugimsk武吉姆斯克() 	feud.Barony
}

type 罗斯拉夫尔RoslavlCounty struct {
	feud.BaseCounty
	亚历山大Aleksandr 	feud.Barony
	霍季姆斯克Chocimsk 	feud.Barony
	伊马斯拉夫尔Imaslavl 	feud.Barony
	坎科尔Kankol 	feud.Barony
	波奇诺克Pochinok 	feud.Barony
	罗斯拉夫尔Roslavl 	feud.Barony
	武吉姆斯克Vugimsk 	feud.Barony
}

func (c *罗斯拉夫尔RoslavlCounty) BAleksandr亚历山大() feud.Barony {
	return c.亚历山大Aleksandr
}
    
func (c *罗斯拉夫尔RoslavlCounty) BChocimsk霍季姆斯克() feud.Barony {
	return c.霍季姆斯克Chocimsk
}
    
func (c *罗斯拉夫尔RoslavlCounty) BImaslavl伊马斯拉夫尔() feud.Barony {
	return c.伊马斯拉夫尔Imaslavl
}
    
func (c *罗斯拉夫尔RoslavlCounty) BKankol坎科尔() feud.Barony {
	return c.坎科尔Kankol
}
    
func (c *罗斯拉夫尔RoslavlCounty) BPochinok波奇诺克() feud.Barony {
	return c.波奇诺克Pochinok
}
    
func (c *罗斯拉夫尔RoslavlCounty) BRoslavl罗斯拉夫尔() feud.Barony {
	return c.罗斯拉夫尔Roslavl
}
    
func (c *罗斯拉夫尔RoslavlCounty) BVugimsk武吉姆斯克() feud.Barony {
	return c.武吉姆斯克Vugimsk
}
    
var CRoslavl罗斯拉夫尔 RoslavlCounty = &罗斯拉夫尔RoslavlCounty{}

func init() {
	f := CRoslavl罗斯拉夫尔.(*罗斯拉夫尔RoslavlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "933",
		Title:     "roslavl",
		TitleName: "罗斯拉夫尔",
		TitleCode: "c_roslavl",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚历山大Aleksandr = BAleksandr亚历山大
	f.亚历山大Aleksandr.SetParent(f)
	
	f.霍季姆斯克Chocimsk = BChocimsk霍季姆斯克
	f.霍季姆斯克Chocimsk.SetParent(f)
	
	f.伊马斯拉夫尔Imaslavl = BImaslavl伊马斯拉夫尔
	f.伊马斯拉夫尔Imaslavl.SetParent(f)
	
	f.坎科尔Kankol = BKankol坎科尔
	f.坎科尔Kankol.SetParent(f)
	
	f.波奇诺克Pochinok = BPochinok波奇诺克
	f.波奇诺克Pochinok.SetParent(f)
	
	f.罗斯拉夫尔Roslavl = BRoslavl罗斯拉夫尔
	f.罗斯拉夫尔Roslavl.SetParent(f)
	
	f.武吉姆斯克Vugimsk = BVugimsk武吉姆斯克
	f.武吉姆斯克Vugimsk.SetParent(f)
	
}
