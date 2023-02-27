package sapmi

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/finnmark"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/kola"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/sapmi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SapmiKingdom interface {
    feud.Kingdom
    DFinnmark芬马克() 	finnmark.FinnmarkDuke
    DKola科拉() 	kola.KolaDuke
    DSapmi拉普兰() 	sapmi.SapmiDuke
}

type 拉普兰SapmiKingdom struct {
	feud.BaseKingdom
	芬马克Finnmark 	finnmark.FinnmarkDuke
	科拉Kola 	kola.KolaDuke
	拉普兰Sapmi 	sapmi.SapmiDuke
}

func (k *拉普兰SapmiKingdom) DFinnmark芬马克() finnmark.FinnmarkDuke {
	return k.芬马克Finnmark
}
    
func (k *拉普兰SapmiKingdom) DKola科拉() kola.KolaDuke {
	return k.科拉Kola
}
    
func (k *拉普兰SapmiKingdom) DSapmi拉普兰() sapmi.SapmiDuke {
	return k.拉普兰Sapmi
}
    
var KSapmi拉普兰 SapmiKingdom = &拉普兰SapmiKingdom{}

func init() {
	f := KSapmi拉普兰.(*拉普兰SapmiKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sapmi",
		TitleName: "拉普兰",
		TitleCode: "k_sapmi",
		Dukes:  map[string]feud.Duke{},
	}

	f.芬马克Finnmark = finnmark.DFinnmark芬马克
	f.芬马克Finnmark.SetParent(f)
	
	f.科拉Kola = kola.DKola科拉
	f.科拉Kola.SetParent(f)
	
	f.拉普兰Sapmi = sapmi.DSapmi拉普兰
	f.拉普兰Sapmi.SetParent(f)
	
}
