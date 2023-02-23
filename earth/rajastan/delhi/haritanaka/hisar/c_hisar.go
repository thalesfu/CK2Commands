package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HisarCounty interface {
	feud.County
	BAgroha阿迦卢诃() feud.Barony
	BAsigarh阿私姞利呬() feud.Barony
	BAsika阿斯迦() feud.Barony
	BHisar醯娑罗() feud.Barony
	BKalanaur迦罗那郁罗() feud.Barony
	BRohtak卢醯呾迦() feud.Barony
	BSurajkund苏罗阇君荼() feud.Barony
}

type 醯娑罗HisarCounty struct {
	feud.BaseCounty
	阿迦卢诃Agroha     feud.Barony
	阿私姞利呬Asigarh   feud.Barony
	阿斯迦Asika       feud.Barony
	醯娑罗Hisar       feud.Barony
	迦罗那郁罗Kalanaur  feud.Barony
	卢醯呾迦Rohtak     feud.Barony
	苏罗阇君荼Surajkund feud.Barony
}

func (c *醯娑罗HisarCounty) BAgroha阿迦卢诃() feud.Barony {
	return c.阿迦卢诃Agroha
}

func (c *醯娑罗HisarCounty) BAsigarh阿私姞利呬() feud.Barony {
	return c.阿私姞利呬Asigarh
}

func (c *醯娑罗HisarCounty) BAsika阿斯迦() feud.Barony {
	return c.阿斯迦Asika
}

func (c *醯娑罗HisarCounty) BHisar醯娑罗() feud.Barony {
	return c.醯娑罗Hisar
}

func (c *醯娑罗HisarCounty) BKalanaur迦罗那郁罗() feud.Barony {
	return c.迦罗那郁罗Kalanaur
}

func (c *醯娑罗HisarCounty) BRohtak卢醯呾迦() feud.Barony {
	return c.卢醯呾迦Rohtak
}

func (c *醯娑罗HisarCounty) BSurajkund苏罗阇君荼() feud.Barony {
	return c.苏罗阇君荼Surajkund
}

var CHisar醯娑罗 HisarCounty = &醯娑罗HisarCounty{}

func init() {
	f := CHisar醯娑罗.(*醯娑罗HisarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1366",
		Title:     "hisar",
		TitleName: "醯娑罗",
		TitleCode: "c_hisar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿迦卢诃Agroha = BAgroha阿迦卢诃
	f.阿迦卢诃Agroha.SetParent(f)

	f.阿私姞利呬Asigarh = BAsigarh阿私姞利呬
	f.阿私姞利呬Asigarh.SetParent(f)

	f.阿斯迦Asika = BAsika阿斯迦
	f.阿斯迦Asika.SetParent(f)

	f.醯娑罗Hisar = BHisar醯娑罗
	f.醯娑罗Hisar.SetParent(f)

	f.迦罗那郁罗Kalanaur = BKalanaur迦罗那郁罗
	f.迦罗那郁罗Kalanaur.SetParent(f)

	f.卢醯呾迦Rohtak = BRohtak卢醯呾迦
	f.卢醯呾迦Rohtak.SetParent(f)

	f.苏罗阇君荼Surajkund = BSurajkund苏罗阇君荼
	f.苏罗阇君荼Surajkund.SetParent(f)

}
