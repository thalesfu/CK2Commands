package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲索ZussowBarony struct {
	feud.BaseBarony
}

var BZussow曲索 feud.Barony = &曲索ZussowBarony{}

func init() {
    f := BZussow曲索.(*曲索ZussowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zussow",
		TitleName: "曲索",
		TitleCode: "b_zussow",
	}
}
