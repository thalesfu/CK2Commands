package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AniCounty interface {
	feud.County
	BAni阿尼() feud.Barony
	BArtashat阿尔塔沙特() feud.Barony
	BKaruts卡尔斯() feud.Barony
	BKarutsberd卡尔斯贝拉德() feud.Barony
	BMidjnaberd米贾那贝德() feud.Barony
	BOshakan奥沙坎() feud.Barony
	BSurpasdvadzadzin圣萨尔基斯() feud.Barony
	BZvartnots兹瓦尔特诺茨() feud.Barony
}

type 阿尼AniCounty struct {
	feud.BaseCounty
	阿尼Ani                 feud.Barony
	阿尔塔沙特Artashat         feud.Barony
	卡尔斯Karuts             feud.Barony
	卡尔斯贝拉德Karutsberd      feud.Barony
	米贾那贝德Midjnaberd       feud.Barony
	奥沙坎Oshakan            feud.Barony
	圣萨尔基斯Surpasdvadzadzin feud.Barony
	兹瓦尔特诺茨Zvartnots       feud.Barony
}

func (c *阿尼AniCounty) BAni阿尼() feud.Barony {
	return c.阿尼Ani
}

func (c *阿尼AniCounty) BArtashat阿尔塔沙特() feud.Barony {
	return c.阿尔塔沙特Artashat
}

func (c *阿尼AniCounty) BKaruts卡尔斯() feud.Barony {
	return c.卡尔斯Karuts
}

func (c *阿尼AniCounty) BKarutsberd卡尔斯贝拉德() feud.Barony {
	return c.卡尔斯贝拉德Karutsberd
}

func (c *阿尼AniCounty) BMidjnaberd米贾那贝德() feud.Barony {
	return c.米贾那贝德Midjnaberd
}

func (c *阿尼AniCounty) BOshakan奥沙坎() feud.Barony {
	return c.奥沙坎Oshakan
}

func (c *阿尼AniCounty) BSurpasdvadzadzin圣萨尔基斯() feud.Barony {
	return c.圣萨尔基斯Surpasdvadzadzin
}

func (c *阿尼AniCounty) BZvartnots兹瓦尔特诺茨() feud.Barony {
	return c.兹瓦尔特诺茨Zvartnots
}

var CAni阿尼 AniCounty = &阿尼AniCounty{}

func init() {
	f := CAni阿尼.(*阿尼AniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "681",
		Title:     "ani",
		TitleName: "阿尼",
		TitleCode: "c_ani",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尼Ani = BAni阿尼
	f.阿尼Ani.SetParent(f)

	f.阿尔塔沙特Artashat = BArtashat阿尔塔沙特
	f.阿尔塔沙特Artashat.SetParent(f)

	f.卡尔斯Karuts = BKaruts卡尔斯
	f.卡尔斯Karuts.SetParent(f)

	f.卡尔斯贝拉德Karutsberd = BKarutsberd卡尔斯贝拉德
	f.卡尔斯贝拉德Karutsberd.SetParent(f)

	f.米贾那贝德Midjnaberd = BMidjnaberd米贾那贝德
	f.米贾那贝德Midjnaberd.SetParent(f)

	f.奥沙坎Oshakan = BOshakan奥沙坎
	f.奥沙坎Oshakan.SetParent(f)

	f.圣萨尔基斯Surpasdvadzadzin = BSurpasdvadzadzin圣萨尔基斯
	f.圣萨尔基斯Surpasdvadzadzin.SetParent(f)

	f.兹瓦尔特诺茨Zvartnots = BZvartnots兹瓦尔特诺茨
	f.兹瓦尔特诺茨Zvartnots.SetParent(f)

}
