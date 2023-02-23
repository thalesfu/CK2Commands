package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 磋梅ChameBarony struct {
	feud.BaseBarony
}

var BChame磋梅 feud.Barony = &磋梅ChameBarony{}

func init() {
	f := BChame磋梅.(*磋梅ChameBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chame",
		TitleName: "磋梅",
		TitleCode: "b_chame",
	}
}
