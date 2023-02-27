package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康加巴KangabaBarony struct {
	feud.BaseBarony
}

var BKangaba康加巴 feud.Barony = &康加巴KangabaBarony{}

func init() {
    f := BKangaba康加巴.(*康加巴KangabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kangaba",
		TitleName: "康加巴",
		TitleCode: "b_kangaba",
	}
}
