package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯陵阇罗KalinjarBarony struct {
	feud.BaseBarony
}

var BKalinjar羯陵阇罗 feud.Barony = &羯陵阇罗KalinjarBarony{}

func init() {
    f := BKalinjar羯陵阇罗.(*羯陵阇罗KalinjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalinjar",
		TitleName: "羯陵阇罗",
		TitleCode: "b_kalinjar",
	}
}
