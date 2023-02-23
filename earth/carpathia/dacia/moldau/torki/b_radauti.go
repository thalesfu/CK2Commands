package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒德乌齐RadautiBarony struct {
	feud.BaseBarony
}

var BRadauti勒德乌齐 feud.Barony = &勒德乌齐RadautiBarony{}

func init() {
	f := BRadauti勒德乌齐.(*勒德乌齐RadautiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radauti",
		TitleName: "勒德乌齐",
		TitleCode: "b_radauti",
	}
}
