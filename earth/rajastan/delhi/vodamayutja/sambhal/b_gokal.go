package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘伽罗GokalBarony struct {
	feud.BaseBarony
}

var BGokal拘伽罗 feud.Barony = &拘伽罗GokalBarony{}

func init() {
	f := BGokal拘伽罗.(*拘伽罗GokalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gokal",
		TitleName: "拘伽罗",
		TitleCode: "b_gokal",
	}
}
