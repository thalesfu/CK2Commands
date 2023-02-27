package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_通托尔UsttuntorBarony struct {
	feud.BaseBarony
}

var BUsttuntor乌斯季_通托尔 feud.Barony = &乌斯季_通托尔UsttuntorBarony{}

func init() {
    f := BUsttuntor乌斯季_通托尔.(*乌斯季_通托尔UsttuntorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usttuntor",
		TitleName: "乌斯季_通托尔",
		TitleCode: "b_usttuntor",
	}
}
