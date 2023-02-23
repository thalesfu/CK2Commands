package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝韦尔BelverBarony struct {
	feud.BaseBarony
}

var BBelver贝韦尔 feud.Barony = &贝韦尔BelverBarony{}

func init() {
	f := BBelver贝韦尔.(*贝韦尔BelverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belver",
		TitleName: "贝韦尔",
		TitleCode: "b_belver",
	}
}
