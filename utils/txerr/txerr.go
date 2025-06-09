package txerr

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Decoder interface {
	Decode(err error) error
}

type errorReasoner interface {
	ErrorData() interface{}
}

type decoder struct {
	ABIs []*abi.ABI
}

func NewDecoder(abis ...*abi.ABI) Decoder {
	return &decoder{ABIs: abis}
}

func (d *decoder) Decode(err error) error {
	originalErr := err
	unknownErr := fmt.Errorf("%w: unknown error", originalErr)

	for ; err != nil; err = errors.Unwrap(err) {
		_, ok := err.(errorReasoner)
		if ok {
			break
		}
	}

	if err == nil {
		return unknownErr
	}

	asErr, isErrorWithData := err.(errorReasoner)
	if !isErrorWithData {
		return unknownErr
	}

	errData, ok := asErr.ErrorData().(string)
	if !ok {
		return unknownErr
	}

	errWithData := fmt.Errorf("err data: %s", errData)

	hexErr, err := hexutil.Decode(errData)
	if err != nil || len(hexErr) < 4 {
		return fmt.Errorf("%w: %w", originalErr, errWithData)
	}

	for _, a := range d.ABIs {
		parsedError, _ := a.ErrorByID([4]byte(hexErr[:4]))
		if parsedError != nil {
			arguments := map[string]interface{}{}
			errorsDecoded := fmt.Errorf("%w: %w", originalErr, errors.New(parsedError.String()))

			unpackingArgErr := parsedError.Inputs.UnpackIntoMap(arguments, hexErr[4:])
			if unpackingArgErr == nil {
				return fmt.Errorf("%w: args %w", errorsDecoded, errors.New(fmt.Sprintf("%+v", arguments)))
			}

			return errorsDecoded
		}
	}

	return fmt.Errorf("%w: %w", originalErr, errWithData)
}
