package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞻波ChambaBarony struct {
	feud.BaseBarony
}

var BChamba瞻波 feud.Barony = &瞻波ChambaBarony{}

func init() {
    f := BChamba瞻波.(*瞻波ChambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chamba",
		TitleName: "瞻波",
		TitleCode: "b_chamba",
	}
}
