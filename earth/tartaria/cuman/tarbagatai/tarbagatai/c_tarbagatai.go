package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarbagataiCounty interface {
	feud.County
	BAlakol阿拉阔勒() feud.Barony
	BBeskol别斯阔勒() feud.Barony
	BTarbagatai塔尔巴哈台() feud.Barony
	BUsharal乌沙拉尔() feud.Barony
	BYntaly恩塔雷() feud.Barony
	BZhanam扎纳姆() feud.Barony
	BZharbulak扎尔布拉克() feud.Barony
}

type 塔尔巴哈台TarbagataiCounty struct {
	feud.BaseCounty
	阿拉阔勒Alakol      feud.Barony
	别斯阔勒Beskol      feud.Barony
	塔尔巴哈台Tarbagatai feud.Barony
	乌沙拉尔Usharal     feud.Barony
	恩塔雷Yntaly       feud.Barony
	扎纳姆Zhanam       feud.Barony
	扎尔布拉克Zharbulak  feud.Barony
}

func (c *塔尔巴哈台TarbagataiCounty) BAlakol阿拉阔勒() feud.Barony {
	return c.阿拉阔勒Alakol
}

func (c *塔尔巴哈台TarbagataiCounty) BBeskol别斯阔勒() feud.Barony {
	return c.别斯阔勒Beskol
}

func (c *塔尔巴哈台TarbagataiCounty) BTarbagatai塔尔巴哈台() feud.Barony {
	return c.塔尔巴哈台Tarbagatai
}

func (c *塔尔巴哈台TarbagataiCounty) BUsharal乌沙拉尔() feud.Barony {
	return c.乌沙拉尔Usharal
}

func (c *塔尔巴哈台TarbagataiCounty) BYntaly恩塔雷() feud.Barony {
	return c.恩塔雷Yntaly
}

func (c *塔尔巴哈台TarbagataiCounty) BZhanam扎纳姆() feud.Barony {
	return c.扎纳姆Zhanam
}

func (c *塔尔巴哈台TarbagataiCounty) BZharbulak扎尔布拉克() feud.Barony {
	return c.扎尔布拉克Zharbulak
}

var CTarbagatai塔尔巴哈台 TarbagataiCounty = &塔尔巴哈台TarbagataiCounty{}

func init() {
	f := CTarbagatai塔尔巴哈台.(*塔尔巴哈台TarbagataiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1876",
		Title:     "tarbagatai",
		TitleName: "塔尔巴哈台",
		TitleCode: "c_tarbagatai",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉阔勒Alakol = BAlakol阿拉阔勒
	f.阿拉阔勒Alakol.SetParent(f)

	f.别斯阔勒Beskol = BBeskol别斯阔勒
	f.别斯阔勒Beskol.SetParent(f)

	f.塔尔巴哈台Tarbagatai = BTarbagatai塔尔巴哈台
	f.塔尔巴哈台Tarbagatai.SetParent(f)

	f.乌沙拉尔Usharal = BUsharal乌沙拉尔
	f.乌沙拉尔Usharal.SetParent(f)

	f.恩塔雷Yntaly = BYntaly恩塔雷
	f.恩塔雷Yntaly.SetParent(f)

	f.扎纳姆Zhanam = BZhanam扎纳姆
	f.扎纳姆Zhanam.SetParent(f)

	f.扎尔布拉克Zharbulak = BZharbulak扎尔布拉克
	f.扎尔布拉克Zharbulak.SetParent(f)

}
