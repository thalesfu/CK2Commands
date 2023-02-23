package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KondanaCounty interface {
	feud.County
	BBhimashankara毗摩商羯罗() feud.Barony
	BChakan查甘() feud.Barony
	BJunir朱尼尔() feud.Barony
	BKondana军荼那() feud.Barony
	BPuna普纳() feud.Barony
	BPurandar富兰陀罗() feud.Barony
	BRaigad拉加德() feud.Barony
	BRajmachi罗阇摩支() feud.Barony
	BSudhagad苏陀迦荼() feud.Barony
	BTikona底拘那() feud.Barony
}

type 军荼那KondanaCounty struct {
	feud.BaseCounty
	毗摩商羯罗Bhimashankara feud.Barony
	查甘Chakan           feud.Barony
	朱尼尔Junir           feud.Barony
	军荼那Kondana         feud.Barony
	普纳Puna             feud.Barony
	富兰陀罗Purandar       feud.Barony
	拉加德Raigad          feud.Barony
	罗阇摩支Rajmachi       feud.Barony
	苏陀迦荼Sudhagad       feud.Barony
	底拘那Tikona          feud.Barony
}

func (c *军荼那KondanaCounty) BBhimashankara毗摩商羯罗() feud.Barony {
	return c.毗摩商羯罗Bhimashankara
}

func (c *军荼那KondanaCounty) BChakan查甘() feud.Barony {
	return c.查甘Chakan
}

func (c *军荼那KondanaCounty) BJunir朱尼尔() feud.Barony {
	return c.朱尼尔Junir
}

func (c *军荼那KondanaCounty) BKondana军荼那() feud.Barony {
	return c.军荼那Kondana
}

func (c *军荼那KondanaCounty) BPuna普纳() feud.Barony {
	return c.普纳Puna
}

func (c *军荼那KondanaCounty) BPurandar富兰陀罗() feud.Barony {
	return c.富兰陀罗Purandar
}

func (c *军荼那KondanaCounty) BRaigad拉加德() feud.Barony {
	return c.拉加德Raigad
}

func (c *军荼那KondanaCounty) BRajmachi罗阇摩支() feud.Barony {
	return c.罗阇摩支Rajmachi
}

func (c *军荼那KondanaCounty) BSudhagad苏陀迦荼() feud.Barony {
	return c.苏陀迦荼Sudhagad
}

func (c *军荼那KondanaCounty) BTikona底拘那() feud.Barony {
	return c.底拘那Tikona
}

var CKondana军荼那 KondanaCounty = &军荼那KondanaCounty{}

func init() {
	f := CKondana军荼那.(*军荼那KondanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1213",
		Title:     "kondana",
		TitleName: "军荼那",
		TitleCode: "c_kondana",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗摩商羯罗Bhimashankara = BBhimashankara毗摩商羯罗
	f.毗摩商羯罗Bhimashankara.SetParent(f)

	f.查甘Chakan = BChakan查甘
	f.查甘Chakan.SetParent(f)

	f.朱尼尔Junir = BJunir朱尼尔
	f.朱尼尔Junir.SetParent(f)

	f.军荼那Kondana = BKondana军荼那
	f.军荼那Kondana.SetParent(f)

	f.普纳Puna = BPuna普纳
	f.普纳Puna.SetParent(f)

	f.富兰陀罗Purandar = BPurandar富兰陀罗
	f.富兰陀罗Purandar.SetParent(f)

	f.拉加德Raigad = BRaigad拉加德
	f.拉加德Raigad.SetParent(f)

	f.罗阇摩支Rajmachi = BRajmachi罗阇摩支
	f.罗阇摩支Rajmachi.SetParent(f)

	f.苏陀迦荼Sudhagad = BSudhagad苏陀迦荼
	f.苏陀迦荼Sudhagad.SetParent(f)

	f.底拘那Tikona = BTikona底拘那
	f.底拘那Tikona.SetParent(f)

}
