package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SambhalCounty interface {
	feud.County
	BAmroha庵卢诃() feud.Barony
	BBachhraon婆掣罗蕴() feud.Barony
	BGokal拘伽罗() feud.Barony
	BSambhal三婆罗() feud.Barony
	BSirsi希尔西() feud.Barony
	BTattandapura多檀陀补罗() feud.Barony
	BUjhani乌诃尼() feud.Barony
}

type 三婆罗SambhalCounty struct {
	feud.BaseCounty
	庵卢诃Amroha         feud.Barony
	婆掣罗蕴Bachhraon     feud.Barony
	拘伽罗Gokal          feud.Barony
	三婆罗Sambhal        feud.Barony
	希尔西Sirsi          feud.Barony
	多檀陀补罗Tattandapura feud.Barony
	乌诃尼Ujhani         feud.Barony
}

func (c *三婆罗SambhalCounty) BAmroha庵卢诃() feud.Barony {
	return c.庵卢诃Amroha
}

func (c *三婆罗SambhalCounty) BBachhraon婆掣罗蕴() feud.Barony {
	return c.婆掣罗蕴Bachhraon
}

func (c *三婆罗SambhalCounty) BGokal拘伽罗() feud.Barony {
	return c.拘伽罗Gokal
}

func (c *三婆罗SambhalCounty) BSambhal三婆罗() feud.Barony {
	return c.三婆罗Sambhal
}

func (c *三婆罗SambhalCounty) BSirsi希尔西() feud.Barony {
	return c.希尔西Sirsi
}

func (c *三婆罗SambhalCounty) BTattandapura多檀陀补罗() feud.Barony {
	return c.多檀陀补罗Tattandapura
}

func (c *三婆罗SambhalCounty) BUjhani乌诃尼() feud.Barony {
	return c.乌诃尼Ujhani
}

var CSambhal三婆罗 SambhalCounty = &三婆罗SambhalCounty{}

func init() {
	f := CSambhal三婆罗.(*三婆罗SambhalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1173",
		Title:     "sambhal",
		TitleName: "三婆罗",
		TitleCode: "c_sambhal",
		Baronies:  map[string]feud.Barony{},
	}

	f.庵卢诃Amroha = BAmroha庵卢诃
	f.庵卢诃Amroha.SetParent(f)

	f.婆掣罗蕴Bachhraon = BBachhraon婆掣罗蕴
	f.婆掣罗蕴Bachhraon.SetParent(f)

	f.拘伽罗Gokal = BGokal拘伽罗
	f.拘伽罗Gokal.SetParent(f)

	f.三婆罗Sambhal = BSambhal三婆罗
	f.三婆罗Sambhal.SetParent(f)

	f.希尔西Sirsi = BSirsi希尔西
	f.希尔西Sirsi.SetParent(f)

	f.多檀陀补罗Tattandapura = BTattandapura多檀陀补罗
	f.多檀陀补罗Tattandapura.SetParent(f)

	f.乌诃尼Ujhani = BUjhani乌诃尼
	f.乌诃尼Ujhani.SetParent(f)

}
