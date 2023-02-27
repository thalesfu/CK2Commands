package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FesCounty interface {
    feud.County
    BBouanane布瓦南() 	feud.Barony
    BElhajeb哈杰卜() 	feud.Barony
    BEnjil因吉勒() 	feud.Barony
    BFes非斯() 	feud.Barony
    BIzerhan伊泽尔昂() 	feud.Barony
    BSubu苏布() 	feud.Barony
    BZerhoun扎霍() 	feud.Barony
}

type 非斯FesCounty struct {
	feud.BaseCounty
	布瓦南Bouanane 	feud.Barony
	哈杰卜Elhajeb 	feud.Barony
	因吉勒Enjil 	feud.Barony
	非斯Fes 	feud.Barony
	伊泽尔昂Izerhan 	feud.Barony
	苏布Subu 	feud.Barony
	扎霍Zerhoun 	feud.Barony
}

func (c *非斯FesCounty) BBouanane布瓦南() feud.Barony {
	return c.布瓦南Bouanane
}
    
func (c *非斯FesCounty) BElhajeb哈杰卜() feud.Barony {
	return c.哈杰卜Elhajeb
}
    
func (c *非斯FesCounty) BEnjil因吉勒() feud.Barony {
	return c.因吉勒Enjil
}
    
func (c *非斯FesCounty) BFes非斯() feud.Barony {
	return c.非斯Fes
}
    
func (c *非斯FesCounty) BIzerhan伊泽尔昂() feud.Barony {
	return c.伊泽尔昂Izerhan
}
    
func (c *非斯FesCounty) BSubu苏布() feud.Barony {
	return c.苏布Subu
}
    
func (c *非斯FesCounty) BZerhoun扎霍() feud.Barony {
	return c.扎霍Zerhoun
}
    
var CFes非斯 FesCounty = &非斯FesCounty{}

func init() {
	f := CFes非斯.(*非斯FesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "840",
		Title:     "fes",
		TitleName: "非斯",
		TitleCode: "c_fes",
		Baronies:  map[string]feud.Barony{},
	}

	f.布瓦南Bouanane = BBouanane布瓦南
	f.布瓦南Bouanane.SetParent(f)
	
	f.哈杰卜Elhajeb = BElhajeb哈杰卜
	f.哈杰卜Elhajeb.SetParent(f)
	
	f.因吉勒Enjil = BEnjil因吉勒
	f.因吉勒Enjil.SetParent(f)
	
	f.非斯Fes = BFes非斯
	f.非斯Fes.SetParent(f)
	
	f.伊泽尔昂Izerhan = BIzerhan伊泽尔昂
	f.伊泽尔昂Izerhan.SetParent(f)
	
	f.苏布Subu = BSubu苏布
	f.苏布Subu.SetParent(f)
	
	f.扎霍Zerhoun = BZerhoun扎霍
	f.扎霍Zerhoun.SetParent(f)
	
}
