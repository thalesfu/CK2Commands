package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦伐耶KadwayaBarony struct {
	feud.BaseBarony
}

var BKadwaya迦伐耶 feud.Barony = &迦伐耶KadwayaBarony{}

func init() {
    f := BKadwaya迦伐耶.(*迦伐耶KadwayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadwaya",
		TitleName: "迦伐耶",
		TitleCode: "b_kadwaya",
	}
}
