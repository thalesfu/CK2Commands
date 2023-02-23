package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科什拉巴德KoshrabotBarony struct {
	feud.BaseBarony
}

var BKoshrabot科什拉巴德 feud.Barony = &科什拉巴德KoshrabotBarony{}

func init() {
	f := BKoshrabot科什拉巴德.(*科什拉巴德KoshrabotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshrabot",
		TitleName: "科什拉巴德",
		TitleCode: "b_koshrabot",
	}
}
