package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温伯恩WimborneBarony struct {
	feud.BaseBarony
}

var BWimborne温伯恩 feud.Barony = &温伯恩WimborneBarony{}

func init() {
    f := BWimborne温伯恩.(*温伯恩WimborneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wimborne",
		TitleName: "温伯恩",
		TitleCode: "b_wimborne",
	}
}
