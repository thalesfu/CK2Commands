package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那乔佗NagaudaBarony struct {
	feud.BaseBarony
}

var BNagauda那乔佗 feud.Barony = &那乔佗NagaudaBarony{}

func init() {
    f := BNagauda那乔佗.(*那乔佗NagaudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagauda",
		TitleName: "那乔佗",
		TitleCode: "b_nagauda",
	}
}
