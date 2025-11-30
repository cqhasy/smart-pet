package service

type DialogService struct{}

func New() *DialogService {
	return new(DialogService)
}

func (s *DialogService) SayHello() string {
	return "喵喵...我是你的彩虹猫桌宠❤️"
}
