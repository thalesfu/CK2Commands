package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格默尔GomorBarony struct {
	feud.BaseBarony
}

var BGomor格默尔 feud.Barony = &格默尔GomorBarony{}

func init() {
	f := BGomor格默尔.(*格默尔GomorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gomor",
		TitleName: "格默尔",
		TitleCode: "b_gomor",
	}
}
