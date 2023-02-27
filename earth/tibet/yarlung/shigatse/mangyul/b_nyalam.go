package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼牙拉木NyalamBarony struct {
	feud.BaseBarony
}

var BNyalam尼牙拉木 feud.Barony = &尼牙拉木NyalamBarony{}

func init() {
    f := BNyalam尼牙拉木.(*尼牙拉木NyalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyalam",
		TitleName: "尼牙拉木",
		TitleCode: "b_nyalam",
	}
}
