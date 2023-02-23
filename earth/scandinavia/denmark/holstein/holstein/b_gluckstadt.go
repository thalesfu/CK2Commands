package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格吕克施塔特GluckstadtBarony struct {
	feud.BaseBarony
}

var BGluckstadt格吕克施塔特 feud.Barony = &格吕克施塔特GluckstadtBarony{}

func init() {
	f := BGluckstadt格吕克施塔特.(*格吕克施塔特GluckstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gluckstadt",
		TitleName: "格吕克施塔特",
		TitleCode: "b_gluckstadt",
	}
}
