package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺达WangdaBarony struct {
	feud.BaseBarony
}

var BWangda旺达 feud.Barony = &旺达WangdaBarony{}

func init() {
    f := BWangda旺达.(*旺达WangdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wangda",
		TitleName: "旺达",
		TitleCode: "b_wangda",
	}
}
