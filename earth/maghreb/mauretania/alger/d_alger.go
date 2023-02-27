package alger

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/alger/al_djazair"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/alger/beni_yanni"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/alger/orania"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlgerDuke interface {
    feud.Duke
    CAl_djazair贾扎伊尔() 	al_djazair.Al_djazairCounty
    CBeni_yanni贝尼耶尼() 	beni_yanni.Beni_yanniCounty
    COrania奥兰() 	orania.OraniaCounty
}

type 阿尔及尔AlgerDuke struct {
	feud.BaseDuke
	贾扎伊尔Al_djazair 	al_djazair.Al_djazairCounty
	贝尼耶尼Beni_yanni 	beni_yanni.Beni_yanniCounty
	奥兰Orania 	orania.OraniaCounty
}

func (d *阿尔及尔AlgerDuke) CAl_djazair贾扎伊尔() al_djazair.Al_djazairCounty {
	return d.贾扎伊尔Al_djazair
}
    
func (d *阿尔及尔AlgerDuke) CBeni_yanni贝尼耶尼() beni_yanni.Beni_yanniCounty {
	return d.贝尼耶尼Beni_yanni
}
    
func (d *阿尔及尔AlgerDuke) COrania奥兰() orania.OraniaCounty {
	return d.奥兰Orania
}
    
var DAlger阿尔及尔 AlgerDuke = &阿尔及尔AlgerDuke{}

func init() {
	f := DAlger阿尔及尔.(*阿尔及尔AlgerDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "alger",
		TitleName: "阿尔及尔",
		TitleCode: "d_alger",
		Counties:  map[string]feud.County{},
	}

	f.贾扎伊尔Al_djazair = al_djazair.CAl_djazair贾扎伊尔
	f.贾扎伊尔Al_djazair.SetParent(f)
	
	f.贝尼耶尼Beni_yanni = beni_yanni.CBeni_yanni贝尼耶尼
	f.贝尼耶尼Beni_yanni.SetParent(f)
	
	f.奥兰Orania = orania.COrania奥兰
	f.奥兰Orania.SetParent(f)
	
}
