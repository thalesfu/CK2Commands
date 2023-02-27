package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉伦登ClarendonBarony struct {
	feud.BaseBarony
}

var BClarendon克拉伦登 feud.Barony = &克拉伦登ClarendonBarony{}

func init() {
    f := BClarendon克拉伦登.(*克拉伦登ClarendonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clarendon",
		TitleName: "克拉伦登",
		TitleCode: "b_clarendon",
	}
}
