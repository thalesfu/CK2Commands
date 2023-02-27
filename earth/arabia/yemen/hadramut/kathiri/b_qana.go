package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀纳QanaBarony struct {
	feud.BaseBarony
}

var BQana喀纳 feud.Barony = &喀纳QanaBarony{}

func init() {
    f := BQana喀纳.(*喀纳QanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qana",
		TitleName: "喀纳",
		TitleCode: "b_qana",
	}
}
