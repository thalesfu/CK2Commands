package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥沙坎OshakanBarony struct {
	feud.BaseBarony
}

var BOshakan奥沙坎 feud.Barony = &奥沙坎OshakanBarony{}

func init() {
    f := BOshakan奥沙坎.(*奥沙坎OshakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oshakan",
		TitleName: "奥沙坎",
		TitleCode: "b_oshakan",
	}
}
