package balata

import (
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/balata/almada"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/balata/lisboa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalataDuke interface {
    feud.Duke
    CAlmada阿尔马达() 	almada.AlmadaCounty
    CLisboa里斯本() 	lisboa.LisboaCounty
}

type 里斯本BalataDuke struct {
	feud.BaseDuke
	阿尔马达Almada 	almada.AlmadaCounty
	里斯本Lisboa 	lisboa.LisboaCounty
}

func (d *里斯本BalataDuke) CAlmada阿尔马达() almada.AlmadaCounty {
	return d.阿尔马达Almada
}
    
func (d *里斯本BalataDuke) CLisboa里斯本() lisboa.LisboaCounty {
	return d.里斯本Lisboa
}
    
var DBalata里斯本 BalataDuke = &里斯本BalataDuke{}

func init() {
	f := DBalata里斯本.(*里斯本BalataDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "balata",
		TitleName: "里斯本",
		TitleCode: "d_balata",
		Counties:  map[string]feud.County{},
	}

	f.阿尔马达Almada = almada.CAlmada阿尔马达
	f.阿尔马达Almada.SetParent(f)
	
	f.里斯本Lisboa = lisboa.CLisboa里斯本
	f.里斯本Lisboa.SetParent(f)
	
}
