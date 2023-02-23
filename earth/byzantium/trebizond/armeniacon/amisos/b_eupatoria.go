package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶夫帕托里亚EupatoriaBarony struct {
	feud.BaseBarony
}

var BEupatoria叶夫帕托里亚 feud.Barony = &叶夫帕托里亚EupatoriaBarony{}

func init() {
	f := BEupatoria叶夫帕托里亚.(*叶夫帕托里亚EupatoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eupatoria",
		TitleName: "叶夫帕托里亚",
		TitleCode: "b_eupatoria",
	}
}
