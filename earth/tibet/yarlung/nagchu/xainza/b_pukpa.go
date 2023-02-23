package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普保PukpaBarony struct {
	feud.BaseBarony
}

var BPukpa普保 feud.Barony = &普保PukpaBarony{}

func init() {
	f := BPukpa普保.(*普保PukpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pukpa",
		TitleName: "普保",
		TitleCode: "b_pukpa",
	}
}
