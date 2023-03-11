package el-arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布阿维吉拉AbuaweigilaBarony struct {
	feud.BaseBarony
}

var BAbuaweigila阿布阿维吉拉 feud.Barony = &阿布阿维吉拉AbuaweigilaBarony{}

func init() {
    f := BAbuaweigila阿布阿维吉拉.(*阿布阿维吉拉AbuaweigilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abuaweigila",
		TitleName: "阿布阿维吉拉",
		TitleCode: "b_abuaweigila",
	}
}
