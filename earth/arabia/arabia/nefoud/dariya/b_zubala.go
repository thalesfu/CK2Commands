package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖巴莱ZubalaBarony struct {
	feud.BaseBarony
}

var BZubala祖巴莱 feud.Barony = &祖巴莱ZubalaBarony{}

func init() {
    f := BZubala祖巴莱.(*祖巴莱ZubalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zubala",
		TitleName: "祖巴莱",
		TitleCode: "b_zubala",
	}
}
