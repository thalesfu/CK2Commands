package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉姆里斯KaramlishBarony struct {
	feud.BaseBarony
}

var BKaramlish卡拉姆里斯 feud.Barony = &卡拉姆里斯KaramlishBarony{}

func init() {
	f := BKaramlish卡拉姆里斯.(*卡拉姆里斯KaramlishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karamlish",
		TitleName: "卡拉姆里斯",
		TitleCode: "b_karamlish",
	}
}
