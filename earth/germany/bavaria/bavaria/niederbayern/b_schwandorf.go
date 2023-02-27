package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施万多夫SchwandorfBarony struct {
	feud.BaseBarony
}

var BSchwandorf施万多夫 feud.Barony = &施万多夫SchwandorfBarony{}

func init() {
    f := BSchwandorf施万多夫.(*施万多夫SchwandorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwandorf",
		TitleName: "施万多夫",
		TitleCode: "b_schwandorf",
	}
}
