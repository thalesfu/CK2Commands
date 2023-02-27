package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈赫GochBarony struct {
	feud.BaseBarony
}

var BGoch戈赫 feud.Barony = &戈赫GochBarony{}

func init() {
    f := BGoch戈赫.(*戈赫GochBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goch",
		TitleName: "戈赫",
		TitleCode: "b_goch",
	}
}
