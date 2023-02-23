package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ForcalquierCounty interface {
	feud.County
	BApt阿普特() feud.Barony
	BBriancon布里昂松() feud.Barony
	BEmbrun昂布兰() feud.Barony
	BForcalquier福卡尔基耶() feud.Barony
	BGap加普() feud.Barony
	BNyons尼永() feud.Barony
	BSisteron锡斯特龙() feud.Barony
	BVaison韦松() feud.Barony
}

type 福卡尔基耶ForcalquierCounty struct {
	feud.BaseCounty
	阿普特Apt           feud.Barony
	布里昂松Briancon     feud.Barony
	昂布兰Embrun        feud.Barony
	福卡尔基耶Forcalquier feud.Barony
	加普Gap            feud.Barony
	尼永Nyons          feud.Barony
	锡斯特龙Sisteron     feud.Barony
	韦松Vaison         feud.Barony
}

func (c *福卡尔基耶ForcalquierCounty) BApt阿普特() feud.Barony {
	return c.阿普特Apt
}

func (c *福卡尔基耶ForcalquierCounty) BBriancon布里昂松() feud.Barony {
	return c.布里昂松Briancon
}

func (c *福卡尔基耶ForcalquierCounty) BEmbrun昂布兰() feud.Barony {
	return c.昂布兰Embrun
}

func (c *福卡尔基耶ForcalquierCounty) BForcalquier福卡尔基耶() feud.Barony {
	return c.福卡尔基耶Forcalquier
}

func (c *福卡尔基耶ForcalquierCounty) BGap加普() feud.Barony {
	return c.加普Gap
}

func (c *福卡尔基耶ForcalquierCounty) BNyons尼永() feud.Barony {
	return c.尼永Nyons
}

func (c *福卡尔基耶ForcalquierCounty) BSisteron锡斯特龙() feud.Barony {
	return c.锡斯特龙Sisteron
}

func (c *福卡尔基耶ForcalquierCounty) BVaison韦松() feud.Barony {
	return c.韦松Vaison
}

var CForcalquier福卡尔基耶 ForcalquierCounty = &福卡尔基耶ForcalquierCounty{}

func init() {
	f := CForcalquier福卡尔基耶.(*福卡尔基耶ForcalquierCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "229",
		Title:     "forcalquier",
		TitleName: "福卡尔基耶",
		TitleCode: "c_forcalquier",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普特Apt = BApt阿普特
	f.阿普特Apt.SetParent(f)

	f.布里昂松Briancon = BBriancon布里昂松
	f.布里昂松Briancon.SetParent(f)

	f.昂布兰Embrun = BEmbrun昂布兰
	f.昂布兰Embrun.SetParent(f)

	f.福卡尔基耶Forcalquier = BForcalquier福卡尔基耶
	f.福卡尔基耶Forcalquier.SetParent(f)

	f.加普Gap = BGap加普
	f.加普Gap.SetParent(f)

	f.尼永Nyons = BNyons尼永
	f.尼永Nyons.SetParent(f)

	f.锡斯特龙Sisteron = BSisteron锡斯特龙
	f.锡斯特龙Sisteron.SetParent(f)

	f.韦松Vaison = BVaison韦松
	f.韦松Vaison.SetParent(f)

}
