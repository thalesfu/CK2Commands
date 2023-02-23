package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌津UzenBarony struct {
	feud.BaseBarony
}

var BUzen乌津 feud.Barony = &乌津UzenBarony{}

func init() {
	f := BUzen乌津.(*乌津UzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uzen",
		TitleName: "乌津",
		TitleCode: "b_uzen",
	}
}
