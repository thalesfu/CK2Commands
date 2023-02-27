package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛沙拉MashalaBarony struct {
	feud.BaseBarony
}

var BMashala玛沙拉 feud.Barony = &玛沙拉MashalaBarony{}

func init() {
    f := BMashala玛沙拉.(*玛沙拉MashalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mashala",
		TitleName: "玛沙拉",
		TitleCode: "b_mashala",
	}
}
