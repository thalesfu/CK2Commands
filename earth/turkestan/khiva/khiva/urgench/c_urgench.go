package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UrgenchCounty interface {
	feud.County
	BAchamayli阿恰迈利() feud.Barony
	BKanlykul坎雷库利() feud.Barony
	BKungrat昆格勒() feud.Barony
	BKunkhodzha昆霍贾() feud.Barony
	BMadminiya马德米尼亚() feud.Barony
	BUlkun乌烈昆() feud.Barony
}

type 马德米尼亚UrgenchCounty struct {
	feud.BaseCounty
	阿恰迈利Achamayli  feud.Barony
	坎雷库利Kanlykul   feud.Barony
	昆格勒Kungrat     feud.Barony
	昆霍贾Kunkhodzha  feud.Barony
	马德米尼亚Madminiya feud.Barony
	乌烈昆Ulkun       feud.Barony
}

func (c *马德米尼亚UrgenchCounty) BAchamayli阿恰迈利() feud.Barony {
	return c.阿恰迈利Achamayli
}

func (c *马德米尼亚UrgenchCounty) BKanlykul坎雷库利() feud.Barony {
	return c.坎雷库利Kanlykul
}

func (c *马德米尼亚UrgenchCounty) BKungrat昆格勒() feud.Barony {
	return c.昆格勒Kungrat
}

func (c *马德米尼亚UrgenchCounty) BKunkhodzha昆霍贾() feud.Barony {
	return c.昆霍贾Kunkhodzha
}

func (c *马德米尼亚UrgenchCounty) BMadminiya马德米尼亚() feud.Barony {
	return c.马德米尼亚Madminiya
}

func (c *马德米尼亚UrgenchCounty) BUlkun乌烈昆() feud.Barony {
	return c.乌烈昆Ulkun
}

var CUrgench马德米尼亚 UrgenchCounty = &马德米尼亚UrgenchCounty{}

func init() {
	f := CUrgench马德米尼亚.(*马德米尼亚UrgenchCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1378",
		Title:     "urgench",
		TitleName: "马德米尼亚",
		TitleCode: "c_urgench",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿恰迈利Achamayli = BAchamayli阿恰迈利
	f.阿恰迈利Achamayli.SetParent(f)

	f.坎雷库利Kanlykul = BKanlykul坎雷库利
	f.坎雷库利Kanlykul.SetParent(f)

	f.昆格勒Kungrat = BKungrat昆格勒
	f.昆格勒Kungrat.SetParent(f)

	f.昆霍贾Kunkhodzha = BKunkhodzha昆霍贾
	f.昆霍贾Kunkhodzha.SetParent(f)

	f.马德米尼亚Madminiya = BMadminiya马德米尼亚
	f.马德米尼亚Madminiya.SetParent(f)

	f.乌烈昆Ulkun = BUlkun乌烈昆
	f.乌烈昆Ulkun.SetParent(f)

}
