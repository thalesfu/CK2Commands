package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊格纳特科沃IgnatkovoBarony struct {
	feud.BaseBarony
}

var BIgnatkovo伊格纳特科沃 feud.Barony = &伊格纳特科沃IgnatkovoBarony{}

func init() {
    f := BIgnatkovo伊格纳特科沃.(*伊格纳特科沃IgnatkovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ignatkovo",
		TitleName: "伊格纳特科沃",
		TitleCode: "b_ignatkovo",
	}
}
