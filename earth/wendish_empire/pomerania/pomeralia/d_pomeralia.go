package pomeralia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pomeralia/bytow"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pomeralia/danzig"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pomeralia/slupsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PomeraliaDuke interface {
    feud.Duke
    CBytow贝图夫() 	bytow.BytowCounty
    CDanzig但泽() 	danzig.DanzigCounty
    CSlupsk斯武普斯克() 	slupsk.SlupskCounty
}

type 波米拉利亚PomeraliaDuke struct {
	feud.BaseDuke
	贝图夫Bytow 	bytow.BytowCounty
	但泽Danzig 	danzig.DanzigCounty
	斯武普斯克Slupsk 	slupsk.SlupskCounty
}

func (d *波米拉利亚PomeraliaDuke) CBytow贝图夫() bytow.BytowCounty {
	return d.贝图夫Bytow
}
    
func (d *波米拉利亚PomeraliaDuke) CDanzig但泽() danzig.DanzigCounty {
	return d.但泽Danzig
}
    
func (d *波米拉利亚PomeraliaDuke) CSlupsk斯武普斯克() slupsk.SlupskCounty {
	return d.斯武普斯克Slupsk
}
    
var DPomeralia波米拉利亚 PomeraliaDuke = &波米拉利亚PomeraliaDuke{}

func init() {
	f := DPomeralia波米拉利亚.(*波米拉利亚PomeraliaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pomeralia",
		TitleName: "波米拉利亚",
		TitleCode: "d_pomeralia",
		Counties:  map[string]feud.County{},
	}

	f.贝图夫Bytow = bytow.CBytow贝图夫
	f.贝图夫Bytow.SetParent(f)
	
	f.但泽Danzig = danzig.CDanzig但泽
	f.但泽Danzig.SetParent(f)
	
	f.斯武普斯克Slupsk = slupsk.CSlupsk斯武普斯克
	f.斯武普斯克Slupsk.SetParent(f)
	
}
