package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绮罗乌QorovulBarony struct {
	feud.BaseBarony
}

var BQorovul绮罗乌 feud.Barony = &绮罗乌QorovulBarony{}

func init() {
    f := BQorovul绮罗乌.(*绮罗乌QorovulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qorovul",
		TitleName: "绮罗乌",
		TitleCode: "b_qorovul",
	}
}
