package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад — это структурный шаблон проектирования, который обеспечивает упрощенный (но ограниченный) интерфейс для сложной системы классов, библиотеки или платформы.
Хотя Фасад снижает общую сложность приложения, он также помогает переместить нежелательные зависимости в одно место.
*/

// Приведу пример - веб-сервис, который предоставляет API для работы с различными сервисами социальных сетей (Facebook, Twitter, Instagram).
// Мы можем использовать паттерн Фасад, чтобы скрыть сложность взаимодействия с каждой из этих социальных сетей за одним простым интерфейсом.

// FacebookService представляет сервис работы с Facebook.
type FacebookService struct{}

func (s *FacebookService) Login() {
	fmt.Println("Вход в Facebook")
}

func (s *FacebookService) Post(message string) {
	fmt.Printf("Пост на Facebook: %s\n", message)
}

// TwitterService представляет сервис работы с Twitter.
type TwitterService struct{}

func (s *TwitterService) Login() {
	fmt.Println("Вход в Twitter")
}

func (s *TwitterService) Tweet(message string) {
	fmt.Printf("Твит: %s\n", message)
}

// InstagramService представляет сервис работы с Instagram.
type InstagramService struct{}

func (s *InstagramService) Login() {
	fmt.Println("Вход в Instagram")
}

func (s *InstagramService) SharePhoto(photo string) {
	fmt.Printf("Публикация фотографии в Instagram: %s\n", photo)
}

// SocialMediaFacade представляет фасад для работы с различными социальными сетями.
type SocialMediaFacade struct {
	facebookService  *FacebookService
	twitterService   *TwitterService
	instagramService *InstagramService
}

func NewSocialMediaFacade() *SocialMediaFacade {
	return &SocialMediaFacade{
		facebookService:  &FacebookService{},
		twitterService:   &TwitterService{},
		instagramService: &InstagramService{},
	}
}

// Здесь и далее методы фасада используют методы сервисов социальных сетей.

func (f *SocialMediaFacade) ShareOnFacebook(message string) {
	f.facebookService.Login()
	f.facebookService.Post(message)
}

func (f *SocialMediaFacade) ShareOnTwitter(message string) {
	f.twitterService.Login()
	f.twitterService.Tweet(message)
}

func (f *SocialMediaFacade) ShareOnInstagram(photo string) {
	f.instagramService.Login()
	f.instagramService.SharePhoto(photo)
}

func main() {
	socialMediaFacade := NewSocialMediaFacade()

	// Публикация на различных социальных сетях с использованием паттерна Фасад.
	socialMediaFacade.ShareOnFacebook("Мой новый пост")
	socialMediaFacade.ShareOnTwitter("Мой новый твит")
	socialMediaFacade.ShareOnInstagram("Моя новая фотография")
}
