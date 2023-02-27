package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉巴尼奥NagybanyaBarony struct {
	feud.BaseBarony
}

var BNagybanya瑙吉巴尼奥 feud.Barony = &瑙吉巴尼奥NagybanyaBarony{}

func init() {
    f := BNagybanya瑙吉巴尼奥.(*瑙吉巴尼奥NagybanyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagybanya",
		TitleName: "瑙吉巴尼奥",
		TitleCode: "b_nagybanya",
	}
}
