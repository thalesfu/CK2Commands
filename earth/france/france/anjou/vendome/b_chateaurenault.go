package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷诺堡ChateaurenaultBarony struct {
	feud.BaseBarony
}

var BChateaurenault雷诺堡 feud.Barony = &雷诺堡ChateaurenaultBarony{}

func init() {
    f := BChateaurenault雷诺堡.(*雷诺堡ChateaurenaultBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaurenault",
		TitleName: "雷诺堡",
		TitleCode: "b_chateaurenault",
	}
}
