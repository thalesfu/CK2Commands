package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GarhwalCounty interface {
    feud.County
    BBadrinath婆陀梨那他() 	feud.Barony
    BDevalghar提伐迦() 	feud.Barony
    BHaridwar诃梨堕罗() 	feud.Barony
    BSrinagar室利那揭罗() 	feud.Barony
    BSudhnagar苏提那迦利() 	feud.Barony
    BTehri代赫里() 	feud.Barony
    BUttarkashi北迦尸() 	feud.Barony
}

type 真日GarhwalCounty struct {
	feud.BaseCounty
	婆陀梨那他Badrinath 	feud.Barony
	提伐迦Devalghar 	feud.Barony
	诃梨堕罗Haridwar 	feud.Barony
	室利那揭罗Srinagar 	feud.Barony
	苏提那迦利Sudhnagar 	feud.Barony
	代赫里Tehri 	feud.Barony
	北迦尸Uttarkashi 	feud.Barony
}

func (c *真日GarhwalCounty) BBadrinath婆陀梨那他() feud.Barony {
	return c.婆陀梨那他Badrinath
}
    
func (c *真日GarhwalCounty) BDevalghar提伐迦() feud.Barony {
	return c.提伐迦Devalghar
}
    
func (c *真日GarhwalCounty) BHaridwar诃梨堕罗() feud.Barony {
	return c.诃梨堕罗Haridwar
}
    
func (c *真日GarhwalCounty) BSrinagar室利那揭罗() feud.Barony {
	return c.室利那揭罗Srinagar
}
    
func (c *真日GarhwalCounty) BSudhnagar苏提那迦利() feud.Barony {
	return c.苏提那迦利Sudhnagar
}
    
func (c *真日GarhwalCounty) BTehri代赫里() feud.Barony {
	return c.代赫里Tehri
}
    
func (c *真日GarhwalCounty) BUttarkashi北迦尸() feud.Barony {
	return c.北迦尸Uttarkashi
}
    
var CGarhwal真日 GarhwalCounty = &真日GarhwalCounty{}

func init() {
	f := CGarhwal真日.(*真日GarhwalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1470",
		Title:     "garhwal",
		TitleName: "真日",
		TitleCode: "c_garhwal",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆陀梨那他Badrinath = BBadrinath婆陀梨那他
	f.婆陀梨那他Badrinath.SetParent(f)
	
	f.提伐迦Devalghar = BDevalghar提伐迦
	f.提伐迦Devalghar.SetParent(f)
	
	f.诃梨堕罗Haridwar = BHaridwar诃梨堕罗
	f.诃梨堕罗Haridwar.SetParent(f)
	
	f.室利那揭罗Srinagar = BSrinagar室利那揭罗
	f.室利那揭罗Srinagar.SetParent(f)
	
	f.苏提那迦利Sudhnagar = BSudhnagar苏提那迦利
	f.苏提那迦利Sudhnagar.SetParent(f)
	
	f.代赫里Tehri = BTehri代赫里
	f.代赫里Tehri.SetParent(f)
	
	f.北迦尸Uttarkashi = BUttarkashi北迦尸
	f.北迦尸Uttarkashi.SetParent(f)
	
}
