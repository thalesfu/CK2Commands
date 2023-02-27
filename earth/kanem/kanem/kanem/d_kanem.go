package kanem

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/kanem/bilma"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/kanem/djimi"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/kanem/manan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanemDuke interface {
    feud.Duke
    CBilma比尔马() 	bilma.BilmaCounty
    CDjimi吉米() 	djimi.DjimiCounty
    CManan马南() 	manan.MananCounty
}

type 加涅姆KanemDuke struct {
	feud.BaseDuke
	比尔马Bilma 	bilma.BilmaCounty
	吉米Djimi 	djimi.DjimiCounty
	马南Manan 	manan.MananCounty
}

func (d *加涅姆KanemDuke) CBilma比尔马() bilma.BilmaCounty {
	return d.比尔马Bilma
}
    
func (d *加涅姆KanemDuke) CDjimi吉米() djimi.DjimiCounty {
	return d.吉米Djimi
}
    
func (d *加涅姆KanemDuke) CManan马南() manan.MananCounty {
	return d.马南Manan
}
    
var DKanem加涅姆 KanemDuke = &加涅姆KanemDuke{}

func init() {
	f := DKanem加涅姆.(*加涅姆KanemDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kanem",
		TitleName: "加涅姆",
		TitleCode: "d_kanem",
		Counties:  map[string]feud.County{},
	}

	f.比尔马Bilma = bilma.CBilma比尔马
	f.比尔马Bilma.SetParent(f)
	
	f.吉米Djimi = djimi.CDjimi吉米
	f.吉米Djimi.SetParent(f)
	
	f.马南Manan = manan.CManan马南
	f.马南Manan.SetParent(f)
	
}
