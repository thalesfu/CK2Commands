package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QayaliqCounty interface {
	feud.County
	BHorgos霍尔果斯() feud.Barony
	BQayaliq海押立() feud.Barony
	BSadyr萨德尔() feud.Barony
	BShagan沙甘() feud.Barony
	BShubar舒巴尔() feud.Barony
}

type 海押立QayaliqCounty struct {
	feud.BaseCounty
	霍尔果斯Horgos feud.Barony
	海押立Qayaliq feud.Barony
	萨德尔Sadyr   feud.Barony
	沙甘Shagan   feud.Barony
	舒巴尔Shubar  feud.Barony
}

func (c *海押立QayaliqCounty) BHorgos霍尔果斯() feud.Barony {
	return c.霍尔果斯Horgos
}

func (c *海押立QayaliqCounty) BQayaliq海押立() feud.Barony {
	return c.海押立Qayaliq
}

func (c *海押立QayaliqCounty) BSadyr萨德尔() feud.Barony {
	return c.萨德尔Sadyr
}

func (c *海押立QayaliqCounty) BShagan沙甘() feud.Barony {
	return c.沙甘Shagan
}

func (c *海押立QayaliqCounty) BShubar舒巴尔() feud.Barony {
	return c.舒巴尔Shubar
}

var CQayaliq海押立 QayaliqCounty = &海押立QayaliqCounty{}

func init() {
	f := CQayaliq海押立.(*海押立QayaliqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1799",
		Title:     "qayaliq",
		TitleName: "海押立",
		TitleCode: "c_qayaliq",
		Baronies:  map[string]feud.Barony{},
	}

	f.霍尔果斯Horgos = BHorgos霍尔果斯
	f.霍尔果斯Horgos.SetParent(f)

	f.海押立Qayaliq = BQayaliq海押立
	f.海押立Qayaliq.SetParent(f)

	f.萨德尔Sadyr = BSadyr萨德尔
	f.萨德尔Sadyr.SetParent(f)

	f.沙甘Shagan = BShagan沙甘
	f.沙甘Shagan.SetParent(f)

	f.舒巴尔Shubar = BShubar舒巴尔
	f.舒巴尔Shubar.SetParent(f)

}
