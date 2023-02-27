package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪景ZhotsenBarony struct {
	feud.BaseBarony
}

var BZhotsen雪景 feud.Barony = &雪景ZhotsenBarony{}

func init() {
    f := BZhotsen雪景.(*雪景ZhotsenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhotsen",
		TitleName: "雪景",
		TitleCode: "b_zhotsen",
	}
}
