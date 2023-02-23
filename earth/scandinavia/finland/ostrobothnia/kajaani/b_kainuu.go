package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯努KainuuBarony struct {
	feud.BaseBarony
}

var BKainuu凯努 feud.Barony = &凯努KainuuBarony{}

func init() {
	f := BKainuu凯努.(*凯努KainuuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kainuu",
		TitleName: "凯努",
		TitleCode: "b_kainuu",
	}
}
