package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 焦婆德JobatBarony struct {
	feud.BaseBarony
}

var BJobat焦婆德 feud.Barony = &焦婆德JobatBarony{}

func init() {
	f := BJobat焦婆德.(*焦婆德JobatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jobat",
		TitleName: "焦婆德",
		TitleCode: "b_jobat",
	}
}
