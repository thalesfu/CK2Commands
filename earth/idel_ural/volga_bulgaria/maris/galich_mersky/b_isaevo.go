package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊萨耶沃IsaevoBarony struct {
	feud.BaseBarony
}

var BIsaevo伊萨耶沃 feud.Barony = &伊萨耶沃IsaevoBarony{}

func init() {
    f := BIsaevo伊萨耶沃.(*伊萨耶沃IsaevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isaevo",
		TitleName: "伊萨耶沃",
		TitleCode: "b_isaevo",
	}
}
