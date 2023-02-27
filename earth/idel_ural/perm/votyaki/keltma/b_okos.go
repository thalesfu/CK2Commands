package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥科斯OkosBarony struct {
	feud.BaseBarony
}

var BOkos奥科斯 feud.Barony = &奥科斯OkosBarony{}

func init() {
    f := BOkos奥科斯.(*奥科斯OkosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okos",
		TitleName: "奥科斯",
		TitleCode: "b_okos",
	}
}
