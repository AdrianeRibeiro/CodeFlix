package usecase

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
)

type PixUseCast struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixkey, err := model.NewPixKey(kind, account, key)
	
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, err
	}

	return pixkey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.Pixkey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
