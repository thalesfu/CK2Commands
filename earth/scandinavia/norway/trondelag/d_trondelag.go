package trondelag

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/trondelag/halogaland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/trondelag/naumadal"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/trondelag/trondelag"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrondelagDuke interface {
	feud.Duke
	CHalogaland霍洛加兰() halogaland.HalogalandCounty
	CNaumadal瑙马达尔() naumadal.NaumadalCounty
	CTrondelag尼达罗斯() trondelag.TrondelagCounty
}

type 尼达罗斯TrondelagDuke struct {
	feud.BaseDuke
	霍洛加兰Halogaland halogaland.HalogalandCounty
	瑙马达尔Naumadal   naumadal.NaumadalCounty
	尼达罗斯Trondelag  trondelag.TrondelagCounty
}

func (d *尼达罗斯TrondelagDuke) CHalogaland霍洛加兰() halogaland.HalogalandCounty {
	return d.霍洛加兰Halogaland
}

func (d *尼达罗斯TrondelagDuke) CNaumadal瑙马达尔() naumadal.NaumadalCounty {
	return d.瑙马达尔Naumadal
}

func (d *尼达罗斯TrondelagDuke) CTrondelag尼达罗斯() trondelag.TrondelagCounty {
	return d.尼达罗斯Trondelag
}

var DTrondelag尼达罗斯 TrondelagDuke = &尼达罗斯TrondelagDuke{}

func init() {
	f := DTrondelag尼达罗斯.(*尼达罗斯TrondelagDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "trondelag",
		TitleName: "尼达罗斯",
		TitleCode: "d_trondelag",
		Counties:  map[string]feud.County{},
	}

	f.霍洛加兰Halogaland = halogaland.CHalogaland霍洛加兰
	f.霍洛加兰Halogaland.SetParent(f)

	f.瑙马达尔Naumadal = naumadal.CNaumadal瑙马达尔
	f.瑙马达尔Naumadal.SetParent(f)

	f.尼达罗斯Trondelag = trondelag.CTrondelag尼达罗斯
	f.尼达罗斯Trondelag.SetParent(f)

}
