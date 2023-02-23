package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VestfoldCounty interface {
	feud.County
	BArendall阿兰达拉() feud.Barony
	BHorten霍滕() feud.Barony
	BKaupang凯于庞() feud.Barony
	BNore努勒() feud.Barony
	BRe雷() feud.Barony
	BSkiringssal斯基林格斯萨() feud.Barony
	BTonsberg滕斯贝格() feud.Barony
	BUvdal乌夫达() feud.Barony
}

type 西福尔VestfoldCounty struct {
	feud.BaseCounty
	阿兰达拉Arendall      feud.Barony
	霍滕Horten          feud.Barony
	凯于庞Kaupang        feud.Barony
	努勒Nore            feud.Barony
	雷Re               feud.Barony
	斯基林格斯萨Skiringssal feud.Barony
	滕斯贝格Tonsberg      feud.Barony
	乌夫达Uvdal          feud.Barony
}

func (c *西福尔VestfoldCounty) BArendall阿兰达拉() feud.Barony {
	return c.阿兰达拉Arendall
}

func (c *西福尔VestfoldCounty) BHorten霍滕() feud.Barony {
	return c.霍滕Horten
}

func (c *西福尔VestfoldCounty) BKaupang凯于庞() feud.Barony {
	return c.凯于庞Kaupang
}

func (c *西福尔VestfoldCounty) BNore努勒() feud.Barony {
	return c.努勒Nore
}

func (c *西福尔VestfoldCounty) BRe雷() feud.Barony {
	return c.雷Re
}

func (c *西福尔VestfoldCounty) BSkiringssal斯基林格斯萨() feud.Barony {
	return c.斯基林格斯萨Skiringssal
}

func (c *西福尔VestfoldCounty) BTonsberg滕斯贝格() feud.Barony {
	return c.滕斯贝格Tonsberg
}

func (c *西福尔VestfoldCounty) BUvdal乌夫达() feud.Barony {
	return c.乌夫达Uvdal
}

var CVestfold西福尔 VestfoldCounty = &西福尔VestfoldCounty{}

func init() {
	f := CVestfold西福尔.(*西福尔VestfoldCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "271",
		Title:     "vestfold",
		TitleName: "西福尔",
		TitleCode: "c_vestfold",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿兰达拉Arendall = BArendall阿兰达拉
	f.阿兰达拉Arendall.SetParent(f)

	f.霍滕Horten = BHorten霍滕
	f.霍滕Horten.SetParent(f)

	f.凯于庞Kaupang = BKaupang凯于庞
	f.凯于庞Kaupang.SetParent(f)

	f.努勒Nore = BNore努勒
	f.努勒Nore.SetParent(f)

	f.雷Re = BRe雷
	f.雷Re.SetParent(f)

	f.斯基林格斯萨Skiringssal = BSkiringssal斯基林格斯萨
	f.斯基林格斯萨Skiringssal.SetParent(f)

	f.滕斯贝格Tonsberg = BTonsberg滕斯贝格
	f.滕斯贝格Tonsberg.SetParent(f)

	f.乌夫达Uvdal = BUvdal乌夫达
	f.乌夫达Uvdal.SetParent(f)

}
