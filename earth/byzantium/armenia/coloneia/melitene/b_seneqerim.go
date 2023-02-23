package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尼格米SeneqerimBarony struct {
	feud.BaseBarony
}

var BSeneqerim瑟尼格米 feud.Barony = &瑟尼格米SeneqerimBarony{}

func init() {
	f := BSeneqerim瑟尼格米.(*瑟尼格米SeneqerimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seneqerim",
		TitleName: "瑟尼格米",
		TitleCode: "b_seneqerim",
	}
}
