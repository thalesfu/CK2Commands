package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格盖拉AqqalaBarony struct {
	feud.BaseBarony
}

var BAqqala阿格盖拉 feud.Barony = &阿格盖拉AqqalaBarony{}

func init() {
    f := BAqqala阿格盖拉.(*阿格盖拉AqqalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqqala",
		TitleName: "阿格盖拉",
		TitleCode: "b_aqqala",
	}
}
