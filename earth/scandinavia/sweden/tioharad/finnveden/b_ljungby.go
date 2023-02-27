package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 永比LjungbyBarony struct {
	feud.BaseBarony
}

var BLjungby永比 feud.Barony = &永比LjungbyBarony{}

func init() {
    f := BLjungby永比.(*永比LjungbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ljungby",
		TitleName: "永比",
		TitleCode: "b_ljungby",
	}
}
