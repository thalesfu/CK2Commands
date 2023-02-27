package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡明KaminBarony struct {
	feud.BaseBarony
}

var BKamin卡明 feud.Barony = &卡明KaminBarony{}

func init() {
    f := BKamin卡明.(*卡明KaminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamin",
		TitleName: "卡明",
		TitleCode: "b_kamin",
	}
}
