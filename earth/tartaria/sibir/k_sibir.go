package sibir

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/om"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/sibir"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/tobol"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/vasyugan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/yugra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SibirKingdom interface {
	feud.Kingdom
	DOm鄂木() om.OmDuke
	DSibir失必儿() sibir.SibirDuke
	DTobol托博尔() tobol.TobolDuke
	DVasyugan瓦休甘() vasyugan.VasyuganDuke
	DYugra尤格拉() yugra.YugraDuke
}

type 失必儿SibirKingdom struct {
	feud.BaseKingdom
	鄂木Om        om.OmDuke
	失必儿Sibir    sibir.SibirDuke
	托博尔Tobol    tobol.TobolDuke
	瓦休甘Vasyugan vasyugan.VasyuganDuke
	尤格拉Yugra    yugra.YugraDuke
}

func (k *失必儿SibirKingdom) DOm鄂木() om.OmDuke {
	return k.鄂木Om
}

func (k *失必儿SibirKingdom) DSibir失必儿() sibir.SibirDuke {
	return k.失必儿Sibir
}

func (k *失必儿SibirKingdom) DTobol托博尔() tobol.TobolDuke {
	return k.托博尔Tobol
}

func (k *失必儿SibirKingdom) DVasyugan瓦休甘() vasyugan.VasyuganDuke {
	return k.瓦休甘Vasyugan
}

func (k *失必儿SibirKingdom) DYugra尤格拉() yugra.YugraDuke {
	return k.尤格拉Yugra
}

var KSibir失必儿 SibirKingdom = &失必儿SibirKingdom{}

func init() {
	f := KSibir失必儿.(*失必儿SibirKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sibir",
		TitleName: "失必儿",
		TitleCode: "k_sibir",
		Dukes:     map[string]feud.Duke{},
	}

	f.鄂木Om = om.DOm鄂木
	f.鄂木Om.SetParent(f)

	f.失必儿Sibir = sibir.DSibir失必儿
	f.失必儿Sibir.SetParent(f)

	f.托博尔Tobol = tobol.DTobol托博尔
	f.托博尔Tobol.SetParent(f)

	f.瓦休甘Vasyugan = vasyugan.DVasyugan瓦休甘
	f.瓦休甘Vasyugan.SetParent(f)

	f.尤格拉Yugra = yugra.DYugra尤格拉
	f.尤格拉Yugra.SetParent(f)

}
