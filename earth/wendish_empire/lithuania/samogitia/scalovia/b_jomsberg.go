package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约姆斯贝格JomsbergBarony struct {
	feud.BaseBarony
}

var BJomsberg约姆斯贝格 feud.Barony = &约姆斯贝格JomsbergBarony{}

func init() {
    f := BJomsberg约姆斯贝格.(*约姆斯贝格JomsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jomsberg",
		TitleName: "约姆斯贝格",
		TitleCode: "b_jomsberg",
	}
}
