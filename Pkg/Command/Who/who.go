package Who

import (
	"Thalox/Pkg/Utils"
	"github.com/bwmarrin/discordgo"
)

func Who(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	response := "" +
		"Discord, oyun toplulukları, arkadaş grupları, ve işbirliği platformları için vazgeçilmez bir iletişim aracı haline geldi. " +
		"Discord sunucuları, özel sohbetler ve çok daha fazlasıyla birlikte, sunucu sahipleri ve üyeler için daha fazla kontrol ve eğlence sunmak amacıyla birçok özel bot geliştirilmektedir." +
		"İşte bu noktada, bleagle ve spear adlı geliştiricilerin KingGo adlı Discord botu devreye giriyor.\n\nKingGo: Kimdir ve Ne Yapar?\n\nKingGo, Discord sunucularınızı daha verimli ve eğlenceli hale getiren özellikler sunan bir Discord botudur. " +
		"bleagle ve spear tarafından geliştirilen bu bot, Discord topluluğunun bir parçası olmak isteyenler için heyecan verici bir yol sunar.\n\nTemel Özellikler\n\nModerasyon Araçları: KingGo, sunucu moderasyonunu kolaylaştırmak için bir dizi araç sunar." +
		"Bu araçlar, kullanıcı yönetimi, mesaj silme, uyarılar ve daha fazlasını içerir.\n\nEğlence ve Oyunlar: Sunucunuzdaki üyelerin eğlencesini artırmak için KingGo, çeşitli eğlence komutları ve oyunlar sunar. Örneğin, kullanıcılar arası bir mini yarışma düzenlemek isterseniz, KingGo'nun \"Trivia\" komutunu kullanabilirsiniz." +
		"\n\nMüzik Çalma: KingGo, sunucunuzda müzik dinlemenizi sağlar. Kullanıcılar, botu bir DJ olarak kullanabilir ve YouTube üzerinden müzik çalabilirler.\n\nÖzel Komutlar ve Kişiselleştirme: KingGo, sunucunuz için özel komutlar oluşturmanıza ve botu sunucunuza uyacak şekilde özelleştirmenize olanak tanır.\n\nEkonomi ve Puan Sistemi: Sunucunuzdaki aktiviteyi artırmak için, KingGo, bir ekonomi ve puan sistemi sunar. Üyeler aktivite gösterdikçe puan kazanabilir ve sunucunuzda rekabeti artırabilirsiniz.\n\nSonuç\n\nbleagle ve spear tarafından geliştirilen KingGo, Discord sunucularınızı daha eğlenceli ve düzenli bir hale getirmek için mükemmel bir seçenektir." +
		"Moderasyon araçlarından eğlenceye ve müziğe kadar geniş bir yelpazede özellikler sunarak sunucu yönetimini ve etkileşimini kolaylaştırır.\n\nBu iki geliştirici, Discord topluluğuna katkıda bulunmak için KingGo'yu yaratırken büyük bir çaba harcadı ve sürekli olarak güncellemelerle sunucu sahiplerine ve üyelere daha fazla seçenek sunmayı hedefliyor." +
		"KingGo, Discord sunucularınızı daha da büyütebilir ve geliştirebilir.\n\nKingGo'nun en son sürümünü Discord sunucunuza ekleyerek, sunucunuzun bir adım önde olmasını sağlayabilirsiniz. Bleagle ve spear gibi yetenekli geliştiricilerin bu botu yaratmış olmaları, KingGo'nun kalitesinin ve güvenilirliğinin garantisidir."

	splitMessage := Utils.SplitText(response, 2000)

	for _, message := range splitMessage {
		s.ChannelMessageSend(m.ChannelID, message)
	}

}
