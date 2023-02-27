package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 背崩BaibungBarony struct {
	feud.BaseBarony
}

var BBaibung背崩 feud.Barony = &背崩BaibungBarony{}

func init() {
    f := BBaibung背崩.(*背崩BaibungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baibung",
		TitleName: "背崩",
		TitleCode: "b_baibung",
	}
}
