package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MontbeliardCounty interface {
    feud.County
    BArbouans阿尔布昂() 	feud.Barony
    BBavans巴旺() 	feud.Barony
    BEtupes埃蒂佩() 	feud.Barony
    BExincourt埃克桑库尔() 	feud.Barony
    BGrand_charmont大沙尔蒙() 	feud.Barony
    BLayre莱尔() 	feud.Barony
    BMontbeliard蒙贝利亚尔() 	feud.Barony
    BValentigney瓦朗蒂涅() 	feud.Barony
}

type 蒙贝利亚尔MontbeliardCounty struct {
	feud.BaseCounty
	阿尔布昂Arbouans 	feud.Barony
	巴旺Bavans 	feud.Barony
	埃蒂佩Etupes 	feud.Barony
	埃克桑库尔Exincourt 	feud.Barony
	大沙尔蒙Grand_charmont 	feud.Barony
	莱尔Layre 	feud.Barony
	蒙贝利亚尔Montbeliard 	feud.Barony
	瓦朗蒂涅Valentigney 	feud.Barony
}

func (c *蒙贝利亚尔MontbeliardCounty) BArbouans阿尔布昂() feud.Barony {
	return c.阿尔布昂Arbouans
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BBavans巴旺() feud.Barony {
	return c.巴旺Bavans
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BEtupes埃蒂佩() feud.Barony {
	return c.埃蒂佩Etupes
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BExincourt埃克桑库尔() feud.Barony {
	return c.埃克桑库尔Exincourt
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BGrand_charmont大沙尔蒙() feud.Barony {
	return c.大沙尔蒙Grand_charmont
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BLayre莱尔() feud.Barony {
	return c.莱尔Layre
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BMontbeliard蒙贝利亚尔() feud.Barony {
	return c.蒙贝利亚尔Montbeliard
}
    
func (c *蒙贝利亚尔MontbeliardCounty) BValentigney瓦朗蒂涅() feud.Barony {
	return c.瓦朗蒂涅Valentigney
}
    
var CMontbeliard蒙贝利亚尔 MontbeliardCounty = &蒙贝利亚尔MontbeliardCounty{}

func init() {
	f := CMontbeliard蒙贝利亚尔.(*蒙贝利亚尔MontbeliardCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1608",
		Title:     "montbeliard",
		TitleName: "蒙贝利亚尔",
		TitleCode: "c_montbeliard",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔布昂Arbouans = BArbouans阿尔布昂
	f.阿尔布昂Arbouans.SetParent(f)
	
	f.巴旺Bavans = BBavans巴旺
	f.巴旺Bavans.SetParent(f)
	
	f.埃蒂佩Etupes = BEtupes埃蒂佩
	f.埃蒂佩Etupes.SetParent(f)
	
	f.埃克桑库尔Exincourt = BExincourt埃克桑库尔
	f.埃克桑库尔Exincourt.SetParent(f)
	
	f.大沙尔蒙Grand_charmont = BGrand_charmont大沙尔蒙
	f.大沙尔蒙Grand_charmont.SetParent(f)
	
	f.莱尔Layre = BLayre莱尔
	f.莱尔Layre.SetParent(f)
	
	f.蒙贝利亚尔Montbeliard = BMontbeliard蒙贝利亚尔
	f.蒙贝利亚尔Montbeliard.SetParent(f)
	
	f.瓦朗蒂涅Valentigney = BValentigney瓦朗蒂涅
	f.瓦朗蒂涅Valentigney.SetParent(f)
	
}
