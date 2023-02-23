package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉凡德SarafandBarony struct {
	feud.BaseBarony
}

var BSarafand萨拉凡德 feud.Barony = &萨拉凡德SarafandBarony{}

func init() {
	f := BSarafand萨拉凡德.(*萨拉凡德SarafandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarafand",
		TitleName: "萨拉凡德",
		TitleCode: "b_sarafand",
	}
}
