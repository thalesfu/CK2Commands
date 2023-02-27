package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马科尔德MaugholdBarony struct {
	feud.BaseBarony
}

var BMaughold马科尔德 feud.Barony = &马科尔德MaugholdBarony{}

func init() {
    f := BMaughold马科尔德.(*马科尔德MaugholdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maughold",
		TitleName: "马科尔德",
		TitleCode: "b_maughold",
	}
}
