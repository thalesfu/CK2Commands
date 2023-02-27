package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商羯罗那伊那罗拘毗罗SankaranayinarkovilBarony struct {
	feud.BaseBarony
}

var BSankaranayinarkovil商羯罗那伊那罗拘毗罗 feud.Barony = &商羯罗那伊那罗拘毗罗SankaranayinarkovilBarony{}

func init() {
    f := BSankaranayinarkovil商羯罗那伊那罗拘毗罗.(*商羯罗那伊那罗拘毗罗SankaranayinarkovilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankaranayinarkovil",
		TitleName: "商羯罗那伊那罗拘毗罗",
		TitleCode: "b_sankaranayinarkovil",
	}
}
