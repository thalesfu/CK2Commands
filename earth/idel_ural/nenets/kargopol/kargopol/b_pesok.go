package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩索克PesokBarony struct {
	feud.BaseBarony
}

var BPesok佩索克 feud.Barony = &佩索克PesokBarony{}

func init() {
    f := BPesok佩索克.(*佩索克PesokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pesok",
		TitleName: "佩索克",
		TitleCode: "b_pesok",
	}
}
