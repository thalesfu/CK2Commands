package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅姆NemBarony struct {
	feud.BaseBarony
}

var BNem涅姆 feud.Barony = &涅姆NemBarony{}

func init() {
    f := BNem涅姆.(*涅姆NemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nem",
		TitleName: "涅姆",
		TitleCode: "b_nem",
	}
}
