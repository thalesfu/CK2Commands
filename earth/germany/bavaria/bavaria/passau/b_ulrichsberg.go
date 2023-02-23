package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔里希斯贝格UlrichsbergBarony struct {
	feud.BaseBarony
}

var BUlrichsberg乌尔里希斯贝格 feud.Barony = &乌尔里希斯贝格UlrichsbergBarony{}

func init() {
	f := BUlrichsberg乌尔里希斯贝格.(*乌尔里希斯贝格UlrichsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulrichsberg",
		TitleName: "乌尔里希斯贝格",
		TitleCode: "b_ulrichsberg",
	}
}
