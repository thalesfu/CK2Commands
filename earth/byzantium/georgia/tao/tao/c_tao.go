package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaoCounty interface {
	feud.County
	BAkhaltsikhe阿哈尔齐赫() feud.Barony
	BBana巴纳() feud.Barony
	BKhakhuli哈胡利() feud.Barony
	BMamrovani马姆罗瓦尼() feud.Barony
	BOshki奥什基() feud.Barony
	BPanaskerti帕纳斯克尔蒂() feud.Barony
	BTortomi托尔图姆() feud.Barony
}

type 陶TaoCounty struct {
	feud.BaseCounty
	阿哈尔齐赫Akhaltsikhe feud.Barony
	巴纳Bana           feud.Barony
	哈胡利Khakhuli      feud.Barony
	马姆罗瓦尼Mamrovani   feud.Barony
	奥什基Oshki         feud.Barony
	帕纳斯克尔蒂Panaskerti feud.Barony
	托尔图姆Tortomi      feud.Barony
}

func (c *陶TaoCounty) BAkhaltsikhe阿哈尔齐赫() feud.Barony {
	return c.阿哈尔齐赫Akhaltsikhe
}

func (c *陶TaoCounty) BBana巴纳() feud.Barony {
	return c.巴纳Bana
}

func (c *陶TaoCounty) BKhakhuli哈胡利() feud.Barony {
	return c.哈胡利Khakhuli
}

func (c *陶TaoCounty) BMamrovani马姆罗瓦尼() feud.Barony {
	return c.马姆罗瓦尼Mamrovani
}

func (c *陶TaoCounty) BOshki奥什基() feud.Barony {
	return c.奥什基Oshki
}

func (c *陶TaoCounty) BPanaskerti帕纳斯克尔蒂() feud.Barony {
	return c.帕纳斯克尔蒂Panaskerti
}

func (c *陶TaoCounty) BTortomi托尔图姆() feud.Barony {
	return c.托尔图姆Tortomi
}

var CTao陶 TaoCounty = &陶TaoCounty{}

func init() {
	f := CTao陶.(*陶TaoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "680",
		Title:     "tao",
		TitleName: "陶",
		TitleCode: "c_tao",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿哈尔齐赫Akhaltsikhe = BAkhaltsikhe阿哈尔齐赫
	f.阿哈尔齐赫Akhaltsikhe.SetParent(f)

	f.巴纳Bana = BBana巴纳
	f.巴纳Bana.SetParent(f)

	f.哈胡利Khakhuli = BKhakhuli哈胡利
	f.哈胡利Khakhuli.SetParent(f)

	f.马姆罗瓦尼Mamrovani = BMamrovani马姆罗瓦尼
	f.马姆罗瓦尼Mamrovani.SetParent(f)

	f.奥什基Oshki = BOshki奥什基
	f.奥什基Oshki.SetParent(f)

	f.帕纳斯克尔蒂Panaskerti = BPanaskerti帕纳斯克尔蒂
	f.帕纳斯克尔蒂Panaskerti.SetParent(f)

	f.托尔图姆Tortomi = BTortomi托尔图姆
	f.托尔图姆Tortomi.SetParent(f)

}
