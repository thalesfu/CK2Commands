package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘兰陀KorandhBarony struct {
	feud.BaseBarony
}

var BKorandh拘兰陀 feud.Barony = &拘兰陀KorandhBarony{}

func init() {
	f := BKorandh拘兰陀.(*拘兰陀KorandhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korandh",
		TitleName: "拘兰陀",
		TitleCode: "b_korandh",
	}
}
