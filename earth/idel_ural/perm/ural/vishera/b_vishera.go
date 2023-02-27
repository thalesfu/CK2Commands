package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维舍拉VisheraBarony struct {
	feud.BaseBarony
}

var BVishera维舍拉 feud.Barony = &维舍拉VisheraBarony{}

func init() {
    f := BVishera维舍拉.(*维舍拉VisheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vishera",
		TitleName: "维舍拉",
		TitleCode: "b_vishera",
	}
}
