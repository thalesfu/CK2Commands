package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌鲁米耶OroumiehBarony struct {
	feud.BaseBarony
}

var BOroumieh乌鲁米耶 feud.Barony = &乌鲁米耶OroumiehBarony{}

func init() {
    f := BOroumieh乌鲁米耶.(*乌鲁米耶OroumiehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oroumieh",
		TitleName: "乌鲁米耶",
		TitleCode: "b_oroumieh",
	}
}
