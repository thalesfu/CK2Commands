package feudal

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

func writeGiveTitle(writer *bufio.Writer, peopleId int, feud feud.Feud) {
	writer.WriteString(fmt.Sprintf("give_title %s %d\n", feud.GetTitleCode(), peopleId))
}
