package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马希斯MahisBarony struct {
	feud.BaseBarony
}

var BMahis马希斯 feud.Barony = &马希斯MahisBarony{}

func init() {
    f := BMahis马希斯.(*马希斯MahisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahis",
		TitleName: "马希斯",
		TitleCode: "b_mahis",
	}
}
