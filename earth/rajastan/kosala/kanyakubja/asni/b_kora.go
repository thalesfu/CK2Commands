package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厥罗KoraBarony struct {
	feud.BaseBarony
}

var BKora厥罗 feud.Barony = &厥罗KoraBarony{}

func init() {
    f := BKora厥罗.(*厥罗KoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kora",
		TitleName: "厥罗",
		TitleCode: "b_kora",
	}
}
