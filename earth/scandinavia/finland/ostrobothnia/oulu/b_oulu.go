package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥卢OuluBarony struct {
	feud.BaseBarony
}

var BOulu奥卢 feud.Barony = &奥卢OuluBarony{}

func init() {
    f := BOulu奥卢.(*奥卢OuluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oulu",
		TitleName: "奥卢",
		TitleCode: "b_oulu",
	}
}
