package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 匈达HundarBarony struct {
	feud.BaseBarony
}

var BHundar匈达 feud.Barony = &匈达HundarBarony{}

func init() {
	f := BHundar匈达.(*匈达HundarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hundar",
		TitleName: "匈达",
		TitleCode: "b_hundar",
	}
}
