package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 墨西拿MessinaBarony struct {
	feud.BaseBarony
}

var BMessina墨西拿 feud.Barony = &墨西拿MessinaBarony{}

func init() {
	f := BMessina墨西拿.(*墨西拿MessinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "messina",
		TitleName: "墨西拿",
		TitleCode: "b_messina",
	}
}
