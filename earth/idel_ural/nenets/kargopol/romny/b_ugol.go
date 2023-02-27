package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌戈尔UgolBarony struct {
	feud.BaseBarony
}

var BUgol乌戈尔 feud.Barony = &乌戈尔UgolBarony{}

func init() {
    f := BUgol乌戈尔.(*乌戈尔UgolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ugol",
		TitleName: "乌戈尔",
		TitleCode: "b_ugol",
	}
}
