package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯吕斯CaylusBarony struct {
	feud.BaseBarony
}

var BCaylus凯吕斯 feud.Barony = &凯吕斯CaylusBarony{}

func init() {
    f := BCaylus凯吕斯.(*凯吕斯CaylusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caylus",
		TitleName: "凯吕斯",
		TitleCode: "b_caylus",
	}
}
