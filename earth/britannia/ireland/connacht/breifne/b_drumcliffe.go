package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜木里菲DrumcliffeBarony struct {
	feud.BaseBarony
}

var BDrumcliffe杜木里菲 feud.Barony = &杜木里菲DrumcliffeBarony{}

func init() {
    f := BDrumcliffe杜木里菲.(*杜木里菲DrumcliffeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drumcliffe",
		TitleName: "杜木里菲",
		TitleCode: "b_drumcliffe",
	}
}
