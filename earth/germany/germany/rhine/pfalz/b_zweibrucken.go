package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨韦布吕肯ZweibruckenBarony struct {
	feud.BaseBarony
}

var BZweibrucken茨韦布吕肯 feud.Barony = &茨韦布吕肯ZweibruckenBarony{}

func init() {
    f := BZweibrucken茨韦布吕肯.(*茨韦布吕肯ZweibruckenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zweibrucken",
		TitleName: "茨韦布吕肯",
		TitleCode: "b_zweibrucken",
	}
}
