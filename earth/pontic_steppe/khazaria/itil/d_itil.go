package itil

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil/itil"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil/saqsin"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil/saray"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil/uzens"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil/yeruslan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ItilDuke interface {
    feud.Duke
    CItil阿得() 	itil.ItilCounty
    CSaqsin撒哈辛() 	saqsin.SaqsinCounty
    CSaray萨赖() 	saray.SarayCounty
    CUzens乌津() 	uzens.UzensCounty
    CYeruslan叶鲁斯兰() 	yeruslan.YeruslanCounty
}

type 阿得ItilDuke struct {
	feud.BaseDuke
	阿得Itil 	itil.ItilCounty
	撒哈辛Saqsin 	saqsin.SaqsinCounty
	萨赖Saray 	saray.SarayCounty
	乌津Uzens 	uzens.UzensCounty
	叶鲁斯兰Yeruslan 	yeruslan.YeruslanCounty
}

func (d *阿得ItilDuke) CItil阿得() itil.ItilCounty {
	return d.阿得Itil
}
    
func (d *阿得ItilDuke) CSaqsin撒哈辛() saqsin.SaqsinCounty {
	return d.撒哈辛Saqsin
}
    
func (d *阿得ItilDuke) CSaray萨赖() saray.SarayCounty {
	return d.萨赖Saray
}
    
func (d *阿得ItilDuke) CUzens乌津() uzens.UzensCounty {
	return d.乌津Uzens
}
    
func (d *阿得ItilDuke) CYeruslan叶鲁斯兰() yeruslan.YeruslanCounty {
	return d.叶鲁斯兰Yeruslan
}
    
var DItil阿得 ItilDuke = &阿得ItilDuke{}

func init() {
	f := DItil阿得.(*阿得ItilDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "itil",
		TitleName: "阿得",
		TitleCode: "d_itil",
		Counties:  map[string]feud.County{},
	}

	f.阿得Itil = itil.CItil阿得
	f.阿得Itil.SetParent(f)
	
	f.撒哈辛Saqsin = saqsin.CSaqsin撒哈辛
	f.撒哈辛Saqsin.SetParent(f)
	
	f.萨赖Saray = saray.CSaray萨赖
	f.萨赖Saray.SetParent(f)
	
	f.乌津Uzens = uzens.CUzens乌津
	f.乌津Uzens.SetParent(f)
	
	f.叶鲁斯兰Yeruslan = yeruslan.CYeruslan叶鲁斯兰
	f.叶鲁斯兰Yeruslan.SetParent(f)
	
}
