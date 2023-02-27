package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比列沃BilievaBarony struct {
	feud.BaseBarony
}

var BBilieva比列沃 feud.Barony = &比列沃BilievaBarony{}

func init() {
    f := BBilieva比列沃.(*比列沃BilievaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilieva",
		TitleName: "比列沃",
		TitleCode: "b_bilieva",
	}
}
