package service

//import (
//	"context"
//	"webook/internal/repository"
//	"webook/internal/service/sms"
//)
//
//type CodeService interface {
//	Send(ctx context.Context, biz, phone string) error
//	Verify(ctx context.Context, biz, phone, inputCode string) error
//}
//type codeService struct {
//	repo repository.CodeRepository
//	sms  sms.Service
//}
//
//func (c *codeService) Send(ctx context.Context, biz, phone string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *codeService) Verify(ctx context.Context, biz, phone, inputCode string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *codeService) NewCOdeService(repo repository.CodeRepository, smsSvc sms.Service) CodeService {
//	return &codeService{
//		repo: repo,
//		sms:  smsSvc,
//	}
//}
