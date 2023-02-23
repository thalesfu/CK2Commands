package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉文纳RavennaBarony struct {
	feud.BaseBarony
}

var BRavenna拉文纳 feud.Barony = &拉文纳RavennaBarony{}

func init() {
	f := BRavenna拉文纳.(*拉文纳RavennaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravenna",
		TitleName: "拉文纳",
		TitleCode: "b_ravenna",
	}
}
