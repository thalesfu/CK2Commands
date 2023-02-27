package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PettauCounty interface {
    feud.County
    BBistrica比斯特里察() 	feud.Barony
    BFriedau弗里道() 	feud.Barony
    BMarburg_pettau马堡() 	feud.Barony
    BPettau佩陶() 	feud.Barony
    BSeitz赛茨() 	feud.Barony
    BStudenitz施蒂德尼茨() 	feud.Barony
    BWeitenstein魏滕施泰因() 	feud.Barony
}

type 佩陶PettauCounty struct {
	feud.BaseCounty
	比斯特里察Bistrica 	feud.Barony
	弗里道Friedau 	feud.Barony
	马堡Marburg_pettau 	feud.Barony
	佩陶Pettau 	feud.Barony
	赛茨Seitz 	feud.Barony
	施蒂德尼茨Studenitz 	feud.Barony
	魏滕施泰因Weitenstein 	feud.Barony
}

func (c *佩陶PettauCounty) BBistrica比斯特里察() feud.Barony {
	return c.比斯特里察Bistrica
}
    
func (c *佩陶PettauCounty) BFriedau弗里道() feud.Barony {
	return c.弗里道Friedau
}
    
func (c *佩陶PettauCounty) BMarburg_pettau马堡() feud.Barony {
	return c.马堡Marburg_pettau
}
    
func (c *佩陶PettauCounty) BPettau佩陶() feud.Barony {
	return c.佩陶Pettau
}
    
func (c *佩陶PettauCounty) BSeitz赛茨() feud.Barony {
	return c.赛茨Seitz
}
    
func (c *佩陶PettauCounty) BStudenitz施蒂德尼茨() feud.Barony {
	return c.施蒂德尼茨Studenitz
}
    
func (c *佩陶PettauCounty) BWeitenstein魏滕施泰因() feud.Barony {
	return c.魏滕施泰因Weitenstein
}
    
var CPettau佩陶 PettauCounty = &佩陶PettauCounty{}

func init() {
	f := CPettau佩陶.(*佩陶PettauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1691",
		Title:     "pettau",
		TitleName: "佩陶",
		TitleCode: "c_pettau",
		Baronies:  map[string]feud.Barony{},
	}

	f.比斯特里察Bistrica = BBistrica比斯特里察
	f.比斯特里察Bistrica.SetParent(f)
	
	f.弗里道Friedau = BFriedau弗里道
	f.弗里道Friedau.SetParent(f)
	
	f.马堡Marburg_pettau = BMarburg_pettau马堡
	f.马堡Marburg_pettau.SetParent(f)
	
	f.佩陶Pettau = BPettau佩陶
	f.佩陶Pettau.SetParent(f)
	
	f.赛茨Seitz = BSeitz赛茨
	f.赛茨Seitz.SetParent(f)
	
	f.施蒂德尼茨Studenitz = BStudenitz施蒂德尼茨
	f.施蒂德尼茨Studenitz.SetParent(f)
	
	f.魏滕施泰因Weitenstein = BWeitenstein魏滕施泰因
	f.魏滕施泰因Weitenstein.SetParent(f)
	
}
