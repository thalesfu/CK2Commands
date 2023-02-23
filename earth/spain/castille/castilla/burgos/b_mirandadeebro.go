package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃布罗河畔米兰达MirandadeebroBarony struct {
	feud.BaseBarony
}

var BMirandadeebro埃布罗河畔米兰达 feud.Barony = &埃布罗河畔米兰达MirandadeebroBarony{}

func init() {
	f := BMirandadeebro埃布罗河畔米兰达.(*埃布罗河畔米兰达MirandadeebroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirandadeebro",
		TitleName: "埃布罗河畔米兰达",
		TitleCode: "b_mirandadeebro",
	}
}
