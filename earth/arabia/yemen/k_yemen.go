package yemen

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/hadramut"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/sanaa"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/socotra"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/taizz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YemenKingdom interface {
	feud.Kingdom
	DHadramut哈得拉姆() hadramut.HadramutDuke
	DSanaa萨那() sanaa.SanaaDuke
	DSocotra索科特拉() socotra.SocotraDuke
	DTaizz塔伊兹() taizz.TaizzDuke
}

type 也门YemenKingdom struct {
	feud.BaseKingdom
	哈得拉姆Hadramut hadramut.HadramutDuke
	萨那Sanaa      sanaa.SanaaDuke
	索科特拉Socotra  socotra.SocotraDuke
	塔伊兹Taizz     taizz.TaizzDuke
}

func (k *也门YemenKingdom) DHadramut哈得拉姆() hadramut.HadramutDuke {
	return k.哈得拉姆Hadramut
}

func (k *也门YemenKingdom) DSanaa萨那() sanaa.SanaaDuke {
	return k.萨那Sanaa
}

func (k *也门YemenKingdom) DSocotra索科特拉() socotra.SocotraDuke {
	return k.索科特拉Socotra
}

func (k *也门YemenKingdom) DTaizz塔伊兹() taizz.TaizzDuke {
	return k.塔伊兹Taizz
}

var KYemen也门 YemenKingdom = &也门YemenKingdom{}

func init() {
	f := KYemen也门.(*也门YemenKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "yemen",
		TitleName: "也门",
		TitleCode: "k_yemen",
		Dukes:     map[string]feud.Duke{},
	}

	f.哈得拉姆Hadramut = hadramut.DHadramut哈得拉姆
	f.哈得拉姆Hadramut.SetParent(f)

	f.萨那Sanaa = sanaa.DSanaa萨那
	f.萨那Sanaa.SetParent(f)

	f.索科特拉Socotra = socotra.DSocotra索科特拉
	f.索科特拉Socotra.SetParent(f)

	f.塔伊兹Taizz = taizz.DTaizz塔伊兹
	f.塔伊兹Taizz.SetParent(f)

}
