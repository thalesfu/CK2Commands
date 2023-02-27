package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都奴何ThanaBarony struct {
	feud.BaseBarony
}

var BThana都奴何 feud.Barony = &都奴何ThanaBarony{}

func init() {
    f := BThana都奴何.(*都奴何ThanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thana",
		TitleName: "都奴何",
		TitleCode: "b_thana",
	}
}
