package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安提俄基亚Antiochia_akroinonBarony struct {
	feud.BaseBarony
}

var BAntiochia_akroinon安提俄基亚 feud.Barony = &安提俄基亚Antiochia_akroinonBarony{}

func init() {
    f := BAntiochia_akroinon安提俄基亚.(*安提俄基亚Antiochia_akroinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antiochia_akroinon",
		TitleName: "安提俄基亚",
		TitleCode: "b_antiochia_akroinon",
	}
}
