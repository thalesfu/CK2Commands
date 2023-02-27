package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新松奇NowysaczBarony struct {
	feud.BaseBarony
}

var BNowysacz新松奇 feud.Barony = &新松奇NowysaczBarony{}

func init() {
    f := BNowysacz新松奇.(*新松奇NowysaczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nowysacz",
		TitleName: "新松奇",
		TitleCode: "b_nowysacz",
	}
}
