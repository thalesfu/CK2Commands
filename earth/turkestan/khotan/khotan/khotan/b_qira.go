package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 质逻QiraBarony struct {
	feud.BaseBarony
}

var BQira质逻 feud.Barony = &质逻QiraBarony{}

func init() {
    f := BQira质逻.(*质逻QiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qira",
		TitleName: "质逻",
		TitleCode: "b_qira",
	}
}
