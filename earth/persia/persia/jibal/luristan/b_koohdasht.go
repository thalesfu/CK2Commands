package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库赫达什特KoohdashtBarony struct {
	feud.BaseBarony
}

var BKoohdasht库赫达什特 feud.Barony = &库赫达什特KoohdashtBarony{}

func init() {
    f := BKoohdasht库赫达什特.(*库赫达什特KoohdashtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koohdasht",
		TitleName: "库赫达什特",
		TitleCode: "b_koohdasht",
	}
}
