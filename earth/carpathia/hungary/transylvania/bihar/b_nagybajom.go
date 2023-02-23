package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉鲍约姆NagybajomBarony struct {
	feud.BaseBarony
}

var BNagybajom瑙吉鲍约姆 feud.Barony = &瑙吉鲍约姆NagybajomBarony{}

func init() {
	f := BNagybajom瑙吉鲍约姆.(*瑙吉鲍约姆NagybajomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagybajom",
		TitleName: "瑙吉鲍约姆",
		TitleCode: "b_nagybajom",
	}
}
