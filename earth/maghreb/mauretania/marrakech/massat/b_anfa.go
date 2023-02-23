package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安法AnfaBarony struct {
	feud.BaseBarony
}

var BAnfa安法 feud.Barony = &安法AnfaBarony{}

func init() {
	f := BAnfa安法.(*安法AnfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anfa",
		TitleName: "安法",
		TitleCode: "b_anfa",
	}
}
