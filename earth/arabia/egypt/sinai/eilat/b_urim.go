package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌里姆UrimBarony struct {
	feud.BaseBarony
}

var BUrim乌里姆 feud.Barony = &乌里姆UrimBarony{}

func init() {
    f := BUrim乌里姆.(*乌里姆UrimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urim",
		TitleName: "乌里姆",
		TitleCode: "b_urim",
	}
}
