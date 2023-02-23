package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeedsCounty interface {
	feud.County
	BConisbrough科尼斯伯勒() feud.Barony
	BDoncaster唐克斯特() feud.Barony
	BLeeds利兹() feud.Barony
	BPontefract庞蒂弗拉克特() feud.Barony
	BSelby塞尔比() feud.Barony
	BSheffield谢菲尔德() feud.Barony
	BSherburn舍本() feud.Barony
}

type 利兹LeedsCounty struct {
	feud.BaseCounty
	科尼斯伯勒Conisbrough feud.Barony
	唐克斯特Doncaster    feud.Barony
	利兹Leeds          feud.Barony
	庞蒂弗拉克特Pontefract feud.Barony
	塞尔比Selby         feud.Barony
	谢菲尔德Sheffield    feud.Barony
	舍本Sherburn       feud.Barony
}

func (c *利兹LeedsCounty) BConisbrough科尼斯伯勒() feud.Barony {
	return c.科尼斯伯勒Conisbrough
}

func (c *利兹LeedsCounty) BDoncaster唐克斯特() feud.Barony {
	return c.唐克斯特Doncaster
}

func (c *利兹LeedsCounty) BLeeds利兹() feud.Barony {
	return c.利兹Leeds
}

func (c *利兹LeedsCounty) BPontefract庞蒂弗拉克特() feud.Barony {
	return c.庞蒂弗拉克特Pontefract
}

func (c *利兹LeedsCounty) BSelby塞尔比() feud.Barony {
	return c.塞尔比Selby
}

func (c *利兹LeedsCounty) BSheffield谢菲尔德() feud.Barony {
	return c.谢菲尔德Sheffield
}

func (c *利兹LeedsCounty) BSherburn舍本() feud.Barony {
	return c.舍本Sherburn
}

var CLeeds利兹 LeedsCounty = &利兹LeedsCounty{}

func init() {
	f := CLeeds利兹.(*利兹LeedsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1942",
		Title:     "leeds",
		TitleName: "利兹",
		TitleCode: "c_leeds",
		Baronies:  map[string]feud.Barony{},
	}

	f.科尼斯伯勒Conisbrough = BConisbrough科尼斯伯勒
	f.科尼斯伯勒Conisbrough.SetParent(f)

	f.唐克斯特Doncaster = BDoncaster唐克斯特
	f.唐克斯特Doncaster.SetParent(f)

	f.利兹Leeds = BLeeds利兹
	f.利兹Leeds.SetParent(f)

	f.庞蒂弗拉克特Pontefract = BPontefract庞蒂弗拉克特
	f.庞蒂弗拉克特Pontefract.SetParent(f)

	f.塞尔比Selby = BSelby塞尔比
	f.塞尔比Selby.SetParent(f)

	f.谢菲尔德Sheffield = BSheffield谢菲尔德
	f.谢菲尔德Sheffield.SetParent(f)

	f.舍本Sherburn = BSherburn舍本
	f.舍本Sherburn.SetParent(f)

}
