package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利斯KlisBarony struct {
	feud.BaseBarony
}

var BKlis克利斯 feud.Barony = &克利斯KlisBarony{}

func init() {
    f := BKlis克利斯.(*克利斯KlisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klis",
		TitleName: "克利斯",
		TitleCode: "b_klis",
	}
}
