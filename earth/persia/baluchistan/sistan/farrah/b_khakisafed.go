package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈基萨法德KhakisafedBarony struct {
	feud.BaseBarony
}

var BKhakisafed哈基萨法德 feud.Barony = &哈基萨法德KhakisafedBarony{}

func init() {
    f := BKhakisafed哈基萨法德.(*哈基萨法德KhakisafedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khakisafed",
		TitleName: "哈基萨法德",
		TitleCode: "b_khakisafed",
	}
}
