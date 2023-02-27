package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗盖KorkaiBarony struct {
	feud.BaseBarony
}

var BKorkai拘罗盖 feud.Barony = &拘罗盖KorkaiBarony{}

func init() {
    f := BKorkai拘罗盖.(*拘罗盖KorkaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korkai",
		TitleName: "拘罗盖",
		TitleCode: "b_korkai",
	}
}
