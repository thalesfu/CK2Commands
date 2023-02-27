package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曳咥ErchisBarony struct {
	feud.BaseBarony
}

var BErchis曳咥 feud.Barony = &曳咥ErchisBarony{}

func init() {
    f := BErchis曳咥.(*曳咥ErchisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erchis",
		TitleName: "曳咥",
		TitleCode: "b_erchis",
	}
}
