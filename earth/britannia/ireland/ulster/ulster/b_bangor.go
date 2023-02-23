package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班戈BangorBarony struct {
	feud.BaseBarony
}

var BBangor班戈 feud.Barony = &班戈BangorBarony{}

func init() {
	f := BBangor班戈.(*班戈BangorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangor",
		TitleName: "班戈",
		TitleCode: "b_bangor",
	}
}
