package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼尔门KhurmenBarony struct {
	feud.BaseBarony
}

var BKhurmen呼尔门 feud.Barony = &呼尔门KhurmenBarony{}

func init() {
    f := BKhurmen呼尔门.(*呼尔门KhurmenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khurmen",
		TitleName: "呼尔门",
		TitleCode: "b_khurmen",
	}
}
