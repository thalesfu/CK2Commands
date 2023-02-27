package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波梯耶罗PathyarBarony struct {
	feud.BaseBarony
}

var BPathyar波梯耶罗 feud.Barony = &波梯耶罗PathyarBarony{}

func init() {
    f := BPathyar波梯耶罗.(*波梯耶罗PathyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pathyar",
		TitleName: "波梯耶罗",
		TitleCode: "b_pathyar",
	}
}
