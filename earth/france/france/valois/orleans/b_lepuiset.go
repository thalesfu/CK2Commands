package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒皮伊塞LepuisetBarony struct {
	feud.BaseBarony
}

var BLepuiset勒皮伊塞 feud.Barony = &勒皮伊塞LepuisetBarony{}

func init() {
    f := BLepuiset勒皮伊塞.(*勒皮伊塞LepuisetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepuiset",
		TitleName: "勒皮伊塞",
		TitleCode: "b_lepuiset",
	}
}
