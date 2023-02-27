package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格林南GreenanBarony struct {
	feud.BaseBarony
}

var BGreenan格林南 feud.Barony = &格林南GreenanBarony{}

func init() {
    f := BGreenan格林南.(*格林南GreenanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "greenan",
		TitleName: "格林南",
		TitleCode: "b_greenan",
	}
}
