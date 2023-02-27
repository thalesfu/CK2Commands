package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特罗维泰WloclawekBarony struct {
	feud.BaseBarony
}

var BWloclawek奥斯特罗维泰 feud.Barony = &奥斯特罗维泰WloclawekBarony{}

func init() {
    f := BWloclawek奥斯特罗维泰.(*奥斯特罗维泰WloclawekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wloclawek",
		TitleName: "奥斯特罗维泰",
		TitleCode: "b_wloclawek",
	}
}
