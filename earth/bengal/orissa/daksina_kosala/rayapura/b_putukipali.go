package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富兜吉波梨PutukipaliBarony struct {
	feud.BaseBarony
}

var BPutukipali富兜吉波梨 feud.Barony = &富兜吉波梨PutukipaliBarony{}

func init() {
    f := BPutukipali富兜吉波梨.(*富兜吉波梨PutukipaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "putukipali",
		TitleName: "富兜吉波梨",
		TitleCode: "b_putukipali",
	}
}
