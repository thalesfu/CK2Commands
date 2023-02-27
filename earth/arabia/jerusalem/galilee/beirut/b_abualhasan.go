package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布_哈桑AbualhasanBarony struct {
	feud.BaseBarony
}

var BAbualhasan阿布_哈桑 feud.Barony = &阿布_哈桑AbualhasanBarony{}

func init() {
    f := BAbualhasan阿布_哈桑.(*阿布_哈桑AbualhasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abualhasan",
		TitleName: "阿布_哈桑",
		TitleCode: "b_abualhasan",
	}
}
