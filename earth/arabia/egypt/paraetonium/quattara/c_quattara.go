package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QuattaraCounty interface {
	feud.County
	BAbdannabi阿卜杜勒纳比() feud.Barony
	BAlamelshwawish阿拉姆沙锡() feud.Barony
	BCaraoasis加拉() feud.Barony
	BDayr代尔() feud.Barony
	BQaretagnes喀雷特阿格尼斯() feud.Barony
	BQuattara盖塔拉() feud.Barony
	BSiwa锡瓦() feud.Barony
	BZiwaelbahari济瓦巴哈里() feud.Barony
}

type 盖塔拉QuattaraCounty struct {
	feud.BaseCounty
	阿卜杜勒纳比Abdannabi     feud.Barony
	阿拉姆沙锡Alamelshwawish feud.Barony
	加拉Caraoasis         feud.Barony
	代尔Dayr              feud.Barony
	喀雷特阿格尼斯Qaretagnes   feud.Barony
	盖塔拉Quattara         feud.Barony
	锡瓦Siwa              feud.Barony
	济瓦巴哈里Ziwaelbahari   feud.Barony
}

func (c *盖塔拉QuattaraCounty) BAbdannabi阿卜杜勒纳比() feud.Barony {
	return c.阿卜杜勒纳比Abdannabi
}

func (c *盖塔拉QuattaraCounty) BAlamelshwawish阿拉姆沙锡() feud.Barony {
	return c.阿拉姆沙锡Alamelshwawish
}

func (c *盖塔拉QuattaraCounty) BCaraoasis加拉() feud.Barony {
	return c.加拉Caraoasis
}

func (c *盖塔拉QuattaraCounty) BDayr代尔() feud.Barony {
	return c.代尔Dayr
}

func (c *盖塔拉QuattaraCounty) BQaretagnes喀雷特阿格尼斯() feud.Barony {
	return c.喀雷特阿格尼斯Qaretagnes
}

func (c *盖塔拉QuattaraCounty) BQuattara盖塔拉() feud.Barony {
	return c.盖塔拉Quattara
}

func (c *盖塔拉QuattaraCounty) BSiwa锡瓦() feud.Barony {
	return c.锡瓦Siwa
}

func (c *盖塔拉QuattaraCounty) BZiwaelbahari济瓦巴哈里() feud.Barony {
	return c.济瓦巴哈里Ziwaelbahari
}

var CQuattara盖塔拉 QuattaraCounty = &盖塔拉QuattaraCounty{}

func init() {
	f := CQuattara盖塔拉.(*盖塔拉QuattaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "803",
		Title:     "quattara",
		TitleName: "盖塔拉",
		TitleCode: "c_quattara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜杜勒纳比Abdannabi = BAbdannabi阿卜杜勒纳比
	f.阿卜杜勒纳比Abdannabi.SetParent(f)

	f.阿拉姆沙锡Alamelshwawish = BAlamelshwawish阿拉姆沙锡
	f.阿拉姆沙锡Alamelshwawish.SetParent(f)

	f.加拉Caraoasis = BCaraoasis加拉
	f.加拉Caraoasis.SetParent(f)

	f.代尔Dayr = BDayr代尔
	f.代尔Dayr.SetParent(f)

	f.喀雷特阿格尼斯Qaretagnes = BQaretagnes喀雷特阿格尼斯
	f.喀雷特阿格尼斯Qaretagnes.SetParent(f)

	f.盖塔拉Quattara = BQuattara盖塔拉
	f.盖塔拉Quattara.SetParent(f)

	f.锡瓦Siwa = BSiwa锡瓦
	f.锡瓦Siwa.SetParent(f)

	f.济瓦巴哈里Ziwaelbahari = BZiwaelbahari济瓦巴哈里
	f.济瓦巴哈里Ziwaelbahari.SetParent(f)

}
