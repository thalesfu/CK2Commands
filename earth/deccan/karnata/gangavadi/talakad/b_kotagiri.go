package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘吒耆厘KotagiriBarony struct {
	feud.BaseBarony
}

var BKotagiri拘吒耆厘 feud.Barony = &拘吒耆厘KotagiriBarony{}

func init() {
    f := BKotagiri拘吒耆厘.(*拘吒耆厘KotagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotagiri",
		TitleName: "拘吒耆厘",
		TitleCode: "b_kotagiri",
	}
}
