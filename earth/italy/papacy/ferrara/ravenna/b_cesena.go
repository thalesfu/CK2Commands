package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切塞纳CesenaBarony struct {
	feud.BaseBarony
}

var BCesena切塞纳 feud.Barony = &切塞纳CesenaBarony{}

func init() {
    f := BCesena切塞纳.(*切塞纳CesenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cesena",
		TitleName: "切塞纳",
		TitleCode: "b_cesena",
	}
}
