package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷克利帕CeskalipaBarony struct {
	feud.BaseBarony
}

var BCeskalipa捷克利帕 feud.Barony = &捷克利帕CeskalipaBarony{}

func init() {
    f := BCeskalipa捷克利帕.(*捷克利帕CeskalipaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ceskalipa",
		TitleName: "捷克利帕",
		TitleCode: "b_ceskalipa",
	}
}
