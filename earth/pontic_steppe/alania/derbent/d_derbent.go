package derbent

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/derbent/balanjar"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/derbent/derbent"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/derbent/durdzukia"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/derbent/semender"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DerbentDuke interface {
    feud.Duke
    CBalanjar巴兰贾尔() 	balanjar.BalanjarCounty
    CDerbent杰尔宾特() 	derbent.DerbentCounty
    CDurdzukia杜尔祖基亚() 	durdzukia.DurdzukiaCounty
    CSemender谢缅杰尔() 	semender.SemenderCounty
}

type 杰尔宾特DerbentDuke struct {
	feud.BaseDuke
	巴兰贾尔Balanjar 	balanjar.BalanjarCounty
	杰尔宾特Derbent 	derbent.DerbentCounty
	杜尔祖基亚Durdzukia 	durdzukia.DurdzukiaCounty
	谢缅杰尔Semender 	semender.SemenderCounty
}

func (d *杰尔宾特DerbentDuke) CBalanjar巴兰贾尔() balanjar.BalanjarCounty {
	return d.巴兰贾尔Balanjar
}
    
func (d *杰尔宾特DerbentDuke) CDerbent杰尔宾特() derbent.DerbentCounty {
	return d.杰尔宾特Derbent
}
    
func (d *杰尔宾特DerbentDuke) CDurdzukia杜尔祖基亚() durdzukia.DurdzukiaCounty {
	return d.杜尔祖基亚Durdzukia
}
    
func (d *杰尔宾特DerbentDuke) CSemender谢缅杰尔() semender.SemenderCounty {
	return d.谢缅杰尔Semender
}
    
var DDerbent杰尔宾特 DerbentDuke = &杰尔宾特DerbentDuke{}

func init() {
	f := DDerbent杰尔宾特.(*杰尔宾特DerbentDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "derbent",
		TitleName: "杰尔宾特",
		TitleCode: "d_derbent",
		Counties:  map[string]feud.County{},
	}

	f.巴兰贾尔Balanjar = balanjar.CBalanjar巴兰贾尔
	f.巴兰贾尔Balanjar.SetParent(f)
	
	f.杰尔宾特Derbent = derbent.CDerbent杰尔宾特
	f.杰尔宾特Derbent.SetParent(f)
	
	f.杜尔祖基亚Durdzukia = durdzukia.CDurdzukia杜尔祖基亚
	f.杜尔祖基亚Durdzukia.SetParent(f)
	
	f.谢缅杰尔Semender = semender.CSemender谢缅杰尔
	f.谢缅杰尔Semender.SetParent(f)
	
}
