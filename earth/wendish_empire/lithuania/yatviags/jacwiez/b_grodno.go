package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗德诺GrodnoBarony struct {
	feud.BaseBarony
}

var BGrodno格罗德诺 feud.Barony = &格罗德诺GrodnoBarony{}

func init() {
    f := BGrodno格罗德诺.(*格罗德诺GrodnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grodno",
		TitleName: "格罗德诺",
		TitleCode: "b_grodno",
	}
}
