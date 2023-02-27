package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 增期ZengqiBarony struct {
	feud.BaseBarony
}

var BZengqi增期 feud.Barony = &增期ZengqiBarony{}

func init() {
    f := BZengqi增期.(*增期ZengqiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zengqi",
		TitleName: "增期",
		TitleCode: "b_zengqi",
	}
}
