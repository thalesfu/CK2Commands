package bihar

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/jharkand"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/kasi"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/magadha"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/tirabhukti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BiharKingdom interface {
    feud.Kingdom
    DJharkand阇罗犍度() 	jharkand.JharkandDuke
    DKasi迦尸() 	kasi.KasiDuke
    DMagadha摩揭陀() 	magadha.MagadhaDuke
    DTirabhukti帝那伏帝() 	tirabhukti.TirabhuktiDuke
}

type 毗诃罗BiharKingdom struct {
	feud.BaseKingdom
	阇罗犍度Jharkand 	jharkand.JharkandDuke
	迦尸Kasi 	kasi.KasiDuke
	摩揭陀Magadha 	magadha.MagadhaDuke
	帝那伏帝Tirabhukti 	tirabhukti.TirabhuktiDuke
}

func (k *毗诃罗BiharKingdom) DJharkand阇罗犍度() jharkand.JharkandDuke {
	return k.阇罗犍度Jharkand
}
    
func (k *毗诃罗BiharKingdom) DKasi迦尸() kasi.KasiDuke {
	return k.迦尸Kasi
}
    
func (k *毗诃罗BiharKingdom) DMagadha摩揭陀() magadha.MagadhaDuke {
	return k.摩揭陀Magadha
}
    
func (k *毗诃罗BiharKingdom) DTirabhukti帝那伏帝() tirabhukti.TirabhuktiDuke {
	return k.帝那伏帝Tirabhukti
}
    
var KBihar毗诃罗 BiharKingdom = &毗诃罗BiharKingdom{}

func init() {
	f := KBihar毗诃罗.(*毗诃罗BiharKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "bihar",
		TitleName: "毗诃罗",
		TitleCode: "k_bihar",
		Dukes:  map[string]feud.Duke{},
	}

	f.阇罗犍度Jharkand = jharkand.DJharkand阇罗犍度
	f.阇罗犍度Jharkand.SetParent(f)
	
	f.迦尸Kasi = kasi.DKasi迦尸
	f.迦尸Kasi.SetParent(f)
	
	f.摩揭陀Magadha = magadha.DMagadha摩揭陀
	f.摩揭陀Magadha.SetParent(f)
	
	f.帝那伏帝Tirabhukti = tirabhukti.DTirabhukti帝那伏帝
	f.帝那伏帝Tirabhukti.SetParent(f)
	
}
