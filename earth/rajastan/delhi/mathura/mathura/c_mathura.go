package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MathuraCounty interface {
	feud.County
	BAgra阿格罗() feud.Barony
	BDhanauli诃努利() feud.Barony
	BGovardhan乔跋檀() feud.Barony
	BKhanwa汗瓦() feud.Barony
	BKrishnajanmabhoomi奎师那诞生地() feud.Barony
	BMankameshwar摩那迦迷湿伐罗() feud.Barony
	BMathura秣菟罗() feud.Barony
	BVrindavan勿邻陀林() feud.Barony
}

type 秣菟罗MathuraCounty struct {
	feud.BaseCounty
	阿格罗Agra                  feud.Barony
	诃努利Dhanauli              feud.Barony
	乔跋檀Govardhan             feud.Barony
	汗瓦Khanwa                 feud.Barony
	奎师那诞生地Krishnajanmabhoomi feud.Barony
	摩那迦迷湿伐罗Mankameshwar      feud.Barony
	秣菟罗Mathura               feud.Barony
	勿邻陀林Vrindavan            feud.Barony
}

func (c *秣菟罗MathuraCounty) BAgra阿格罗() feud.Barony {
	return c.阿格罗Agra
}

func (c *秣菟罗MathuraCounty) BDhanauli诃努利() feud.Barony {
	return c.诃努利Dhanauli
}

func (c *秣菟罗MathuraCounty) BGovardhan乔跋檀() feud.Barony {
	return c.乔跋檀Govardhan
}

func (c *秣菟罗MathuraCounty) BKhanwa汗瓦() feud.Barony {
	return c.汗瓦Khanwa
}

func (c *秣菟罗MathuraCounty) BKrishnajanmabhoomi奎师那诞生地() feud.Barony {
	return c.奎师那诞生地Krishnajanmabhoomi
}

func (c *秣菟罗MathuraCounty) BMankameshwar摩那迦迷湿伐罗() feud.Barony {
	return c.摩那迦迷湿伐罗Mankameshwar
}

func (c *秣菟罗MathuraCounty) BMathura秣菟罗() feud.Barony {
	return c.秣菟罗Mathura
}

func (c *秣菟罗MathuraCounty) BVrindavan勿邻陀林() feud.Barony {
	return c.勿邻陀林Vrindavan
}

var CMathura秣菟罗 MathuraCounty = &秣菟罗MathuraCounty{}

func init() {
	f := CMathura秣菟罗.(*秣菟罗MathuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1359",
		Title:     "mathura",
		TitleName: "秣菟罗",
		TitleCode: "c_mathura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格罗Agra = BAgra阿格罗
	f.阿格罗Agra.SetParent(f)

	f.诃努利Dhanauli = BDhanauli诃努利
	f.诃努利Dhanauli.SetParent(f)

	f.乔跋檀Govardhan = BGovardhan乔跋檀
	f.乔跋檀Govardhan.SetParent(f)

	f.汗瓦Khanwa = BKhanwa汗瓦
	f.汗瓦Khanwa.SetParent(f)

	f.奎师那诞生地Krishnajanmabhoomi = BKrishnajanmabhoomi奎师那诞生地
	f.奎师那诞生地Krishnajanmabhoomi.SetParent(f)

	f.摩那迦迷湿伐罗Mankameshwar = BMankameshwar摩那迦迷湿伐罗
	f.摩那迦迷湿伐罗Mankameshwar.SetParent(f)

	f.秣菟罗Mathura = BMathura秣菟罗
	f.秣菟罗Mathura.SetParent(f)

	f.勿邻陀林Vrindavan = BVrindavan勿邻陀林
	f.勿邻陀林Vrindavan.SetParent(f)

}
