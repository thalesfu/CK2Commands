package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姞利瑟泥湿伐罗GrishneshwarBarony struct {
	feud.BaseBarony
}

var BGrishneshwar姞利瑟泥湿伐罗 feud.Barony = &姞利瑟泥湿伐罗GrishneshwarBarony{}

func init() {
	f := BGrishneshwar姞利瑟泥湿伐罗.(*姞利瑟泥湿伐罗GrishneshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grishneshwar",
		TitleName: "姞利瑟泥湿伐罗",
		TitleCode: "b_grishneshwar",
	}
}
