package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝图夫BytowBarony struct {
	feud.BaseBarony
}

var BBytow贝图夫 feud.Barony = &贝图夫BytowBarony{}

func init() {
    f := BBytow贝图夫.(*贝图夫BytowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bytow",
		TitleName: "贝图夫",
		TitleCode: "b_bytow",
	}
}
