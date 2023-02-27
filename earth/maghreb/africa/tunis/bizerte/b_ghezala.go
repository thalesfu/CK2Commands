package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格赫扎拉GhezalaBarony struct {
	feud.BaseBarony
}

var BGhezala格赫扎拉 feud.Barony = &格赫扎拉GhezalaBarony{}

func init() {
    f := BGhezala格赫扎拉.(*格赫扎拉GhezalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghezala",
		TitleName: "格赫扎拉",
		TitleCode: "b_ghezala",
	}
}
