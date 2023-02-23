package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PangongCounty interface {
	feud.County
	BAksayqin阿克赛钦() feud.Barony
	BKhumak库尔纳克() feud.Barony
	BKizilyeza克什尔村() feud.Barony
	BPangong班公() feud.Barony
	BSpanggur斯潘古尔() feud.Barony
	BSumdo苏姆多() feud.Barony
	BSurigh萨利吉勒() feud.Barony
}

type 班公PangongCounty struct {
	feud.BaseCounty
	阿克赛钦Aksayqin  feud.Barony
	库尔纳克Khumak    feud.Barony
	克什尔村Kizilyeza feud.Barony
	班公Pangong     feud.Barony
	斯潘古尔Spanggur  feud.Barony
	苏姆多Sumdo      feud.Barony
	萨利吉勒Surigh    feud.Barony
}

func (c *班公PangongCounty) BAksayqin阿克赛钦() feud.Barony {
	return c.阿克赛钦Aksayqin
}

func (c *班公PangongCounty) BKhumak库尔纳克() feud.Barony {
	return c.库尔纳克Khumak
}

func (c *班公PangongCounty) BKizilyeza克什尔村() feud.Barony {
	return c.克什尔村Kizilyeza
}

func (c *班公PangongCounty) BPangong班公() feud.Barony {
	return c.班公Pangong
}

func (c *班公PangongCounty) BSpanggur斯潘古尔() feud.Barony {
	return c.斯潘古尔Spanggur
}

func (c *班公PangongCounty) BSumdo苏姆多() feud.Barony {
	return c.苏姆多Sumdo
}

func (c *班公PangongCounty) BSurigh萨利吉勒() feud.Barony {
	return c.萨利吉勒Surigh
}

var CPangong班公 PangongCounty = &班公PangongCounty{}

func init() {
	f := CPangong班公.(*班公PangongCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1488",
		Title:     "pangong",
		TitleName: "班公",
		TitleCode: "c_pangong",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克赛钦Aksayqin = BAksayqin阿克赛钦
	f.阿克赛钦Aksayqin.SetParent(f)

	f.库尔纳克Khumak = BKhumak库尔纳克
	f.库尔纳克Khumak.SetParent(f)

	f.克什尔村Kizilyeza = BKizilyeza克什尔村
	f.克什尔村Kizilyeza.SetParent(f)

	f.班公Pangong = BPangong班公
	f.班公Pangong.SetParent(f)

	f.斯潘古尔Spanggur = BSpanggur斯潘古尔
	f.斯潘古尔Spanggur.SetParent(f)

	f.苏姆多Sumdo = BSumdo苏姆多
	f.苏姆多Sumdo.SetParent(f)

	f.萨利吉勒Surigh = BSurigh萨利吉勒
	f.萨利吉勒Surigh.SetParent(f)

}
