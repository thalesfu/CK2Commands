package anti-atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿盖达勒AguedalBarony struct {
	feud.BaseBarony
}

var BAguedal阿盖达勒 feud.Barony = &阿盖达勒AguedalBarony{}

func init() {
    f := BAguedal阿盖达勒.(*阿盖达勒AguedalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aguedal",
		TitleName: "阿盖达勒",
		TitleCode: "b_aguedal",
	}
}
