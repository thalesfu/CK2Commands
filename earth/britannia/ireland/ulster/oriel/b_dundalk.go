package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓多克DundalkBarony struct {
	feud.BaseBarony
}

var BDundalk邓多克 feud.Barony = &邓多克DundalkBarony{}

func init() {
	f := BDundalk邓多克.(*邓多克DundalkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dundalk",
		TitleName: "邓多克",
		TitleCode: "b_dundalk",
	}
}
