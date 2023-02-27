package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛涅CoulogneBarony struct {
	feud.BaseBarony
}

var BCoulogne库洛涅 feud.Barony = &库洛涅CoulogneBarony{}

func init() {
    f := BCoulogne库洛涅.(*库洛涅CoulogneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coulogne",
		TitleName: "库洛涅",
		TitleCode: "b_coulogne",
	}
}
