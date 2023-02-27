package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎佩尔QuimperBarony struct {
	feud.BaseBarony
}

var BQuimper坎佩尔 feud.Barony = &坎佩尔QuimperBarony{}

func init() {
    f := BQuimper坎佩尔.(*坎佩尔QuimperBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quimper",
		TitleName: "坎佩尔",
		TitleCode: "b_quimper",
	}
}
