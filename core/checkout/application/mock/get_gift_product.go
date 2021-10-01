package mock

import applicationCatalog "gmarcial/eCommerce/core/catalog/application"

//GetGiftProductUseCase interface to mock the GetGiftProductUseCase.
type GetGiftProductUseCase struct {
	ExecuteMock func() (*applicationCatalog.GiftProductObtained, error)
}

//Execute mocked the use case
func (useCase *GetGiftProductUseCase) Execute() (*applicationCatalog.GiftProductObtained, error) {
	return useCase.ExecuteMock()
}
