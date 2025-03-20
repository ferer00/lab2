package handler

import (
	"io"
	"lab2/pkg/compute" // Імпорт пакету compute
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := strings.TrimSpace(string(data))
	result, err := compute.PrefixToInfix(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result + "\n"))
	return err
}