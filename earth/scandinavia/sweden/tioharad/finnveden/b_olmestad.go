package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄尔默斯塔德OlmestadBarony struct {
	feud.BaseBarony
}

var BOlmestad厄尔默斯塔德 feud.Barony = &厄尔默斯塔德OlmestadBarony{}

func init() {
	f := BOlmestad厄尔默斯塔德.(*厄尔默斯塔德OlmestadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olmestad",
		TitleName: "厄尔默斯塔德",
		TitleCode: "b_olmestad",
	}
}
