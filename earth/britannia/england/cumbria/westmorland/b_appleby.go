package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普尔比ApplebyBarony struct {
	feud.BaseBarony
}

var BAppleby阿普尔比 feud.Barony = &阿普尔比ApplebyBarony{}

func init() {
    f := BAppleby阿普尔比.(*阿普尔比ApplebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "appleby",
		TitleName: "阿普尔比",
		TitleCode: "b_appleby",
	}
}
