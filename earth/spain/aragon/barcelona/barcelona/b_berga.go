package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔加BergaBarony struct {
	feud.BaseBarony
}

var BBerga贝尔加 feud.Barony = &贝尔加BergaBarony{}

func init() {
	f := BBerga贝尔加.(*贝尔加BergaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berga",
		TitleName: "贝尔加",
		TitleCode: "b_berga",
	}
}
