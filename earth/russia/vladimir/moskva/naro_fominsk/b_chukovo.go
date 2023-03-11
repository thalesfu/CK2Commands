package naro_fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘科沃ChukovoBarony struct {
	feud.BaseBarony
}

var BChukovo丘科沃 feud.Barony = &丘科沃ChukovoBarony{}

func init() {
    f := BChukovo丘科沃.(*丘科沃ChukovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chukovo",
		TitleName: "丘科沃",
		TitleCode: "b_chukovo",
	}
}
