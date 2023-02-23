package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀阿富尔QaafurBarony struct {
	feud.BaseBarony
}

var BQaafur喀阿富尔 feud.Barony = &喀阿富尔QaafurBarony{}

func init() {
	f := BQaafur喀阿富尔.(*喀阿富尔QaafurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaafur",
		TitleName: "喀阿富尔",
		TitleCode: "b_qaafur",
	}
}
