package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉纳达GranadaBarony struct {
	feud.BaseBarony
}

var BGranada格拉纳达 feud.Barony = &格拉纳达GranadaBarony{}

func init() {
	f := BGranada格拉纳达.(*格拉纳达GranadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "granada",
		TitleName: "格拉纳达",
		TitleCode: "b_granada",
	}
}
