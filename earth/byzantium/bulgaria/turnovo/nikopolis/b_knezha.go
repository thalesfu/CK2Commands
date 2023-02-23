package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克内扎KnezhaBarony struct {
	feud.BaseBarony
}

var BKnezha克内扎 feud.Barony = &克内扎KnezhaBarony{}

func init() {
	f := BKnezha克内扎.(*克内扎KnezhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knezha",
		TitleName: "克内扎",
		TitleCode: "b_knezha",
	}
}
