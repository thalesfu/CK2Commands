package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乾牟伽KanmukalBarony struct {
	feud.BaseBarony
}

var BKanmukal乾牟伽 feud.Barony = &乾牟伽KanmukalBarony{}

func init() {
	f := BKanmukal乾牟伽.(*乾牟伽KanmukalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanmukal",
		TitleName: "乾牟伽",
		TitleCode: "b_kanmukal",
	}
}
