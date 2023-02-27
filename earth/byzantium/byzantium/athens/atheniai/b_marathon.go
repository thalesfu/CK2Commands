package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉松MarathonBarony struct {
	feud.BaseBarony
}

var BMarathon马拉松 feud.Barony = &马拉松MarathonBarony{}

func init() {
    f := BMarathon马拉松.(*马拉松MarathonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marathon",
		TitleName: "马拉松",
		TitleCode: "b_marathon",
	}
}
