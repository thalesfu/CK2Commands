package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗德诺HrodnaBarony struct {
	feud.BaseBarony
}

var BHrodna格罗德诺 feud.Barony = &格罗德诺HrodnaBarony{}

func init() {
    f := BHrodna格罗德诺.(*格罗德诺HrodnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hrodna",
		TitleName: "格罗德诺",
		TitleCode: "b_hrodna",
	}
}
