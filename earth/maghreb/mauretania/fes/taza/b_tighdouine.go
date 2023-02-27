package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂格杜因TighdouineBarony struct {
	feud.BaseBarony
}

var BTighdouine蒂格杜因 feud.Barony = &蒂格杜因TighdouineBarony{}

func init() {
    f := BTighdouine蒂格杜因.(*蒂格杜因TighdouineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tighdouine",
		TitleName: "蒂格杜因",
		TitleCode: "b_tighdouine",
	}
}
