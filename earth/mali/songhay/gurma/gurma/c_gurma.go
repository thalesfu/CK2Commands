package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurmaCounty interface {
	feud.County
	BBilanga比朗加() feud.Barony
	BGurma古尔玛() feud.Barony
	BHalogo哈洛戈() feud.Barony
	BHounde洪代() feud.Barony
	BNungu努恩古() feud.Barony
	BTenkudugo滕科多戈() feud.Barony
	BWagulo瓦古洛() feud.Barony
	BYrgo伊尔戈() feud.Barony
	BZenande泽南德() feud.Barony
	BZorgho佐尔戈() feud.Barony
}

type 古尔玛GurmaCounty struct {
	feud.BaseCounty
	比朗加Bilanga    feud.Barony
	古尔玛Gurma      feud.Barony
	哈洛戈Halogo     feud.Barony
	洪代Hounde      feud.Barony
	努恩古Nungu      feud.Barony
	滕科多戈Tenkudugo feud.Barony
	瓦古洛Wagulo     feud.Barony
	伊尔戈Yrgo       feud.Barony
	泽南德Zenande    feud.Barony
	佐尔戈Zorgho     feud.Barony
}

func (c *古尔玛GurmaCounty) BBilanga比朗加() feud.Barony {
	return c.比朗加Bilanga
}

func (c *古尔玛GurmaCounty) BGurma古尔玛() feud.Barony {
	return c.古尔玛Gurma
}

func (c *古尔玛GurmaCounty) BHalogo哈洛戈() feud.Barony {
	return c.哈洛戈Halogo
}

func (c *古尔玛GurmaCounty) BHounde洪代() feud.Barony {
	return c.洪代Hounde
}

func (c *古尔玛GurmaCounty) BNungu努恩古() feud.Barony {
	return c.努恩古Nungu
}

func (c *古尔玛GurmaCounty) BTenkudugo滕科多戈() feud.Barony {
	return c.滕科多戈Tenkudugo
}

func (c *古尔玛GurmaCounty) BWagulo瓦古洛() feud.Barony {
	return c.瓦古洛Wagulo
}

func (c *古尔玛GurmaCounty) BYrgo伊尔戈() feud.Barony {
	return c.伊尔戈Yrgo
}

func (c *古尔玛GurmaCounty) BZenande泽南德() feud.Barony {
	return c.泽南德Zenande
}

func (c *古尔玛GurmaCounty) BZorgho佐尔戈() feud.Barony {
	return c.佐尔戈Zorgho
}

var CGurma古尔玛 GurmaCounty = &古尔玛GurmaCounty{}

func init() {
	f := CGurma古尔玛.(*古尔玛GurmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "928",
		Title:     "gurma",
		TitleName: "古尔玛",
		TitleCode: "c_gurma",
		Baronies:  map[string]feud.Barony{},
	}

	f.比朗加Bilanga = BBilanga比朗加
	f.比朗加Bilanga.SetParent(f)

	f.古尔玛Gurma = BGurma古尔玛
	f.古尔玛Gurma.SetParent(f)

	f.哈洛戈Halogo = BHalogo哈洛戈
	f.哈洛戈Halogo.SetParent(f)

	f.洪代Hounde = BHounde洪代
	f.洪代Hounde.SetParent(f)

	f.努恩古Nungu = BNungu努恩古
	f.努恩古Nungu.SetParent(f)

	f.滕科多戈Tenkudugo = BTenkudugo滕科多戈
	f.滕科多戈Tenkudugo.SetParent(f)

	f.瓦古洛Wagulo = BWagulo瓦古洛
	f.瓦古洛Wagulo.SetParent(f)

	f.伊尔戈Yrgo = BYrgo伊尔戈
	f.伊尔戈Yrgo.SetParent(f)

	f.泽南德Zenande = BZenande泽南德
	f.泽南德Zenande.SetParent(f)

	f.佐尔戈Zorgho = BZorgho佐尔戈
	f.佐尔戈Zorgho.SetParent(f)

}
