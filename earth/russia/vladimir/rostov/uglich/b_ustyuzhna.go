package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯秋日纳UstyuzhnaBarony struct {
	feud.BaseBarony
}

var BUstyuzhna乌斯秋日纳 feud.Barony = &乌斯秋日纳UstyuzhnaBarony{}

func init() {
    f := BUstyuzhna乌斯秋日纳.(*乌斯秋日纳UstyuzhnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustyuzhna",
		TitleName: "乌斯秋日纳",
		TitleCode: "b_ustyuzhna",
	}
}
