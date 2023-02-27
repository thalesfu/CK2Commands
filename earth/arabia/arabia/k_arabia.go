package arabia

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/amman"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/oman"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArabiaKingdom interface {
    feud.Kingdom
    DAmman巴林() 	amman.AmmanDuke
    DArabia_petrae阿拉比亚() 	arabia_petrae.Arabia_petraeDuke
    DMedina汉志() 	medina.MedinaDuke
    DNefoud内夫得() 	nefoud.NefoudDuke
    DOman阿曼() 	oman.OmanDuke
}

type 阿拉伯ArabiaKingdom struct {
	feud.BaseKingdom
	巴林Amman 	amman.AmmanDuke
	阿拉比亚Arabia_petrae 	arabia_petrae.Arabia_petraeDuke
	汉志Medina 	medina.MedinaDuke
	内夫得Nefoud 	nefoud.NefoudDuke
	阿曼Oman 	oman.OmanDuke
}

func (k *阿拉伯ArabiaKingdom) DAmman巴林() amman.AmmanDuke {
	return k.巴林Amman
}
    
func (k *阿拉伯ArabiaKingdom) DArabia_petrae阿拉比亚() arabia_petrae.Arabia_petraeDuke {
	return k.阿拉比亚Arabia_petrae
}
    
func (k *阿拉伯ArabiaKingdom) DMedina汉志() medina.MedinaDuke {
	return k.汉志Medina
}
    
func (k *阿拉伯ArabiaKingdom) DNefoud内夫得() nefoud.NefoudDuke {
	return k.内夫得Nefoud
}
    
func (k *阿拉伯ArabiaKingdom) DOman阿曼() oman.OmanDuke {
	return k.阿曼Oman
}
    
var KArabia阿拉伯 ArabiaKingdom = &阿拉伯ArabiaKingdom{}

func init() {
	f := KArabia阿拉伯.(*阿拉伯ArabiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "arabia",
		TitleName: "阿拉伯",
		TitleCode: "k_arabia",
		Dukes:  map[string]feud.Duke{},
	}

	f.巴林Amman = amman.DAmman巴林
	f.巴林Amman.SetParent(f)
	
	f.阿拉比亚Arabia_petrae = arabia_petrae.DArabia_petrae阿拉比亚
	f.阿拉比亚Arabia_petrae.SetParent(f)
	
	f.汉志Medina = medina.DMedina汉志
	f.汉志Medina.SetParent(f)
	
	f.内夫得Nefoud = nefoud.DNefoud内夫得
	f.内夫得Nefoud.SetParent(f)
	
	f.阿曼Oman = oman.DOman阿曼
	f.阿曼Oman.SetParent(f)
	
}
