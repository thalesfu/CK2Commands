package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖塔拉QuattaraBarony struct {
	feud.BaseBarony
}

var BQuattara盖塔拉 feud.Barony = &盖塔拉QuattaraBarony{}

func init() {
    f := BQuattara盖塔拉.(*盖塔拉QuattaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quattara",
		TitleName: "盖塔拉",
		TitleCode: "b_quattara",
	}
}
