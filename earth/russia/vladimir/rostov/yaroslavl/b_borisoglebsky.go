package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍里索格列布斯基BorisoglebskyBarony struct {
	feud.BaseBarony
}

var BBorisoglebsky鲍里索格列布斯基 feud.Barony = &鲍里索格列布斯基BorisoglebskyBarony{}

func init() {
    f := BBorisoglebsky鲍里索格列布斯基.(*鲍里索格列布斯基BorisoglebskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borisoglebsky",
		TitleName: "鲍里索格列布斯基",
		TitleCode: "b_borisoglebsky",
	}
}
