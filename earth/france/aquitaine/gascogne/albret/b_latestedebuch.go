package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉泰斯特德比克LatestedebuchBarony struct {
	feud.BaseBarony
}

var BLatestedebuch拉泰斯特德比克 feud.Barony = &拉泰斯特德比克LatestedebuchBarony{}

func init() {
    f := BLatestedebuch拉泰斯特德比克.(*拉泰斯特德比克LatestedebuchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latestedebuch",
		TitleName: "拉泰斯特德比克",
		TitleCode: "b_latestedebuch",
	}
}
