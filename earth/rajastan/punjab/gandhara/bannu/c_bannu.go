package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BannuCounty interface {
	feud.County
	BBannu般奴() feud.Barony
	BDinkot邓科特() feud.Barony
	BHangu汉古() feud.Barony
	BHindolia信度梨耶() feud.Barony
	BKalabagh加拉伯克() feud.Barony
	BKarrak卡拉克() feud.Barony
	BKohat科哈特() feud.Barony
}

type 般奴BannuCounty struct {
	feud.BaseCounty
	般奴Bannu      feud.Barony
	邓科特Dinkot    feud.Barony
	汉古Hangu      feud.Barony
	信度梨耶Hindolia feud.Barony
	加拉伯克Kalabagh feud.Barony
	卡拉克Karrak    feud.Barony
	科哈特Kohat     feud.Barony
}

func (c *般奴BannuCounty) BBannu般奴() feud.Barony {
	return c.般奴Bannu
}

func (c *般奴BannuCounty) BDinkot邓科特() feud.Barony {
	return c.邓科特Dinkot
}

func (c *般奴BannuCounty) BHangu汉古() feud.Barony {
	return c.汉古Hangu
}

func (c *般奴BannuCounty) BHindolia信度梨耶() feud.Barony {
	return c.信度梨耶Hindolia
}

func (c *般奴BannuCounty) BKalabagh加拉伯克() feud.Barony {
	return c.加拉伯克Kalabagh
}

func (c *般奴BannuCounty) BKarrak卡拉克() feud.Barony {
	return c.卡拉克Karrak
}

func (c *般奴BannuCounty) BKohat科哈特() feud.Barony {
	return c.科哈特Kohat
}

var CBannu般奴 BannuCounty = &般奴BannuCounty{}

func init() {
	f := CBannu般奴.(*般奴BannuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1191",
		Title:     "bannu",
		TitleName: "般奴",
		TitleCode: "c_bannu",
		Baronies:  map[string]feud.Barony{},
	}

	f.般奴Bannu = BBannu般奴
	f.般奴Bannu.SetParent(f)

	f.邓科特Dinkot = BDinkot邓科特
	f.邓科特Dinkot.SetParent(f)

	f.汉古Hangu = BHangu汉古
	f.汉古Hangu.SetParent(f)

	f.信度梨耶Hindolia = BHindolia信度梨耶
	f.信度梨耶Hindolia.SetParent(f)

	f.加拉伯克Kalabagh = BKalabagh加拉伯克
	f.加拉伯克Kalabagh.SetParent(f)

	f.卡拉克Karrak = BKarrak卡拉克
	f.卡拉克Karrak.SetParent(f)

	f.科哈特Kohat = BKohat科哈特
	f.科哈特Kohat.SetParent(f)

}
