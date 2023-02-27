package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑竭罗洲SagardwipBarony struct {
	feud.BaseBarony
}

var BSagardwip娑竭罗洲 feud.Barony = &娑竭罗洲SagardwipBarony{}

func init() {
    f := BSagardwip娑竭罗洲.(*娑竭罗洲SagardwipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagardwip",
		TitleName: "娑竭罗洲",
		TitleCode: "b_sagardwip",
	}
}
