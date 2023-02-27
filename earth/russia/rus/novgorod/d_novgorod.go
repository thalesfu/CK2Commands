package novgorod

import (
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod/bezhetsky_verh"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod/borovichi"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod/kholm"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod/novgorod"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod/rusa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NovgorodDuke interface {
    feud.Duke
    CBezhetsky_verh拉多加() 	bezhetsky_verh.Bezhetsky_verhCounty
    CBorovichi博罗维奇() 	borovichi.BorovichiCounty
    CKholm霍尔姆() 	kholm.KholmCounty
    CNovgorod诺夫哥罗德() 	novgorod.NovgorodCounty
    CRusa鲁萨() 	rusa.RusaCounty
}

type 诺夫哥罗德NovgorodDuke struct {
	feud.BaseDuke
	拉多加Bezhetsky_verh 	bezhetsky_verh.Bezhetsky_verhCounty
	博罗维奇Borovichi 	borovichi.BorovichiCounty
	霍尔姆Kholm 	kholm.KholmCounty
	诺夫哥罗德Novgorod 	novgorod.NovgorodCounty
	鲁萨Rusa 	rusa.RusaCounty
}

func (d *诺夫哥罗德NovgorodDuke) CBezhetsky_verh拉多加() bezhetsky_verh.Bezhetsky_verhCounty {
	return d.拉多加Bezhetsky_verh
}
    
func (d *诺夫哥罗德NovgorodDuke) CBorovichi博罗维奇() borovichi.BorovichiCounty {
	return d.博罗维奇Borovichi
}
    
func (d *诺夫哥罗德NovgorodDuke) CKholm霍尔姆() kholm.KholmCounty {
	return d.霍尔姆Kholm
}
    
func (d *诺夫哥罗德NovgorodDuke) CNovgorod诺夫哥罗德() novgorod.NovgorodCounty {
	return d.诺夫哥罗德Novgorod
}
    
func (d *诺夫哥罗德NovgorodDuke) CRusa鲁萨() rusa.RusaCounty {
	return d.鲁萨Rusa
}
    
var DNovgorod诺夫哥罗德 NovgorodDuke = &诺夫哥罗德NovgorodDuke{}

func init() {
	f := DNovgorod诺夫哥罗德.(*诺夫哥罗德NovgorodDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "novgorod",
		TitleName: "诺夫哥罗德",
		TitleCode: "d_novgorod",
		Counties:  map[string]feud.County{},
	}

	f.拉多加Bezhetsky_verh = bezhetsky_verh.CBezhetsky_verh拉多加
	f.拉多加Bezhetsky_verh.SetParent(f)
	
	f.博罗维奇Borovichi = borovichi.CBorovichi博罗维奇
	f.博罗维奇Borovichi.SetParent(f)
	
	f.霍尔姆Kholm = kholm.CKholm霍尔姆
	f.霍尔姆Kholm.SetParent(f)
	
	f.诺夫哥罗德Novgorod = novgorod.CNovgorod诺夫哥罗德
	f.诺夫哥罗德Novgorod.SetParent(f)
	
	f.鲁萨Rusa = rusa.CRusa鲁萨
	f.鲁萨Rusa.SetParent(f)
	
}
