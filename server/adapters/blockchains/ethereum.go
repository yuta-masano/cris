package blockchains

import "github.com/ktr0731/cris/server/usecases/dto"

type EthereumAdapter struct{}

func (a *EthereumAdapter) FindTxByHash(hash string) (*dto.Transaction, error) {
	return nil, nil
}
