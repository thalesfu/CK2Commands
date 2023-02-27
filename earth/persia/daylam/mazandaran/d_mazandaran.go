package mazandaran

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/mazandaran/alamut"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/mazandaran/gurgan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/mazandaran/qwivir"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/mazandaran/tabaristan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MazandaranDuke interface {
    feud.Duke
    CAlamut阿拉穆特() 	alamut.AlamutCounty
    CGurgan戈尔甘() 	gurgan.GurganCounty
    CQwivir奎维尔() 	qwivir.QwivirCounty
    CTabaristan陀拔斯单() 	tabaristan.TabaristanCounty
}

type 陀拔斯单MazandaranDuke struct {
	feud.BaseDuke
	阿拉穆特Alamut 	alamut.AlamutCounty
	戈尔甘Gurgan 	gurgan.GurganCounty
	奎维尔Qwivir 	qwivir.QwivirCounty
	陀拔斯单Tabaristan 	tabaristan.TabaristanCounty
}

func (d *陀拔斯单MazandaranDuke) CAlamut阿拉穆特() alamut.AlamutCounty {
	return d.阿拉穆特Alamut
}
    
func (d *陀拔斯单MazandaranDuke) CGurgan戈尔甘() gurgan.GurganCounty {
	return d.戈尔甘Gurgan
}
    
func (d *陀拔斯单MazandaranDuke) CQwivir奎维尔() qwivir.QwivirCounty {
	return d.奎维尔Qwivir
}
    
func (d *陀拔斯单MazandaranDuke) CTabaristan陀拔斯单() tabaristan.TabaristanCounty {
	return d.陀拔斯单Tabaristan
}
    
var DMazandaran陀拔斯单 MazandaranDuke = &陀拔斯单MazandaranDuke{}

func init() {
	f := DMazandaran陀拔斯单.(*陀拔斯单MazandaranDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mazandaran",
		TitleName: "陀拔斯单",
		TitleCode: "d_mazandaran",
		Counties:  map[string]feud.County{},
	}

	f.阿拉穆特Alamut = alamut.CAlamut阿拉穆特
	f.阿拉穆特Alamut.SetParent(f)
	
	f.戈尔甘Gurgan = gurgan.CGurgan戈尔甘
	f.戈尔甘Gurgan.SetParent(f)
	
	f.奎维尔Qwivir = qwivir.CQwivir奎维尔
	f.奎维尔Qwivir.SetParent(f)
	
	f.陀拔斯单Tabaristan = tabaristan.CTabaristan陀拔斯单
	f.陀拔斯单Tabaristan.SetParent(f)
	
}
