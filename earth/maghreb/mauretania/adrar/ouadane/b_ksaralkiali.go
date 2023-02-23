package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基阿里堡KsaralkialiBarony struct {
	feud.BaseBarony
}

var BKsaralkiali基阿里堡 feud.Barony = &基阿里堡KsaralkialiBarony{}

func init() {
	f := BKsaralkiali基阿里堡.(*基阿里堡KsaralkialiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ksaralkiali",
		TitleName: "基阿里堡",
		TitleCode: "b_ksaralkiali",
	}
}
