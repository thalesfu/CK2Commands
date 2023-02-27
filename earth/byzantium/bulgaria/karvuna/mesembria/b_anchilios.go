package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安基里奥斯AnchiliosBarony struct {
	feud.BaseBarony
}

var BAnchilios安基里奥斯 feud.Barony = &安基里奥斯AnchiliosBarony{}

func init() {
    f := BAnchilios安基里奥斯.(*安基里奥斯AnchiliosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anchilios",
		TitleName: "安基里奥斯",
		TitleCode: "b_anchilios",
	}
}
