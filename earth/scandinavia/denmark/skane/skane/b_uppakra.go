package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌波克拉UppakraBarony struct {
	feud.BaseBarony
}

var BUppakra乌波克拉 feud.Barony = &乌波克拉UppakraBarony{}

func init() {
    f := BUppakra乌波克拉.(*乌波克拉UppakraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uppakra",
		TitleName: "乌波克拉",
		TitleCode: "b_uppakra",
	}
}
