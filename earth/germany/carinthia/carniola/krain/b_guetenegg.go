package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古埃特内格GueteneggBarony struct {
	feud.BaseBarony
}

var BGuetenegg古埃特内格 feud.Barony = &古埃特内格GueteneggBarony{}

func init() {
	f := BGuetenegg古埃特内格.(*古埃特内格GueteneggBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guetenegg",
		TitleName: "古埃特内格",
		TitleCode: "b_guetenegg",
	}
}
