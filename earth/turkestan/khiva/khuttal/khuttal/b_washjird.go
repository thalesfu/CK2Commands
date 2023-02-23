package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦希尔德WashjirdBarony struct {
	feud.BaseBarony
}

var BWashjird瓦希尔德 feud.Barony = &瓦希尔德WashjirdBarony{}

func init() {
	f := BWashjird瓦希尔德.(*瓦希尔德WashjirdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "washjird",
		TitleName: "瓦希尔德",
		TitleCode: "b_washjird",
	}
}
