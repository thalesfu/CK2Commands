package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 上沃洛乔克Vyshny_volochyokBarony struct {
	feud.BaseBarony
}

var BVyshny_volochyok上沃洛乔克 feud.Barony = &上沃洛乔克Vyshny_volochyokBarony{}

func init() {
    f := BVyshny_volochyok上沃洛乔克.(*上沃洛乔克Vyshny_volochyokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyshny_volochyok",
		TitleName: "上沃洛乔克",
		TitleCode: "b_vyshny_volochyok",
	}
}
