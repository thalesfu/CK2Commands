package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴奇BacsBarony struct {
	feud.BaseBarony
}

var BBacs巴奇 feud.Barony = &巴奇BacsBarony{}

func init() {
    f := BBacs巴奇.(*巴奇BacsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bacs",
		TitleName: "巴奇",
		TitleCode: "b_bacs",
	}
}
