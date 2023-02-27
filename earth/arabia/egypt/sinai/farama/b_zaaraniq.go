package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎拉尼克ZaaraniqBarony struct {
	feud.BaseBarony
}

var BZaaraniq扎拉尼克 feud.Barony = &扎拉尼克ZaaraniqBarony{}

func init() {
    f := BZaaraniq扎拉尼克.(*扎拉尼克ZaaraniqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaaraniq",
		TitleName: "扎拉尼克",
		TitleCode: "b_zaaraniq",
	}
}
