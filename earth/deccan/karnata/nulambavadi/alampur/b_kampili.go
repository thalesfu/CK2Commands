package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剑毕离KampiliBarony struct {
	feud.BaseBarony
}

var BKampili剑毕离 feud.Barony = &剑毕离KampiliBarony{}

func init() {
    f := BKampili剑毕离.(*剑毕离KampiliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kampili",
		TitleName: "剑毕离",
		TitleCode: "b_kampili",
	}
}
