package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙蒂永ChatillonBarony struct {
	feud.BaseBarony
}

var BChatillon沙蒂永 feud.Barony = &沙蒂永ChatillonBarony{}

func init() {
    f := BChatillon沙蒂永.(*沙蒂永ChatillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatillon",
		TitleName: "沙蒂永",
		TitleCode: "b_chatillon",
	}
}
