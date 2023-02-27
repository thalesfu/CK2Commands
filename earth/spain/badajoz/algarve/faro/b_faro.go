package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法鲁FaroBarony struct {
	feud.BaseBarony
}

var BFaro法鲁 feud.Barony = &法鲁FaroBarony{}

func init() {
    f := BFaro法鲁.(*法鲁FaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faro",
		TitleName: "法鲁",
		TitleCode: "b_faro",
	}
}
