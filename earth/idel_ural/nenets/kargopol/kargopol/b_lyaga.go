package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利亚加LyagaBarony struct {
	feud.BaseBarony
}

var BLyaga利亚加 feud.Barony = &利亚加LyagaBarony{}

func init() {
    f := BLyaga利亚加.(*利亚加LyagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyaga",
		TitleName: "利亚加",
		TitleCode: "b_lyaga",
	}
}
