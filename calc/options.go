package calc

import (
    "log"
    "github.com/pkg/errors"
)

type CalculatorServiceOption func(*CalculatorService) error

func Debug(input bool) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        this.Debug = input
        return nil
    }
}

func RemoveBroadcastAddress(input bool) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        this.RemoveBroadcastAddress = input
        return nil
    }
}

func RemoveNetworkAddress(input bool) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        this.RemoveNetworkAddress = input
        return nil
    }
}

func InputPath(input string) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        if input == "" {
            return errors.New("invalid inputPath")
        }

        this.InputPath = input
        return nil
    }
}

func Logger(input *log.Logger) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        if input == nil {
            return errors.New("invalid logger")
        }

        this.Logger = input
        return nil
    }
}
