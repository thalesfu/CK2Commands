package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄斯特沃格OstervagenBarony struct {
	feud.BaseBarony
}

var BOstervagen厄斯特沃格 feud.Barony = &厄斯特沃格OstervagenBarony{}

func init() {
    f := BOstervagen厄斯特沃格.(*厄斯特沃格OstervagenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostervagen",
		TitleName: "厄斯特沃格",
		TitleCode: "b_ostervagen",
	}
}
