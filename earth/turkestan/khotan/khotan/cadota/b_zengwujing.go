package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曾屋迳ZengwujingBarony struct {
	feud.BaseBarony
}

var BZengwujing曾屋迳 feud.Barony = &曾屋迳ZengwujingBarony{}

func init() {
	f := BZengwujing曾屋迳.(*曾屋迳ZengwujingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zengwujing",
		TitleName: "曾屋迳",
		TitleCode: "b_zengwujing",
	}
}
