package kashmir

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/kashmir"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/pamir"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/uttaranchal"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KashmirKingdom interface {
    feud.Kingdom
    DKashmir迦湿弥罗() 	kashmir.KashmirDuke
    DPamir播密() 	pamir.PamirDuke
    DUttaranchal北安遮罗() 	uttaranchal.UttaranchalDuke
}

type 迦湿弥罗KashmirKingdom struct {
	feud.BaseKingdom
	迦湿弥罗Kashmir 	kashmir.KashmirDuke
	播密Pamir 	pamir.PamirDuke
	北安遮罗Uttaranchal 	uttaranchal.UttaranchalDuke
}

func (k *迦湿弥罗KashmirKingdom) DKashmir迦湿弥罗() kashmir.KashmirDuke {
	return k.迦湿弥罗Kashmir
}
    
func (k *迦湿弥罗KashmirKingdom) DPamir播密() pamir.PamirDuke {
	return k.播密Pamir
}
    
func (k *迦湿弥罗KashmirKingdom) DUttaranchal北安遮罗() uttaranchal.UttaranchalDuke {
	return k.北安遮罗Uttaranchal
}
    
var KKashmir迦湿弥罗 KashmirKingdom = &迦湿弥罗KashmirKingdom{}

func init() {
	f := KKashmir迦湿弥罗.(*迦湿弥罗KashmirKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "kashmir",
		TitleName: "迦湿弥罗",
		TitleCode: "k_kashmir",
		Dukes:  map[string]feud.Duke{},
	}

	f.迦湿弥罗Kashmir = kashmir.DKashmir迦湿弥罗
	f.迦湿弥罗Kashmir.SetParent(f)
	
	f.播密Pamir = pamir.DPamir播密
	f.播密Pamir.SetParent(f)
	
	f.北安遮罗Uttaranchal = uttaranchal.DUttaranchal北安遮罗
	f.北安遮罗Uttaranchal.SetParent(f)
	
}
