package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔伯费尔德ElberfeldBarony struct {
	feud.BaseBarony
}

var BElberfeld埃尔伯费尔德 feud.Barony = &埃尔伯费尔德ElberfeldBarony{}

func init() {
	f := BElberfeld埃尔伯费尔德.(*埃尔伯费尔德ElberfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elberfeld",
		TitleName: "埃尔伯费尔德",
		TitleCode: "b_elberfeld",
	}
}
