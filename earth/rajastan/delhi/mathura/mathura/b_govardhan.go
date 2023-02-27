package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔跋檀GovardhanBarony struct {
	feud.BaseBarony
}

var BGovardhan乔跋檀 feud.Barony = &乔跋檀GovardhanBarony{}

func init() {
    f := BGovardhan乔跋檀.(*乔跋檀GovardhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "govardhan",
		TitleName: "乔跋檀",
		TitleCode: "b_govardhan",
	}
}
