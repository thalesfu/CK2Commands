package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾斯克JaskBarony struct {
	feud.BaseBarony
}

var BJask贾斯克 feud.Barony = &贾斯克JaskBarony{}

func init() {
	f := BJask贾斯克.(*贾斯克JaskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jask",
		TitleName: "贾斯克",
		TitleCode: "b_jask",
	}
}
