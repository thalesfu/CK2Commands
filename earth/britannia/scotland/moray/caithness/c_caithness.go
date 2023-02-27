package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CaithnessCounty interface {
    feud.County
    BDornoch多诺赫() 	feud.Barony
    BDunbeath邓比斯() 	feud.Barony
    BDunrobin邓罗宾() 	feud.Barony
    BFreswick弗雷西克() 	feud.Barony
    BGolspie戈尔斯皮() 	feud.Barony
    BLatheron拉瑟伦() 	feud.Barony
    BThurso瑟索() 	feud.Barony
    BWick威克() 	feud.Barony
}

type 凯斯内斯CaithnessCounty struct {
	feud.BaseCounty
	多诺赫Dornoch 	feud.Barony
	邓比斯Dunbeath 	feud.Barony
	邓罗宾Dunrobin 	feud.Barony
	弗雷西克Freswick 	feud.Barony
	戈尔斯皮Golspie 	feud.Barony
	拉瑟伦Latheron 	feud.Barony
	瑟索Thurso 	feud.Barony
	威克Wick 	feud.Barony
}

func (c *凯斯内斯CaithnessCounty) BDornoch多诺赫() feud.Barony {
	return c.多诺赫Dornoch
}
    
func (c *凯斯内斯CaithnessCounty) BDunbeath邓比斯() feud.Barony {
	return c.邓比斯Dunbeath
}
    
func (c *凯斯内斯CaithnessCounty) BDunrobin邓罗宾() feud.Barony {
	return c.邓罗宾Dunrobin
}
    
func (c *凯斯内斯CaithnessCounty) BFreswick弗雷西克() feud.Barony {
	return c.弗雷西克Freswick
}
    
func (c *凯斯内斯CaithnessCounty) BGolspie戈尔斯皮() feud.Barony {
	return c.戈尔斯皮Golspie
}
    
func (c *凯斯内斯CaithnessCounty) BLatheron拉瑟伦() feud.Barony {
	return c.拉瑟伦Latheron
}
    
func (c *凯斯内斯CaithnessCounty) BThurso瑟索() feud.Barony {
	return c.瑟索Thurso
}
    
func (c *凯斯内斯CaithnessCounty) BWick威克() feud.Barony {
	return c.威克Wick
}
    
var CCaithness凯斯内斯 CaithnessCounty = &凯斯内斯CaithnessCounty{}

func init() {
	f := CCaithness凯斯内斯.(*凯斯内斯CaithnessCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "37",
		Title:     "caithness",
		TitleName: "凯斯内斯",
		TitleCode: "c_caithness",
		Baronies:  map[string]feud.Barony{},
	}

	f.多诺赫Dornoch = BDornoch多诺赫
	f.多诺赫Dornoch.SetParent(f)
	
	f.邓比斯Dunbeath = BDunbeath邓比斯
	f.邓比斯Dunbeath.SetParent(f)
	
	f.邓罗宾Dunrobin = BDunrobin邓罗宾
	f.邓罗宾Dunrobin.SetParent(f)
	
	f.弗雷西克Freswick = BFreswick弗雷西克
	f.弗雷西克Freswick.SetParent(f)
	
	f.戈尔斯皮Golspie = BGolspie戈尔斯皮
	f.戈尔斯皮Golspie.SetParent(f)
	
	f.拉瑟伦Latheron = BLatheron拉瑟伦
	f.拉瑟伦Latheron.SetParent(f)
	
	f.瑟索Thurso = BThurso瑟索
	f.瑟索Thurso.SetParent(f)
	
	f.威克Wick = BWick威克
	f.威克Wick.SetParent(f)
	
}
