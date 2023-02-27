package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍里索格列布斯克BorisoglebskBarony struct {
	feud.BaseBarony
}

var BBorisoglebsk鲍里索格列布斯克 feud.Barony = &鲍里索格列布斯克BorisoglebskBarony{}

func init() {
    f := BBorisoglebsk鲍里索格列布斯克.(*鲍里索格列布斯克BorisoglebskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borisoglebsk",
		TitleName: "鲍里索格列布斯克",
		TitleCode: "b_borisoglebsk",
	}
}
