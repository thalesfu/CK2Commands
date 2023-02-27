package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉胡纳GhunanBarony struct {
	feud.BaseBarony
}

var BGhunan吉胡纳 feud.Barony = &吉胡纳GhunanBarony{}

func init() {
    f := BGhunan吉胡纳.(*吉胡纳GhunanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghunan",
		TitleName: "吉胡纳",
		TitleCode: "b_ghunan",
	}
}
