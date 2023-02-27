package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩日采PryzyceBarony struct {
	feud.BaseBarony
}

var BPryzyce佩日采 feud.Barony = &佩日采PryzyceBarony{}

func init() {
    f := BPryzyce佩日采.(*佩日采PryzyceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pryzyce",
		TitleName: "佩日采",
		TitleCode: "b_pryzyce",
	}
}
