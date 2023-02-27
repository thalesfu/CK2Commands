package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温什科姆WinchcombeBarony struct {
	feud.BaseBarony
}

var BWinchcombe温什科姆 feud.Barony = &温什科姆WinchcombeBarony{}

func init() {
    f := BWinchcombe温什科姆.(*温什科姆WinchcombeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "winchcombe",
		TitleName: "温什科姆",
		TitleCode: "b_winchcombe",
	}
}
