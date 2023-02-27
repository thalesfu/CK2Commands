package songhay

import (
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/gurma"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/songhay"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SonghayKingdom interface {
    feud.Kingdom
    DGurma古尔马() 	gurma.GurmaDuke
    DSonghay桑海() 	songhay.SonghayDuke
}

type 桑海SonghayKingdom struct {
	feud.BaseKingdom
	古尔马Gurma 	gurma.GurmaDuke
	桑海Songhay 	songhay.SonghayDuke
}

func (k *桑海SonghayKingdom) DGurma古尔马() gurma.GurmaDuke {
	return k.古尔马Gurma
}
    
func (k *桑海SonghayKingdom) DSonghay桑海() songhay.SonghayDuke {
	return k.桑海Songhay
}
    
var KSonghay桑海 SonghayKingdom = &桑海SonghayKingdom{}

func init() {
	f := KSonghay桑海.(*桑海SonghayKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "songhay",
		TitleName: "桑海",
		TitleCode: "k_songhay",
		Dukes:  map[string]feud.Duke{},
	}

	f.古尔马Gurma = gurma.DGurma古尔马
	f.古尔马Gurma.SetParent(f)
	
	f.桑海Songhay = songhay.DSonghay桑海
	f.桑海Songhay.SetParent(f)
	
}
