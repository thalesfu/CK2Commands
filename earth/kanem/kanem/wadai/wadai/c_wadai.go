package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WadaiCounty interface {
	feud.County
	BAbeche阿贝歇() feud.Barony
	BAsafik阿萨菲克() feud.Barony
	BGourmoga古尔莫加() feud.Barony
	BKoulkoule库尔库莱() feud.Barony
	BNimro尼姆罗() feud.Barony
	BWara瓦拉() feud.Barony
	BYue于埃() feud.Barony
}

type 瓦达伊WadaiCounty struct {
	feud.BaseCounty
	阿贝歇Abeche     feud.Barony
	阿萨菲克Asafik    feud.Barony
	古尔莫加Gourmoga  feud.Barony
	库尔库莱Koulkoule feud.Barony
	尼姆罗Nimro      feud.Barony
	瓦拉Wara        feud.Barony
	于埃Yue         feud.Barony
}

func (c *瓦达伊WadaiCounty) BAbeche阿贝歇() feud.Barony {
	return c.阿贝歇Abeche
}

func (c *瓦达伊WadaiCounty) BAsafik阿萨菲克() feud.Barony {
	return c.阿萨菲克Asafik
}

func (c *瓦达伊WadaiCounty) BGourmoga古尔莫加() feud.Barony {
	return c.古尔莫加Gourmoga
}

func (c *瓦达伊WadaiCounty) BKoulkoule库尔库莱() feud.Barony {
	return c.库尔库莱Koulkoule
}

func (c *瓦达伊WadaiCounty) BNimro尼姆罗() feud.Barony {
	return c.尼姆罗Nimro
}

func (c *瓦达伊WadaiCounty) BWara瓦拉() feud.Barony {
	return c.瓦拉Wara
}

func (c *瓦达伊WadaiCounty) BYue于埃() feud.Barony {
	return c.于埃Yue
}

var CWadai瓦达伊 WadaiCounty = &瓦达伊WadaiCounty{}

func init() {
	f := CWadai瓦达伊.(*瓦达伊WadaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1770",
		Title:     "wadai",
		TitleName: "瓦达伊",
		TitleCode: "c_wadai",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿贝歇Abeche = BAbeche阿贝歇
	f.阿贝歇Abeche.SetParent(f)

	f.阿萨菲克Asafik = BAsafik阿萨菲克
	f.阿萨菲克Asafik.SetParent(f)

	f.古尔莫加Gourmoga = BGourmoga古尔莫加
	f.古尔莫加Gourmoga.SetParent(f)

	f.库尔库莱Koulkoule = BKoulkoule库尔库莱
	f.库尔库莱Koulkoule.SetParent(f)

	f.尼姆罗Nimro = BNimro尼姆罗
	f.尼姆罗Nimro.SetParent(f)

	f.瓦拉Wara = BWara瓦拉
	f.瓦拉Wara.SetParent(f)

	f.于埃Yue = BYue于埃
	f.于埃Yue.SetParent(f)

}
