package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾布林JabrinBarony struct {
	feud.BaseBarony
}

var BJabrin贾布林 feud.Barony = &贾布林JabrinBarony{}

func init() {
	f := BJabrin贾布林.(*贾布林JabrinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabrin",
		TitleName: "贾布林",
		TitleCode: "b_jabrin",
	}
}
