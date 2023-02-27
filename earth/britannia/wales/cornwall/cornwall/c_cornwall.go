package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CornwallCounty interface {
    feud.County
    BBodmin博德明() 	feud.Barony
    BLaunceston朗斯顿() 	feud.Barony
    BLiskeard利斯卡德() 	feud.Barony
    BRestormel里斯托默尔() 	feud.Barony
    BSt_germans圣杰曼斯() 	feud.Barony
    BSt_michael圣迈克尔() 	feud.Barony
    BTintagel廷塔杰尔() 	feud.Barony
    BTruro特鲁罗() 	feud.Barony
}

type 康沃尔CornwallCounty struct {
	feud.BaseCounty
	博德明Bodmin 	feud.Barony
	朗斯顿Launceston 	feud.Barony
	利斯卡德Liskeard 	feud.Barony
	里斯托默尔Restormel 	feud.Barony
	圣杰曼斯St_germans 	feud.Barony
	圣迈克尔St_michael 	feud.Barony
	廷塔杰尔Tintagel 	feud.Barony
	特鲁罗Truro 	feud.Barony
}

func (c *康沃尔CornwallCounty) BBodmin博德明() feud.Barony {
	return c.博德明Bodmin
}
    
func (c *康沃尔CornwallCounty) BLaunceston朗斯顿() feud.Barony {
	return c.朗斯顿Launceston
}
    
func (c *康沃尔CornwallCounty) BLiskeard利斯卡德() feud.Barony {
	return c.利斯卡德Liskeard
}
    
func (c *康沃尔CornwallCounty) BRestormel里斯托默尔() feud.Barony {
	return c.里斯托默尔Restormel
}
    
func (c *康沃尔CornwallCounty) BSt_germans圣杰曼斯() feud.Barony {
	return c.圣杰曼斯St_germans
}
    
func (c *康沃尔CornwallCounty) BSt_michael圣迈克尔() feud.Barony {
	return c.圣迈克尔St_michael
}
    
func (c *康沃尔CornwallCounty) BTintagel廷塔杰尔() feud.Barony {
	return c.廷塔杰尔Tintagel
}
    
func (c *康沃尔CornwallCounty) BTruro特鲁罗() feud.Barony {
	return c.特鲁罗Truro
}
    
var CCornwall康沃尔 CornwallCounty = &康沃尔CornwallCounty{}

func init() {
	f := CCornwall康沃尔.(*康沃尔CornwallCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "31",
		Title:     "cornwall",
		TitleName: "康沃尔",
		TitleCode: "c_cornwall",
		Baronies:  map[string]feud.Barony{},
	}

	f.博德明Bodmin = BBodmin博德明
	f.博德明Bodmin.SetParent(f)
	
	f.朗斯顿Launceston = BLaunceston朗斯顿
	f.朗斯顿Launceston.SetParent(f)
	
	f.利斯卡德Liskeard = BLiskeard利斯卡德
	f.利斯卡德Liskeard.SetParent(f)
	
	f.里斯托默尔Restormel = BRestormel里斯托默尔
	f.里斯托默尔Restormel.SetParent(f)
	
	f.圣杰曼斯St_germans = BSt_germans圣杰曼斯
	f.圣杰曼斯St_germans.SetParent(f)
	
	f.圣迈克尔St_michael = BSt_michael圣迈克尔
	f.圣迈克尔St_michael.SetParent(f)
	
	f.廷塔杰尔Tintagel = BTintagel廷塔杰尔
	f.廷塔杰尔Tintagel.SetParent(f)
	
	f.特鲁罗Truro = BTruro特鲁罗
	f.特鲁罗Truro.SetParent(f)
	
}
