package Help

import (
	"Thalox/Pkg/Utils"
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	test := "**Komutlar**\n\n🚀 " +
		"`>komut1`: Bu komutu kullanarak şunu yapabilirsiniz.\nÖrnek kullanım: `>komut1 arguman1 arguman2`\n\n🌟 " +
		"`>komut2`: Bu komutu kullanarak şunu yapabilirsiniz.\nÖrnek kullanım: `>komut2 arguman`\n\n🎉 " +
		"`>komut3`: Bu komutu kullanarak şunu yapabilirsiniz.\nÖrnek kullanım: `>komut3`\n\n❓ " +
		"Daha fazla bilgi için [web sitemizi](http://thalox.net) ziyaret edebilirsiniz.\n"
	splitMessage := Utils.SplitText(test, 2000)

	for _, message := range splitMessage {
		s.ChannelMessageSend(m.ChannelID, message)
	}
}
