package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温波WenboBarony struct {
	feud.BaseBarony
}

var BWenbo温波 feud.Barony = &温波WenboBarony{}

func init() {
    f := BWenbo温波.(*温波WenboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wenbo",
		TitleName: "温波",
		TitleCode: "b_wenbo",
	}
}
