package entity

type ValidatedPayment struct {
	Payment
	isValidated bool
}

func (vp *ValidatedPayment) IsValid() bool {
	return vp.isValidated
}

func NewValidatedPayment(payment *Payment) (*ValidatedPayment, error) {
	if err := payment.validate(); err != nil {
		return nil, err
	}

	return &ValidatedPayment{
		Payment:     *payment,
		isValidated: true,
	}, nil
}
