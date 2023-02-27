package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜布罗夫诺DubrownaBarony struct {
	feud.BaseBarony
}

var BDubrowna杜布罗夫诺 feud.Barony = &杜布罗夫诺DubrownaBarony{}

func init() {
    f := BDubrowna杜布罗夫诺.(*杜布罗夫诺DubrownaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubrowna",
		TitleName: "杜布罗夫诺",
		TitleCode: "b_dubrowna",
	}
}
