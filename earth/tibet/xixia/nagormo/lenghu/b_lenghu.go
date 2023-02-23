package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 冷湖LenghuBarony struct {
	feud.BaseBarony
}

var BLenghu冷湖 feud.Barony = &冷湖LenghuBarony{}

func init() {
	f := BLenghu冷湖.(*冷湖LenghuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lenghu",
		TitleName: "冷湖",
		TitleCode: "b_lenghu",
	}
}
