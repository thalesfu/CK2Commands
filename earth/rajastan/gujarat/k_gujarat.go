package gujarat

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/anartta"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/gurjara_mandala"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/lata"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/saurashtra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GujaratKingdom interface {
    feud.Kingdom
    DAnartta阿捺多() 	anartta.AnarttaDuke
    DGurjara_mandala瞿折罗曼荼罗() 	gurjara_mandala.Gurjara_mandalaDuke
    DLata罗吒() 	lata.LataDuke
    DSaurashtra苏剌侘() 	saurashtra.SaurashtraDuke
}

type 瞿折罗GujaratKingdom struct {
	feud.BaseKingdom
	阿捺多Anartta 	anartta.AnarttaDuke
	瞿折罗曼荼罗Gurjara_mandala 	gurjara_mandala.Gurjara_mandalaDuke
	罗吒Lata 	lata.LataDuke
	苏剌侘Saurashtra 	saurashtra.SaurashtraDuke
}

func (k *瞿折罗GujaratKingdom) DAnartta阿捺多() anartta.AnarttaDuke {
	return k.阿捺多Anartta
}
    
func (k *瞿折罗GujaratKingdom) DGurjara_mandala瞿折罗曼荼罗() gurjara_mandala.Gurjara_mandalaDuke {
	return k.瞿折罗曼荼罗Gurjara_mandala
}
    
func (k *瞿折罗GujaratKingdom) DLata罗吒() lata.LataDuke {
	return k.罗吒Lata
}
    
func (k *瞿折罗GujaratKingdom) DSaurashtra苏剌侘() saurashtra.SaurashtraDuke {
	return k.苏剌侘Saurashtra
}
    
var KGujarat瞿折罗 GujaratKingdom = &瞿折罗GujaratKingdom{}

func init() {
	f := KGujarat瞿折罗.(*瞿折罗GujaratKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "gujarat",
		TitleName: "瞿折罗",
		TitleCode: "k_gujarat",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿捺多Anartta = anartta.DAnartta阿捺多
	f.阿捺多Anartta.SetParent(f)
	
	f.瞿折罗曼荼罗Gurjara_mandala = gurjara_mandala.DGurjara_mandala瞿折罗曼荼罗
	f.瞿折罗曼荼罗Gurjara_mandala.SetParent(f)
	
	f.罗吒Lata = lata.DLata罗吒
	f.罗吒Lata.SetParent(f)
	
	f.苏剌侘Saurashtra = saurashtra.DSaurashtra苏剌侘
	f.苏剌侘Saurashtra.SetParent(f)
	
}
