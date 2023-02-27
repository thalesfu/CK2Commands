package athens

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/athens/atheniai"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/athens/demetrias"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/athens/hellas"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/athens/thessalia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AthensDuke interface {
    feud.Duke
    CAtheniai雅典() 	atheniai.AtheniaiCounty
    CDemetrias德米特里阿斯() 	demetrias.DemetriasCounty
    CHellas底比斯() 	hellas.HellasCounty
    CThessalia色萨利() 	thessalia.ThessaliaCounty
}

type 希腊AthensDuke struct {
	feud.BaseDuke
	雅典Atheniai 	atheniai.AtheniaiCounty
	德米特里阿斯Demetrias 	demetrias.DemetriasCounty
	底比斯Hellas 	hellas.HellasCounty
	色萨利Thessalia 	thessalia.ThessaliaCounty
}

func (d *希腊AthensDuke) CAtheniai雅典() atheniai.AtheniaiCounty {
	return d.雅典Atheniai
}
    
func (d *希腊AthensDuke) CDemetrias德米特里阿斯() demetrias.DemetriasCounty {
	return d.德米特里阿斯Demetrias
}
    
func (d *希腊AthensDuke) CHellas底比斯() hellas.HellasCounty {
	return d.底比斯Hellas
}
    
func (d *希腊AthensDuke) CThessalia色萨利() thessalia.ThessaliaCounty {
	return d.色萨利Thessalia
}
    
var DAthens希腊 AthensDuke = &希腊AthensDuke{}

func init() {
	f := DAthens希腊.(*希腊AthensDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "athens",
		TitleName: "希腊",
		TitleCode: "d_athens",
		Counties:  map[string]feud.County{},
	}

	f.雅典Atheniai = atheniai.CAtheniai雅典
	f.雅典Atheniai.SetParent(f)
	
	f.德米特里阿斯Demetrias = demetrias.CDemetrias德米特里阿斯
	f.德米特里阿斯Demetrias.SetParent(f)
	
	f.底比斯Hellas = hellas.CHellas底比斯
	f.底比斯Hellas.SetParent(f)
	
	f.色萨利Thessalia = thessalia.CThessalia色萨利
	f.色萨利Thessalia.SetParent(f)
	
}
