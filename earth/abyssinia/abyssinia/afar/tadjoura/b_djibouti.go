package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉布提DjiboutiBarony struct {
	feud.BaseBarony
}

var BDjibouti吉布提 feud.Barony = &吉布提DjiboutiBarony{}

func init() {
	f := BDjibouti吉布提.(*吉布提DjiboutiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djibouti",
		TitleName: "吉布提",
		TitleCode: "b_djibouti",
	}
}
