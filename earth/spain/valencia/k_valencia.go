package valencia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/murcia"
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/valencia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValenciaKingdom interface {
    feud.Kingdom
    DMurcia穆尔西亚() 	murcia.MurciaDuke
    DValencia瓦伦西亚() 	valencia.ValenciaDuke
}

type 瓦伦西亚ValenciaKingdom struct {
	feud.BaseKingdom
	穆尔西亚Murcia 	murcia.MurciaDuke
	瓦伦西亚Valencia 	valencia.ValenciaDuke
}

func (k *瓦伦西亚ValenciaKingdom) DMurcia穆尔西亚() murcia.MurciaDuke {
	return k.穆尔西亚Murcia
}
    
func (k *瓦伦西亚ValenciaKingdom) DValencia瓦伦西亚() valencia.ValenciaDuke {
	return k.瓦伦西亚Valencia
}
    
var KValencia瓦伦西亚 ValenciaKingdom = &瓦伦西亚ValenciaKingdom{}

func init() {
	f := KValencia瓦伦西亚.(*瓦伦西亚ValenciaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "valencia",
		TitleName: "瓦伦西亚",
		TitleCode: "k_valencia",
		Dukes:  map[string]feud.Duke{},
	}

	f.穆尔西亚Murcia = murcia.DMurcia穆尔西亚
	f.穆尔西亚Murcia.SetParent(f)
	
	f.瓦伦西亚Valencia = valencia.DValencia瓦伦西亚
	f.瓦伦西亚Valencia.SetParent(f)
	
}
