package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼补罗ManpuraBarony struct {
	feud.BaseBarony
}

var BManpura曼补罗 feud.Barony = &曼补罗ManpuraBarony{}

func init() {
    f := BManpura曼补罗.(*曼补罗ManpuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manpura",
		TitleName: "曼补罗",
		TitleCode: "b_manpura",
	}
}
