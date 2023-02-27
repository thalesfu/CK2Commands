package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查陀尔坎德ChatorkhandBarony struct {
	feud.BaseBarony
}

var BChatorkhand查陀尔坎德 feud.Barony = &查陀尔坎德ChatorkhandBarony{}

func init() {
    f := BChatorkhand查陀尔坎德.(*查陀尔坎德ChatorkhandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatorkhand",
		TitleName: "查陀尔坎德",
		TitleCode: "b_chatorkhand",
	}
}
