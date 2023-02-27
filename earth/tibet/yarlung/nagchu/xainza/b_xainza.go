package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申扎XainzaBarony struct {
	feud.BaseBarony
}

var BXainza申扎 feud.Barony = &申扎XainzaBarony{}

func init() {
    f := BXainza申扎.(*申扎XainzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xainza",
		TitleName: "申扎",
		TitleCode: "b_xainza",
	}
}
