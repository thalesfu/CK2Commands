package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅胡丹JahudanBarony struct {
	feud.BaseBarony
}

var BJahudan雅胡丹 feud.Barony = &雅胡丹JahudanBarony{}

func init() {
	f := BJahudan雅胡丹.(*雅胡丹JahudanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahudan",
		TitleName: "雅胡丹",
		TitleCode: "b_jahudan",
	}
}
