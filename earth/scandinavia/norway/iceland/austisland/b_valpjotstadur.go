package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔肖夫斯塔泽ValpjotstadurBarony struct {
	feud.BaseBarony
}

var BValpjotstadur瓦尔肖夫斯塔泽 feud.Barony = &瓦尔肖夫斯塔泽ValpjotstadurBarony{}

func init() {
    f := BValpjotstadur瓦尔肖夫斯塔泽.(*瓦尔肖夫斯塔泽ValpjotstadurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valpjotstadur",
		TitleName: "瓦尔肖夫斯塔泽",
		TitleCode: "b_valpjotstadur",
	}
}
