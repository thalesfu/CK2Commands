package beloozero

import (
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero/bezichy"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero/chlisselbourg"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero/tikhvine"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero/torzhok"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero/zaozerye"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeloozeroDuke interface {
    feud.Duke
    CBezichy别日奇() 	bezichy.BezichyCounty
    CChlisselbourg什利谢利堡() 	chlisselbourg.ChlisselbourgCounty
    CTikhvine季赫温() 	tikhvine.TikhvineCounty
    CTorzhok托尔若克() 	torzhok.TorzhokCounty
    CZaozerye扎奥焦里耶() 	zaozerye.ZaozeryeCounty
}

type 韦普斯BeloozeroDuke struct {
	feud.BaseDuke
	别日奇Bezichy 	bezichy.BezichyCounty
	什利谢利堡Chlisselbourg 	chlisselbourg.ChlisselbourgCounty
	季赫温Tikhvine 	tikhvine.TikhvineCounty
	托尔若克Torzhok 	torzhok.TorzhokCounty
	扎奥焦里耶Zaozerye 	zaozerye.ZaozeryeCounty
}

func (d *韦普斯BeloozeroDuke) CBezichy别日奇() bezichy.BezichyCounty {
	return d.别日奇Bezichy
}
    
func (d *韦普斯BeloozeroDuke) CChlisselbourg什利谢利堡() chlisselbourg.ChlisselbourgCounty {
	return d.什利谢利堡Chlisselbourg
}
    
func (d *韦普斯BeloozeroDuke) CTikhvine季赫温() tikhvine.TikhvineCounty {
	return d.季赫温Tikhvine
}
    
func (d *韦普斯BeloozeroDuke) CTorzhok托尔若克() torzhok.TorzhokCounty {
	return d.托尔若克Torzhok
}
    
func (d *韦普斯BeloozeroDuke) CZaozerye扎奥焦里耶() zaozerye.ZaozeryeCounty {
	return d.扎奥焦里耶Zaozerye
}
    
var DBeloozero韦普斯 BeloozeroDuke = &韦普斯BeloozeroDuke{}

func init() {
	f := DBeloozero韦普斯.(*韦普斯BeloozeroDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "beloozero",
		TitleName: "韦普斯",
		TitleCode: "d_beloozero",
		Counties:  map[string]feud.County{},
	}

	f.别日奇Bezichy = bezichy.CBezichy别日奇
	f.别日奇Bezichy.SetParent(f)
	
	f.什利谢利堡Chlisselbourg = chlisselbourg.CChlisselbourg什利谢利堡
	f.什利谢利堡Chlisselbourg.SetParent(f)
	
	f.季赫温Tikhvine = tikhvine.CTikhvine季赫温
	f.季赫温Tikhvine.SetParent(f)
	
	f.托尔若克Torzhok = torzhok.CTorzhok托尔若克
	f.托尔若克Torzhok.SetParent(f)
	
	f.扎奥焦里耶Zaozerye = zaozerye.CZaozerye扎奥焦里耶
	f.扎奥焦里耶Zaozerye.SetParent(f)
	
}
