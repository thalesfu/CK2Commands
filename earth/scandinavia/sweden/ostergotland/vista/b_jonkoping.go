package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延雪平JonkopingBarony struct {
	feud.BaseBarony
}

var BJonkoping延雪平 feud.Barony = &延雪平JonkopingBarony{}

func init() {
	f := BJonkoping延雪平.(*延雪平JonkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jonkoping",
		TitleName: "延雪平",
		TitleCode: "b_jonkoping",
	}
}
