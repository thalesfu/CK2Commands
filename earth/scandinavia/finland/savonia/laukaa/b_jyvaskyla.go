package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于韦斯屈莱JyvaskylaBarony struct {
	feud.BaseBarony
}

var BJyvaskyla于韦斯屈莱 feud.Barony = &于韦斯屈莱JyvaskylaBarony{}

func init() {
    f := BJyvaskyla于韦斯屈莱.(*于韦斯屈莱JyvaskylaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jyvaskyla",
		TitleName: "于韦斯屈莱",
		TitleCode: "b_jyvaskyla",
	}
}
