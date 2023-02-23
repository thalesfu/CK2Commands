package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 爱诺斯AinosBarony struct {
	feud.BaseBarony
}

var BAinos爱诺斯 feud.Barony = &爱诺斯AinosBarony{}

func init() {
	f := BAinos爱诺斯.(*爱诺斯AinosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainos",
		TitleName: "爱诺斯",
		TitleCode: "b_ainos",
	}
}
