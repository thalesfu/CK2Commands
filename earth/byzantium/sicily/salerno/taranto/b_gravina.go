package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉维纳GravinaBarony struct {
	feud.BaseBarony
}

var BGravina格拉维纳 feud.Barony = &格拉维纳GravinaBarony{}

func init() {
	f := BGravina格拉维纳.(*格拉维纳GravinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gravina",
		TitleName: "格拉维纳",
		TitleCode: "b_gravina",
	}
}
