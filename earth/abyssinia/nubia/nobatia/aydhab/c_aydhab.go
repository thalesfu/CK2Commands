package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AydhabCounty interface {
	feud.County
	BAydhab阿伊迪哈卜() feud.Barony
	BElba艾尔巴() feud.Barony
	BHalayeb哈拉伊卜() feud.Barony
	BMarsaalam马萨阿拉姆() feud.Barony
	BShalateen沙腊腾() feud.Barony
	BTroglodytica特罗葛罗蒂提卡() feud.Barony
	BZabargad宰拜尔杰德() feud.Barony
}

type 阿伊迪哈卜AydhabCounty struct {
	feud.BaseCounty
	阿伊迪哈卜Aydhab         feud.Barony
	艾尔巴Elba             feud.Barony
	哈拉伊卜Halayeb         feud.Barony
	马萨阿拉姆Marsaalam      feud.Barony
	沙腊腾Shalateen        feud.Barony
	特罗葛罗蒂提卡Troglodytica feud.Barony
	宰拜尔杰德Zabargad       feud.Barony
}

func (c *阿伊迪哈卜AydhabCounty) BAydhab阿伊迪哈卜() feud.Barony {
	return c.阿伊迪哈卜Aydhab
}

func (c *阿伊迪哈卜AydhabCounty) BElba艾尔巴() feud.Barony {
	return c.艾尔巴Elba
}

func (c *阿伊迪哈卜AydhabCounty) BHalayeb哈拉伊卜() feud.Barony {
	return c.哈拉伊卜Halayeb
}

func (c *阿伊迪哈卜AydhabCounty) BMarsaalam马萨阿拉姆() feud.Barony {
	return c.马萨阿拉姆Marsaalam
}

func (c *阿伊迪哈卜AydhabCounty) BShalateen沙腊腾() feud.Barony {
	return c.沙腊腾Shalateen
}

func (c *阿伊迪哈卜AydhabCounty) BTroglodytica特罗葛罗蒂提卡() feud.Barony {
	return c.特罗葛罗蒂提卡Troglodytica
}

func (c *阿伊迪哈卜AydhabCounty) BZabargad宰拜尔杰德() feud.Barony {
	return c.宰拜尔杰德Zabargad
}

var CAydhab阿伊迪哈卜 AydhabCounty = &阿伊迪哈卜AydhabCounty{}

func init() {
	f := CAydhab阿伊迪哈卜.(*阿伊迪哈卜AydhabCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1320",
		Title:     "aydhab",
		TitleName: "阿伊迪哈卜",
		TitleCode: "c_aydhab",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伊迪哈卜Aydhab = BAydhab阿伊迪哈卜
	f.阿伊迪哈卜Aydhab.SetParent(f)

	f.艾尔巴Elba = BElba艾尔巴
	f.艾尔巴Elba.SetParent(f)

	f.哈拉伊卜Halayeb = BHalayeb哈拉伊卜
	f.哈拉伊卜Halayeb.SetParent(f)

	f.马萨阿拉姆Marsaalam = BMarsaalam马萨阿拉姆
	f.马萨阿拉姆Marsaalam.SetParent(f)

	f.沙腊腾Shalateen = BShalateen沙腊腾
	f.沙腊腾Shalateen.SetParent(f)

	f.特罗葛罗蒂提卡Troglodytica = BTroglodytica特罗葛罗蒂提卡
	f.特罗葛罗蒂提卡Troglodytica.SetParent(f)

	f.宰拜尔杰德Zabargad = BZabargad宰拜尔杰德
	f.宰拜尔杰德Zabargad.SetParent(f)

}
