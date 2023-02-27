package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法乌新堡Chateauneuf_du_faouBarony struct {
	feud.BaseBarony
}

var BChateauneuf_du_faou法乌新堡 feud.Barony = &法乌新堡Chateauneuf_du_faouBarony{}

func init() {
    f := BChateauneuf_du_faou法乌新堡.(*法乌新堡Chateauneuf_du_faouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateauneuf_du_faou",
		TitleName: "法乌新堡",
		TitleCode: "b_chateauneuf_du_faou",
	}
}
