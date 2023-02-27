package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德勒格尔杭爱DelgerkhangaiBarony struct {
	feud.BaseBarony
}

var BDelgerkhangai德勒格尔杭爱 feud.Barony = &德勒格尔杭爱DelgerkhangaiBarony{}

func init() {
    f := BDelgerkhangai德勒格尔杭爱.(*德勒格尔杭爱DelgerkhangaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delgerkhangai",
		TitleName: "德勒格尔杭爱",
		TitleCode: "b_delgerkhangai",
	}
}
