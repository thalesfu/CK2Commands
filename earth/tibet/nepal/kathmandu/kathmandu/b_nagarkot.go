package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那揭罗拘吒NagarkotBarony struct {
	feud.BaseBarony
}

var BNagarkot那揭罗拘吒 feud.Barony = &那揭罗拘吒NagarkotBarony{}

func init() {
    f := BNagarkot那揭罗拘吒.(*那揭罗拘吒NagarkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagarkot",
		TitleName: "那揭罗拘吒",
		TitleCode: "b_nagarkot",
	}
}
