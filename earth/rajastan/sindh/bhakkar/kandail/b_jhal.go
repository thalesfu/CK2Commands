package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔JhalBarony struct {
	feud.BaseBarony
}

var BJhal切尔 feud.Barony = &切尔JhalBarony{}

func init() {
	f := BJhal切尔.(*切尔JhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhal",
		TitleName: "切尔",
		TitleCode: "b_jhal",
	}
}
