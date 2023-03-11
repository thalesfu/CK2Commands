package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 助勒基福勒Di_l_kiflBarony struct {
	feud.BaseBarony
}

var BDi_l_kifl助勒基福勒 feud.Barony = &助勒基福勒Di_l_kiflBarony{}

func init() {
    f := BDi_l_kifl助勒基福勒.(*助勒基福勒Di_l_kiflBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "di_l_kifl",
		TitleName: "助勒基福勒",
		TitleCode: "b_di-l-kifl",
	}
}
