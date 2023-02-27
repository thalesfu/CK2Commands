package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艺德KhedBarony struct {
	feud.BaseBarony
}

var BKhed艺德 feud.Barony = &艺德KhedBarony{}

func init() {
    f := BKhed艺德.(*艺德KhedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khed",
		TitleName: "艺德",
		TitleCode: "b_khed",
	}
}
