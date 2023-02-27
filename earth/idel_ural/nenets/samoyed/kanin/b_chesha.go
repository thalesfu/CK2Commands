package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔沙CheshaBarony struct {
	feud.BaseBarony
}

var BChesha乔沙 feud.Barony = &乔沙CheshaBarony{}

func init() {
    f := BChesha乔沙.(*乔沙CheshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chesha",
		TitleName: "乔沙",
		TitleCode: "b_chesha",
	}
}
