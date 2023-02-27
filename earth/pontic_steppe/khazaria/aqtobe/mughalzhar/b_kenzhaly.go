package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯扎雷KenzhalyBarony struct {
	feud.BaseBarony
}

var BKenzhaly肯扎雷 feud.Barony = &肯扎雷KenzhalyBarony{}

func init() {
    f := BKenzhaly肯扎雷.(*肯扎雷KenzhalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenzhaly",
		TitleName: "肯扎雷",
		TitleCode: "b_kenzhaly",
	}
}
