package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉菲拉堡QasrfarfraBarony struct {
	feud.BaseBarony
}

var BQasrfarfra费拉菲拉堡 feud.Barony = &费拉菲拉堡QasrfarfraBarony{}

func init() {
    f := BQasrfarfra费拉菲拉堡.(*费拉菲拉堡QasrfarfraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasrfarfra",
		TitleName: "费拉菲拉堡",
		TitleCode: "b_qasrfarfra",
	}
}
