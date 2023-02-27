package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔姆利达TamridaBarony struct {
	feud.BaseBarony
}

var BTamrida塔姆利达 feud.Barony = &塔姆利达TamridaBarony{}

func init() {
    f := BTamrida塔姆利达.(*塔姆利达TamridaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamrida",
		TitleName: "塔姆利达",
		TitleCode: "b_tamrida",
	}
}
