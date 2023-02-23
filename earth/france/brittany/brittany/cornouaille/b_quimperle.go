package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎佩莱QuimperleBarony struct {
	feud.BaseBarony
}

var BQuimperle坎佩莱 feud.Barony = &坎佩莱QuimperleBarony{}

func init() {
	f := BQuimperle坎佩莱.(*坎佩莱QuimperleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quimperle",
		TitleName: "坎佩莱",
		TitleCode: "b_quimperle",
	}
}
