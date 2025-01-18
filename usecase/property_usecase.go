type PropertyUseCase struct {
	propertyRepo    PropertyRepository
	transactionRepo TransactionRepository
}

func NewPropertyUseCase(pr PropertyRepository, tr TransactionRepository) *PropertyUseCase {
	return &PropertyUseCase{
		propertyRepo:    pr,
		transactionRepo: tr,
	}
}

func (pu *PropertyUseCase) CreateProperty(ctx context.Context, property *Property) error {
	if err := validateProperty(property); err != nil {
		return err
	}

	return pu.propertyRepo.Create(ctx, property)
}