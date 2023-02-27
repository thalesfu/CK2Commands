package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣让德洛讷StjeandelosneBarony struct {
	feud.BaseBarony
}

var BStjeandelosne圣让德洛讷 feud.Barony = &圣让德洛讷StjeandelosneBarony{}

func init() {
    f := BStjeandelosne圣让德洛讷.(*圣让德洛讷StjeandelosneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stjeandelosne",
		TitleName: "圣让德洛讷",
		TitleCode: "b_stjeandelosne",
	}
}
