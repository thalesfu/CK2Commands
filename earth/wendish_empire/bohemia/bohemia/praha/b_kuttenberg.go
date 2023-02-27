package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库滕贝格KuttenbergBarony struct {
	feud.BaseBarony
}

var BKuttenberg库滕贝格 feud.Barony = &库滕贝格KuttenbergBarony{}

func init() {
    f := BKuttenberg库滕贝格.(*库滕贝格KuttenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuttenberg",
		TitleName: "库滕贝格",
		TitleCode: "b_kuttenberg",
	}
}
