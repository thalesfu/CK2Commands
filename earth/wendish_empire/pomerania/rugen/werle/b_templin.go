package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕普林TemplinBarony struct {
	feud.BaseBarony
}

var BTemplin滕普林 feud.Barony = &滕普林TemplinBarony{}

func init() {
    f := BTemplin滕普林.(*滕普林TemplinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "templin",
		TitleName: "滕普林",
		TitleCode: "b_templin",
	}
}
