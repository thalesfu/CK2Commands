package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PannagalluCounty interface {
	feud.County
	BBodla佛罗() feud.Barony
	BCharakonda车罗军荼() feud.Barony
	BKunda军陀() feud.Barony
	BMalikpur摩力补罗() feud.Barony
	BMutudugu穆图杜古() feud.Barony
	BNagarkurnool纳加尔库尔诺奥尔() feud.Barony
	BPannagallu般那伽楼() feud.Barony
}

type 般那伽楼PannagalluCounty struct {
	feud.BaseCounty
	佛罗Bodla              feud.Barony
	车罗军荼Charakonda       feud.Barony
	军陀Kunda              feud.Barony
	摩力补罗Malikpur         feud.Barony
	穆图杜古Mutudugu         feud.Barony
	纳加尔库尔诺奥尔Nagarkurnool feud.Barony
	般那伽楼Pannagallu       feud.Barony
}

func (c *般那伽楼PannagalluCounty) BBodla佛罗() feud.Barony {
	return c.佛罗Bodla
}

func (c *般那伽楼PannagalluCounty) BCharakonda车罗军荼() feud.Barony {
	return c.车罗军荼Charakonda
}

func (c *般那伽楼PannagalluCounty) BKunda军陀() feud.Barony {
	return c.军陀Kunda
}

func (c *般那伽楼PannagalluCounty) BMalikpur摩力补罗() feud.Barony {
	return c.摩力补罗Malikpur
}

func (c *般那伽楼PannagalluCounty) BMutudugu穆图杜古() feud.Barony {
	return c.穆图杜古Mutudugu
}

func (c *般那伽楼PannagalluCounty) BNagarkurnool纳加尔库尔诺奥尔() feud.Barony {
	return c.纳加尔库尔诺奥尔Nagarkurnool
}

func (c *般那伽楼PannagalluCounty) BPannagallu般那伽楼() feud.Barony {
	return c.般那伽楼Pannagallu
}

var CPannagallu般那伽楼 PannagalluCounty = &般那伽楼PannagalluCounty{}

func init() {
	f := CPannagallu般那伽楼.(*般那伽楼PannagalluCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1211",
		Title:     "pannagallu",
		TitleName: "般那伽楼",
		TitleCode: "c_pannagallu",
		Baronies:  map[string]feud.Barony{},
	}

	f.佛罗Bodla = BBodla佛罗
	f.佛罗Bodla.SetParent(f)

	f.车罗军荼Charakonda = BCharakonda车罗军荼
	f.车罗军荼Charakonda.SetParent(f)

	f.军陀Kunda = BKunda军陀
	f.军陀Kunda.SetParent(f)

	f.摩力补罗Malikpur = BMalikpur摩力补罗
	f.摩力补罗Malikpur.SetParent(f)

	f.穆图杜古Mutudugu = BMutudugu穆图杜古
	f.穆图杜古Mutudugu.SetParent(f)

	f.纳加尔库尔诺奥尔Nagarkurnool = BNagarkurnool纳加尔库尔诺奥尔
	f.纳加尔库尔诺奥尔Nagarkurnool.SetParent(f)

	f.般那伽楼Pannagallu = BPannagallu般那伽楼
	f.般那伽楼Pannagallu.SetParent(f)

}
