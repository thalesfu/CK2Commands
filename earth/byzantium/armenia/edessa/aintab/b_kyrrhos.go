package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛斯KyrrhosBarony struct {
	feud.BaseBarony
}

var BKyrrhos库洛斯 feud.Barony = &库洛斯KyrrhosBarony{}

func init() {
    f := BKyrrhos库洛斯.(*库洛斯KyrrhosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyrrhos",
		TitleName: "库洛斯",
		TitleCode: "b_kyrrhos",
	}
}
