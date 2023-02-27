package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃普施泰因EppsteinBarony struct {
	feud.BaseBarony
}

var BEppstein埃普施泰因 feud.Barony = &埃普施泰因EppsteinBarony{}

func init() {
    f := BEppstein埃普施泰因.(*埃普施泰因EppsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eppstein",
		TitleName: "埃普施泰因",
		TitleCode: "b_eppstein",
	}
}
