package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普热德布日PrzedborzBarony struct {
	feud.BaseBarony
}

var BPrzedborz普热德布日 feud.Barony = &普热德布日PrzedborzBarony{}

func init() {
    f := BPrzedborz普热德布日.(*普热德布日PrzedborzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "przedborz",
		TitleName: "普热德布日",
		TitleCode: "b_przedborz",
	}
}
