package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉布夫GrabowBarony struct {
	feud.BaseBarony
}

var BGrabow格拉布夫 feud.Barony = &格拉布夫GrabowBarony{}

func init() {
    f := BGrabow格拉布夫.(*格拉布夫GrabowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grabow",
		TitleName: "格拉布夫",
		TitleCode: "b_grabow",
	}
}
