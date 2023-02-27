package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷斯科沃LyskovoBarony struct {
	feud.BaseBarony
}

var BLyskovo雷斯科沃 feud.Barony = &雷斯科沃LyskovoBarony{}

func init() {
    f := BLyskovo雷斯科沃.(*雷斯科沃LyskovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyskovo",
		TitleName: "雷斯科沃",
		TitleCode: "b_lyskovo",
	}
}
