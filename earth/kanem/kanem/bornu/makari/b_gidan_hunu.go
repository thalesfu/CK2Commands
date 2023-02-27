package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉丹胡努Gidan_hunuBarony struct {
	feud.BaseBarony
}

var BGidan_hunu吉丹胡努 feud.Barony = &吉丹胡努Gidan_hunuBarony{}

func init() {
    f := BGidan_hunu吉丹胡努.(*吉丹胡努Gidan_hunuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gidan_hunu",
		TitleName: "吉丹胡努",
		TitleCode: "b_gidan_hunu",
	}
}
