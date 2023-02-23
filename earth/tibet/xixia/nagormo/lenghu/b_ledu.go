package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乐都LeduBarony struct {
	feud.BaseBarony
}

var BLedu乐都 feud.Barony = &乐都LeduBarony{}

func init() {
	f := BLedu乐都.(*乐都LeduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ledu",
		TitleName: "乐都",
		TitleCode: "b_ledu",
	}
}
