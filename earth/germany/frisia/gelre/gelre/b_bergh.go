package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔赫BerghBarony struct {
	feud.BaseBarony
}

var BBergh贝尔赫 feud.Barony = &贝尔赫BerghBarony{}

func init() {
    f := BBergh贝尔赫.(*贝尔赫BerghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergh",
		TitleName: "贝尔赫",
		TitleCode: "b_bergh",
	}
}
