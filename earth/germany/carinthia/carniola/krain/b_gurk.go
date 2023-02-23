package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔克GurkBarony struct {
	feud.BaseBarony
}

var BGurk古尔克 feud.Barony = &古尔克GurkBarony{}

func init() {
	f := BGurk古尔克.(*古尔克GurkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurk",
		TitleName: "古尔克",
		TitleCode: "b_gurk",
	}
}
