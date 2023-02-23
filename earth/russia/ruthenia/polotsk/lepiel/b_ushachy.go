package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌沙奇UshachyBarony struct {
	feud.BaseBarony
}

var BUshachy乌沙奇 feud.Barony = &乌沙奇UshachyBarony{}

func init() {
	f := BUshachy乌沙奇.(*乌沙奇UshachyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ushachy",
		TitleName: "乌沙奇",
		TitleCode: "b_ushachy",
	}
}
