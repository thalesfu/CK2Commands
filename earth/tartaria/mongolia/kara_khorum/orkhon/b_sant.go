package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑特SantBarony struct {
	feud.BaseBarony
}

var BSant桑特 feud.Barony = &桑特SantBarony{}

func init() {
    f := BSant桑特.(*桑特SantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sant",
		TitleName: "桑特",
		TitleCode: "b_sant",
	}
}
