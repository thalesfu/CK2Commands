package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲尔斯滕贝格FurstenbergBarony struct {
	feud.BaseBarony
}

var BFurstenberg菲尔斯滕贝格 feud.Barony = &菲尔斯滕贝格FurstenbergBarony{}

func init() {
    f := BFurstenberg菲尔斯滕贝格.(*菲尔斯滕贝格FurstenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furstenberg",
		TitleName: "菲尔斯滕贝格",
		TitleCode: "b_furstenberg",
	}
}
