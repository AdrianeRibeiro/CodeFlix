package usecase

import (
	"github.com/AdrianeRibeiro/CodePix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixkey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixkey)

	if pixkey.ID == "" {
		return nil, err
	}

	return pixkey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
