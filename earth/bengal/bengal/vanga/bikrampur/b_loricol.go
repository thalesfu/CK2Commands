package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 若利孔LoricolBarony struct {
	feud.BaseBarony
}

var BLoricol若利孔 feud.Barony = &若利孔LoricolBarony{}

func init() {
    f := BLoricol若利孔.(*若利孔LoricolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loricol",
		TitleName: "若利孔",
		TitleCode: "b_loricol",
	}
}
