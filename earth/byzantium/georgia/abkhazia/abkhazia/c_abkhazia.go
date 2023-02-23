package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbkhaziaCounty interface {
	feud.County
	BAbaatha阿巴特哈() feud.Barony
	BAnakopia阿纳科皮亚() feud.Barony
	BBzyb布兹皮() feud.Barony
	BGagra加格拉() feud.Barony
	BPitsunda皮聪达() feud.Barony
	BTskhoumi茨胡米() feud.Barony
	BZugdidi祖格迪迪() feud.Barony
}

type 阿布哈兹AbkhaziaCounty struct {
	feud.BaseCounty
	阿巴特哈Abaatha   feud.Barony
	阿纳科皮亚Anakopia feud.Barony
	布兹皮Bzyb       feud.Barony
	加格拉Gagra      feud.Barony
	皮聪达Pitsunda   feud.Barony
	茨胡米Tskhoumi   feud.Barony
	祖格迪迪Zugdidi   feud.Barony
}

func (c *阿布哈兹AbkhaziaCounty) BAbaatha阿巴特哈() feud.Barony {
	return c.阿巴特哈Abaatha
}

func (c *阿布哈兹AbkhaziaCounty) BAnakopia阿纳科皮亚() feud.Barony {
	return c.阿纳科皮亚Anakopia
}

func (c *阿布哈兹AbkhaziaCounty) BBzyb布兹皮() feud.Barony {
	return c.布兹皮Bzyb
}

func (c *阿布哈兹AbkhaziaCounty) BGagra加格拉() feud.Barony {
	return c.加格拉Gagra
}

func (c *阿布哈兹AbkhaziaCounty) BPitsunda皮聪达() feud.Barony {
	return c.皮聪达Pitsunda
}

func (c *阿布哈兹AbkhaziaCounty) BTskhoumi茨胡米() feud.Barony {
	return c.茨胡米Tskhoumi
}

func (c *阿布哈兹AbkhaziaCounty) BZugdidi祖格迪迪() feud.Barony {
	return c.祖格迪迪Zugdidi
}

var CAbkhazia阿布哈兹 AbkhaziaCounty = &阿布哈兹AbkhaziaCounty{}

func init() {
	f := CAbkhazia阿布哈兹.(*阿布哈兹AbkhaziaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "600",
		Title:     "abkhazia",
		TitleName: "阿布哈兹",
		TitleCode: "c_abkhazia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴特哈Abaatha = BAbaatha阿巴特哈
	f.阿巴特哈Abaatha.SetParent(f)

	f.阿纳科皮亚Anakopia = BAnakopia阿纳科皮亚
	f.阿纳科皮亚Anakopia.SetParent(f)

	f.布兹皮Bzyb = BBzyb布兹皮
	f.布兹皮Bzyb.SetParent(f)

	f.加格拉Gagra = BGagra加格拉
	f.加格拉Gagra.SetParent(f)

	f.皮聪达Pitsunda = BPitsunda皮聪达
	f.皮聪达Pitsunda.SetParent(f)

	f.茨胡米Tskhoumi = BTskhoumi茨胡米
	f.茨胡米Tskhoumi.SetParent(f)

	f.祖格迪迪Zugdidi = BZugdidi祖格迪迪
	f.祖格迪迪Zugdidi.SetParent(f)

}
