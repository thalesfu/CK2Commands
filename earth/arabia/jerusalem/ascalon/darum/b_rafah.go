package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉法RafahBarony struct {
	feud.BaseBarony
}

var BRafah拉法 feud.Barony = &拉法RafahBarony{}

func init() {
    f := BRafah拉法.(*拉法RafahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rafah",
		TitleName: "拉法",
		TitleCode: "b_rafah",
	}
}
